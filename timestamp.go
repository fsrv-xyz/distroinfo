package distroinfo

import (
	"strings"
	"time"
)

type TimeStamp struct {
	time.Time
}

func (t *TimeStamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}
	var err error
	t.Time, err = time.Parse("2006-01-02", s)
	return err
}
