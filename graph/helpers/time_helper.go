package helpers

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func ESI_time_to_millis(esi_time string) int64 {
	log.WithFields(log.Fields{"time": esi_time}).Debugf("Parsing Time : %s", esi_time)
	t, err := time.Parse("2006-01-02T15:04:05-0700", esi_time)
	if err != nil {
		log.WithFields(log.Fields{"time": esi_time}).Errorf("Failed to Parse Time with Error : %v", err)
		return 43200000
	}

	log.Debugf("Responding with time: %s ", strconv.FormatInt(t.UTC().UnixMilli()-time.Now().UnixMilli(), 10))
	return t.UTC().UnixMilli() - time.Now().UnixMilli()
}
