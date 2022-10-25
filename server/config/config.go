package config

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:55
 * @File: config.go
 * @Description: //TODO $
 * @Version:
 */

type Config struct {
	Server Server `yaml:"server" json:"server"`
	Log    Logger `yaml:"log" json:"log"`
}
