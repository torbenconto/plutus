package tests

import "github.com/torbenconto/plutus/historical"

var historicalServerData = []byte(`
	{
  "chart": {
    "result": [
      {
        "indicators": {
          "quote": [
            {
              "close": [
                12.0,
                14.6,
                13.2
              ],
              "open": [
                16.0,
                7.4,
                9.62
              ],
              "volume": [
                212010,
                43021,
                350511
              ],
              "high": [
                13.6,
                15.2,
                7.2
              ],
              "low": [
                10.0,
                12.2,
                2.31
              ]
            }
          ]
        },
        "timestamp": [
          1615891200,
          1615977600,
          1616064000
        ]
      }
    ]
  }
}`)

var historicalTestCases = []historical.PricePoint{
	{1615891200, 16.0, 12.0, 13.6, 10.0, 212010},
	{1615977600, 7.4, 14.6, 15.2, 12.2, 43021},
	{1616064000, 9.62, 13.2, 7.2, 2.31, 350511},
}
