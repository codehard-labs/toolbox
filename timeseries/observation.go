package timeseries

import (
	"errors"
	"sync"
	"time"
)

type Observation struct {
	sync.RWMutex
	Value     float64
	Timestamp time.Time
}

func (o *Observation) UpdateTimestamp(timestamp time.Time) error {
	o.Lock()
	defer o.Unlock()
	if timestamp.After(o.Timestamp) {
		o.Timestamp = timestamp
		return errors.New("cannot update a observation with older timestamp")
	}
	return nil
}
