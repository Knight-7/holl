//Author  :     knight
//Date    :     2020/07/13 17:37:27
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     生成全局唯一sessionID

package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

//GenerateSessionID 生成全局唯一sessionID
func GenerateSessionID() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	randNum := rand.Int63()
	return mymd5(strconv.FormatInt(nano, 10)) + mymd5(strconv.FormatInt(randNum, 10))
}

func mymd5(t string) string {
	hasMD := md5.New()
	io.WriteString(hasMD, t)
	return fmt.Sprintf("%x", hasMD.Sum(nil))
}