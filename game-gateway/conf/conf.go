package conf

import "github.com/jinzhu/configor"

var GatewayConf struct {
	Host string `default:"127.0.0.1:8765"`
}

func GatewayConfInit() {
	configor.Load(&GatewayConf, "")
}
