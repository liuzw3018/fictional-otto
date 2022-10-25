package config

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:55
 * @File: server.go
 * @Description: //TODO $
 * @Version:
 */

type Server struct {
	Listen string `yaml:"listen" json:"listen"`
	Mode   string `yaml:"mode" json:"mode"`
}
