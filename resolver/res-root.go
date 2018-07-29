package resolver

import "test/graph"

type Root struct{}

func (Root) Query() graph.QueryResolver {
	return query{}
}

func (Root) Repository() graph.RepositoryResolver {
	return repository{}
}

