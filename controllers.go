package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/flock3/ip2country/config"
	maxminddb "github.com/oschwald/maxminddb-golang"
)

// Controller struct
type Controller struct {
	Config config.Config
}

// FetchCountry fetches the country by a given IP Address
func (c Controller) FetchCountry(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {
	db, err := maxminddb.Open("test-data/test-data/GeoIP2-City-Test.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP("81.2.69.142")

	var record struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	} // Or any appropriate struct

	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(record.Country.ISOCode)
	// Output:
	// GB

	return
}

// UpdateDB will go and fetch the latest maxmind country database and store it
func (c Controller) UpdateDB(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {

	gzf, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gzf.Read()

	return
}
