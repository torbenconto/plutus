package plutus

// StockDataProvider is an interface for different data providers.
type StockDataProvider interface {
	Populate(s *Stock, apiKey ...string) (*Stock, error)
}
