package helpers

import (
	"strconv"
	"testing"
	"time"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

func TestEsiTtlToMillisSuccess(t *testing.T) {
	now := time.Now().UTC()
	now = now.Add(time.Millisecond * 50000)
	nowString := now.Format(time.RFC1123)
	result := EsiTtlToMillis(nowString)
	if result > 0 {
		return
	} else {
		t.Errorf("TTL return was not greater than 0: %s", strconv.FormatInt(result, 10))
	}
}

func TestEsiTtlToMillisFail(t *testing.T) {
	now := time.Now().UTC()
	now = now.Add(time.Millisecond * 50000)
	nowString := now.Format(time.RFC1123)
	nowString = nowString + "GARBAGE"
	result := EsiTtlToMillis(nowString)
	if result == 43200000 {
		return
	} else {
		t.Errorf("TTL return was not greater than 0: %s", strconv.FormatInt(result, 10))
	}
}

func init() {
	configuration.TestConfigPath = "c:\\Users\\Bryan\\Documents\\workspace\\eve-graphql-go\\graph\\helpers\\config.yml"
}
