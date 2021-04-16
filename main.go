package main

import (
	"log"
	"ovh-availability/api"
	"ovh-availability/database"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	var runParsing bool = true
	if runParsing {
		// token := os.Getenv("DISCORD_TOKEN")

		config := GetConfig()

		// Suppose we have a function call `f(s)`. Here's how
		// we'd call that in the usual way, running it
		// synchronously.

		// To invoke this function in a goroutine, use
		// `go f(s)`. This new goroutine will execute
		// concurrently with the calling one.

		// var info []models.ServerInfo
		// tick := time.Tick(5 * time.Second)
		// for range tick {
		// 	fmt.Println("Tick")
		// 	info = api.Parsing(config)
		// }
		// log.Println(info)

		ticker := time.NewTicker(time.Second * 1).C
		go func() {
			for {
				var infos []string
				select {
				case <-ticker:
					localinfo := api.Parsing(config)
					var info []string
					for _, v := range localinfo {
						if !testEq(info, infos) {
							info = append(info, v.ID, v.Availability, v.Region)
							log.Println(infos)
						}
						infos = info

					}
					// if testEq(info, localinfo) {
					// 	info = localinfo
					// 	log.Println(info)
					// }

				}
			}

		}()

		time.Sleep(time.Second * 10)
	}

}
