package common

import "time"

type TimeWindow struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (tw *TimeWindow) StartEpochSeconds() int64 {
	return tw.Start.Unix()
}

func (tw *TimeWindow) EndEpochSeconds() int64 {	
	return tw.End.Unix()
}