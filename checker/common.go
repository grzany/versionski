package checker

// CheckResponse is a response for a Check request.
type CheckResponse struct {
	// Current is current latest version on source.
	Current string

	// Outdate is true when target version is less than Curernt on source.
	Outdated bool

	// Latest is true when target version is equal to Current on source.
	Latest bool
}

type FetchResponse struct {
	Version string
}

// Source is the interface that every version information source must implement.
type Source interface {

	// Validate checks if params satisfy Fetch() for given source
	Validate() error

	// Fetch is called in Check to fetch information from remote sources.
	// After fetching, it will convert it into common expression (FetchResponse)
	Fetch() (*FetchResponse, error)
}
