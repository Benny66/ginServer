package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type ModelTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
	zone        = "Asia/Shanghai"
)

func (t *ModelTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = ModelTime(now)
	return
}

func (t ModelTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t ModelTime) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t ModelTime) local() time.Time {
	loc, _ := time.LoadLocation(zone)
	return time.Time(t).In(loc)
}

func (t ModelTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (t *ModelTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = ModelTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
