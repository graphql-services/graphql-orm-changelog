package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Apply ...
func (s ChangelogChangeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("changelog_changes"), sorts, joins)
}

// ApplyWithAlias ...
func (s ChangelogChangeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Column != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("column"), Direction: s.Column.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ColumnMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("column") + ")", Direction: s.ColumnMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ColumnMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("column") + ")", Direction: s.ColumnMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.OldValue != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("oldValue"), Direction: s.OldValue.String()}
		*sorts = append(*sorts, sort)
	}

	if s.OldValueMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("oldValue") + ")", Direction: s.OldValueMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.OldValueMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("oldValue") + ")", Direction: s.OldValueMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NewValue != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("newValue"), Direction: s.NewValue.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NewValueMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("newValue") + ")", Direction: s.NewValueMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NewValueMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("newValue") + ")", Direction: s.NewValueMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LogID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("logId"), Direction: s.LogID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.LogIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("logId") + ")", Direction: s.LogIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LogIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("logId") + ")", Direction: s.LogIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Log != nil {
		_alias := alias + "_log"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("changelogs"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+dialect.Quote(alias)+"."+dialect.Quote("logId"))
		err := s.Log.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply ...
func (s ChangelogSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("changelogs"), sorts, joins)
}

// ApplyWithAlias ...
func (s ChangelogSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Entity != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("entity"), Direction: s.Entity.String()}
		*sorts = append(*sorts, sort)
	}

	if s.EntityMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("entity") + ")", Direction: s.EntityMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.EntityMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("entity") + ")", Direction: s.EntityMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.EntityID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("entityID"), Direction: s.EntityID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.EntityIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("entityID") + ")", Direction: s.EntityIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.EntityIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("entityID") + ")", Direction: s.EntityIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PrincipalID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("principalID"), Direction: s.PrincipalID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.PrincipalIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("principalID") + ")", Direction: s.PrincipalIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.PrincipalIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("principalID") + ")", Direction: s.PrincipalIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Date != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("date"), Direction: s.Date.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DateMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("date") + ")", Direction: s.DateMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DateMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("date") + ")", Direction: s.DateMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Changes != nil {
		_alias := alias + "_changes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("changelog_changes"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("logId")+" = "+dialect.Quote(alias)+".id")
		err := s.Changes.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
