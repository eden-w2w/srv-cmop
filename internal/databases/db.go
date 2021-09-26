package databases

import (
	"github.com/eden-framework/sqlx"
)

var Config = struct {
	DBTest *sqlx.Database
}{
	DBTest: &sqlx.Database{},
}
