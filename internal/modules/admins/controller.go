package admins

import (
	"crypto"
	"crypto/sha256"
	"encoding/hex"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/srv-cmop/internal/contants/errors"
	"github.com/eden-w2w/srv-cmop/internal/databases"
	"github.com/eden-w2w/srv-cmop/internal/global"
	"github.com/eden-w2w/srv-cmop/internal/id_generator"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var controller *Controller

func GetController() *Controller {
	if controller == nil {
		controller = newController(global.Config.MasterDB, global.Config.PasswordSalt, global.Config.TokenExpired)
	}
	return controller
}

type Controller struct {
	db              sqlx.DBExecutor
	expiredDuration time.Duration
	salt            string
}

func newController(db sqlx.DBExecutor, salt string, d time.Duration) *Controller {
	return &Controller{db: db, salt: salt, expiredDuration: d}
}

func (c Controller) LoginCheck(params LoginParams) (*databases.Administrators, error) {
	model := &databases.Administrators{UserName: params.UserName}
	err := model.FetchByUserName(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, errors.AdminNotFound
		}
		logrus.Errorf("[LoginCheck] model.FetchByUserName err: %v, params: %+v", err, params)
		return nil, errors.InternalError
	}

	if c.password(params.Password) != model.Password {
		return nil, errors.InvalidUserNamePassword
	}
	return model, nil
}

func (c Controller) password(password string) string {
	hasher := sha256.New()
	hash := hasher.Sum([]byte(password + c.salt))
	return hex.EncodeToString(hash)
}

func (c Controller) GetAdminByToken(token string) (*databases.Administrators, error) {
	model := &databases.Administrators{
		Token: token,
	}
	err := model.FetchByToken(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, errors.AdminNotFound
		}
		logrus.Errorf("[GetAdminByToken] model.FetchByToken err: %v, token: %s", err, token)
		return nil, errors.InternalError
	}
	return model, nil
}

func (c Controller) RefreshToken(id uint64) (*databases.Administrators, error) {
	token := c.generateToken(id)
	model := &databases.Administrators{
		AdministratorsID: id,
		Token:            token,
		ExpiredAt:        datatypes.MySQLTimestamp(time.Now().Add(c.expiredDuration)),
	}
	err := model.UpdateByAdministratorsIDWithStruct(c.db)
	if err != nil {
		logrus.Errorf("[RefreshToken] model.UpdateByAdministratorsIDWithStruct(c.db) err: %v, adminID: %d", err, id)
		return nil, errors.InternalError
	}

	err = model.FetchByAdministratorsID(c.db)
	if err != nil {
		logrus.Errorf("[RefreshToken] model.FetchByAdministratorsID(c.db) err: %v, adminID: %d", err, id)
		return nil, errors.InternalError
	}

	return model, nil
}

func (c Controller) CreateAdmin(params LoginParams) (*databases.Administrators, error) {
	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	model := &databases.Administrators{
		AdministratorsID: id,
		UserName:         params.UserName,
		Password:         c.password(params.Password),
	}
	err := model.Create(c.db)
	if err != nil {
		logrus.Errorf("[CreateAdmin] model.Create err: %v, params: %+v", err, params)
		return nil, errors.InternalError
	}

	model, err = c.RefreshToken(id)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (c Controller) generateToken(userID uint64) string {
	id := strconv.FormatUint(userID, 10)
	t := strconv.FormatInt(time.Now().UnixNano(), 10)
	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(id + t))
	hash := sha256.Sum(nil)
	return hex.EncodeToString(hash)
}
