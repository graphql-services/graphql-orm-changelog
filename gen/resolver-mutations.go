package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/novacloudcz/graphql-orm/events"
	"github.com/novacloudcz/graphql-orm/resolvers"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

func (r *GeneratedMutationResolver) CreateChangelogChange(ctx context.Context, input map[string]interface{}) (item *ChangelogChange, err error) {
	return r.Handlers.CreateChangelogChange(ctx, r.GeneratedResolver, input)
}
func CreateChangelogChangeHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *ChangelogChange, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &ChangelogChange{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

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
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
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

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateChangelogChange(ctx context.Context, id string, input map[string]interface{}) (item *ChangelogChange, err error) {
	return r.Handlers.UpdateChangelogChange(ctx, r.GeneratedResolver, id, input)
}
func UpdateChangelogChangeHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *ChangelogChange, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ChangelogChange{}
	now := time.Now()
	tx := r.DB.db.Begin()

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
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
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

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteChangelogChange(ctx context.Context, id string) (item *ChangelogChange, err error) {
	return r.Handlers.DeleteChangelogChange(ctx, r.GeneratedResolver, id)
}
func DeleteChangelogChangeHandler(ctx context.Context, r *GeneratedResolver, id string) (item *ChangelogChange, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &ChangelogChange{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "ChangelogChange",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, "changelog_changes.id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllChangelogChanges(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllChangelogChanges(ctx, r.GeneratedResolver)
}
func DeleteAllChangelogChangesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&ChangelogChange{}).Error
	return err == nil, err
}

func (r *GeneratedMutationResolver) CreateChangelog(ctx context.Context, input map[string]interface{}) (item *Changelog, err error) {
	return r.Handlers.CreateChangelog(ctx, r.GeneratedResolver, input)
}
func CreateChangelogHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Changelog, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Changelog{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.DB.db.Begin()

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
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
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

	if ids, ok := input["changesIds"].([]interface{}); ok {
		items := []ChangelogChange{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Changes")
		association.Replace(items)
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func (r *GeneratedMutationResolver) UpdateChangelog(ctx context.Context, id string, input map[string]interface{}) (item *Changelog, err error) {
	return r.Handlers.UpdateChangelog(ctx, r.GeneratedResolver, id, input)
}
func UpdateChangelogHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Changelog, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Changelog{}
	now := time.Now()
	tx := r.DB.db.Begin()

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
		return
	}

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
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

	if ids, ok := input["changesIds"].([]interface{}); ok {
		items := []ChangelogChange{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Changes")
		association.Replace(items)
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		err = r.EventController.SendEvent(ctx, &event)
		// data, _ := json.Marshal(event)
		// fmt.Println("?",string(data))
	}

	return
}
func (r *GeneratedMutationResolver) DeleteChangelog(ctx context.Context, id string) (item *Changelog, err error) {
	return r.Handlers.DeleteChangelog(ctx, r.GeneratedResolver, id)
}
func DeleteChangelogHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Changelog, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Changelog{}
	now := time.Now()
	tx := r.DB.db.Begin()

	err = resolvers.GetItem(ctx, tx, item, &id)
	if err != nil {
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Changelog",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, "changelogs.id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = r.EventController.SendEvent(ctx, &event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllChangelogs(ctx context.Context) (bool, error) {
	return r.Handlers.DeleteAllChangelogs(ctx, r.GeneratedResolver)
}
func DeleteAllChangelogsHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	err := r.DB.db.Delete(&Changelog{}).Error
	return err == nil, err
}
