package settlement

import (
	"fmt"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/modules/settlement_flow"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"time"
)

func TaskSettlement() {
	var name = ""
	if global.Config.SettlementConfig.SettlementType == enums.SETTLEMENT_TYPE__WEEK {
		_, week := time.Now().ISOWeek()
		name = fmt.Sprintf("第%d周", week)
	} else {
		name = datatypes.MySQLTimestamp(time.Now().AddDate(0, -1, 0)).Format("2006.01")
	}
	_ = settlement_flow.GetController().RunTaskSettlement(name)
}
