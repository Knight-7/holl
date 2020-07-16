//Author  :     knight
//Date    :     2020/07/13 23:20:46
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     获得当前时间

package util

import (
	"os"
	"path/filepath"
	"time"
)

const (
	formatTime = "2006-01-02 15:04:05"
)

//GetLocalTime 获取当前本地时间
func GetLocalTime() *time.Time {
	dir, _ := os.Getwd()
	zipPath := filepath.Dir(dir) + "\\config\\data.zip"
	os.Setenv("ZONEINFO", zipPath)
	local, _ := time.LoadLocation("Local")
	localTime, _ := time.ParseInLocation(formatTime, time.Now().Format(formatTime), local)
	return &localTime
}