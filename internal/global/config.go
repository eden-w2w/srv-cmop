package global

import (
	"github.com/eden-framework/courier/transport_grpc"
	"github.com/eden-framework/courier/transport_http"
	"github.com/eden-framework/eden-framework/pkg/client/mysql"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/eden-w2w/srv-cmop/internal/databases"
)

type SnowflakeConfig struct {
	Epoch      int64
	BaseNodeID int64
	NodeCount  int64
	NodeBits   uint8
	StepBits   uint8
}

var Config = struct {
	LogLevel logrus.Level

	// db
	MasterDB *mysql.MySQL
	SlaveDB  *mysql.MySQL

	// administrator
	GRPCServer *transport_grpc.ServeGRPC
	HTTPServer *transport_http.ServeHTTP

	SnowflakeConfig

	PasswordSalt string
	TokenExpired time.Duration
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
	SnowflakeConfig: SnowflakeConfig{
		Epoch:      1288351723598,
		BaseNodeID: 2,
		NodeCount:  100,
		NodeBits:   10,
		StepBits:   12,
	},
}
