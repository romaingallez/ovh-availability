package api

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"ovh-availability/models"
// 	"ovh-availability/utils"
// )

// func GetType(config models.Config) []string {
// 	// Open our jsonFile
// 	fileName := "api/server_types.json"
// 	file, err := os.Open(fileName) // For read access.
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Successfully Opened ", fileName)

// 	bb, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var servers models.ServerTypes
// 	json.Unmarshal(bb, &servers)

// 	// log.Println(serverName)
// 	var refs []string
// 	for _, offer := range servers.Offers {
// 		_, found := utils.Find(config.Servers, offer.Name)
// 		if found {
// 			// log.Println(found, offer.Name, offer.Ref)
// 			refs = append(refs, offer.Ref)

// 		}
// 	}
// 	return refs
// }
