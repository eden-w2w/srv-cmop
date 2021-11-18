package global

import (
	"github.com/eden-framework/courier/transport_grpc"
	"github.com/eden-framework/courier/transport_http"
	"github.com/eden-framework/eden-framework/pkg/client/mysql"
	"github.com/eden-w2w/lib-modules/clients/gaode"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/eden-w2w/lib-modules/modules/settlement_flow"
	"github.com/eden-w2w/lib-modules/modules/wechat"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/eden-w2w/lib-modules/databases"
)

type UploaderConfig struct {
	Type         string
	Endpoint     string
	AccessKey    string
	AccessSecret string
	BucketName   string
}

var Config = struct {
	LogLevel logrus.Level

	// db
	MasterDB *mysql.MySQL
	SlaveDB  *mysql.MySQL

	// administrator
	GRPCServer *transport_grpc.ServeGRPC
	HTTPServer *transport_http.ServeHTTP

	id_generator.SnowflakeConfig

	// 登录设置
	PasswordSalt string
	TokenExpired time.Duration

	// 订单超时时间
	OrderExpireIn time.Duration
	// 订单超时任务配置
	CancelExpiredOrderTask string
	// 对账任务配置
	ReconciliationTask string
	// 预售单任务配置
	BookingFlowsTask string

	// 上传设置
	Uploader UploaderConfig

	// 结算设置
	settlement_flow.SettlementConfig

	ClientGaode *gaode.GaodeClient

	// wechat config
	Wechat wechat.Wechat
}{
	LogLevel: logrus.DebugLevel,

	MasterDB: &mysql.MySQL{Database: databases.Config.DB},
	SlaveDB:  &mysql.MySQL{Database: databases.Config.DB},

	GRPCServer: &transport_grpc.ServeGRPC{
		Port: 8900,
	},
	HTTPServer: &transport_http.ServeHTTP{
		Port:     8800,
		WithCORS: true,
	},
	SnowflakeConfig: id_generator.SnowflakeConfig{
		Epoch:      1288351723598,
		BaseNodeID: 2,
		NodeCount:  100,
		NodeBits:   10,
		StepBits:   12,
	},
	SettlementConfig: settlement_flow.SettlementConfig{
		SettlementType: enums.SETTLEMENT_TYPE__WEEK,
		SettlementDate: 1,
		SettlementRules: []settlement_flow.SettlementRule{
			{
				MinSales:   0,
				MaxSales:   500000,
				Proportion: 0.1,
			},
			{
				MinSales:   500000,
				MaxSales:   5000000,
				Proportion: 0.15,
			},
			{
				MinSales:   5000000,
				MaxSales:   ^uint64(0),
				Proportion: 0.2,
			},
		},
		// 结算等待7天，可能涉及7天内退货
		SettlementDuration: 7 * 24 * time.Hour,
	},
	ClientGaode: &gaode.GaodeClient{},
}
