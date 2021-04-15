package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"ovh-availability/models"
	"ovh-availability/service"
	"time"
)

func Parsing(config models.Config, token string) {
	var msg string
	for _, server := range config.Servers {
		apiUrl := fmt.Sprintf("https://api.ovh.com/1.0/dedicated/server/availabilities?country=FR&hardware=%s", server)
		client := http.Client{
			Timeout: time.Second * 5, // Timeout after 2 seconds
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
			log.Fatal(jsonErr)
		}

		for _, Availability := range Availabilities {
			if Availability.Region == models.Region(config.Region) {
				for _, Datacenter := range Availability.Datacenters {
					if Datacenter.Availability != "unavailable" {
						url := fmt.Sprintf("https://www.kimsufi.com/fr/commande/kimsufi.xml?reference=%s", Availability.Hardware)
						msg = fmt.Sprintf("%s\n%s\n%s\n%s", Availability.Hardware, Datacenter.Availability, Datacenter.Datacenter, url)

					}
				}
			}

		}

	}
	service.DoNotify(msg, token)

}
