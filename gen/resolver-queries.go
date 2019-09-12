package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/novacloudcz/graphql-orm/resolvers"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryChangelogChangeHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ChangelogChangeFilterType
}

func (r *GeneratedQueryResolver) ChangelogChange(ctx context.Context, id *string, q *string, filter *ChangelogChangeFilterType) (*ChangelogChange, error) {
	opts := QueryChangelogChangeHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryChangelogChange(ctx, r.GeneratedResolver, opts)
}
func QueryChangelogChangeHandler(ctx context.Context, r *GeneratedResolver, opts QueryChangelogChangeHandlerOptions) (*ChangelogChange, error) {
	query := ChangelogChangeQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ChangelogChangeResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("changelog_changes")+".id = ?", *opts.ID)
	}

	var items []*ChangelogChange
	err := rt.GetItems(ctx, qb, TableName("changelog_changes"), &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "ChangelogChange"}
	}
	return items[0], err
}

type QueryChangelogChangesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []ChangelogChangeSortType
	Filter *ChangelogChangeFilterType
}

func (r *GeneratedQueryResolver) ChangelogChanges(ctx context.Context, offset *int, limit *int, q *string, sort []ChangelogChangeSortType, filter *ChangelogChangeFilterType) (*ChangelogChangeResultType, error) {
	opts := QueryChangelogChangesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryChangelogChanges(ctx, r.GeneratedResolver, opts)
}
func QueryChangelogChangesHandler(ctx context.Context, r *GeneratedResolver, opts QueryChangelogChangesHandlerOptions) (*ChangelogChangeResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := ChangelogChangeQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &ChangelogChangeResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedChangelogChangeResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedChangelogChangeResultTypeResolver) Items(ctx context.Context, obj *ChangelogChangeResultType) (items []*ChangelogChange, err error) {
	err = obj.GetItems(ctx, r.DB.db, TableName("changelog_changes"), &items)
	return
}

func (r *GeneratedChangelogChangeResultTypeResolver) Count(ctx context.Context, obj *ChangelogChangeResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &ChangelogChange{})
}

type GeneratedChangelogChangeResolver struct{ *GeneratedResolver }

func (r *GeneratedChangelogChangeResolver) Log(ctx context.Context, obj *ChangelogChange) (res *Changelog, err error) {
	return r.Handlers.ChangelogChangeLog(ctx, r, obj)
}
func ChangelogChangeLogHandler(ctx context.Context, r *GeneratedChangelogChangeResolver, obj *ChangelogChange) (res *Changelog, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.LogID != nil {
		item, _err := loaders["Changelog"].Load(ctx, dataloader.StringKey(*obj.LogID))()
		res, _ = item.(*Changelog)

		err = _err
	}

	return
}

type QueryChangelogHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ChangelogFilterType
}

func (r *GeneratedQueryResolver) Changelog(ctx context.Context, id *string, q *string, filter *ChangelogFilterType) (*Changelog, error) {
	opts := QueryChangelogHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryChangelog(ctx, r.GeneratedResolver, opts)
}
func QueryChangelogHandler(ctx context.Context, r *GeneratedResolver, opts QueryChangelogHandlerOptions) (*Changelog, error) {
	query := ChangelogQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ChangelogResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where(TableName("changelogs")+".id = ?", *opts.ID)
	}

	var items []*Changelog
	err := rt.GetItems(ctx, qb, TableName("changelogs"), &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Changelog"}
	}
	return items[0], err
}

type QueryChangelogsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []ChangelogSortType
	Filter *ChangelogFilterType
}

func (r *GeneratedQueryResolver) Changelogs(ctx context.Context, offset *int, limit *int, q *string, sort []ChangelogSortType, filter *ChangelogFilterType) (*ChangelogResultType, error) {
	opts := QueryChangelogsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryChangelogs(ctx, r.GeneratedResolver, opts)
}
func QueryChangelogsHandler(ctx context.Context, r *GeneratedResolver, opts QueryChangelogsHandlerOptions) (*ChangelogResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := ChangelogQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &ChangelogResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedChangelogResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedChangelogResultTypeResolver) Items(ctx context.Context, obj *ChangelogResultType) (items []*Changelog, err error) {
	err = obj.GetItems(ctx, r.DB.db, TableName("changelogs"), &items)
	return
}

func (r *GeneratedChangelogResultTypeResolver) Count(ctx context.Context, obj *ChangelogResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Changelog{})
}

type GeneratedChangelogResolver struct{ *GeneratedResolver }

func (r *GeneratedChangelogResolver) Changes(ctx context.Context, obj *Changelog) (res []*ChangelogChange, err error) {
	return r.Handlers.ChangelogChanges(ctx, r, obj)
}
func ChangelogChangesHandler(ctx context.Context, r *GeneratedChangelogResolver, obj *Changelog) (res []*ChangelogChange, err error) {

	items := []*ChangelogChange{}
	err = r.DB.Query().Model(obj).Related(&items, "Changes").Error
	res = items

	return
}

func (r *GeneratedChangelogResolver) ChangesIds(ctx context.Context, obj *Changelog) (ids []string, err error) {
	ids = []string{}

	items := []*ChangelogChange{}
	err = r.DB.Query().Model(obj).Select(TableName("changelog_changes")+".id").Related(&items, "Changes").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}
