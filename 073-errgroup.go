package main

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
)

/*
 * Errgroup is an efficient way to handle errors for concurrent tasks.
 * It runs tasks until completion or until any one errors out, and if it does it cancels the context.
 */

func handlerErrgroup(ctx context.Context, circles []int) ([]int, error) {
	results := make([]int, len(circles))
	g, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		i := i
		circle := circle

		g.Go(func() error {
			result, err := processCircle(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func processCircle(ctx context.Context, circle int) (int, error) {
	return 0, errors.New("Kaboom!")
}
