package reconciliation

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type TradeBill struct {
	// 交易时间
	TradeTime datatypes.MySQLTimestamp
	// 公众账号ID（小程序AppID）
	AppID string
	// 商户号
	MerchantID string
	// 特别商户号
	SpecialMerchantID string
	// 设备号
	DeviceID string
	// 微信订单号
	PaymentID string
	// 商户订单号
	PaymentFlowID uint64
	// 用户标识
	OpenID string
	// 交易类型
	Type string
	// 交易状态
	Status enums.WechatTradeState
	// 付款银行
	BankName string
	// 货币种类
	CurrencyType string
	// 应结订单金额
	ConcludedOrderPrice int
	// 代金券金额
	CouponPrice int
	// 微信退款单号
	RefundID string
	// 商户退款单号
	RefundFlowID string
	// 退款金额
	RefundPrice int
	// 充值券退款金额
	CouponRefundPrice int
	// 退款类型
	RefundChannel enums.RefundChannel
	// 退款状态
	RefundStatus enums.RefundStatus
	// 商品名称
	GoodsName string
	// 商户数据包
	MerchantData string
	// 手续费
	ServiceCharge int
	// 费率
	Rate string
	// 订单金额
	OrderPrice int
	// 申请退款金额
	ActualRefundPrice int
	// 费率备注
	RateRemark string
}

type TradeBillParser struct {
	rawData []byte
	errors  []string
}

func NewTradeBillParser(data []byte) *TradeBillParser {
	return &TradeBillParser{rawData: data, errors: make([]string, 0)}
}

func (p TradeBillParser) Errors() string {
	return strings.Join(p.errors, ", ")
}

func (p *TradeBillParser) Iterator(check func(bill *TradeBill, last bool) (err error)) error {
	bytesReader := bytes.NewReader(p.rawData)
	reader := bufio.NewReader(bytesReader)
	_, _, _ = reader.ReadLine()
	var lineNo = uint32(1)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			logrus.Errorf("[TradeBillParser].Iterator reader.ReadLine() line: %d, err: %v", lineNo, err)
			return err
		}

		rows := strings.Split(string(line), ",")
		if len(rows) < 27 {
			// TODO 最后两行处理

			// 执行最后检查
			err = check(nil, true)
			if err != nil {
				logrus.Warningf("[TradeBillParser].Iterator check last, err: %v", err)
				return err
			}
			break
		}

		tradeTime, err := datatypes.ParseTimestampFromStringWithLayout(
			strings.TrimLeft(rows[0], "`"),
			"2006-01-02 15:04:05",
		)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator tradeTime datatypes.ParseTimestampFromStringWithLayout line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[0], "`"),
			)
			return err
		}
		paymentFlowID, err := strconv.ParseUint(strings.TrimLeft(rows[6], "`"), 10, 64)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator paymentFlowID strconv.ParseUint line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[6], "`"),
			)
			return err
		}
		status, err := enums.ParseWechatTradeStateFromString(strings.TrimLeft(rows[9], "`"))
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator status enums.ParseWechatTradeStateFromString line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[9], "`"),
			)
			return err
		}
		concludedOrderPrice, err := strconv.ParseFloat(strings.TrimLeft(rows[12], "`"), 64)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator concludedOrderPrice strconv.ParseFloat line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[12], "`"),
			)
			return err
		}
		couponPrice, err := strconv.ParseFloat(strings.TrimLeft(rows[13], "`"), 64)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator couponPrice strconv.ParseFloat line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[13], "`"),
			)
			return err
		}
		refundPrice, err := strconv.ParseFloat(strings.TrimLeft(rows[16], "`"), 64)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator refundPrice strconv.ParseFloat line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[16], "`"),
			)
			return err
		}
		couponRefundPrice, err := strconv.ParseFloat(strings.TrimLeft(rows[17], "`"), 64)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator couponRefundPrice strconv.ParseFloat line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[17], "`"),
			)
			return err
		}
		refundChannel, err := enums.ParseRefundChannelFromString(strings.TrimLeft(rows[18], "`"))
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator refundChannel enums.ParseRefundChannelFromString line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[18], "`"),
			)
			return err
		}
		refundStatus, err := enums.ParseRefundStatusFromString(strings.TrimLeft(rows[19], "`"))
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator refundStatus enums.ParseRefundStatusFromString line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[19], "`"),
			)
			return err
		}
		serviceCharge, err := strconv.ParseFloat(strings.TrimLeft(rows[22], "`"), 64)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator serviceCharge strconv.ParseFloat line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[22], "`"),
			)
			return err
		}
		orderPrice, err := strconv.ParseFloat(strings.TrimLeft(rows[24], "`"), 64)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator orderPrice strconv.ParseFloat line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[24], "`"),
			)
			return err
		}
		actualRefundPrice, err := strconv.ParseFloat(strings.TrimLeft(rows[25], "`"), 64)
		if err != nil {
			logrus.Errorf(
				"[TradeBillParser].Iterator actualRefundPrice strconv.ParseFloat line: %d, err: %v, str: %s",
				lineNo,
				err,
				strings.TrimLeft(rows[25], "`"),
			)
			return err
		}

		bill := &TradeBill{
			TradeTime:           tradeTime,
			AppID:               strings.TrimLeft(rows[1], "`"),
			MerchantID:          strings.TrimLeft(rows[2], "`"),
			SpecialMerchantID:   strings.TrimLeft(rows[3], "`0"),
			DeviceID:            strings.TrimLeft(rows[4], "`"),
			PaymentID:           strings.TrimLeft(rows[5], "`"),
			PaymentFlowID:       paymentFlowID,
			OpenID:              strings.TrimLeft(rows[7], "`"),
			Type:                strings.TrimLeft(rows[8], "`"),
			Status:              status,
			BankName:            strings.TrimLeft(rows[10], "`"),
			CurrencyType:        strings.TrimLeft(rows[11], "`"),
			ConcludedOrderPrice: int(concludedOrderPrice * 100),
			CouponPrice:         int(couponPrice * 100),
			RefundID:            strings.TrimLeft(rows[14], "`"),
			RefundFlowID:        strings.TrimLeft(rows[15], "`"),
			RefundPrice:         int(refundPrice * 100),
			CouponRefundPrice:   int(couponRefundPrice * 100),
			RefundChannel:       refundChannel,
			RefundStatus:        refundStatus,
			GoodsName:           strings.TrimLeft(rows[20], "`"),
			MerchantData:        strings.TrimLeft(rows[21], "`"),
			ServiceCharge:       int(serviceCharge * 100),
			Rate:                strings.TrimLeft(rows[23], "`"),
			OrderPrice:          int(orderPrice * 100),
			ActualRefundPrice:   int(actualRefundPrice * 100),
			RateRemark:          strings.TrimLeft(rows[26], "`"),
		}

		err = check(bill, false)
		if err != nil {
			logrus.Warningf("[TradeBillParser].Iterator check line: %d, err: %v, bill: %+v", lineNo, err, bill)
			p.errors = append(p.errors, fmt.Sprintf("[line: %d, err: %v]", lineNo, err))
		}

		lineNo++
	}

	return nil
}
