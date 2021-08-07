package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/ast"
)

// GeneratedQueryResolver ...
type GeneratedQueryResolver struct{ *GeneratedResolver }

// QueryChangelogChangeHandlerOptions ...
type QueryChangelogChangeHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ChangelogChangeFilterType
}

// ChangelogChange ...
func (r *GeneratedQueryResolver) ChangelogChange(ctx context.Context, id *string, q *string, filter *ChangelogChangeFilterType) (*ChangelogChange, error) {
	opts := QueryChangelogChangeHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryChangelogChange(ctx, r.GeneratedResolver, opts)
}

// QueryChangelogChangeHandler ...
func QueryChangelogChangeHandler(ctx context.Context, r *GeneratedResolver, opts QueryChangelogChangeHandlerOptions) (*ChangelogChange, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := ChangelogChangeQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ChangelogChangeResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("changelog_changes")+".id = ?", *opts.ID)
	}

	var items []*ChangelogChange
	giOpts := GetItemsOptions{
		Alias:      TableName("changelog_changes"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryChangelogChangesHandlerOptions ...
type QueryChangelogChangesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ChangelogChangeSortType
	Filter *ChangelogChangeFilterType
}

// ChangelogChanges ...
func (r *GeneratedQueryResolver) ChangelogChanges(ctx context.Context, offset *int, limit *int, q *string, sort []*ChangelogChangeSortType, filter *ChangelogChangeFilterType) (*ChangelogChangeResultType, error) {
	opts := QueryChangelogChangesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryChangelogChanges(ctx, r.GeneratedResolver, opts)
}

// ChangelogChangesItems ...
func (r *GeneratedResolver) ChangelogChangesItems(ctx context.Context, opts QueryChangelogChangesHandlerOptions) (res []*ChangelogChange, err error) {
	resultType, err := r.Handlers.QueryChangelogChanges(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("changelog_changes"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// ChangelogChangesCount ...
func (r *GeneratedResolver) ChangelogChangesCount(ctx context.Context, opts QueryChangelogChangesHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryChangelogChanges(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("changelog_changes"),
	}, &ChangelogChange{})
}

// QueryChangelogChangesHandler ...
func QueryChangelogChangesHandler(ctx context.Context, r *GeneratedResolver, opts QueryChangelogChangesHandlerOptions) (*ChangelogChangeResultType, error) {
	query := ChangelogChangeQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ChangelogChangeResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedChangelogChangeResultTypeResolver ...
type GeneratedChangelogChangeResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedChangelogChangeResultTypeResolver) Items(ctx context.Context, obj *ChangelogChangeResultType) (items []*ChangelogChange, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("changelog_changes"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), opts, &items)

	uniqueItems := []*ChangelogChange{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedChangelogChangeResultTypeResolver) Count(ctx context.Context, obj *ChangelogChangeResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("changelog_changes"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &ChangelogChange{})
}

// Aggregations ...
func (r *GeneratedChangelogChangeResultTypeResolver) Aggregations(ctx context.Context, obj *ChangelogChangeResultType) (res *ChangelogChangeResultAggregations, err error) {
	aggregationsMap := map[string]GetAggregationsAggregationField{

		"columnMax":   GetAggregationsAggregationField{"column", "Max"},
		"columnMin":   GetAggregationsAggregationField{"column", "Min"},
		"newValueMax": GetAggregationsAggregationField{"newValue", "Max"},
		"newValueMin": GetAggregationsAggregationField{"newValue", "Min"},
		"oldValueMax": GetAggregationsAggregationField{"oldValue", "Max"},
		"oldValueMin": GetAggregationsAggregationField{"oldValue", "Min"},
	}
	aggregationFieldsMap := map[string]string{

		"columnMax":   "column",
		"columnMin":   "column",
		"newValueMax": "newValue",
		"newValueMin": "newValue",
		"oldValueMax": "oldValue",
		"oldValueMin": "oldValue",
	}
	fieldsMap := map[string]struct{}{}
	fields := []string{}
	aggregationFields := []GetAggregationsAggregationField{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		aggregationFields = append(aggregationFields, aggregationsMap[f.Field.Name])
		fieldsMap[aggregationFieldsMap[f.Field.Name]] = struct{}{}
	}
	for key, _ := range fieldsMap {
		fields = append(fields, key)
	}

	opts := GetAggregationsOptions{
		Alias:             TableName("changelog_changes"),
		Fields:            fields,
		AggregationFields: aggregationFields,
	}
	res = &ChangelogChangeResultAggregations{}
	err = obj.GetAggregations(ctx, r.GetDB(ctx), opts, &ChangelogChange{}, res)
	return
}

type GeneratedChangelogChangeResolver struct{ *GeneratedResolver }

// Log ...
func (r *GeneratedChangelogChangeResolver) Log(ctx context.Context, obj *ChangelogChange) (res *Changelog, err error) {
	return r.Handlers.ChangelogChangeLog(ctx, r.GeneratedResolver, obj)
}

// LogHandler ...
func ChangelogChangeLogHandler(ctx context.Context, r *GeneratedResolver, obj *ChangelogChange) (res *Changelog, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.LogID != nil {
		item, _err := loaders["Changelog"].Load(ctx, dataloader.StringKey(*obj.LogID))()
		res, _ = item.(*Changelog)

		err = _err
	}

	return
}

// QueryChangelogHandlerOptions ...
type QueryChangelogHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *ChangelogFilterType
}

// Changelog ...
func (r *GeneratedQueryResolver) Changelog(ctx context.Context, id *string, q *string, filter *ChangelogFilterType) (*Changelog, error) {
	opts := QueryChangelogHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryChangelog(ctx, r.GeneratedResolver, opts)
}

// QueryChangelogHandler ...
func QueryChangelogHandler(ctx context.Context, r *GeneratedResolver, opts QueryChangelogHandlerOptions) (*Changelog, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := ChangelogQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &ChangelogResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if opts.ID != nil {
		qb = qb.Where(TableName("changelogs")+".id = ?", *opts.ID)
	}

	var items []*Changelog
	giOpts := GetItemsOptions{
		Alias: TableName("changelogs"),
		Preloaders: []string{
			"Changes",
		},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

// QueryChangelogsHandlerOptions ...
type QueryChangelogsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*ChangelogSortType
	Filter *ChangelogFilterType
}

// Changelogs ...
func (r *GeneratedQueryResolver) Changelogs(ctx context.Context, offset *int, limit *int, q *string, sort []*ChangelogSortType, filter *ChangelogFilterType) (*ChangelogResultType, error) {
	opts := QueryChangelogsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryChangelogs(ctx, r.GeneratedResolver, opts)
}

// ChangelogsItems ...
func (r *GeneratedResolver) ChangelogsItems(ctx context.Context, opts QueryChangelogsHandlerOptions) (res []*Changelog, err error) {
	resultType, err := r.Handlers.QueryChangelogs(ctx, r, opts)
	if err != nil {
		return
	}
	err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("changelogs"),
	}, &res)
	if err != nil {
		return
	}
	return
}

// ChangelogsCount ...
func (r *GeneratedResolver) ChangelogsCount(ctx context.Context, opts QueryChangelogsHandlerOptions) (count int, err error) {
	resultType, err := r.Handlers.QueryChangelogs(ctx, r, opts)
	if err != nil {
		return
	}
	return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
		Alias: TableName("changelogs"),
	}, &Changelog{})
}

// QueryChangelogsHandler ...
func QueryChangelogsHandler(ctx context.Context, r *GeneratedResolver, opts QueryChangelogsHandlerOptions) (*ChangelogResultType, error) {
	query := ChangelogQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &ChangelogResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedChangelogResultTypeResolver ...
type GeneratedChangelogResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedChangelogResultTypeResolver) Items(ctx context.Context, obj *ChangelogResultType) (items []*Changelog, err error) {
	opts := GetItemsOptions{
		Alias: TableName("changelogs"),
		Preloaders: []string{
			"Changes",
		},
	}
	err = obj.GetItems(ctx, r.GetDB(ctx), opts, &items)

	for _, item := range items {

		item.ChangesPreloaded = true
	}

	uniqueItems := []*Changelog{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedChangelogResultTypeResolver) Count(ctx context.Context, obj *ChangelogResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias: TableName("changelogs"),
		Preloaders: []string{
			"Changes",
		},
	}
	return obj.GetCount(ctx, r.GetDB(ctx), opts, &Changelog{})
}

// Aggregations ...
func (r *GeneratedChangelogResultTypeResolver) Aggregations(ctx context.Context, obj *ChangelogResultType) (res *ChangelogResultAggregations, err error) {
	aggregationsMap := map[string]GetAggregationsAggregationField{

		"entityIDMax":    GetAggregationsAggregationField{"entityID", "Max"},
		"entityIDMin":    GetAggregationsAggregationField{"entityID", "Min"},
		"entityMax":      GetAggregationsAggregationField{"entity", "Max"},
		"entityMin":      GetAggregationsAggregationField{"entity", "Min"},
		"principalIDMax": GetAggregationsAggregationField{"principalID", "Max"},
		"principalIDMin": GetAggregationsAggregationField{"principalID", "Min"},
	}
	aggregationFieldsMap := map[string]string{

		"entityIDMax":    "entityID",
		"entityIDMin":    "entityID",
		"entityMax":      "entity",
		"entityMin":      "entity",
		"principalIDMax": "principalID",
		"principalIDMin": "principalID",
	}
	fieldsMap := map[string]struct{}{}
	fields := []string{}
	aggregationFields := []GetAggregationsAggregationField{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		aggregationFields = append(aggregationFields, aggregationsMap[f.Field.Name])
		fieldsMap[aggregationFieldsMap[f.Field.Name]] = struct{}{}
	}
	for key, _ := range fieldsMap {
		fields = append(fields, key)
	}

	opts := GetAggregationsOptions{
		Alias:             TableName("changelogs"),
		Fields:            fields,
		AggregationFields: aggregationFields,
	}
	res = &ChangelogResultAggregations{}
	err = obj.GetAggregations(ctx, r.GetDB(ctx), opts, &Changelog{}, res)
	return
}

type GeneratedChangelogResolver struct{ *GeneratedResolver }

// Changes ...
func (r *GeneratedChangelogResolver) Changes(ctx context.Context, obj *Changelog) (res []*ChangelogChange, err error) {
	return r.Handlers.ChangelogChanges(ctx, r.GeneratedResolver, obj)
}

// ChangesHandler ...
func ChangelogChangesHandler(ctx context.Context, r *GeneratedResolver, obj *Changelog) (res []*ChangelogChange, err error) {

	if obj.ChangesPreloaded {
		res = obj.Changes
	} else {

		items := []*ChangelogChange{}
		db := r.GetDB(ctx)
		err = db.Model(obj).Related(&items, "Changes").Error
		res = items

	}

	return
}

// ChangesIds ...
func (r *GeneratedChangelogResolver) ChangesIds(ctx context.Context, obj *Changelog) (ids []string, err error) {
	ids = []string{}

	items := []*ChangelogChange{}
	db := r.GetDB(ctx)
	err = db.Model(obj).Select(TableName("changelog_changes")+".id").Related(&items, "Changes").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// ChangesConnection ...
func (r *GeneratedChangelogResolver) ChangesConnection(ctx context.Context, obj *Changelog, offset *int, limit *int, q *string, sort []*ChangelogChangeSortType, filter *ChangelogChangeFilterType) (res *ChangelogChangeResultType, err error) {
	f := &ChangelogChangeFilterType{
		Log: &ChangelogFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &ChangelogChangeFilterType{
			And: []*ChangelogChangeFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryChangelogChangesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryChangelogChanges(ctx, r.GeneratedResolver, opts)
}
