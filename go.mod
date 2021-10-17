module github.com/eden-w2w/srv-cmop

go 1.16

replace k8s.io/client-go => k8s.io/client-go v0.18.8

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.1.10+incompatible
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/eden-framework/context v0.0.3
	github.com/eden-framework/courier v1.0.5
	github.com/eden-framework/eden-framework v1.2.4-0.20211017075623-cd25f6b59e98
	github.com/eden-framework/sqlx v0.0.1
	github.com/eden-w2w/lib-modules v0.0.5
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
)
