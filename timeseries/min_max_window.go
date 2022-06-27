package timeseries

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type MinMaxWindow struct {
	sync.RWMutex
	window   time.Duration
	minStack []*Observation
	maxStack []*Observation
}

func NewMinMaxWindow(window time.Duration) *MinMaxWindow {
	return &MinMaxWindow{
		window:   window,
		minStack: make([]*Observation, 0),
		maxStack: make([]*Observation, 0),
	}
}

func (w *MinMaxWindow) AddNewObservation(o *Observation) error {
	w.Lock()
	defer w.Unlock()
	if time.Since(o.Timestamp) > w.window {
		return errors.New("timestamp out of window")
	}
	w.pushIntoMaxStack(o)
	w.pushIntoMinStack(o)
	return nil
}

func (w *MinMaxWindow) pushIntoMaxStack(o *Observation) {
	l := len(w.maxStack)
	for i := l - 1; i >= 0; i-- {
		if w.maxStack[i].Value <= o.Value {
			continue
		} else {
			w.maxStack = append(w.maxStack[:i+1], o)
			return
		}
	}
	w.maxStack = []*Observation{o}
}

func (w *MinMaxWindow) pushIntoMinStack(o *Observation) {
	l := len(w.minStack)
	for i := l - 1; i >= 0; i-- {
		if w.minStack[i].Value >= o.Value {
			continue
		} else {
			w.minStack = append(w.minStack[:i+1], o)
			return
		}
	}
	w.minStack = []*Observation{o}
}

func (w *MinMaxWindow) GetMaxObservation() *Observation {
	w.Lock()
	defer w.Unlock()
	l := len(w.maxStack)
	for i := 0; i < l; i++ {
		if time.Since(w.maxStack[i].Timestamp) > w.window {
			continue
		} else {
			w.maxStack = w.maxStack[i:l]
			return w.maxStack[0]
		}
	}
	return nil
}

func (w *MinMaxWindow) GetMinObservation() *Observation {
	w.Lock()
	defer w.Unlock()
	l := len(w.minStack)
	for i := 0; i < l; i++ {
		if time.Since(w.minStack[i].Timestamp) > w.window {
			continue
		} else {
			w.minStack = w.minStack[i:l]
			return w.minStack[0]
		}
	}
	return nil
}

func (w *MinMaxWindow) Print() {
	w.RLock()
	defer w.RUnlock()
	fmt.Println("Max Stack:")
	for _, o := range w.maxStack {
		fmt.Printf("Value: %f, Timestamp: %s\n", o.Value, o.Timestamp)
	}
	fmt.Println("Min Stack:")
	for _, o := range w.minStack {
		fmt.Printf("Value: %f, Timestamp: %s\n", o.Value, o.Timestamp)
	}
}
