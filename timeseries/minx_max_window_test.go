package timeseries

import (
	"fmt"
	"testing"
	"time"
)

func TestMinMaxWindow(t *testing.T) {
	w := NewMinMaxWindow(80 * time.Millisecond)
	values := []float64{1.3, 0.8, 1.9, -13, -9.2, 11, 1.4, 8.9, 7.2, 0.8}
	ticker := time.NewTicker(15 * time.Millisecond)
	i := 0
	for _ = range ticker.C {
		if i < len(values) {
			w.AddNewObservation(&Observation{Value: values[i], Timestamp: time.Now()})
			i++
		}
		if i >= len(values) {
			break
		}
	}
	fmt.Println("min:", w.GetMinObservation().Value, time.Now())
	fmt.Println("max:", w.GetMaxObservation().Value, time.Now())
	w.Print()
}
