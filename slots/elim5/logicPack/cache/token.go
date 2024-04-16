package cache

import (
	"elim5/enum"
	"elim5/model/business"
	"elim5/utils"
	"strconv"
	"time"
)

const (
	tokenExpire = 12 * time.Hour //token老化时间
)

func ParseUid(token string) int {
	//i := strings.Index(token, "-")
	//if i == -1 {
	//	return 0
	//}
	//s := token[:i]
	//return utils.IdParse(s)
	return 0
}

func GetTokenKey(id uint) string {
	return "{token}:" + strconv.Itoa(int(id))
}

// SetToken 两种情况 一种给商户授权token 一种自身测试商户的token
//func SetToken(user *business.User) (string, error) {
//	u, _ := uuid.NewV4()
//	user.Token = utils.IdFmt(user.ID) + "-" + u.String()
//	err := utils.SetCache(GetTokenKey(user.ID), user, tokenExpire)
//	if err != nil {
//		global.GVA_LOG.Error("SetToken Error: ", zap.Error(err))
//		return "", err
//	}
//	return user.Token, nil
//}

// SetCustomToken 自定义token
func SetCustomToken(user *business.User) error {
	return utils.SetCache(GetTokenKey(user.ID), user, tokenExpire)
}

func GetUser(id uint) (user *business.User, err error) {
	err = utils.GetCache(GetTokenKey(id), &user)
	return user, err
}

func GetUserByToken(token string) (*business.User, error) {
	id := ParseUid(token)
	if id == 0 {
		return nil, enum.ErrTokenInvalid
	}
	user, err := GetUser(uint(id))
	if err != nil {
		return nil, enum.ErrTokenInvalid
	}
	if user.Token != token {
		return nil, enum.ErrTokenInvalid
	}
	return user, nil
}
