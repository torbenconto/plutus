package quote

import (
	"fmt"
	"time"
)

// Stream Return a constant stream of updated data from the quote using goroutines
// - https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go (example on how to use)
func (q *Quote) Stream(delay time.Duration) <-chan *Quote {
	stream := make(chan *Quote)

	delayT := delay

	go func() {
		defer close(stream)

		for {
			data, err := q.Populate()
			if err != nil {
				fmt.Println("Error fetching data for quote: ", err)
			}

			stream <- data

			time.Sleep(delayT)
		}
	}()

	return stream
}
