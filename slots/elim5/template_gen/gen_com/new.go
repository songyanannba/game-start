package gen_com

import (
	"elim5/logicPack/template"
)

func NewTestGameInfo(gTem *GenTemplate) (info *template.SpinInfo, err error) {
	info, err = template.NewGameInfo(gTem, int(gTem.TemGen.Type))
	return
}
