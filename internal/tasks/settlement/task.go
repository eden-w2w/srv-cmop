package settlement

import (
	"fmt"
	"github.com/eden-w2w/lib-modules/modules/settlement_flow"
	"time"
)

func TaskSettlement() {
	_, week := time.Now().ISOWeek()
	_ = settlement_flow.GetController().RunTaskSettlement(fmt.Sprintf("第%d周", week))
}
