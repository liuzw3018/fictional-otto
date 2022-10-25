package test

import (
	"github.com/liuzw3018/otto/server/pkg/utils"
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
	timeStr := "1.5h"
	duration, err := utils.ParseDuration(timeStr)
	if err != nil {
		panic(err)
	}
	log.Println(duration)
}
