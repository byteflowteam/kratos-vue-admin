package oss

import (
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/byteflowteam/kratos-vue-admin/app/admin/internal/conf"
)

type aliyunClient struct {
	log *log.Helper

	client *oss.Client
	bucket *oss.Bucket
	domain string
}

func newAliyunClient(log *log.Helper, config *conf.OssConfig) (*aliyunClient, error) {
	client, err := oss.New(config.Endpoint, config.AccessKey, config.AccessSecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		return nil, err
	}
	return &aliyunClient{
		log:    log,
		client: client,
		bucket: bucket,
		domain: config.ImgDomain,
	}, nil
}

func (c *aliyunClient) UploadFile(file multipart.File, path string) (string, error) {
	err := c.bucket.PutObject(path, file)
	return c.domain, err
}
