package test

import (
	"github.com/eden-w2w/lib-modules/databases"
	"testing"
)

var adminModel *databases.Administrators
var orderUserModel *databases.User
var promotionUserModel *databases.User
var orderModel *databases.Order
var logisticsModel *databases.OrderLogistics
var paymentFlowModel *databases.PaymentFlow
var promotionFlowModel []databases.PromotionFlow

func TestAll(t *testing.T) {
	t.Run("testCreateAdmin", testCreateAdmin)

}
