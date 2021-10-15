package v0

import (
	"bytes"
	"context"
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/httpx"
	"github.com/eden-w2w/srv-cmop/internal/contants/errors"
	"github.com/eden-w2w/srv-cmop/pkg/uploader"
	"github.com/sirupsen/logrus"
)

func init() {
	AuthRouter.Register(courier.NewRouter(Upload{}))
}

type UploadRequest struct {
	Key  string `in:"body" json:"key"`
	Data []byte `in:"body" json:"data"`
}

type UploadResponse struct {
	Url string `json:"url"`
}

// Upload 上传
type Upload struct {
	httpx.MethodPost
	Body UploadRequest `in:"body"`
}

func (req Upload) Path() string {
	return "/upload"
}

func (req Upload) Output(ctx context.Context) (result interface{}, err error) {
	r := bytes.NewReader(req.Body.Data)
	url, err := uploader.GetManager().PutObject(req.Body.Key, r)
	if err != nil {
		logrus.Errorf("[Upload] uploader.GetManager().PutObject err: %v, key: %s", err, req.Body.Key)
		return nil, errors.InternalError
	}
	return &UploadResponse{Url: url}, nil
}
