package parallelhttp

import (
	"context"
	"sync"
)

type Fetcher interface {
	Fetch(ctx context.Context, url string) (string, error)
}

func FetchAll(ctx context.Context, f Fetcher, urls []string) ([]string, error) {
	if len(urls) == 0 {
		return []string{}, nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	results := make([]string, len(urls))
	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		firstErr error
	)

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			if ctx.Err() != nil {
				return
			}
			body, err := f.Fetch(ctx, url)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err
					cancel()
				}
				mu.Unlock()
				return
			}
			results[i] = body
		}(i, url)
	}

	wg.Wait()
	if firstErr != nil {
		return nil, firstErr
	}
	return results, nil
}
