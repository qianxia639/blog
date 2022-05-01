package command

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Timestamp time.Time

func (ts *Timestamp) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("can not convert %v to timestamp", t)
	}

	*ts = Timestamp(t)
	return nil
}

func (ts Timestamp) Value() (driver.Value, error) {
	var t = time.Time(ts)
	if t.UnixNano() == time.Now().UnixNano() {
		return nil, nil
	}
	return t, nil
}
