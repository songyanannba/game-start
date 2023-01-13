package conf

import "github.com/jinzhu/configor"

var PlayerConf struct {
	Host      string `default:"127.0.0.1:8765"`
	ServiceId string `default:"service_syn"`
}

func PlayerConfInit() {

	configor.Load(&PlayerConf, "")
}
