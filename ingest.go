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
		fmt.Printf("%s\n", self.Marshal())
		fmt.Println(reader.Buffered())
	}
}
