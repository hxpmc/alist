package woc_space

import (
	"context"
	"errors"
	"github.com/alist-org/alist/v3/drivers/base"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/go-resty/resty/v2"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

// do others that not defined in Driver interface

func (d *WocSpace) request(pathname string, method string, callback base.ReqCallback, resp interface{}) ([]byte, error) {
	//u := d.conf.api + pathname
	u := "https://api.woc.space" + pathname
	req := base.RestyClient.R()
	req.SetHeaders(map[string]string{
		"Accept":        "application/json, text/plain, */*",
		"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
		"Authorization": "Bearer " + d.Token,
	})

	if callback != nil {
		callback(req)
	}
	if resp != nil {
		req.SetResult(resp)
	}
	var e Resp
	req.SetError(&e)
	res, err := req.Execute(method, u)
	if err != nil {
		return nil, err
	}

	if e.Code >= 200 {
		return nil, errors.New(e.Message)
	}
	return res.Body(), nil
}

func (d *WocSpace) GetSpaces() ([]Space, error) {
	res := make([]Space, 0)
	var resp SpaceResp
	_, err := d.request("/space/mine", http.MethodGet, func(req *resty.Request) {
	}, &resp)
	if err != nil {
		return nil, err
	}
	res = append(res, resp.Data.Spaces...)
	return res, nil
}

func (d *WocSpace) GetAssets(id string) ([]Asset, error) {
	assets := make([]Asset, 0)
	page := 0
	//size := 30
	query := map[string]string{}
	var resp AssetResp
	for {
		query["page"] = strconv.Itoa(page)
		_, err := d.request("/space/"+id+"/assets", http.MethodGet, func(req *resty.Request) {
			req.SetQueryParams(query)
		}, &resp)
		if err != nil {
			return nil, err
		}
		if page >= resp.Data.TotalPages {
			break
		}
		assets = append(assets, resp.Data.SpaceEntities...)
		page++
	}

	return assets, nil
}
func (d *WocSpace) createSpace(name string) error {
	formData := map[string]string{
		"name": name,
	}
	_, err := d.request("/space/create", http.MethodPost, func(req *resty.Request) {
		req.SetFormData(formData)
	}, nil)
	return err
}
func (d *WocSpace) renameSpace(spaceGuid string, fileName string) error {

	formData := map[string]string{
		"name": fileName,
		"guid": spaceGuid,
	}
	_, err := d.request("/space/update_space_name", http.MethodPost, func(req *resty.Request) {
		req.SetFormData(formData)
	}, nil)
	return err
}
func (d *WocSpace) renameAsset(spaceGuid string, entityGuid string, fileName string) error {

	formData := map[string]string{
		"spaceGuid":  spaceGuid,
		"entityGuid": entityGuid,
		"fileName":   fileName,
	}
	_, err := d.request("/space/rename_entity", http.MethodPost, func(req *resty.Request) {
		req.SetFormData(formData)
	}, nil)
	return err
}
func (d *WocSpace) removeSpace(spaceGuid string) error {
	formData := map[string]string{
		"spaceGuid": spaceGuid,
	}
	_, err := d.request("/space/remove", http.MethodPost, func(req *resty.Request) {
		req.SetFormData(formData)
	}, nil)
	return err
}
func (d *WocSpace) removeAsset(spaceGuid string, entityGuid string) error {
	formData := map[string]string{
		"spaceGuid":  spaceGuid,
		"entityGuid": entityGuid,
	}
	_, err := d.request("/space/entities_remove", http.MethodPost, func(req *resty.Request) {
		req.SetFormData(formData)
	}, nil)
	return err
}
func (d *WocSpace) putFile(ctx context.Context, dstDir model.Obj, stream model.FileStreamer) error {
	tempFile, err := stream.CacheFullInTempFile()
	if err != nil {
		return err
	}
	defer func() {
		_ = tempFile.Close()
	}()

	///initial_file_entity
	baseName := filepath.Base(stream.GetName())
	name := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	extensionName := strings.TrimPrefix(filepath.Ext(baseName), ".")
	var upInitResp UpInitResp
	initialFormData := map[string]string{
		"fileName":      name,
		"extensionName": extensionName,
		"size":          strconv.FormatInt(stream.GetSize(), 10),
		"spaceGuid":     dstDir.GetID(),
	}
	_, err1 := d.request("/space/initial_file_entity", http.MethodPost, func(req *resty.Request) {
		req.SetFormData(initialFormData)
	}, &upInitResp)
	if err1 != nil {
		return err
	}
	respData := upInitResp.Data
	if respData.Supplier == "QI_NIU" {
		_, err := d.upClient.R().SetMultipartFormData(map[string]string{
			"token": respData.Token,
			"key":   respData.Key,
			"fname": stream.GetName(),
		}).SetMultipartField("file", stream.GetName(), stream.GetMimetype(), tempFile).
			Post("https://upload.qiniup.com/")
		if err != nil {
			return err
		}
	} else if upInitResp.Data.Supplier == "HUO_SHAN" {
		uploadTos(respData, tempFile)

	} else {
		return nil
	}

	_, err3 := d.request("/space/file_entity_uploaded", http.MethodPost, func(req *resty.Request) {
		req.SetFormData(map[string]string{
			"spaceGuid": dstDir.GetID(),
			"fileGuid":  upInitResp.Data.Guid,
			"size":      strconv.FormatInt(stream.GetSize(), 10),
		})
	}, nil)
	return err3
}
func (d *WocSpace) getDownLoadUrl(spaceGuid string, entityGuid string) (DownloadResp, error) {
	var resp DownloadResp
	formData := map[string]string{
		"spaceGuid": spaceGuid,
		"guids":     entityGuid,
	}
	_, err := d.request("/space/download", http.MethodPost, func(req *resty.Request) {
		req.SetFormData(formData)
	}, &resp)

	return resp, err
}
func uploadTos(upInit UpInit, file io.Reader) error {
	cred := tos.NewStaticCredentials(upInit.AccessKeyId, upInit.SecretAccessKey)
	cred.WithSecurityToken(upInit.Token)
	// 初始化客户端
	client, err := tos.NewClientV2(upInit.EndPoint, tos.WithRegion(upInit.Region), tos.WithCredentials(cred))
	// 上传对象 Body ， 以 string 对象为例
	if err != nil {
		return err
	}
	// 上传对象
	_, err1 := client.PutObjectV2(context.Background(), &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: upInit.Bucket,
			Key:    upInit.Key,
		},
		Content: file,
	})
	if err1 != nil {
		return err
	}
	return nil
}
