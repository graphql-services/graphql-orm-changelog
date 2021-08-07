package src

import (
	"context"

	"github.com/graphql-services/graphql-orm-changelog/gen"
)

// NewResolver ...
func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{gen.NewGeneratedResolver(handlers, db, ec)}
}

// Resolver ...
type Resolver struct {
	*gen.GeneratedResolver
}

// MutationResolver ...
type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

// BeginTransaction ...
func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
		tx := r.GeneratedResolver.GetDB(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

// QueryResolver ...
type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

// Mutation ...
func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}

// Query ...
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

// ChangelogChangeResultTypeResolver ...
type ChangelogChangeResultTypeResolver struct {
	*gen.GeneratedChangelogChangeResultTypeResolver
}

// ChangelogChangeResultType ...
func (r *Resolver) ChangelogChangeResultType() gen.ChangelogChangeResultTypeResolver {
	return &ChangelogChangeResultTypeResolver{&gen.GeneratedChangelogChangeResultTypeResolver{r.GeneratedResolver}}
}

// ChangelogChangeResolver ...
type ChangelogChangeResolver struct {
	*gen.GeneratedChangelogChangeResolver
}

// ChangelogChange ...
func (r *Resolver) ChangelogChange() gen.ChangelogChangeResolver {
	return &ChangelogChangeResolver{&gen.GeneratedChangelogChangeResolver{r.GeneratedResolver}}
}

// ChangelogResultTypeResolver ...
type ChangelogResultTypeResolver struct {
	*gen.GeneratedChangelogResultTypeResolver
}

// ChangelogResultType ...
func (r *Resolver) ChangelogResultType() gen.ChangelogResultTypeResolver {
	return &ChangelogResultTypeResolver{&gen.GeneratedChangelogResultTypeResolver{r.GeneratedResolver}}
}

// ChangelogResolver ...
type ChangelogResolver struct {
	*gen.GeneratedChangelogResolver
}

// Changelog ...
func (r *Resolver) Changelog() gen.ChangelogResolver {
	return &ChangelogResolver{&gen.GeneratedChangelogResolver{r.GeneratedResolver}}
}
