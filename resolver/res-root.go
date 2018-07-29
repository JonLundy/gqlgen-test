package resolver

import "test/graph"

type Root struct{}

func (Root) Query() graph.QueryResolver {
	panic("implement me")
}

func (Root) Repository() graph.RepositoryResolver {
	panic("implement me")
}

