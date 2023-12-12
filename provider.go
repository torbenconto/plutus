package plutus

// StockDataProvider is an interface for different data providers.
type StockDataProvider interface {
	Populate(*Stock) (*Stock, error)
}
