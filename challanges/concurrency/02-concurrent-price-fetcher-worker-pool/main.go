package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

func FetchAll(ctx context.Context, ids []string, maxConcurrent int) (map[string]int, error) {
	if maxConcurrent < 1 {
		maxConcurrent = 1
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	results := make(map[string]int, len(ids))
	var mu sync.Mutex

	jobs := make(chan string)
	errCh := make(chan error, 1) // buffered: allow single non-blocking send
	var wg sync.WaitGroup

	// Start worker pool
	for i := 0; i < maxConcurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range jobs {
				select {
				case <-childCtx.Done():
					return
				default:
				}

				price, err := mockFetch(childCtx, id)
				if err != nil {
					select {
					case errCh <- fmt.Errorf("fetch failed for %s: %w", id, err):
						cancel() // cancel others on first error
					default:
					}
					return
				}

				if childCtx.Err() == nil {
					mu.Lock()
					results[id] = price
					mu.Unlock()
				}
			}
		}()
	}

	// Feed jobs
	go func() {
		defer close(jobs)
		for _, id := range ids {
			select {
			case <-childCtx.Done():
				return
			case jobs <- id:
			}
		}
	}()

	// Wait for workers in background and close errCh once done
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Wait for either an error or all workers finishing cleanly
	select {
	case err := <-errCh:
		if err != nil {
			return nil, err
		}
	case <-childCtx.Done():
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
	}

	return results, nil
}

func mockFetch(ctx context.Context, id string) (int, error) {
	if strings.HasPrefix(id, "bad:") {
		return 0, errors.New("remote error")
	}
	if strings.HasPrefix(id, "slow:") {
		<-ctx.Done()
		return 0, ctx.Err()
	}
	sum := 0
	for _, r := range id {
		sum += int(r)
	}
	return len(id) + (sum % 17), nil
}
