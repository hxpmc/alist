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
	Guid     string  `json:"guid"`
	FileId   string  `json:"fileId"`
	Name     string  `json:"name"`
	FileSize float64 `json:"fileSize"`
	Key      string  `json:"key"`
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
type UpInitResp struct {
	Resp
	Data struct {
		Token           string `json:"token"`
		Supplier        string `json:"supplier"`
		Key             string `json:"key"`
		Guid            string `json:"guid"`
		AccessKeyId     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
		EndPoint        string `json:"endPoint"`
		Region          string `json:"region"`
		Bucket          string `json:"bucket"`
	} `json:"data"`
}
type AssetsObj struct {
	model.Object
	SpaceId string `json:"spaceId"`
}

func fileToAssetsObj(asset Asset, spaceId string) *AssetsObj {
	return &AssetsObj{
		SpaceId: spaceId,
		Object: model.Object{
			Name:     asset.Name,
			ID:       asset.Guid,
			Size:     int64(asset.FileSize),
			IsFolder: false,
		},
	}
}
