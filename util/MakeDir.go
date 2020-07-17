//Author  :     knight
//Date    :     2020/07/16 21:38:11
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     创建临时的文件夹

package util

import (
	"log"
	"os"
)

var tmpDir string

func init() {
	dir, _ := os.Getwd()
	tmpDir = dir + "\\tmp"
}

//MakeTmpDir 创建临时的文件夹
func MakeTmpDir() error {
	if err := os.Mkdir(tmpDir, 644); err != nil {
		return err
	}
	log.Println("Create Tmp Dir success")
	return nil
}

//RemoveTmpDir 删除临时的文件夹
func RemoveTmpDir() error {
	if err := os.Remove(tmpDir); err != nil {
		return err
	}
	log.Println("Remove Tmp Dir sucesss")
	return nil
}