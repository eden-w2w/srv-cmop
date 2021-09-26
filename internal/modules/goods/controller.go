package goods

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"github.com/sirupsen/logrus"
	"sync"

	"github.com/eden-w2w/srv-cmop/internal/contants/errors"
	"github.com/eden-w2w/srv-cmop/internal/databases"
)

var controller *Controller

func GetController() *Controller {
	if controller == nil {
		controller = newController(global.Config.MasterDB)
	}
	return controller
}

// Controller 商品控制器，兼顾库存管理的能力
type Controller struct {
	db       sqlx.DBExecutor
	managers map[uint64]sync.Mutex
}

func newController(db sqlx.DBExecutor) *Controller {
	goods := databases.Goods{}
	goodsList, err := goods.List(db, nil)
	if err != nil {
		logrus.Panicf("goods.newController goods.List(db, nil) err: %v", err)
	}
	managers := make(map[uint64]sync.Mutex)
	for _, g := range goodsList {
		managers[g.GoodsID] = sync.Mutex{}
	}
	return &Controller{db: db, managers: managers}
}

func (c Controller) GetGoods(p GetGoodsParams) ([]databases.Goods, error) {
	m := databases.Goods{}
	goods, err := m.List(c.db, p.Conditions(c.db), p.Additions()...)
	if err != nil {
		logrus.Errorf("[GetGoods] m.List err: %v, params: %+v", err, p)
		return nil, errors.InternalError
	}
	return goods, nil
}

func (c Controller) GetGoodsByID(goodsID uint64) (*databases.Goods, error) {
	m := &databases.Goods{GoodsID: goodsID}
	err := m.FetchByGoodsID(c.db)
	if err != nil {
		logrus.Errorf("[GetGood] m.FetchByGoodsID err: %v, goodsID: %d", err, goodsID)
		return nil, errors.InternalError
	}
	return m, nil
}

func (c Controller) LockInventory(db sqlx.DBExecutor, goodsID uint64, amount uint32) error {
	if locker, ok := c.managers[goodsID]; ok {
		locker.Lock()
		defer locker.Unlock()

		goods := databases.Goods{GoodsID: goodsID}
		err := goods.FetchByGoodsID(db)
		if err != nil {
			logrus.Errorf("[LockInventory] goods.FetchByGoodsID(db) err: %v, goodsID: %d", err, goodsID)
			return errors.InternalError
		}

		inventory := goods.Inventory - uint64(amount)
		f := builder.FieldValues{
			"Inventory": inventory,
		}
		err = goods.UpdateByGoodsIDWithMap(db, f)
		if err != nil {
			logrus.Errorf("[LockInventory] goods.UpdateByGoodsIDWithStruct(db) err: %v, goodsID: %d, fields: %+v", err, goodsID, f)
			return errors.InternalError
		}

		return nil
	}

	logrus.Errorf("[LockInventory] goodsID not found, goodsID: %d", goodsID)
	return errors.NotFound
}

func (c Controller) UnlockInventory(db sqlx.DBExecutor, goodsID uint64, amount uint32) error {
	if locker, ok := c.managers[goodsID]; ok {
		locker.Lock()
		defer locker.Unlock()

		goods := databases.Goods{GoodsID: goodsID}
		err := goods.FetchByGoodsID(db)
		if err != nil {
			logrus.Errorf("[UnlockInventory] goods.FetchByGoodsID(db) err: %v, goodsID: %d", err, goodsID)
			return errors.InternalError
		}

		inventory := goods.Inventory + uint64(amount)
		f := builder.FieldValues{
			"Inventory": inventory,
		}
		err = goods.UpdateByGoodsIDWithMap(db, f)
		if err != nil {
			logrus.Errorf("[LockInventory] goods.UpdateByGoodsIDWithStruct(db) err: %v, goodsID: %d, fields: %+v", err, goodsID, f)
			return errors.InternalError
		}

		return nil
	}

	logrus.Errorf("[UnlockInventory] goodsID not found, goodsID: %d", goodsID)
	return errors.NotFound
}

func (c Controller) UpdateGoods(goodsID uint64, params UpdateGoodsParams) error {
	model := &databases.Goods{GoodsID: goodsID}
	if params.Name != "" {
		model.Name = params.Name
	}
	if params.Comment != "" {
		model.Comment = params.Comment
	}
	if params.DispatchAddr != "" {
		model.DispatchAddr = params.DispatchAddr
	}
	if params.Sales != 0 {
		model.Sales = params.Sales
	}
	if params.MainPicture != "" {
		model.MainPicture = params.MainPicture
	}
	if len(params.Pictures) > 0 {
		model.Pictures = params.Pictures
	}
	if len(params.Specifications) > 0 {
		model.Specifications = params.Specifications
	}
	if len(params.Activities) > 0 {
		model.Activities = params.Activities
	}
	if params.LogisticPolicy != "" {
		model.LogisticPolicy = params.LogisticPolicy
	}
	if params.Price != 0 {
		model.Price = params.Price
	}
	if params.Inventory != 0 {
		model.Inventory = params.Inventory
	}
	if params.Detail != "" {
		model.Detail = params.Detail
	}
	err := model.UpdateByGoodsIDWithStruct(c.db)
	if err != nil {
		logrus.Errorf("[UpdateGoods] model.UpdateByGoodsIDWithStruct err: %v, params: %+v", err, params)
		return errors.InternalError
	}
	return nil
}
