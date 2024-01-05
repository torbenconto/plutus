package plutus

// Internal struct for polygon api provider
type p_PolygonApiProvider struct {
	apiKey string
}

// Expose variable containing provider
var PolygonApiProvider *p_PolygonApiProvider = &p_PolygonApiProvider{}

// Populate fills in the fields of the Stock struct with data from polygon.io (requires api key).
func (p *p_PolygonApiProvider) Populate(s *Stock, apiKey ...string) (*Stock, error) {
	if len(apiKey) > 0 {
		p.apiKey = apiKey[0]
	} else {
		return nil, ErrNoAPIKey
	}

	return s, nil
}
