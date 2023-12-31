package plutus

import "time"

// Return a constant stream of updated data from the stock using goroutines
// - https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go (example on how to use)
func (s *Stock) Stream(delay int) <-chan *Stock {
	stream := make(chan *Stock)

	delayT := time.Duration(delay) * time.Millisecond

	go func() {
		defer close(stream)

		for {
			// This is very bad i think
			data, _ := NewStock(s.Ticker, s.Provider)

			stream <- data

			time.Sleep(delayT)
		}
	}()

	return stream
}
