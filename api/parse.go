package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"ovh-availability/models"
	"time"
)

func Parsing(config models.Config) (ServersInfo []models.ServerInfo) {
	// var msg string
	// log.Println(serverinfo)
	var Availabilities models.Availabilities
	var serverinfo models.ServerInfo
	for _, server := range config.Servers {
		log.Println(server)
		apiUrl := fmt.Sprintf("https://api.ovh.com/1.0/dedicated/server/availabilities?country=FR&hardware=%s", server)
		// apiUrl := fmt.Sprintf("http://localhost:8000/mod_response_%s.json", server)
		client := http.Client{
			Timeout: time.Second * 10, // Timeout after 2 seconds
		}
		req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
		res, getErr := client.Do(req)
		if getErr != nil {
			log.Fatal(getErr)
		}
		if res.Body != nil {
			defer res.Body.Close()
		}
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)

		}
		jsonErr := json.Unmarshal(body, &Availabilities)
		if jsonErr != nil {
			log.Println(jsonErr)
		}
		// log.Println(Availabilities)
		for _, Availability := range Availabilities {
			if Availability.Region == models.Region(config.Region) {
				for _, Datacenter := range Availability.Datacenters {
					if Datacenter.Availability != "unavailable" {
						if Datacenter.Datacenter == "default" {
							url := fmt.Sprintf("https://www.kimsufi.com/fr/commande/kimsufi.xml?reference=%s", Availability.Hardware)
							serverinfo.Availability = string(Datacenter.Availability)
							serverinfo.ID = Availability.Hardware
							serverinfo.Region = string(Datacenter.Datacenter)
							serverinfo.Url = url
							ServersInfo = append(ServersInfo, serverinfo)
						}

					}

				}
			}

		}

	}
	return ServersInfo
}
