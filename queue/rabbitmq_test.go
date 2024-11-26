package queue

import (
	"github.com/newdee/aipaper-util/config/business/common"
	"testing"

	"github.com/newdee/aipaper-util/config"
)

func TestSendMsg(t *testing.T) {
	err := config.Register(config.Common, common.MapEnvToConfig, config.DevEnv)
	if err != nil {
		t.Errorf("register failed, err:%v", err)
		return
	}

	cfg, err := common.GetMsgQueueConfig()
	if err != nil || cfg == nil {
		t.Errorf("get redis config failed, err:%v", err)
		return
	}
	t.Logf("msg queue config:+%v", cfg)
	err = InitMQ(*cfg)
	if err != nil {
		t.Errorf("init mq failed, err:%v", err)
		return
	}
	err = SendMsg("aipaper-queue", "test_msg")
	if err != nil {
		t.Errorf("send msg failed, err:%v", err)
	}
}
