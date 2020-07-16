//Author  :     knight
//Date    :     2020/07/13 17:15:33
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     微信解密工具

package util

import (
	"github.com/xlstudio/wxbizdatacrypt"
)

//Decrypt 解密
func Decrypt(appID, sessionKey, encryptedData, iv string) (string, error) {

	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: appID, SessionKey: sessionKey}
	result, err := pc.Decrypt(encryptedData, iv, true)
	if err != nil {
		return "",  err
	}
	return result.(string), nil
}