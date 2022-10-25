package initialize

import (
	"github.com/liuzw3018/otto/server/pkg/global"
	"github.com/spf13/viper"
	"log"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 18:02
 * @File: sys_config.go
 * @Description: //TODO $
 * @Version:
 */

func init() {
	global.OttoViper = viper.New()
	global.OttoViper.AddConfigPath("./config/")
	global.OttoViper.AddConfigPath("/Users/liuzhineng/GolandProjects/otto/etc")
	global.OttoViper.SetConfigName("server")
	global.OttoViper.SetConfigType("yaml")
	if err := global.OttoViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("配置文件没有找到，请检查")
			return
		} else {
			panic(err)
		}
	}
	if err := global.OttoViper.Unmarshal(&global.OttoConfig); err != nil {
		log.Println("配置反序列化失败")
		panic(err)
	}

	// 在配置初始化完毕后开始初始化其他组件
	initLogrus()
}
