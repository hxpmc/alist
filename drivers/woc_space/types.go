package woc_space

import "github.com/alist-org/alist/v3/internal/model"

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Space struct {
	Guid string  `json:"guid"`
	Name string  `json:"name"`
	Size float64 `json:"size"`
}

type Asset struct {
	Guid          string  `json:"guid"`
	FileId        string  `json:"fileId"`
	Name          string  `json:"name"`
	FileSize      float64 `json:"fileSize"`
	Key           string  `json:"key"`
	ExtensionName string  `json:"extensionName"`
	MineType      string  `json:"mimeType"`
}

type SpaceResp struct {
	Resp
	Data struct {
		Spaces []Space `json:"spaces"`
	} `json:"data"`
}
type AssetResp struct {
	Resp
	Data struct {
		CurrentPage   int     `json:"currentPage"`
		TotalPages    int     `json:"totalPages"`
		SpaceEntities []Asset `json:"spaceEntities"`
	} `json:"data"`
}

type UpInit struct {
	Token           string `json:"token"`
	Supplier        string `json:"supplier"`
	Key             string `json:"key"`
	Guid            string `json:"guid"`
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	EndPoint        string `json:"endPoint"`
	Region          string `json:"region"`
	Bucket          string `json:"bucket"`
}

type UpInitResp struct {
	Resp
	Data UpInit `json:"data"`
}

type DownloadResp struct {
	Resp
	Data struct {
		DownloadItems []struct {
			FileName  string  `json:"fileName"`
			Extension string  `json:"extension"`
			Key       string  `json:"key"`
			Size      float64 `json:"size"`
		} `json:"downloadItems"`
	} `json:"data"`
}

type AssetsObj struct {
	model.Object
	SpaceId string `json:"spaceId"`
	Key     string `json:"key"`
}

func fileToAssetsObj(asset Asset, spaceId string) *AssetsObj {
	return &AssetsObj{
		SpaceId: spaceId,
		Key:     asset.Key,
		Object: model.Object{
			Name:     asset.Name + "." + asset.ExtensionName,
			ID:       asset.Guid,
			Size:     int64(asset.FileSize),
			IsFolder: false,
		},
	}
}
