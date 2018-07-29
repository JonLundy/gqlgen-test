package resolver

import "test/graph"

type repository struct{}

func (repository) AllBugs(ctx context.Context, obj *graph.Repository, input graph.ConnectionInput) (graph.BugConnection, error) {
	panic("implement me")
}

func (repository) Bug(ctx context.Context, obj *graph.Repository, prefix string) (*graph.Bug, error) {
	panic("implement me")
}


