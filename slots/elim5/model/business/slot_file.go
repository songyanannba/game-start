package business

type GameFile struct {
	//global.GVA_MODEL
	ID     uint   `json:"id" form:"id" gorm:"column:id;primary_key;auto_increment;comment:主键id;size:32;"`
	GameId uint   `json:"game_id" form:"game_id" gorm:"column:game_id;default:0;comment:游戏id;size:32;"`
	Name   string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:255;"`
	Path   string `json:"path" form:"path" gorm:"column:path;default:0;comment:路径;size:255;"`
	Type   int    `json:"type" form:"type" gorm:"column:type;default:0;comment:文件类型;size:32;"`
	//Pid    uint   `json:"pid" form:"pid" gorm:"column:pid;default:0;comment:父级id;size:32;"`
	//Ext    string `json:"ext" form:"ext" gorm:"column:ext;default:0;comment:扩展名;size:255;"`
}

// TableName GameFile 表名
func (GameFile) TableName() string {
	return "b_game_file"
}

type GameFileRes struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
