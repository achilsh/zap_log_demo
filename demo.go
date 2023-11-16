package main

import (
	"context"

	log_wrapper "github.com/achilsh/zap_log_demo/demo_one"
)

func demo_one() {
	log_wrapper.NewZapLogHandler()
	log_wrapper.ZError(context.Background(), "this is test....")
}

func demo_two() {
	log_wrapper.NewSugaredZapLogHandler(nil)
	log_wrapper.Infof(context.Background(), "this is test 2...")
}

func other_caller() {
	log_wrapper.Infof(context.Background(), "is inner caller..")
}

func demo_three() {
	cfgStr := "{\"path_file\":\"./test.log\",\"file_max_size_mb\":10,\"old_file_remain_day\":1,\"old_file_nums\":2,\"old_file_compress\":false,\"log_level\":\"info\"}"
	cfg := log_wrapper.ParseCfg(cfgStr)
	log_wrapper.NewSugaredZapLogHandler(cfg)
	log_wrapper.Infof(context.Background(), "this is test 3...", "idfadfa")
	////
	log_wrapper.Info(context.Background(), "is os ok....")
	other_caller()

}
func main() {
	//demo_one()
	//demo_two()

	demo_three()
}
