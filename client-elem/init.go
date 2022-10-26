package client_elem

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func InitClients() {
	ClientList = make([]Client, NrClients, NrClients)

	file, err := os.Open("./jsonConfig/clients.json")
	if err != nil {
		log.Fatal("Error opening clients.json", err)
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &ClientList)
	if err != nil {
		log.Fatal("Error unmarshaling clients.json", err)
	}

}

func MakeNewClients() {

	for i := range NotifClientManag {
		ClientList[i-1] = Client{
			Id: i,
		}
		go ClientList[i-1].OrderOnline()
	}
}
