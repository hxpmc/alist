package woc_space

import (
	"errors"
	"github.com/alist-org/alist/v3/drivers/base"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// do others that not defined in Driver interface

func (d *WocSpace) request(pathname string, method string, callback base.ReqCallback, resp interface{}) ([]byte, error) {
	//u := d.conf.api + pathname
	u := "https://api.woc.space" + pathname
	req := base.RestyClient.R()
	req.SetHeaders(map[string]string{
		"Accept":        "application/json, text/plain, */*",
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
	var resp SpaceResp
	_, err := d.request("/space/mine", http.MethodGet, func(req *resty.Request) {
	}, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.Space, nil
}

func (d *WocSpace) GetAssets() ([]Space, error) {
	files := make([]Space, 0)
	var resp SpaceResp
	_, err := d.request("/space/mine", http.MethodGet, func(req *resty.Request) {
	}, &resp)
	if err != nil {
		return nil, err
	}
	return files, nil
}
