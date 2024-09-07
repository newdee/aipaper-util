package redis

import (
	"context"
	"testing"
	"time"

	"github.com/newdee/util/database"
	"github.com/stretchr/testify/require"
)

func TestInitRedis(t *testing.T) {
	conf := database.DBConfig{
		Host:     "124.222.112.40",
		Port:     3379,
		Password: "aipp@24",
	}

	err := InitRedis(&conf)
	require.NoErrorf(t, err, "init redis failed, err:%v", err)

	cmd := GetGlobalClient().Set(context.TODO(), "key1", "val1", time.Minute*5)
	t.Logf("%v", cmd)

	get := GetGlobalClient().Get(context.TODO(), "key1")
	t.Logf("%v", get.String())

}
