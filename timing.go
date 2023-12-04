package plutus

import "time"

func Stream(s *Stock, delay int) <-chan string {
	stream := make(chan string)

	delayT := time.Duration(delay) * time.Millisecond

	go func() {
		defer close(stream)

		for {
			data := "Your constant data"

			stream <- data

			time.Sleep(delayT)
		}
	}()

	return stream
}
