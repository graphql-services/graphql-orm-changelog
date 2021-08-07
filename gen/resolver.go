package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

// ResolutionHandlers ...
type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateChangelogChange     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ChangelogChange, err error)
	UpdateChangelogChange     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ChangelogChange, err error)
	DeleteChangelogChange     func(ctx context.Context, r *GeneratedResolver, id string) (item *ChangelogChange, err error)
	DeleteAllChangelogChanges func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryChangelogChange      func(ctx context.Context, r *GeneratedResolver, opts QueryChangelogChangeHandlerOptions) (*ChangelogChange, error)
	QueryChangelogChanges     func(ctx context.Context, r *GeneratedResolver, opts QueryChangelogChangesHandlerOptions) (*ChangelogChangeResultType, error)

	ChangelogChangeLog func(ctx context.Context, r *GeneratedResolver, obj *ChangelogChange) (res *Changelog, err error)

	CreateChangelog     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Changelog, err error)
	UpdateChangelog     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Changelog, err error)
	DeleteChangelog     func(ctx context.Context, r *GeneratedResolver, id string) (item *Changelog, err error)
	DeleteAllChangelogs func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryChangelog      func(ctx context.Context, r *GeneratedResolver, opts QueryChangelogHandlerOptions) (*Changelog, error)
	QueryChangelogs     func(ctx context.Context, r *GeneratedResolver, opts QueryChangelogsHandlerOptions) (*ChangelogResultType, error)

	ChangelogChanges func(ctx context.Context, r *GeneratedResolver, obj *Changelog) (res []*ChangelogChange, err error)
}

// DefaultResolutionHandlers ...
func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

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

// GeneratedResolver ...
type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	db              *DB
	EventController *EventController
}

// NewGeneratedResolver ...
func NewGeneratedResolver(handlers ResolutionHandlers, db *DB, ec *EventController) *GeneratedResolver {
	return &GeneratedResolver{Handlers: handlers, db: db, EventController: ec}
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.db.Query()
	}
	return db
}
