package sms

import (
	"fmt"
	"github.com/newdee/aipaper-util/log"
	"math/rand"
	"time"

	unisms "github.com/apistd/uni-go-sdk/sms"
	"github.com/newdee/aipaper-util/config/business/common"
)

var globalSmsCli *unisms.UniSMSClient

// Param 短信参数
type Param struct {
	Signature  string `json:"signature"`
	TemplateID string `json:"template_id"`
	TTL        string `json:"ttl"`
}

func init() {
	globalSmsCli = unisms.NewClient(
		"P7te1AsoN6fZcp6k8hNTkyd3RCAbzZLUVrxvCQ6dNFfRiHbmL")
}

func generateSMSCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

// Send 发送短信
func Send(phoneNum ...string) (string, error) {
	if globalSmsCli == nil {
		return "", fmt.Errorf("global sms client is nil")
	}
	cfg, err := common.GetSMSConfig()
	if err != nil {
		log.Errorf("get sms config failed, err:%v", err)
		return "", err
	}
	smsCode := generateSMSCode()
	message := unisms.BuildMessage()
	message.SetTo(phoneNum...)
	message.SetSignature(cfg.Signature)
	message.SetTemplateId(cfg.TemplateID)
	message.SetTemplateData(map[string]string{
		"code": smsCode,
		"ttl":  cfg.TTL,
	})

	response, err := globalSmsCli.Send(message)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Printf("%v\n", response)
	return smsCode, nil
}
