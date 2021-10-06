package external

import (
	"encoding/json"
	"homework-server/contract"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetDataJson() map[string]contract.User {
	resp, err := http.Get("https://jsonkeeper.com/b/DMXK")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var mapUser map[string]contract.User
	dec := json.NewDecoder(strings.NewReader(string(body)))

	for {
		if err := dec.Decode(&mapUser); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	return mapUser
}
