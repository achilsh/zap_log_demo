package demo_one

import (
	"encoding/json"
	"fmt"
)

type LogConfig struct {
	LogPath          string `json:"path_file"`
	FileMaxSizeMB    int64  `json:"file_max_size_mb"`
	OldFileRemainDay int32  `json:"old_file_remain_day"`
	OldFileNums      int32  `json:"old_file_nums"`
	OldCompress      bool   `json:"old_file_compress"`
	LogLevel         string `json:"log_level"`
	LogStdout        bool   `json:"log_std_out"`
}

func ParseCfg(cfg string) *LogConfig {
	var cfgLog *LogConfig = &LogConfig{}
	err := json.Unmarshal([]byte(cfg), cfgLog)
	if err != nil {
		panic(fmt.Sprintf("parse cfg fail, err: %v", err))
	}

	return cfgLog
}
