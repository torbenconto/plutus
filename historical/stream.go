package historical

import (
	"fmt"
	"time"
)

// Stream Return a constant stream of updated data from the historical data using goroutines
// - https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go (example on how to use)
func (h *Historical) Stream(delay time.Duration) <-chan *Historical {
	stream := make(chan *Historical)

	delayT := delay

	go func() {
		defer close(stream)

		for {
			data, err := h.Populate()
			if err != nil {
				fmt.Println("Error fetching data for historical: ", err)
			}

			stream <- data

			time.Sleep(delayT)
		}
	}()

	return stream
}
