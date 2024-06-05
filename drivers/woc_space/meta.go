package woc_space

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	driver.RootPath
	Token string `json:"token" required:"true"`
}

var config = driver.Config{
	Name:        "WocSpace",
	DefaultRoot: "/",
}

type Conf struct {
	ua      string
	referer string
	api     string
	pr      string
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &WocSpace{}
	})
}
