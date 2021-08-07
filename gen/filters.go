package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// IsEmpty ...
func (f *ChangelogChangeFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply ...
func (f *ChangelogChangeFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("changelog_changes"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias ...
func (f *ChangelogChangeFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Log != nil {
		_alias := alias + "_log"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("changelogs"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+dialect.Quote(alias)+"."+dialect.Quote("logId"))
		err := f.Log.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
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

	if f.ColumnNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" NOT IN (?)")
		values = append(values, f.ColumnNotIn)
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

	if f.ColumnNull != nil {
		if *f.ColumnNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("column")+" IS NOT NULL")
		}
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

	if f.OldValueNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" NOT IN (?)")
		values = append(values, f.OldValueNotIn)
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

	if f.OldValueNull != nil {
		if *f.OldValueNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("oldValue")+" IS NOT NULL")
		}
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

	if f.NewValueNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" NOT IN (?)")
		values = append(values, f.NewValueNotIn)
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

	if f.NewValueNull != nil {
		if *f.NewValueNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("newValue")+" IS NOT NULL")
		}
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

	if f.LogIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" NOT IN (?)")
		values = append(values, f.LogIDNotIn)
	}

	if f.LogIDNull != nil {
		if *f.LogIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("logId")+" IS NOT NULL")
		}
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent ...
func (f *ChangelogChangeFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.ColumnMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") = ?")
		values = append(values, f.ColumnMin)
	}

	if f.ColumnMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") = ?")
		values = append(values, f.ColumnMax)
	}

	if f.ColumnMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") != ?")
		values = append(values, f.ColumnMinNe)
	}

	if f.ColumnMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") != ?")
		values = append(values, f.ColumnMaxNe)
	}

	if f.ColumnMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") > ?")
		values = append(values, f.ColumnMinGt)
	}

	if f.ColumnMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") > ?")
		values = append(values, f.ColumnMaxGt)
	}

	if f.ColumnMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") < ?")
		values = append(values, f.ColumnMinLt)
	}

	if f.ColumnMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") < ?")
		values = append(values, f.ColumnMaxLt)
	}

	if f.ColumnMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") >= ?")
		values = append(values, f.ColumnMinGte)
	}

	if f.ColumnMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") >= ?")
		values = append(values, f.ColumnMaxGte)
	}

	if f.ColumnMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") <= ?")
		values = append(values, f.ColumnMinLte)
	}

	if f.ColumnMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") <= ?")
		values = append(values, f.ColumnMaxLte)
	}

	if f.ColumnMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") IN (?)")
		values = append(values, f.ColumnMinIn)
	}

	if f.ColumnMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") IN (?)")
		values = append(values, f.ColumnMaxIn)
	}

	if f.ColumnMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") NOT IN (?)")
		values = append(values, f.ColumnMinNotIn)
	}

	if f.ColumnMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") NOT IN (?)")
		values = append(values, f.ColumnMaxNotIn)
	}

	if f.ColumnMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ColumnMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ColumnMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ColumnMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ColumnMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ColumnMinPrefix))
	}

	if f.ColumnMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ColumnMaxPrefix))
	}

	if f.ColumnMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("column")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ColumnMinSuffix))
	}

	if f.ColumnMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("column")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ColumnMaxSuffix))
	}

	if f.OldValueMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") = ?")
		values = append(values, f.OldValueMin)
	}

	if f.OldValueMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") = ?")
		values = append(values, f.OldValueMax)
	}

	if f.OldValueMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") != ?")
		values = append(values, f.OldValueMinNe)
	}

	if f.OldValueMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") != ?")
		values = append(values, f.OldValueMaxNe)
	}

	if f.OldValueMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") > ?")
		values = append(values, f.OldValueMinGt)
	}

	if f.OldValueMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") > ?")
		values = append(values, f.OldValueMaxGt)
	}

	if f.OldValueMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") < ?")
		values = append(values, f.OldValueMinLt)
	}

	if f.OldValueMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") < ?")
		values = append(values, f.OldValueMaxLt)
	}

	if f.OldValueMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") >= ?")
		values = append(values, f.OldValueMinGte)
	}

	if f.OldValueMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") >= ?")
		values = append(values, f.OldValueMaxGte)
	}

	if f.OldValueMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") <= ?")
		values = append(values, f.OldValueMinLte)
	}

	if f.OldValueMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") <= ?")
		values = append(values, f.OldValueMaxLte)
	}

	if f.OldValueMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") IN (?)")
		values = append(values, f.OldValueMinIn)
	}

	if f.OldValueMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") IN (?)")
		values = append(values, f.OldValueMaxIn)
	}

	if f.OldValueMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") NOT IN (?)")
		values = append(values, f.OldValueMinNotIn)
	}

	if f.OldValueMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") NOT IN (?)")
		values = append(values, f.OldValueMaxNotIn)
	}

	if f.OldValueMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.OldValueMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.OldValueMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.OldValueMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.OldValueMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.OldValueMinPrefix))
	}

	if f.OldValueMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.OldValueMaxPrefix))
	}

	if f.OldValueMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("oldValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.OldValueMinSuffix))
	}

	if f.OldValueMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("oldValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.OldValueMaxSuffix))
	}

	if f.NewValueMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") = ?")
		values = append(values, f.NewValueMin)
	}

	if f.NewValueMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") = ?")
		values = append(values, f.NewValueMax)
	}

	if f.NewValueMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") != ?")
		values = append(values, f.NewValueMinNe)
	}

	if f.NewValueMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") != ?")
		values = append(values, f.NewValueMaxNe)
	}

	if f.NewValueMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") > ?")
		values = append(values, f.NewValueMinGt)
	}

	if f.NewValueMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") > ?")
		values = append(values, f.NewValueMaxGt)
	}

	if f.NewValueMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") < ?")
		values = append(values, f.NewValueMinLt)
	}

	if f.NewValueMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") < ?")
		values = append(values, f.NewValueMaxLt)
	}

	if f.NewValueMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") >= ?")
		values = append(values, f.NewValueMinGte)
	}

	if f.NewValueMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") >= ?")
		values = append(values, f.NewValueMaxGte)
	}

	if f.NewValueMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") <= ?")
		values = append(values, f.NewValueMinLte)
	}

	if f.NewValueMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") <= ?")
		values = append(values, f.NewValueMaxLte)
	}

	if f.NewValueMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") IN (?)")
		values = append(values, f.NewValueMinIn)
	}

	if f.NewValueMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") IN (?)")
		values = append(values, f.NewValueMaxIn)
	}

	if f.NewValueMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") NOT IN (?)")
		values = append(values, f.NewValueMinNotIn)
	}

	if f.NewValueMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") NOT IN (?)")
		values = append(values, f.NewValueMaxNotIn)
	}

	if f.NewValueMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NewValueMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NewValueMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NewValueMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NewValueMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NewValueMinPrefix))
	}

	if f.NewValueMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NewValueMaxPrefix))
	}

	if f.NewValueMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("newValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NewValueMinSuffix))
	}

	if f.NewValueMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("newValue")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NewValueMaxSuffix))
	}

	if f.LogIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("logId")+") = ?")
		values = append(values, f.LogIDMin)
	}

	if f.LogIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("logId")+") = ?")
		values = append(values, f.LogIDMax)
	}

	if f.LogIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("logId")+") != ?")
		values = append(values, f.LogIDMinNe)
	}

	if f.LogIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("logId")+") != ?")
		values = append(values, f.LogIDMaxNe)
	}

	if f.LogIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("logId")+") > ?")
		values = append(values, f.LogIDMinGt)
	}

	if f.LogIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("logId")+") > ?")
		values = append(values, f.LogIDMaxGt)
	}

	if f.LogIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("logId")+") < ?")
		values = append(values, f.LogIDMinLt)
	}

	if f.LogIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("logId")+") < ?")
		values = append(values, f.LogIDMaxLt)
	}

	if f.LogIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("logId")+") >= ?")
		values = append(values, f.LogIDMinGte)
	}

	if f.LogIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("logId")+") >= ?")
		values = append(values, f.LogIDMaxGte)
	}

	if f.LogIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("logId")+") <= ?")
		values = append(values, f.LogIDMinLte)
	}

	if f.LogIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("logId")+") <= ?")
		values = append(values, f.LogIDMaxLte)
	}

	if f.LogIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("logId")+") IN (?)")
		values = append(values, f.LogIDMinIn)
	}

	if f.LogIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("logId")+") IN (?)")
		values = append(values, f.LogIDMaxIn)
	}

	if f.LogIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("logId")+") NOT IN (?)")
		values = append(values, f.LogIDMinNotIn)
	}

	if f.LogIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("logId")+") NOT IN (?)")
		values = append(values, f.LogIDMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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

// IsEmpty ...
func (f *ChangelogFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply ...
func (f *ChangelogFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("changelogs"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias ...
func (f *ChangelogFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Changes != nil {
		_alias := alias + "_changes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("changelog_changes"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("logId")+" = "+dialect.Quote(alias)+".id")
		err := f.Changes.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
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

	if f.IDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" NOT IN (?)")
		values = append(values, f.IDNotIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
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

	if f.EntityNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" NOT IN (?)")
		values = append(values, f.EntityNotIn)
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

	if f.EntityNull != nil {
		if *f.EntityNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("entity")+" IS NOT NULL")
		}
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

	if f.EntityIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" NOT IN (?)")
		values = append(values, f.EntityIDNotIn)
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

	if f.EntityIDNull != nil {
		if *f.EntityIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("entityID")+" IS NOT NULL")
		}
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

	if f.PrincipalIDNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" NOT IN (?)")
		values = append(values, f.PrincipalIDNotIn)
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

	if f.PrincipalIDNull != nil {
		if *f.PrincipalIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("principalID")+" IS NOT NULL")
		}
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

	if f.TypeNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" NOT IN (?)")
		values = append(values, f.TypeNotIn)
	}

	if f.TypeNull != nil {
		if *f.TypeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("type")+" IS NOT NULL")
		}
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

	if f.DateNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" NOT IN (?)")
		values = append(values, f.DateNotIn)
	}

	if f.DateNull != nil {
		if *f.DateNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IS NOT NULL")
		}
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

	if f.UpdatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" NOT IN (?)")
		values = append(values, f.UpdatedAtNotIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
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

	if f.CreatedAtNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" NOT IN (?)")
		values = append(values, f.CreatedAtNotIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
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

	if f.UpdatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" NOT IN (?)")
		values = append(values, f.UpdatedByNotIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
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

	if f.CreatedByNotIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" NOT IN (?)")
		values = append(values, f.CreatedByNotIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent ...
func (f *ChangelogFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.IDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMinNotIn)
	}

	if f.IDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") NOT IN (?)")
		values = append(values, f.IDMaxNotIn)
	}

	if f.EntityMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") = ?")
		values = append(values, f.EntityMin)
	}

	if f.EntityMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") = ?")
		values = append(values, f.EntityMax)
	}

	if f.EntityMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") != ?")
		values = append(values, f.EntityMinNe)
	}

	if f.EntityMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") != ?")
		values = append(values, f.EntityMaxNe)
	}

	if f.EntityMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") > ?")
		values = append(values, f.EntityMinGt)
	}

	if f.EntityMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") > ?")
		values = append(values, f.EntityMaxGt)
	}

	if f.EntityMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") < ?")
		values = append(values, f.EntityMinLt)
	}

	if f.EntityMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") < ?")
		values = append(values, f.EntityMaxLt)
	}

	if f.EntityMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") >= ?")
		values = append(values, f.EntityMinGte)
	}

	if f.EntityMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") >= ?")
		values = append(values, f.EntityMaxGte)
	}

	if f.EntityMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") <= ?")
		values = append(values, f.EntityMinLte)
	}

	if f.EntityMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") <= ?")
		values = append(values, f.EntityMaxLte)
	}

	if f.EntityMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") IN (?)")
		values = append(values, f.EntityMinIn)
	}

	if f.EntityMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") IN (?)")
		values = append(values, f.EntityMaxIn)
	}

	if f.EntityMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") NOT IN (?)")
		values = append(values, f.EntityMinNotIn)
	}

	if f.EntityMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") NOT IN (?)")
		values = append(values, f.EntityMaxNotIn)
	}

	if f.EntityMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EntityMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EntityMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EntityMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EntityMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EntityMinPrefix))
	}

	if f.EntityMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EntityMaxPrefix))
	}

	if f.EntityMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entity")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EntityMinSuffix))
	}

	if f.EntityMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entity")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EntityMaxSuffix))
	}

	if f.EntityIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") = ?")
		values = append(values, f.EntityIDMin)
	}

	if f.EntityIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") = ?")
		values = append(values, f.EntityIDMax)
	}

	if f.EntityIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") != ?")
		values = append(values, f.EntityIDMinNe)
	}

	if f.EntityIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") != ?")
		values = append(values, f.EntityIDMaxNe)
	}

	if f.EntityIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") > ?")
		values = append(values, f.EntityIDMinGt)
	}

	if f.EntityIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") > ?")
		values = append(values, f.EntityIDMaxGt)
	}

	if f.EntityIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") < ?")
		values = append(values, f.EntityIDMinLt)
	}

	if f.EntityIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") < ?")
		values = append(values, f.EntityIDMaxLt)
	}

	if f.EntityIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") >= ?")
		values = append(values, f.EntityIDMinGte)
	}

	if f.EntityIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") >= ?")
		values = append(values, f.EntityIDMaxGte)
	}

	if f.EntityIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") <= ?")
		values = append(values, f.EntityIDMinLte)
	}

	if f.EntityIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") <= ?")
		values = append(values, f.EntityIDMaxLte)
	}

	if f.EntityIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") IN (?)")
		values = append(values, f.EntityIDMinIn)
	}

	if f.EntityIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") IN (?)")
		values = append(values, f.EntityIDMaxIn)
	}

	if f.EntityIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") NOT IN (?)")
		values = append(values, f.EntityIDMinNotIn)
	}

	if f.EntityIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") NOT IN (?)")
		values = append(values, f.EntityIDMaxNotIn)
	}

	if f.EntityIDMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EntityIDMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EntityIDMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EntityIDMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EntityIDMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EntityIDMinPrefix))
	}

	if f.EntityIDMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EntityIDMaxPrefix))
	}

	if f.EntityIDMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("entityID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EntityIDMinSuffix))
	}

	if f.EntityIDMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("entityID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EntityIDMaxSuffix))
	}

	if f.PrincipalIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") = ?")
		values = append(values, f.PrincipalIDMin)
	}

	if f.PrincipalIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") = ?")
		values = append(values, f.PrincipalIDMax)
	}

	if f.PrincipalIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") != ?")
		values = append(values, f.PrincipalIDMinNe)
	}

	if f.PrincipalIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") != ?")
		values = append(values, f.PrincipalIDMaxNe)
	}

	if f.PrincipalIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") > ?")
		values = append(values, f.PrincipalIDMinGt)
	}

	if f.PrincipalIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") > ?")
		values = append(values, f.PrincipalIDMaxGt)
	}

	if f.PrincipalIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") < ?")
		values = append(values, f.PrincipalIDMinLt)
	}

	if f.PrincipalIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") < ?")
		values = append(values, f.PrincipalIDMaxLt)
	}

	if f.PrincipalIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") >= ?")
		values = append(values, f.PrincipalIDMinGte)
	}

	if f.PrincipalIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") >= ?")
		values = append(values, f.PrincipalIDMaxGte)
	}

	if f.PrincipalIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") <= ?")
		values = append(values, f.PrincipalIDMinLte)
	}

	if f.PrincipalIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") <= ?")
		values = append(values, f.PrincipalIDMaxLte)
	}

	if f.PrincipalIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") IN (?)")
		values = append(values, f.PrincipalIDMinIn)
	}

	if f.PrincipalIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") IN (?)")
		values = append(values, f.PrincipalIDMaxIn)
	}

	if f.PrincipalIDMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") NOT IN (?)")
		values = append(values, f.PrincipalIDMinNotIn)
	}

	if f.PrincipalIDMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") NOT IN (?)")
		values = append(values, f.PrincipalIDMaxNotIn)
	}

	if f.PrincipalIDMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PrincipalIDMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PrincipalIDMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.PrincipalIDMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.PrincipalIDMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PrincipalIDMinPrefix))
	}

	if f.PrincipalIDMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.PrincipalIDMaxPrefix))
	}

	if f.PrincipalIDMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("principalID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PrincipalIDMinSuffix))
	}

	if f.PrincipalIDMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("principalID")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.PrincipalIDMaxSuffix))
	}

	if f.TypeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") = ?")
		values = append(values, f.TypeMin)
	}

	if f.TypeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") = ?")
		values = append(values, f.TypeMax)
	}

	if f.TypeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") != ?")
		values = append(values, f.TypeMinNe)
	}

	if f.TypeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") != ?")
		values = append(values, f.TypeMaxNe)
	}

	if f.TypeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") > ?")
		values = append(values, f.TypeMinGt)
	}

	if f.TypeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") > ?")
		values = append(values, f.TypeMaxGt)
	}

	if f.TypeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") < ?")
		values = append(values, f.TypeMinLt)
	}

	if f.TypeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") < ?")
		values = append(values, f.TypeMaxLt)
	}

	if f.TypeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") >= ?")
		values = append(values, f.TypeMinGte)
	}

	if f.TypeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") >= ?")
		values = append(values, f.TypeMaxGte)
	}

	if f.TypeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") <= ?")
		values = append(values, f.TypeMinLte)
	}

	if f.TypeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") <= ?")
		values = append(values, f.TypeMaxLte)
	}

	if f.TypeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") IN (?)")
		values = append(values, f.TypeMinIn)
	}

	if f.TypeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") IN (?)")
		values = append(values, f.TypeMaxIn)
	}

	if f.TypeMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("type")+") NOT IN (?)")
		values = append(values, f.TypeMinNotIn)
	}

	if f.TypeMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("type")+") NOT IN (?)")
		values = append(values, f.TypeMaxNotIn)
	}

	if f.DateMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") = ?")
		values = append(values, f.DateMin)
	}

	if f.DateMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") = ?")
		values = append(values, f.DateMax)
	}

	if f.DateMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") != ?")
		values = append(values, f.DateMinNe)
	}

	if f.DateMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") != ?")
		values = append(values, f.DateMaxNe)
	}

	if f.DateMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") > ?")
		values = append(values, f.DateMinGt)
	}

	if f.DateMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") > ?")
		values = append(values, f.DateMaxGt)
	}

	if f.DateMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") < ?")
		values = append(values, f.DateMinLt)
	}

	if f.DateMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") < ?")
		values = append(values, f.DateMaxLt)
	}

	if f.DateMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") >= ?")
		values = append(values, f.DateMinGte)
	}

	if f.DateMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") >= ?")
		values = append(values, f.DateMaxGte)
	}

	if f.DateMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") <= ?")
		values = append(values, f.DateMinLte)
	}

	if f.DateMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") <= ?")
		values = append(values, f.DateMaxLte)
	}

	if f.DateMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") IN (?)")
		values = append(values, f.DateMinIn)
	}

	if f.DateMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") IN (?)")
		values = append(values, f.DateMaxIn)
	}

	if f.DateMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") NOT IN (?)")
		values = append(values, f.DateMinNotIn)
	}

	if f.DateMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") NOT IN (?)")
		values = append(values, f.DateMaxNotIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.UpdatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMinNotIn)
	}

	if f.UpdatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") NOT IN (?)")
		values = append(values, f.UpdatedAtMaxNotIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.CreatedAtMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMinNotIn)
	}

	if f.CreatedAtMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") NOT IN (?)")
		values = append(values, f.CreatedAtMaxNotIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.UpdatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMinNotIn)
	}

	if f.UpdatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") NOT IN (?)")
		values = append(values, f.UpdatedByMaxNotIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	if f.CreatedByMinNotIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMinNotIn)
	}

	if f.CreatedByMaxNotIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") NOT IN (?)")
		values = append(values, f.CreatedByMaxNotIn)
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
