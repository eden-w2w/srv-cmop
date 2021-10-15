package uploader

import "io"

var m *Manager

func GetManager() *Manager {
	if m == nil {
		m = &Manager{}
	}
	return m
}

type uploader interface {
	PutObject(key string, io io.Reader) (string, error)
}

type Manager struct {
	driver uploader
}

func (m *Manager) Init(driverType string, endpoint, accessKey, accessSecret, bucket string) {
	if driverType == "oss" {
		m.driver = NewAliyunOssDriver(WithEndpoint(endpoint), WithAccessKeySecret(accessKey, accessSecret), WithBucketName(bucket))
	}
}

func (m Manager) PutObject(key string, io io.Reader) (string, error) {
	return m.driver.PutObject(key, io)
}
