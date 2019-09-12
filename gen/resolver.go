package gen

import (
	"context"

	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	CreateChangelogChange     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ChangelogChange, err error)
	UpdateChangelogChange     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ChangelogChange, err error)
	DeleteChangelogChange     func(ctx context.Context, r *GeneratedResolver, id string) (item *ChangelogChange, err error)
	DeleteAllChangelogChanges func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryChangelogChange      func(ctx context.Context, r *GeneratedResolver, opts QueryChangelogChangeHandlerOptions) (*ChangelogChange, error)
	QueryChangelogChanges     func(ctx context.Context, r *GeneratedResolver, opts QueryChangelogChangesHandlerOptions) (*ChangelogChangeResultType, error)

	ChangelogChangeLog func(ctx context.Context, r *GeneratedChangelogChangeResolver, obj *ChangelogChange) (res *Changelog, err error)

	CreateChangelog     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Changelog, err error)
	UpdateChangelog     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Changelog, err error)
	DeleteChangelog     func(ctx context.Context, r *GeneratedResolver, id string) (item *Changelog, err error)
	DeleteAllChangelogs func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryChangelog      func(ctx context.Context, r *GeneratedResolver, opts QueryChangelogHandlerOptions) (*Changelog, error)
	QueryChangelogs     func(ctx context.Context, r *GeneratedResolver, opts QueryChangelogsHandlerOptions) (*ChangelogResultType, error)

	ChangelogChanges func(ctx context.Context, r *GeneratedChangelogResolver, obj *Changelog) (res []*ChangelogChange, err error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{

		CreateChangelogChange:     CreateChangelogChangeHandler,
		UpdateChangelogChange:     UpdateChangelogChangeHandler,
		DeleteChangelogChange:     DeleteChangelogChangeHandler,
		DeleteAllChangelogChanges: DeleteAllChangelogChangesHandler,
		QueryChangelogChange:      QueryChangelogChangeHandler,
		QueryChangelogChanges:     QueryChangelogChangesHandler,

		ChangelogChangeLog: ChangelogChangeLogHandler,

		CreateChangelog:     CreateChangelogHandler,
		UpdateChangelog:     UpdateChangelogHandler,
		DeleteChangelog:     DeleteChangelogHandler,
		DeleteAllChangelogs: DeleteAllChangelogsHandler,
		QueryChangelog:      QueryChangelogHandler,
		QueryChangelogs:     QueryChangelogsHandler,

		ChangelogChanges: ChangelogChangesHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *events.EventController
}
