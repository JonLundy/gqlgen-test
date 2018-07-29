package resolver

import "test/graph"

type query struct{}

func (query) DefaultRepository(ctx context.Context) (*graph.Repository, error) {
	panic("implement me")
}

func (query) Repository(ctx context.Context, id string) (*graph.Repository, error) {
	panic("implement me")
}


