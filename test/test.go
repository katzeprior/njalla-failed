package main

import (
	"context"
	"log"
	"time"

	"github.com/libdns/libdns"
	"github.com/libdns/njalla"
)

func main() {
	p := njalla.Provider{APIToken: "6be76c59196b6a6383df2ade4f59c8275d50c8f0"}
	// log.Println(p.GetRecords(context.Background(), "0dayz.rocks."))
	log.Println(p.SetRecords(context.Background(), "0dayz.rocks.", []libdns.Record{libdns.Record{
		ID:    "763232",
		Type:  "A",
		Name:  "@",
		Value: "127.0.0.2",
		TTL:   time.Duration(time.Duration(time.Hour.Seconds() * 3).Seconds()),
	}}))
}
