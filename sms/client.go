package sms

import (
	"fmt"
	"math/rand"
	"time"

	unisms "github.com/apistd/uni-go-sdk/sms"
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
	smsCode := generateSMSCode()

	message := unisms.BuildMessage()
	message.SetTo(phoneNum...)
	message.SetSignature("浅思科技")
	message.SetTemplateId("pub_verif_ttl3")
	message.SetTemplateData(map[string]string{
		"code": smsCode,
		"ttl":  "1",
	})

	response, err := globalSmsCli.Send(message)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Printf("%v\n", response)
	return smsCode, nil
}
