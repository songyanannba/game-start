package business

import (
	"elim5/global"
	"elim5/model/common"
	"elim5/utils"
	"elim5/utils/helper"
	"errors"
	"strconv"
	"strings"
)

// Merchant 结构体
type Merchant struct {
	global.GVA_MODEL
	Agent      string `json:"agent" form:"agent" gorm:"column:agent;comment:agent;size:30;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:name;size:255;"`
	Currency   string `json:"currency" form:"currency" gorm:"column:currency;comment:currency;size:30;"`
	Type       uint8  `json:"type" form:"type" gorm:"column:type;default:1;comment:type;size:8;"`
	ApiUrl     string `json:"apiUrl" form:"apiUrl" gorm:"column:api_url;comment:api_url;size:255;"`
	Appkey     string `json:"appkey" form:"appkey" gorm:"column:appkey;comment:appKey;size:50;"`
	Secret     string `json:"secret" form:"secret" gorm:"column:secret;comment:secret;size:50;"`
	ProviderId string `json:"providerId" form:"providerId" gorm:"column:provider_id;comment:provider_id;size:50;"`
	CoinType   uint   `json:"coinType" form:"coinType" gorm:"column:coin_type;default:0;comment:coin_type;size:32;"`
	Status     uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:status;size:8;"`
	Remark     string `json:"remark" form:"remark" gorm:"type:text;column:remark;comment:remark;"`
	Rtps       string `json:"rtps" form:"rtps" gorm:"column:rtps;comment:rtps;size:255;"`
	SetId      uint   `json:"setId" form:"setId" gorm:"column:set_id;default:0;comment:set_id;size:32;"`
	SetName    string `json:"setName" form:"setName" gorm:"column:set_name;comment:set_name;size:255;"`

	common.MerchantExtraParams
}

// TableName Merchant 表名
func (Merchant) TableName() string {
	return "b_merchant"
}

func MerchantIsExist(merchant *Merchant) bool {
	var count int64
	q := global.GVA_DB.Table(merchant.TableName()).Select("id").
		Where("(name = ? || agent = ?)",
			merchant.Name, merchant.Agent)
	if merchant.ID > 0 {
		q.Where("id != ?", merchant.ID)
	}
	q.Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func (m *Merchant) FormatCurrency() error {
	m.Currency = utils.FormatCommandStr(m.Currency)
	if m.Currency == "" {
		return errors.New("currency is empty")
	}
	m.Currency = strings.ToUpper(m.Currency)
	return nil
}

func (m *Merchant) FormatRtps() error {
	_, err := strconv.Atoi(m.RtpGear)
	if err != nil {
		return errors.New("rtp gear must int")
	}

	for _, rtp := range strings.Split(m.Rtps, ",") {
		_, err = strconv.Atoi(rtp)
		if err != nil {
			return errors.New("rtp options format error")
		}
	}
	return nil
}

// CheckCurrency 检查商户是否支持该币种
func (m Merchant) CheckCurrency(currency string) bool {
	supportCurrency := strings.Split(m.Currency, " ")
	return helper.CaseInsensitiveInArr(currency, supportCurrency)
}
