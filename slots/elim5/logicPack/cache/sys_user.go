package cache

import (
	"elim5/global"
	"elim5/model/system"
	"elim5/utils/helper"
	"time"
)

func GetSysUserKey[T int | uint](id T) string {
	return "sys_user_" + helper.Itoa(id)
}

func SetSysUser(user *system.SysUser) {
	global.BlackCache.Set(GetSysUserKey(user.ID), user, 60*time.Minute)
}

func GetSysUser[T int | uint](id T) (*system.SysUser, bool) {
	s, ok := global.BlackCache.Get(GetSysUserKey(id))
	if ok {
		return s.(*system.SysUser), ok
	}
	var user system.SysUser
	err := global.GVA_READ_DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, false
	}
	SetSysUser(&user)
	return &user, true
}

func DelSysUser[T int | uint](id T) {
	global.BlackCache.Delete(GetSysUserKey(id))
}

//func IsMerchant(c *gin.Context, id *uint) error {
//	info := utils.GetUserInfo(c)
//	if info.AuthorityId == 10 {
//		user, ok := GetSysUser(info.ID)
//		if !ok || user.Enable == 2 {
//			return errors.New("account has been disabled")
//		}
//		*id = user.MerchantId
//		return nil
//	}
//	return nil
//}
