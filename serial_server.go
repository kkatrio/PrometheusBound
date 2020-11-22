package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/tarm/serial"
	"strings"
)

func send_serial(event string) {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
		s, err = serial.OpenPort(c)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("port opened")

	n, err := s.Write([]byte(event))
	if err != nil {
		log.Fatal(err)
	}
	_ = n // avoid declared but not used error
	s.Close()
}

type Message struct {
	Status string
}

func main() {

	http.HandleFunc("/alarm", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("body:", string(body))

		var m Message
		dec := json.NewDecoder(strings.NewReader(string(body))) // should be able to use r.body directly
		err = dec.Decode(&m)
		if err != nil {
			log.Fatalf("could not decode", err)
		}
		log.Printf("alert status: %s\nsending serial\n", m.Status)

		if m.Status == "firing" {
			send_serial("1")
			fmt.Println("fired")
		} else if m.Status == "resolved" {
			send_serial("0")
			fmt.Println("resolved")
		} else {
			log.Println("status ?")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
