package ticker

import (
	"RestAPI/my_service/database"
	"RestAPI/my_service/nats"
	"RestAPI/pkg/logger"
	"github.com/rs/zerolog/log"

	"fmt"
	"time"
)

var l = logger.Get()

func Ticker() {
	ticker := time.NewTicker(10000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				database.Connect()

				err, _ := nats.Open()
				if err != nil {

				}

				log.Info().Msgf("Tick, all right.")

			}
		}
	}()

	time.Sleep(300000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
