package parallelhttp

import "context"

// Fetcher loads a resource body for a URL (HTTP client, fake, etc.).
type Fetcher interface {
	Fetch(ctx context.Context, url string) (string, error)
}

// FetchAll returns one body per URL in the same order as urls.
func FetchAll(ctx context.Context, f Fetcher, urls []string) ([]string, error) {
	// FIX ME: fetch concurrently; preserve order; cancel on first error
	out := make([]string, len(urls))
	return out, nil
}
