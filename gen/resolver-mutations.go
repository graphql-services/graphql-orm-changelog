package gen

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/novacloudcz/graphql-orm/events"
)

// GeneratedMutationResolver ...
type GeneratedMutationResolver struct{ *GeneratedResolver }

// MutationEvents ...
type MutationEvents struct {
	Events []events.Event
}

// EnrichContextWithMutations ...
func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.GetDB(ctx).Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}

// FinishMutationContext ...
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := r.GetDB(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}

// RollbackMutationContext ...
func RollbackMutationContext(ctx context.Context, r *GeneratedResolver) error {
	tx := r.GetDB(ctx)
	return tx.Rollback().Error
}

// GetMutationEventStore ...
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}

// AddMutationEvent ...
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

// CreateChangelogChange ...
func (r *GeneratedMutationResolver) CreateChangelogChange(ctx context.Context, input map[string]interface{}) (item *ChangelogChange, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateChangelogChange(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateChangelogChangeHandler ...
func CreateChangelogChangeHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ChangelogChange, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &ChangelogChange{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "ChangelogChange",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ChangelogChangeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["column"]; ok && (item.Column != changes.Column) {
		item.Column = changes.Column

		event.AddNewValue("column", changes.Column)
	}

	if _, ok := input["oldValue"]; ok && (item.OldValue != changes.OldValue) && (item.OldValue == nil || changes.OldValue == nil || *item.OldValue != *changes.OldValue) {
		item.OldValue = changes.OldValue

		event.AddNewValue("oldValue", changes.OldValue)
	}

	if _, ok := input["newValue"]; ok && (item.NewValue != changes.NewValue) && (item.NewValue == nil || changes.NewValue == nil || *item.NewValue != *changes.NewValue) {
		item.NewValue = changes.NewValue

		event.AddNewValue("newValue", changes.NewValue)
	}

	if _, ok := input["logId"]; ok && (item.LogID != changes.LogID) && (item.LogID == nil || changes.LogID == nil || *item.LogID != *changes.LogID) {
		item.LogID = changes.LogID

		event.AddNewValue("logId", changes.LogID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateChangelogChange ...
func (r *GeneratedMutationResolver) UpdateChangelogChange(ctx context.Context, id string, input map[string]interface{}) (item *ChangelogChange, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateChangelogChange(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateChangelogChangeHandler ...
func UpdateChangelogChangeHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ChangelogChange, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ChangelogChange{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "ChangelogChange",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ChangelogChangeChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["column"]; ok && (item.Column != changes.Column) {
		event.AddOldValue("column", item.Column)
		event.AddNewValue("column", changes.Column)
		item.Column = changes.Column
	}

	if _, ok := input["oldValue"]; ok && (item.OldValue != changes.OldValue) && (item.OldValue == nil || changes.OldValue == nil || *item.OldValue != *changes.OldValue) {
		event.AddOldValue("oldValue", item.OldValue)
		event.AddNewValue("oldValue", changes.OldValue)
		item.OldValue = changes.OldValue
	}

	if _, ok := input["newValue"]; ok && (item.NewValue != changes.NewValue) && (item.NewValue == nil || changes.NewValue == nil || *item.NewValue != *changes.NewValue) {
		event.AddOldValue("newValue", item.NewValue)
		event.AddNewValue("newValue", changes.NewValue)
		item.NewValue = changes.NewValue
	}

	if _, ok := input["logId"]; ok && (item.LogID != changes.LogID) && (item.LogID == nil || changes.LogID == nil || *item.LogID != *changes.LogID) {
		event.AddOldValue("logId", item.LogID)
		event.AddNewValue("logId", changes.LogID)
		item.LogID = changes.LogID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteChangelogChange ...
func (r *GeneratedMutationResolver) DeleteChangelogChange(ctx context.Context, id string) (item *ChangelogChange, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteChangelogChange(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteChangelogChangeHandler ...
func DeleteChangelogChangeHandler(ctx context.Context, r *GeneratedResolver, id string) (item *ChangelogChange, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ChangelogChange{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "ChangelogChange",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("changelog_changes")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllChangelogChanges ...
func (r *GeneratedMutationResolver) DeleteAllChangelogChanges(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllChangelogChanges(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllChangelogChangesHandler ...
func DeleteAllChangelogChangesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&ChangelogChange{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateChangelog ...
func (r *GeneratedMutationResolver) CreateChangelog(ctx context.Context, input map[string]interface{}) (item *Changelog, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateChangelog(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateChangelogHandler ...
func CreateChangelogHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Changelog, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Changelog{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Changelog",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ChangelogChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["entity"]; ok && (item.Entity != changes.Entity) {
		item.Entity = changes.Entity

		event.AddNewValue("entity", changes.Entity)
	}

	if _, ok := input["entityID"]; ok && (item.EntityID != changes.EntityID) {
		item.EntityID = changes.EntityID

		event.AddNewValue("entityID", changes.EntityID)
	}

	if _, ok := input["principalID"]; ok && (item.PrincipalID != changes.PrincipalID) && (item.PrincipalID == nil || changes.PrincipalID == nil || *item.PrincipalID != *changes.PrincipalID) {
		item.PrincipalID = changes.PrincipalID

		event.AddNewValue("principalID", changes.PrincipalID)
	}

	if _, ok := input["type"]; ok && (item.Type != changes.Type) {
		item.Type = changes.Type

		event.AddNewValue("type", changes.Type)
	}

	if _, ok := input["date"]; ok && (item.Date != changes.Date) {
		item.Date = changes.Date

		event.AddNewValue("date", changes.Date)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["changesIds"]; exists {
		err = tx.Model(ChangelogChange{}).Where("id IN (?)", ids).Update("logId", item.ID).Error
		if err != nil {
			return
		}
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateChangelog ...
func (r *GeneratedMutationResolver) UpdateChangelog(ctx context.Context, id string, input map[string]interface{}) (item *Changelog, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateChangelog(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateChangelogHandler ...
func UpdateChangelogHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Changelog, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Changelog{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Changelog",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes ChangelogChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["entity"]; ok && (item.Entity != changes.Entity) {
		event.AddOldValue("entity", item.Entity)
		event.AddNewValue("entity", changes.Entity)
		item.Entity = changes.Entity
	}

	if _, ok := input["entityID"]; ok && (item.EntityID != changes.EntityID) {
		event.AddOldValue("entityID", item.EntityID)
		event.AddNewValue("entityID", changes.EntityID)
		item.EntityID = changes.EntityID
	}

	if _, ok := input["principalID"]; ok && (item.PrincipalID != changes.PrincipalID) && (item.PrincipalID == nil || changes.PrincipalID == nil || *item.PrincipalID != *changes.PrincipalID) {
		event.AddOldValue("principalID", item.PrincipalID)
		event.AddNewValue("principalID", changes.PrincipalID)
		item.PrincipalID = changes.PrincipalID
	}

	if _, ok := input["type"]; ok && (item.Type != changes.Type) {
		event.AddOldValue("type", item.Type)
		event.AddNewValue("type", changes.Type)
		item.Type = changes.Type
	}

	if _, ok := input["date"]; ok && (item.Date != changes.Date) {
		event.AddOldValue("date", item.Date)
		event.AddNewValue("date", changes.Date)
		item.Date = changes.Date
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["changesIds"]; exists {
		items := []ChangelogChange{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Changes")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteChangelog ...
func (r *GeneratedMutationResolver) DeleteChangelog(ctx context.Context, id string) (item *Changelog, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteChangelog(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteChangelogHandler ...
func DeleteChangelogHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Changelog, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Changelog{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Changelog",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("changelogs")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllChangelogs ...
func (r *GeneratedMutationResolver) DeleteAllChangelogs(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllChangelogs(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllChangelogsHandler ...
func DeleteAllChangelogsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	// delete all resolvers are primarily used for
	if os.Getenv("ENABLE_DELETE_ALL_RESOLVERS") == "" {
		return false, fmt.Errorf("delete all resolver is not enabled (ENABLE_DELETE_ALL_RESOLVERS not specified)")
	}
	tx := r.GetDB(ctx)
	err := tx.Delete(&Changelog{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
