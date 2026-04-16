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
	var (
		mu    sync.Mutex
		wg    sync.WaitGroup
		sem   = make(chan struct{}, maxConcurrent)
		errCh = make(chan error, 1) // buffered so send never blocks
	)

outer:
	for _, id := range ids {
		select {
		case <-childCtx.Done():
			break outer
		default:
		}

		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			select {
			case sem <- struct{}{}:
			case <-childCtx.Done():
				return
			}
			defer func() { <-sem }()

			price, err := mockFetch(childCtx, id)
			if err != nil {
				select {
				case errCh <- fmt.Errorf("fetch failed for %s: %w", id, err):
					cancel() // trigger context cancel on first error
				default:
					// ignore subsequent errors if one already sent
				}
				return
			}

			if childCtx.Err() == nil {
				mu.Lock()
				results[id] = price
				mu.Unlock()
			}
		}(id)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	// wait for either an error or full completion
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
