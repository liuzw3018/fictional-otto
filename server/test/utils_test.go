package test

import (
	"github.com/liuzw3018/otto/server/pkg/utils"
	"github.com/liuzw3018/saber/lib"
	"log"
	"testing"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 18:24
 * @File: utils_test.go
 * @Description: //TODO $
 * @Version:
 */

func TestUtils(t *testing.T) {
	ea, err := utils.ParseDuration(lib.GetConf("jwt.expires_at").(string))
	if err != nil {

	}
	log.Println(ea)
	log.Println()
}
