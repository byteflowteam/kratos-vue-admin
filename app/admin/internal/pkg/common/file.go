package common

import (
	"context"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func FormFile(ctx context.Context) (file multipart.File, handler *multipart.FileHeader, err error) {
	request, ok := http.RequestFromServerContext(ctx)
	if !ok {
		err = errors.New("get file fail")
		return
	}
	return request.FormFile("file")
}

func IsAllowedFileExt(ext string) bool {
	// 判断文件名后缀是否在允许的范围内
	allowedExtList := []string{".jpg", ".jpeg", ".png", ".gif"}
	for _, allowedExt := range allowedExtList {
		if strings.EqualFold(ext, allowedExt) {
			return true
		}
	}
	return false
}
