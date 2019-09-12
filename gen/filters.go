package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *ChangelogChangeFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ChangelogChangeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("changelog_changes"), wheres, values, joins)
}
func (f *ChangelogChangeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Log != nil {
		_alias := alias + "_log"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("changelogs")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("logId"))
		err := f.Log.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ChangelogChangeFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Column != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" = ?")
		values = append(values, f.Column)
	}
	if f.ColumnNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" != ?")
		values = append(values, f.ColumnNe)
	}
	if f.ColumnGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" > ?")
		values = append(values, f.ColumnGt)
	}
	if f.ColumnLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" < ?")
		values = append(values, f.ColumnLt)
	}
	if f.ColumnGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" >= ?")
		values = append(values, f.ColumnGte)
	}
	if f.ColumnLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" <= ?")
		values = append(values, f.ColumnLte)
	}
	if f.ColumnIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" IN (?)")
		values = append(values, f.ColumnIn)
	}
	if f.ColumnLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ColumnLike, "?", "_", -1), "*", "%", -1))
	}
	if f.ColumnPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ColumnPrefix))
	}
	if f.ColumnSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ColumnSuffix))
	}

	if f.OldValue != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" = ?")
		values = append(values, f.OldValue)
	}
	if f.OldValueNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" != ?")
		values = append(values, f.OldValueNe)
	}
	if f.OldValueGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" > ?")
		values = append(values, f.OldValueGt)
	}
	if f.OldValueLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" < ?")
		values = append(values, f.OldValueLt)
	}
	if f.OldValueGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" >= ?")
		values = append(values, f.OldValueGte)
	}
	if f.OldValueLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" <= ?")
		values = append(values, f.OldValueLte)
	}
	if f.OldValueIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" IN (?)")
		values = append(values, f.OldValueIn)
	}
	if f.OldValueLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.OldValueLike, "?", "_", -1), "*", "%", -1))
	}
	if f.OldValuePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.OldValuePrefix))
	}
	if f.OldValueSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.OldValueSuffix))
	}

	if f.NewValue != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" = ?")
		values = append(values, f.NewValue)
	}
	if f.NewValueNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" != ?")
		values = append(values, f.NewValueNe)
	}
	if f.NewValueGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" > ?")
		values = append(values, f.NewValueGt)
	}
	if f.NewValueLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" < ?")
		values = append(values, f.NewValueLt)
	}
	if f.NewValueGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" >= ?")
		values = append(values, f.NewValueGte)
	}
	if f.NewValueLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" <= ?")
		values = append(values, f.NewValueLte)
	}
	if f.NewValueIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" IN (?)")
		values = append(values, f.NewValueIn)
	}
	if f.NewValueLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NewValueLike, "?", "_", -1), "*", "%", -1))
	}
	if f.NewValuePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NewValuePrefix))
	}
	if f.NewValueSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NewValueSuffix))
	}

	if f.LogID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" = ?")
		values = append(values, f.LogID)
	}
	if f.LogIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" != ?")
		values = append(values, f.LogIDNe)
	}
	if f.LogIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" > ?")
		values = append(values, f.LogIDGt)
	}
	if f.LogIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" < ?")
		values = append(values, f.LogIDLt)
	}
	if f.LogIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" >= ?")
		values = append(values, f.LogIDGte)
	}
	if f.LogIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" <= ?")
		values = append(values, f.LogIDLte)
	}
	if f.LogIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" IN (?)")
		values = append(values, f.LogIDIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ChangelogChangeFilterType) AndWith(f2 ...*ChangelogChangeFilterType) *ChangelogChangeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ChangelogChangeFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ChangelogChangeFilterType) OrWith(f2 ...*ChangelogChangeFilterType) *ChangelogChangeFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ChangelogChangeFilterType{
		Or: append(_f2, f),
	}
}

func (f *ChangelogFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	values := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &values, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0
}
func (f *ChangelogFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, values *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("changelogs"), wheres, values, joins)
}
func (f *ChangelogFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _values := f.WhereContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*values = append(*values, _values...)

	if f.Or != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			err := or.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, "("+strings.Join(cs, " OR ")+")")
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		cs := []string{}
		vs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &cs, &vs, &js)
			if err != nil {
				return err
			}
		}
		if len(cs) > 0 {
			*wheres = append(*wheres, strings.Join(cs, " AND "))
		}
		*values = append(*values, vs...)
		*joins = append(*joins, js...)
	}

	if f.Changes != nil {
		_alias := alias + "_changes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("changelog_changes")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("logId")+" = "+dialect.Quote(alias)+".id")
		err := f.Changes.ApplyWithAlias(ctx, dialect, _alias, wheres, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *ChangelogFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}
	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}
	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}
	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}
	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}
	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}
	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.Entity != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" = ?")
		values = append(values, f.Entity)
	}
	if f.EntityNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" != ?")
		values = append(values, f.EntityNe)
	}
	if f.EntityGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" > ?")
		values = append(values, f.EntityGt)
	}
	if f.EntityLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" < ?")
		values = append(values, f.EntityLt)
	}
	if f.EntityGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" >= ?")
		values = append(values, f.EntityGte)
	}
	if f.EntityLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" <= ?")
		values = append(values, f.EntityLte)
	}
	if f.EntityIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" IN (?)")
		values = append(values, f.EntityIn)
	}
	if f.EntityLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EntityLike, "?", "_", -1), "*", "%", -1))
	}
	if f.EntityPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EntityPrefix))
	}
	if f.EntitySuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EntitySuffix))
	}

	if f.EntityID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" = ?")
		values = append(values, f.EntityID)
	}
	if f.EntityIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" != ?")
		values = append(values, f.EntityIDNe)
	}
	if f.EntityIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" > ?")
		values = append(values, f.EntityIDGt)
	}
	if f.EntityIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" < ?")
		values = append(values, f.EntityIDLt)
	}
	if f.EntityIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" >= ?")
		values = append(values, f.EntityIDGte)
	}
	if f.EntityIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" <= ?")
		values = append(values, f.EntityIDLte)
	}
	if f.EntityIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" IN (?)")
		values = append(values, f.EntityIDIn)
	}
	if f.EntityIDLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EntityIDLike, "?", "_", -1), "*", "%", -1))
	}
	if f.EntityIDPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EntityIDPrefix))
	}
	if f.EntityIDSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EntityIDSuffix))
	}

	if f.PrincipalID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" = ?")
		values = append(values, f.PrincipalID)
	}
	if f.PrincipalIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" != ?")
		values = append(values, f.PrincipalIDNe)
	}
	if f.PrincipalIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" > ?")
		values = append(values, f.PrincipalIDGt)
	}
	if f.PrincipalIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" < ?")
		values = append(values, f.PrincipalIDLt)
	}
	if f.PrincipalIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" >= ?")
		values = append(values, f.PrincipalIDGte)
	}
	if f.PrincipalIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" <= ?")
		values = append(values, f.PrincipalIDLte)
	}
	if f.PrincipalIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" IN (?)")
		values = append(values, f.PrincipalIDIn)
	}
	if f.PrincipalIDLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PrincipalIDLike, "?", "_", -1), "*", "%", -1))
	}
	if f.PrincipalIDPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PrincipalIDPrefix))
	}
	if f.PrincipalIDSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PrincipalIDSuffix))
	}

	if f.Type != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" = ?")
		values = append(values, f.Type)
	}
	if f.TypeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" != ?")
		values = append(values, f.TypeNe)
	}
	if f.TypeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" > ?")
		values = append(values, f.TypeGt)
	}
	if f.TypeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" < ?")
		values = append(values, f.TypeLt)
	}
	if f.TypeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" >= ?")
		values = append(values, f.TypeGte)
	}
	if f.TypeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" <= ?")
		values = append(values, f.TypeLte)
	}
	if f.TypeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IN (?)")
		values = append(values, f.TypeIn)
	}

	if f.Date != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" = ?")
		values = append(values, f.Date)
	}
	if f.DateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" != ?")
		values = append(values, f.DateNe)
	}
	if f.DateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" > ?")
		values = append(values, f.DateGt)
	}
	if f.DateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" < ?")
		values = append(values, f.DateLt)
	}
	if f.DateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" >= ?")
		values = append(values, f.DateGte)
	}
	if f.DateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" <= ?")
		values = append(values, f.DateLte)
	}
	if f.DateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IN (?)")
		values = append(values, f.DateIn)
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}
	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}
	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}
	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}
	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}
	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}
	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}
	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}
	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}
	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}
	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}
	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}
	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}
	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}
	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}
	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}
	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}
	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}
	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}
	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}
	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}
	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}
	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}
	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}
	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *ChangelogFilterType) AndWith(f2 ...*ChangelogFilterType) *ChangelogFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ChangelogFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *ChangelogFilterType) OrWith(f2 ...*ChangelogFilterType) *ChangelogFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &ChangelogFilterType{
		Or: append(_f2, f),
	}
}
