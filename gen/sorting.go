package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s ChangelogChangeSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("changelog_changes"), sorts, joins)
}
func (s ChangelogChangeSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+"id "+s.ID.String())
	}

	if s.Column != nil {
		*sorts = append(*sorts, aliasPrefix+"column "+s.Column.String())
	}

	if s.OldValue != nil {
		*sorts = append(*sorts, aliasPrefix+"oldValue "+s.OldValue.String())
	}

	if s.NewValue != nil {
		*sorts = append(*sorts, aliasPrefix+"newValue "+s.NewValue.String())
	}

	if s.LogID != nil {
		*sorts = append(*sorts, aliasPrefix+"logId "+s.LogID.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedAt "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"createdAt "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedBy "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"createdBy "+s.CreatedBy.String())
	}

	if s.Log != nil {
		_alias := alias + "_log"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("changelogs"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("logId"))
		err := s.Log.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s ChangelogSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]string, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("changelogs"), sorts, joins)
}
func (s ChangelogSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]string, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		*sorts = append(*sorts, aliasPrefix+"id "+s.ID.String())
	}

	if s.Entity != nil {
		*sorts = append(*sorts, aliasPrefix+"entity "+s.Entity.String())
	}

	if s.EntityID != nil {
		*sorts = append(*sorts, aliasPrefix+"entityID "+s.EntityID.String())
	}

	if s.PrincipalID != nil {
		*sorts = append(*sorts, aliasPrefix+"principalID "+s.PrincipalID.String())
	}

	if s.Date != nil {
		*sorts = append(*sorts, aliasPrefix+"date "+s.Date.String())
	}

	if s.UpdatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedAt "+s.UpdatedAt.String())
	}

	if s.CreatedAt != nil {
		*sorts = append(*sorts, aliasPrefix+"createdAt "+s.CreatedAt.String())
	}

	if s.UpdatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"updatedBy "+s.UpdatedBy.String())
	}

	if s.CreatedBy != nil {
		*sorts = append(*sorts, aliasPrefix+"createdBy "+s.CreatedBy.String())
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
