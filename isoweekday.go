package timeext

import (
	"encoding/json"
	"errors"
	"strings"
)

type ISOWeekday int

const (
	All ISOWeekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var ErrInvalidWeekday = errors.New("Invalid weekday.")

func (wd *ISOWeekday) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		switch strings.ToLower(s) {
		case "all", "":
			*wd = All
		case "monday":
			*wd = Monday
		case "tuesday":
			*wd = Tuesday
		case "wednesday":
			*wd = Wednesday
		case "thursday":
			*wd = Thursday
		case "friday":
			*wd = Friday
		case "saturday":
			*wd = Saturday
		case "sunday":
			*wd = Sunday
		default:
			return ErrInvalidWeekday
		}
		return nil
	} else {
		return err
	}
}

func (wd ISOWeekday) MarshalJSON() ([]byte, error) {
	switch wd {
	case All:
		return json.Marshal("All")
	case Monday:
		return json.Marshal("Monday")
	case Tuesday:
		return json.Marshal("Tuesday")
	case Wednesday:
		return json.Marshal("Wednesday")
	case Thursday:
		return json.Marshal("Thursday")
	case Friday:
		return json.Marshal("Friday")
	case Saturday:
		return json.Marshal("Saturday")
	case Sunday:
		return json.Marshal("Sunday")
	default:
		return nil, ErrInvalidWeekday
	}
}
