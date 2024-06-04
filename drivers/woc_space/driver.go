package woc_space

import (
	"context"
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/errs"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/pkg/utils"
)

type WocSpace struct {
	model.Storage
	Addition
	conf Conf
}

func (d WocSpace) Config() driver.Config {
	return config
}

func (d WocSpace) GetAddition() driver.Additional {
	return &d.Addition
}

func (d WocSpace) Init(ctx context.Context) error {
	// TODO login / refresh token
	//op.MustSaveDriverStorage(d)
	return nil
}

func (d WocSpace) Drop(ctx context.Context) error {
	return nil
}

func (d WocSpace) List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error) {
	// TODO return the files list, required
	spaces, err := d.GetSpaces()
	if err != nil {
		return nil, err
	}
	return utils.SliceConvert(spaces, func(f Space) (model.Obj, error) {
		return &model.Object{
			Name:     f.Name,
			ID:       f.Guid,
			Size:     int64(f.Size),
			IsFolder: true,
		}, nil
	})
}

func (d WocSpace) Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error) {
	// TODO return link of file, required
	return nil, errs.NotImplement
}

func (d WocSpace) MakeDir(ctx context.Context, parentDir model.Obj, dirName string) (model.Obj, error) {
	// TODO create folder, optional
	return nil, errs.NotImplement
}

func (d WocSpace) Move(ctx context.Context, srcObj, dstDir model.Obj) (model.Obj, error) {
	// TODO move obj, optional
	return nil, errs.NotImplement
}

func (d WocSpace) Rename(ctx context.Context, srcObj model.Obj, newName string) (model.Obj, error) {
	// TODO rename obj, optional
	return nil, errs.NotImplement
}

func (d WocSpace) Copy(ctx context.Context, srcObj, dstDir model.Obj) (model.Obj, error) {
	// TODO copy obj, optional
	return nil, errs.NotImplement
}

func (d WocSpace) Remove(ctx context.Context, obj model.Obj) error {
	// TODO remove obj, optional
	return errs.NotImplement
}

func (d WocSpace) Put(ctx context.Context, dstDir model.Obj, stream model.FileStreamer, up driver.UpdateProgress) (model.Obj, error) {
	// TODO upload file, optional
	return nil, errs.NotImplement
}

//func (d WocSpace) Other(ctx context.Context, args model.OtherArgs) (interface{}, error) {
//	return nil, errs.NotSupport
//}

var _ driver.Driver = (*WocSpace)(nil)
