package main

import (
	"context"
	. "zap_demo/demo_one"
)

func demo_one() {
	NewZapLogHandler()
	ZError(context.Background(), "this is test....")
}

func demo_two() {
	NewSugaredZapLogHandler(nil)
	Infof(context.Background(), "this is test 2...")
}

func demo_three() {
	cfgStr := "{\"path_file\":\"./test.log\",\"file_max_size_mb\":10,\"old_file_remain_day\":1,\"old_file_nums\":2,\"old_file_compress\":false,\"log_level\":\"info\"}"
	cfg := ParseCfg(cfgStr)
	NewSugaredZapLogHandler(cfg)
	Infof(context.Background(), "this is test 3...", "idfadfa")
	////
	Info(context.Background(), "is os ok....")

}
func main() {
	//demo_one()
	//demo_two()

	demo_three()
}
