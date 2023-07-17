package main

import (
	"fmt"
	"log"
	"time"
)

type Event struct {
	ID   string    `json:"id"`
	Time time.Time `json:"time"`
}

type DoorEvent struct {
	Event
	Action string `json:"action"`
}

type TemperatureEvent struct {
	Event
	Value float64 `json:"value"`
}

func NewDoorEvent(id string, time time.Time, action string) (*DoorEvent, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}
	evnt := DoorEvent{
		Event:  Event{id, time},
		Action: action,
	}

	return &evnt, nil
}

func main() {
	evnt, err := NewDoorEvent("front door", time.Now(), "open")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", evnt)
	//go run events.go
	//&{Event:{ID:front door Time:2023-07-16 21:58:24.051883 -0700 PDT m=+0.000268668} Action:open}
}
