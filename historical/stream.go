package historical

import "time"

// Stream Return a constant stream of updated data from the historical data using goroutines
// - https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go (example on how to use)
func (h *Historical) Stream(delay int) <-chan *Historical {
	stream := make(chan *Historical)

	delayT := time.Duration(delay) * time.Millisecond

	go func() {
		defer close(stream)

		for {
			data, _ := h.Populate()

			stream <- data

			time.Sleep(delayT)
		}
	}()

	return stream
}
