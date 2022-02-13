package helpers

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

func TestEsiTTLToMillisSuccess(t *testing.T) {
	now := time.Now().UTC()
	now = now.Add(time.Millisecond * 50000)
	nowString := now.Format(time.RFC1123)
	result := EsiTTLToMillis(context.Background(), nowString)
	if result > 0 {
		return
	}
	t.Errorf("TTL return was not greater than 0: %s", strconv.FormatInt(result, 10))
}

func TestEsiTTLToMillisFail(t *testing.T) {
	now := time.Now().UTC()
	now = now.Add(time.Millisecond * 50000)
	nowString := now.Format(time.RFC1123)
	nowString = nowString + "GARBAGE"
	result := EsiTTLToMillis(context.Background(), nowString)
	if result == 43200000 {
		return
	}
	t.Errorf("TTL return was not greater than 0: %s", strconv.FormatInt(result, 10))
}

func init() {
	configuration.TestConfigPath = "c:\\Users\\Bryan\\Documents\\workspace\\eve-graphql-go\\graph\\helpers\\config.yml"
}
