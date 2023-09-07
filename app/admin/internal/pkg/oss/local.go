package oss

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
)

type localClient struct {
	log *log.Helper
	dir string
}

func newLocalClient(log *log.Helper, config *conf.OssLocalConfig) (*localClient, error) {
	return &localClient{log: log, dir: config.Dir}, nil
}

func (c *localClient) UploadFile(file multipart.File, path string) (string, error) {
	path = filepath.Join(c.dir, path)
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0744); err != nil {
		return "", err
	}
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return "/", err
}
