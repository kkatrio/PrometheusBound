package main

import (
        "fmt"
		"log"
		"net/http"
		"html"
        "github.com/tarm/serial"
)

func send_serial(event string) {
	c := &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("port opened")

	n, err := s.Write([]byte(event))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("n: ", n)
	s.Close()
}

func main() {

	http.HandleFunc("/alarm", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sending: %q", html.EscapeString(r.URL.Path))
		send_serial("on")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
