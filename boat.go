package main

import (
	"encoding/json"
)

type Boat struct {
	Id          int32   `json:"id"`
	Rudder      int32   `json:"rudder"`
	Depth       float32 `json:"depth"`
	MainSail    int32   `json:"mainsail"`
	Jib         int32   `json:"jib"`
	Volts       float32 `json:"volts`
	Amperes     float32 `json:"amperes"`
	JoulesTotal float32 `json:"joules_total"`
	joulesTrip  float32 `json:"joules_trip"`
	Heading     float32 `json:"heading"`
	Pitch       float32 `json:"pitch,omitempty"`
	Roll        float32 `json:"roll"`
}

func (b *Boat) Marshal() *[]byte {
	encoded, _ := json.Marshal(b)
	return &encoded
}

func NewBoat() *Boat {
	b := Boat{}
	return &b
}

type ack struct {
	Id int32 `json:"id"`
}

func (a *ack) Marshal() *[]byte {
	encoded, _ := json.Marshal(a)
	return &encoded
}
func newAck() *ack {
	a := ack{}
	return &a
}
