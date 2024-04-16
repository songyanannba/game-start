package template

import (
	"elim5/logicPack/base"
	"elim5/utils/helper"
	"github.com/samber/lo"
)

type Verify struct {
	site   *base.Tag
	verify map[[2]int]bool
	sites  []*base.Tag
}

func NewVerify() *Verify {
	return &Verify{
		site:   &base.Tag{},
		verify: map[[2]int]bool{},
		sites:  make([]*base.Tag, 0),
	}
}

// SetSite 设置当前的查找的标签
func (v *Verify) SetSite(tag *base.Tag) bool {
	if v.Verify(tag) {
		return false
	}

	//第一个不能是wild
	if len(v.sites) == 0 {
		if !tag.IsWild {
			goto end
		} else {
			return false
		}
	}

	//不是wild
	if !tag.IsWild {
		if tag.Name == "" || tag.Name != v.site.Name {
			return false
		}
	} else {
		//是wild
		if !lo.Contains(tag.Include, v.site.Name) {
			//变换标签不包含当前的
			return false
		}
	}

end:
	//设置当前的查找的标签
	if tag.IsWild {
		nowTag := tag.Copy()
		nowTag.Name = v.site.Name
		v.site = nowTag.Copy()
	} else {
		v.site = tag.Copy()
	}

	if ok := v.verify[[2]int{tag.X, tag.Y}]; ok {
		return false
	}
	v.verify[[2]int{tag.X, tag.Y}] = true
	v.sites = append(v.sites, v.site.Copy())
	return true
}

// Verify 是否已经使用过
func (v *Verify) Verify(tag *base.Tag) bool {
	if v.verify[[2]int{tag.X, tag.Y}] && !tag.IsWild {
		return true
	}
	return false
}

// GetSites  获取当前的查找的标签
func (v *Verify) GetSites() []*base.Tag {
	return helper.CopyList(v.sites)
}

// Restart 重置
func (v *Verify) Restart() {
	v.site = &base.Tag{}
	v.sites = make([]*base.Tag, 0)
}

func (v *Verify) ResetVerify(s *SpinInfo) {
	//如果是wild 删除使用标示
	for ints, _ := range v.verify {
		if s.Display[ints[0]][ints[1]].IsWild {
			delete(v.verify, ints)
		}
	}
}
