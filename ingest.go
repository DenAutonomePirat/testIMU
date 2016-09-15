package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tarm/serial"
	"log"
)

func ingest() {
	self := NewBoat()

	config := &serial.Config{
		Name: "/dev/ttyACM1",
		Baud: 115200,
	}
	arduino, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	defer arduino.Close()

	for {
		reader := bufio.NewReader(arduino)
		token, err := reader.ReadBytes('\x0a')
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(token, self)
		if err != nil {
			fmt.Println(".")
		}

	}
}

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
