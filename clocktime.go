package timeext

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var errInvalidType = errors.New("invalid type")
var ErrInvalidTimeFormat = errors.New("Invalid time format.")

type ClockTime struct {
	Hour, Min, Sec int
}

func Clock(t time.Time) ClockTime {
	return ClockTime{
		t.Hour(),
		t.Minute(),
		t.Second(),
	}
}

func (ct *ClockTime) Scan(value interface{}) error {
	if value == nil {
		ct.Hour = 0
		ct.Min = 0
		ct.Sec = 0
		return nil
	}
	var s string
	switch v := value.(type) {
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return errInvalidType
	}
	parts := strings.Split(s, ":")
	ct.Hour, _ = strconv.Atoi(parts[0])
	ct.Min, _ = strconv.Atoi(parts[1])
	ct.Sec, _ = strconv.Atoi(parts[2])
	return nil
}

func (ct ClockTime) Value() (driver.Value, error) {
	return ct.String(), nil
}

func (ct ClockTime) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", ct.Hour, ct.Min, ct.Sec)
}

func (ct *ClockTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		parts := strings.Split(s, ":")
		ct.Hour, ct.Min, ct.Sec = 0, 0, 0
		ct.Hour, err = strconv.Atoi(parts[0])
		if err != nil {
			return ErrInvalidTimeFormat
		}
		if len(parts) > 1 {
			ct.Min, err = strconv.Atoi(parts[1])
			if err != nil {
				return ErrInvalidTimeFormat
			}
		}
		if len(parts) > 2 {
			ct.Sec, err = strconv.Atoi(parts[2])
			if err != nil {
				return ErrInvalidTimeFormat
			}
		}
		return nil
	} else {
		return err
	}
}

func (ct ClockTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.String())
}

func (ct ClockTime) SecondOfDay() int {
	return ct.Sec + (ct.Min * 60) + (ct.Hour * 3600)
}
