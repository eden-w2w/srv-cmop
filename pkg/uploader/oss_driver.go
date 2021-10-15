package uploader

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"io"
	"net/url"
)

type AliyunOSSDriver struct {
	endpoint     string
	accessKey    string
	accessSecret string
	bucketName   string
	bucketHost   string
	client       *oss.Client
	bucket       *oss.Bucket
}

func NewAliyunOssDriver(opts ...func(*AliyunOSSDriver)) *AliyunOSSDriver {
	d := &AliyunOSSDriver{}
	for _, opt := range opts {
		opt(d)
	}

	if d.endpoint == "" || d.accessKey == "" || d.accessSecret == "" || d.bucketName == "" {
		logrus.Panic("[NewAliyunOssDriver] need endpoint/accessKey/accessSecret/bucketName")
	}

	var err error
	d.client, err = oss.New(d.endpoint, d.accessKey, d.accessSecret, oss.InsecureSkipVerify(true))
	if err != nil {
		logrus.Panicf("[NewAliyunOssDriver] oss.New err: %v", err)
	}
	d.bucket, err = d.client.Bucket(d.bucketName)
	if err != nil {
		logrus.Panicf("[NewAliyunOssDriver] d.client.Bucket(d.bucketName) err: %v, bucket: %s", err, d.bucketName)
	}

	endpoint, err := url.Parse(d.endpoint)
	if err != nil {
		logrus.Panicf("[NewAliyunOssDriver] url.Parse(d.endpoint) err: %v, endpoint: %s", err, d.endpoint)
	}
	d.bucketHost = fmt.Sprintf("%s://%s.%s", endpoint.Scheme, d.bucketName, endpoint.Host)
	logrus.Infof("[AliyunOSSDriver] initialized, bucketHost: %s", d.bucketHost)
	return d
}

func WithEndpoint(endpoint string) func(*AliyunOSSDriver) {
	return func(d *AliyunOSSDriver) {
		d.endpoint = endpoint
	}
}

func WithAccessKeySecret(accessKey, accessSecret string) func(*AliyunOSSDriver) {
	return func(d *AliyunOSSDriver) {
		d.accessKey = accessKey
		d.accessSecret = accessSecret
	}
}

func WithBucketName(bucket string) func(*AliyunOSSDriver) {
	return func(d *AliyunOSSDriver) {
		d.bucketName = bucket
	}
}

func (d AliyunOSSDriver) PutObject(key string, io io.Reader) (string, error) {
	err := d.bucket.PutObject(key, io)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", d.bucketHost, key), nil
}
