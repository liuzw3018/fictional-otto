package config

/**
 * @Author: liu zw
 * @Date: 2022/10/20 18:37
 * @File: logger.go
 * @Description: //TODO $
 * @Version:
 */

type Logger struct {
	Driver       string `yaml:"driver" json:"driver"`
	Path         string `yaml:"path" json:"path"`
	Split        string `yaml:"split" json:"split"`
	FileName     string `yaml:"fileName" json:"fileName"`
	MaxRemainCnt uint   `yaml:"maxRemainCnt" json:"maxRemainCnt"`
	Level        string `yaml:"level" json:"level"`
}
