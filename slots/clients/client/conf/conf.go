package conf

import "github.com/jinzhu/configor"

const HOST = "1127.0.0.1:9002"
const PATH = "/v1/logicPack/Slot6SpinConn"

var CliConf struct {
	Host string `default:"127.0.0.1:9002"`
	Port string `default:"9002"`
}

func CliConfInit() {

	configor.Load(&CliConf, "")
}
