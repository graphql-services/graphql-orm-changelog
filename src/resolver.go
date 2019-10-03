package src

import (
	"github.com/novacloudcz/graphql-orm-changelog/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func New(db *gen.DB, ec *events.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	return resolver
}
