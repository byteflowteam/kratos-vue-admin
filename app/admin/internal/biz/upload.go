package biz

import (
	"mime/multipart"
)

type OssRepo interface {
	UploadFile(file multipart.File, path string) (string, error)
}
