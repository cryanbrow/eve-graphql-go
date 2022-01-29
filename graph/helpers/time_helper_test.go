package helpers

import (
	"strconv"
	"testing"
	"time"

	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
)

func TestESI_ttl_to_millis_success(t *testing.T) {
	now := time.Now().UTC()
	now = now.Add(time.Millisecond * 50000)
	now_string := now.Format(time.RFC1123)
	result := ESI_ttl_to_millis(now_string)
	if result > 0 {
		return
	} else {
		t.Errorf("TTL return was not greater than 0: %s", strconv.FormatInt(result, 10))
	}
}

func TestESI_ttl_to_millis_fail(t *testing.T) {
	now := time.Now().UTC()
	now = now.Add(time.Millisecond * 50000)
	now_string := now.Format(time.RFC1123)
	now_string = now_string + "GARBAGE"
	result := ESI_ttl_to_millis(now_string)
	if result == 43200000 {
		return
	} else {
		t.Errorf("TTL return was not greater than 0: %s", strconv.FormatInt(result, 10))
	}
}

func init() {
	configuration.TestConfigPath = "c:\\Users\\Bryan\\Documents\\workspace\\eve-graphql-go\\graph\\helpers\\config.yml"
}
