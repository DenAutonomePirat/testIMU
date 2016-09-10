package main

import (
	//"encoding/json"
	"fmt"
	"github.com/tarm/serial"
	//"io"
	"log"
)

func main() {

	c := &serial.Config{
		Name: "/dev/ttyACM0",
		Baud: 115200,
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	buf := make([]byte, 128)

	dec := json.NewDecoder()

	for {
		var self Boat
		if err := dec.Decode(&self); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", self.Heading)
	}

	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", buf[:n])
	}
}

/*
	options := serial.OpenOptions{
		PortName:        "/dev/ttyACM0",
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()
*/

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
