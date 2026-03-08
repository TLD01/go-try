package types

import (
	"fmt"
	"time"
)
type ISO8601Time time.Time


func (t ISO8601Time) MarshalJSON() ([]byte, error) {
	// Format: yyyy-MM-ddTHH:mm:ss.fffZ
	// We use .000 for exactly 3 decimal places
	formatted := time.Time(t).UTC().Format("2006-01-02T15:04:05.000Z")
	return fmt.Appendf(nil, "%q", formatted), nil
}

func (t *ISO8601Time) UnmarshalJSON(data []byte) error {
	// Remove quotes from the JSON string
	s := string(data)
	if s == "null" {
		return nil
	}
	tt, err := time.Parse("\"2006-01-02T15:04:05.000Z\"", s)	
	if err != nil {
		return err
	}
	*t = ISO8601Time(tt)
	return nil
}
