package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"ovh-availability/database"
	"ovh-availability/models"
	"time"
)

func Parsing(config models.Config, token string) {
	db := database.DB
	var serverinfo models.ServerInfo
	// var msg string
	db.First(&serverinfo)
	log.Println(serverinfo)
	var localServerInfo models.ServerInfo
	for _, server := range config.Servers {
		// apiUrl := fmt.Sprintf("https://api.ovh.com/1.0/dedicated/server/availabilities?country=FR&hardware=%s", server)
		apiUrl := fmt.Sprintf("http://localhost:8000/mod_response_%s.json", server)
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
		var Availabilities models.Availabilities
		jsonErr := json.Unmarshal(body, &Availabilities)
		if jsonErr != nil {
			log.Println(jsonErr)
		}
		log.Println(Availabilities)
		for _, Availability := range Availabilities {
			if Availability.Region == models.Region(config.Region) {
				for _, Datacenter := range Availability.Datacenters {
					if Datacenter.Availability != "unavailable" {

						url := fmt.Sprintf("https://www.kimsufi.com/fr/commande/kimsufi.xml?reference=%s", Availability.Hardware)
						localServerInfo.Availability = string(Datacenter.Availability)
						localServerInfo.ID = Availability.Hardware
						localServerInfo.Region = string(Datacenter.Datacenter)
						localServerInfo.Url = url

						log.Println(serverinfo.Availability, Availability.Hardware)
						serverinfo = localServerInfo

					} else {
						localServerInfo.Availability = string(Datacenter.Availability)
						localServerInfo.ID = Availability.Hardware
						localServerInfo.Region = string(Datacenter.Datacenter)
						localServerInfo.Url = string(Datacenter.Availability)
						serverinfo = localServerInfo
						// log.Println("no server available sleeping for 15s")
						// time.Sleep(15 * time.Second)
					}
					// log.Println(db.First(&serverinfo, server))
					db.Save(&serverinfo)

					// if localServerInfo.Availability != serverinfo.Availability {
					// 	msg := fmt.Sprintf("\n%s %s %s\n%s", serverinfo.Availability, serverinfo.ID, serverinfo.Region, serverinfo.Url)
					// 	log.Println(msg)

					// 	service.DoNotify(msg, token)
					// }

				}
			}

		}

	}
	// db.Save(&serverinfo)

	// log.Println("wait for 5s")
	// time.Sleep(5 * time.Second)

	// service.DoNotify(msg, token)
}
