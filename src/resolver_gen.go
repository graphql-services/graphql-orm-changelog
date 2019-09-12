package src

import (
	"github.com/novacloudcz/graphql-orm-changelog/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func NewResolver(db *gen.DB, ec *events.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type ChangelogChangeResultTypeResolver struct {
	*gen.GeneratedChangelogChangeResultTypeResolver
}

func (r *Resolver) ChangelogChangeResultType() gen.ChangelogChangeResultTypeResolver {
	return &ChangelogChangeResultTypeResolver{&gen.GeneratedChangelogChangeResultTypeResolver{r.GeneratedResolver}}
}

type ChangelogChangeResolver struct {
	*gen.GeneratedChangelogChangeResolver
}

func (r *Resolver) ChangelogChange() gen.ChangelogChangeResolver {
	return &ChangelogChangeResolver{&gen.GeneratedChangelogChangeResolver{r.GeneratedResolver}}
}

type ChangelogResultTypeResolver struct {
	*gen.GeneratedChangelogResultTypeResolver
}

func (r *Resolver) ChangelogResultType() gen.ChangelogResultTypeResolver {
	return &ChangelogResultTypeResolver{&gen.GeneratedChangelogResultTypeResolver{r.GeneratedResolver}}
}

type ChangelogResolver struct {
	*gen.GeneratedChangelogResolver
}

func (r *Resolver) Changelog() gen.ChangelogResolver {
	return &ChangelogResolver{&gen.GeneratedChangelogResolver{r.GeneratedResolver}}
}
