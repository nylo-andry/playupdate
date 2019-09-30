package worker

import (
	"log"
	"sync"

	"github.com/nylo-andry/playupdate"
)

func startWorker(inputs <-chan string, wg *sync.WaitGroup, updateService playupdate.UpdateService) {
	defer wg.Done()

	for mac := range inputs {
		err := updateService.Update(mac)
		if err != nil {
			log.Printf("Could not update mac [%s]: %s", mac, err)
		}
	}
}
