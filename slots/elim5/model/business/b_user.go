package business

import (
	"elim5/enum"
	"elim5/global"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User 结构体
type User struct {
	global.GVA_MODEL
	Username   string `json:"username" form:"username" gorm:"index;column:username;comment:用户名;size:30;"`
	Password   string `json:"password,omitempty" form:"password" gorm:"column:password;comment:密码;size:200;"`
	Uuid       string `json:"uuid,omitempty" form:"uuid" gorm:"index;column:uuid;comment:UUID;"`
	NickName   string `json:"nickName,omitempty" form:"nickName" gorm:"column:nick_name;comment:昵称;size:50;"`
	Phone      string `json:"phone,omitempty" form:"phone" gorm:"column:phone;comment:手机号;size:30;"`
	Email      string `json:"email,omitempty" form:"email" gorm:"column:email;comment:邮箱;size:100;"`
	HeaderImg  string `json:"headerImg,omitempty" form:"headerImg" gorm:"column:header_img;comment:头像;size:255;"`
	Status     uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态 1正常 2冻结;size:8;"`
	MerchantId uint   `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:商户ID;size:32;"`

	Type        uint8  `json:"type" form:"type" gorm:"column:type;default:1;comment:类型;size:8;"`
	Ip          string `json:"ip,omitempty" form:"ip" gorm:"column:ip;comment:注册IP;size:30;"`
	LastIp      string `json:"lastIp,omitempty" form:"lastIp" gorm:"column:last_ip;comment:最后登录IP;size:30;"`
	Amount      int64  `json:"amount,omitempty" form:"amount" gorm:"column:amount;default:0;comment:金额;size:64;"`
	Online      uint8  `json:"online,omitempty" form:"online" gorm:"column:online;default:2;comment:是否在线;size:8;"`
	Currency    string `json:"currency" form:"currency" gorm:"column:currency;default:USD;comment:货币;size:12;"`
	IndicateNum uint8  `json:"indicateNum" form:"indicateNum" gorm:"column:indicate_num;default:0;comment:ab测 随机编号;size:8;"`
	Token       string `json:"token" form:"token" gorm:"-"`
	Rtp         int    `json:"rtp" form:"rtp" gorm:"-"`
}

// TableName User 表名
func (User) TableName() string {
	return "b_user"
}

type UserBaseInfo struct {
	ID       uint   `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
}

func (u User) CheckInfo() error {
	if u.MerchantId == 0 {
		return errors.New("商户ID不能为空")
	}
	if len(u.Username) < 3 || len(u.NickName) < 3 {
		return errors.New("用户名、昵称长度不能小于3")
	}
	if len(u.Password) < 3 {
		return errors.New("密码长度不能小于3")
	}
	q := global.GVA_READ_DB.Select("id").Where("username = ? and merchant_id = ?", u.Username, u.MerchantId)
	if u.ID != 0 {
		q = q.Where("id != ?", u.ID)
	}
	err := q.First(&User{}).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	return nil
}

func (UserBaseInfo) TableName() string {
	return "b_user"
}

func (u *User) GetAmount() (int64, error) {
	if u.ID == 0 {
		return 0, nil
	}
	var err error
	u.Amount, err = GetUserAmountByDb(u.ID)
	return u.Amount, err
}

func GetMerchantUser(merchantId uint, username string, columns ...string) (*User, error) {
	user := &User{}
	q := global.GVA_READ_DB.Model(&User{})
	if len(columns) > 0 {
		q = q.Select(columns)
	}
	err := q.Where("username = ? and merchant_id = ?", username, merchantId).
		First(user).Error
	return user, err
}

func GetUserAmountByDb(uid uint) (int64, error) {
	var amount int64
	err := global.GVA_READ_DB.Model(&User{}).Where("id = ?", uid).Pluck("amount", &amount).Error
	return amount, err
}

// ChangeBalanceAndGetDetail 更新余额并获取变更信息
func ChangeBalanceAndGetDetail(tx *gorm.DB, userID uint, changeAmount int64, outAll bool, fn func(tx *gorm.DB, b, c, a int64) error) (beforeAmount, realChange, amount int64, err error) {
	transaction := func(tx *gorm.DB) error {
		tableName := User{}.TableName()
		// 锁定记录并只获取amount字段
		if err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Table(tableName).Select("amount").Where("id = ?", userID).Row().Scan(&beforeAmount); err != nil {
			return errors.New("user not found")
		}

		if outAll {
			// 移出所有余额
			realChange = beforeAmount
			amount = 0

		} else {
			// 移出指定余额
			realChange = changeAmount
			amount = beforeAmount + realChange
			if amount < 0 {
				return errors.New("not sufficient funds")
			}
		}

		if realChange == 0 {
			return nil
		}

		// 更新记录
		if err = tx.Table(tableName).Where("id = ?", userID).Update("amount", amount).Error; err != nil {
			return err
		}

		// 执行回调
		if fn != nil {
			if err = fn(tx, beforeAmount, realChange, amount); err != nil {
				return err
			}
		}

		return nil
	}
	if tx == nil {
		err = global.GVA_DB.Transaction(transaction)
	} else {
		err = transaction(tx)
	}

	if err != nil {
		global.GVA_LOG.Error("ChangeBalanceAndGetDetail err: " + err.Error())
		return 0, 0, 0, err
	}

	return
}

func (u *User) GetCurrency() {
	if u.ID == 0 {
		u.Currency = enum.FUN
		return
	}
	var currency string
	global.GVA_DB.Model(&User{}).Where("id = ?", u.ID).Pluck("currency", &currency)
	if currency == "" {
		u.Currency = enum.FUN
		return
	}
	u.Currency = currency
	return
}

func GetUser(id uint) (*User, error) {
	user := &User{}
	err := global.GVA_READ_DB.First(user, id).Error
	return user, err
}
