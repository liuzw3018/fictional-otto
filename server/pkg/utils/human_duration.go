package utils

import (
	"strconv"
	"strings"
	"time"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 18:24
 * @File: human_duration.go
 * @Description: //TODO $
 * @Version:
 */

func ParseDuration(d string) (time.Duration, error) {
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.HasSuffix(d, "d") {
		h := strings.TrimSuffix(d, "d")
		hour, _ := strconv.Atoi(h)
		dr = time.Hour * 24 * time.Duration(hour)
		return dr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
