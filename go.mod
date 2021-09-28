module github.com/eden-w2w/srv-cmop

go 1.16

replace (
	github.com/eden-w2w/lib-modules => ../lib-modules
	k8s.io/client-go => k8s.io/client-go v0.18.8
)

require (
	github.com/eden-framework/context v0.0.3
	github.com/eden-framework/courier v1.0.5
	github.com/eden-framework/eden-framework v1.2.1
	github.com/eden-framework/sqlx v0.0.1
	github.com/eden-w2w/lib-modules v0.0.0-20210928012527-375167560905
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
)
