package stock

import (
	"github.com/torbenconto/plutus/config"
	"github.com/torbenconto/plutus/internal/tests"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var dividendServerData = []byte(`{
  "data": {
    "dividendHeaderValues": [
      {
        "label": "Ex-Dividend Date",
        "value": "05/09/2024"
      },
      {
        "label": "Dividend Yield",
        "value": "15.51%"
      },
      {
        "label": "Annual Dividend",
        "value": "$1.88"
      },
      {
        "label": "P/E Ratio",
        "value": "7.23"
      }
    ],
    "exDividendDate": "05/09/2024",
    "dividendPaymentDate": "05/24/2024",
    "yield": "15.51%",
    "annualizedDividend": "1.88",
    "payoutRatio": "7.23",
    "dividends": {
      "asOf": null,
      "headers": {
        "exOrEffDate": "Ex/EFF Date",
        "type": "Type",
        "amount": "Cash Amount",
        "declarationDate": "Declaration Date",
        "recordDate": "Record Date",
        "paymentDate": "Payment Date"
      },
      "rows": [
        {
          "exOrEffDate": "05/09/2024",
          "type": "Cash",
          "amount": "$0.47",
          "declarationDate": "04/30/2024",
          "recordDate": "05/10/2024",
          "paymentDate": "05/24/2024",
          "currency": "USD"
        },
        {
          "exOrEffDate": "02/09/2024",
          "type": "Cash",
          "amount": "$0.47",
          "declarationDate": "02/01/2024",
          "recordDate": "02/12/2024",
          "paymentDate": "02/28/2024",
          "currency": "USD"
        },
        {
          "exOrEffDate": "11/10/2023",
          "type": "Cash",
          "amount": "$0.46",
          "declarationDate": "11/02/2023",
          "recordDate": "11/13/2023",
          "paymentDate": "11/28/2023",
          "currency": "USD"
        },
        {
          "exOrEffDate": "08/14/2023",
          "type": "Cash",
          "amount": "$0.45",
          "declarationDate": "08/01/2023",
          "recordDate": "08/15/2023",
          "paymentDate": "08/31/2023",
          "currency": "USD"
        },
        {
          "exOrEffDate": "05/12/2023",
          "type": "Cash",
          "amount": "$0.45",
          "declarationDate": "05/02/2023",
          "recordDate": "05/15/2023",
          "paymentDate": "05/31/2023",
          "currency": "USD"
        },
        {
          "exOrEffDate": "03/06/2023",
          "type": "Cash",
          "amount": "$0.45",
          "declarationDate": "02/23/2023",
          "recordDate": "03/07/2023",
          "paymentDate": "03/21/2023",
          "currency": "USD"
        },
        {
          "exOrEffDate": "11/07/2022",
          "type": "Cash",
          "amount": "$0.36",
          "declarationDate": "10/27/2022",
          "recordDate": "11/08/2022",
          "paymentDate": "11/22/2022",
          "currency": "USD"
        },
        {
          "exOrEffDate": "08/08/2022",
          "type": "Cash",
          "amount": "$0.33",
          "declarationDate": "07/28/2022",
          "recordDate": "08/09/2022",
          "paymentDate": "08/23/2022",
          "currency": "USD"
        },
        {
          "exOrEffDate": "05/09/2022",
          "type": "Cash",
          "amount": "$0.30",
          "declarationDate": "04/28/2022",
          "recordDate": "05/10/2022",
          "paymentDate": "05/24/2022",
          "currency": "USD"
        },
        {
          "exOrEffDate": "03/07/2022",
          "type": "Cash",
          "amount": "$0.27",
          "declarationDate": "02/24/2022",
          "recordDate": "03/08/2022",
          "paymentDate": "03/22/2022",
          "currency": "USD"
        },
        {
          "exOrEffDate": "11/05/2021",
          "type": "Cash",
          "amount": "$0.25",
          "declarationDate": "10/28/2021",
          "recordDate": "11/08/2021",
          "paymentDate": "11/22/2021",
          "currency": "USD"
        }
      ]
    }
  },
  "message": null,
  "status": {
    "rCode": 200,
    "bCodeMessage": null,
    "developerMessage": null
  }
}`)

var timeData, _ = time.Parse("1/02/2006", "05/09/2024")

var dividendTestCases = []struct {
	field string
	value interface{}
}{
	{"ExDividendDate", timeData},
	{"DividendYield", 15.51},
	{"AnnualDividendAmount", 1.88},
}

func TestDividend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(dividendServerData)
		if err != nil {
			t.Error("Error writing response")
		}
	}))
	defer server.Close()

	stock, err := NewDividendInfo("RWAY", config.Config{
		Url: server.URL,
	})
	if err != nil {
		t.Error("Error fetching data for stock", err)
	}

	for _, tc := range dividendTestCases {
		if fieldValue := tests.GetField(stock, tc.field); fieldValue != tc.value {
			t.Errorf("Expected %s to be %v, got %v", tc.field, tc.value, fieldValue)
		}
	}
}

func TestNasdaqDividendApi(t *testing.T) {
	// Use a well known dividend stock
	stock, err := NewDividendInfo("T")
	if err != nil {
		t.Error("Error fetching data for stock", err)
	}

	if stock == nil {
		t.Error("Stock is nil")
	}
}
