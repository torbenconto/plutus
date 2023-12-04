package plutus

import "time"

func Stream(s *Stock, delay int) <-chan *Stock {
	stream := make(chan *Stock)

	delayT := time.Duration(delay) * time.Millisecond

	go func() {
		defer close(stream)

		for {
			s.Populate()
			data := s

			stream <- data

			time.Sleep(delayT)
		}
	}()

	return stream
}
