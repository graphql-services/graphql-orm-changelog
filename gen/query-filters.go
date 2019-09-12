package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/ast"
)

type ChangelogChangeQueryFilter struct {
	Query *string
}

func (qf *ChangelogChangeQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("changelog_changes"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *ChangelogChangeQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["column"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("column")+" LIKE ? OR %[1]s"+dialect.Quote("column")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["oldValue"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("oldValue")+" LIKE ? OR %[1]s"+dialect.Quote("oldValue")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["newValue"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("newValue")+" LIKE ? OR %[1]s"+dialect.Quote("newValue")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if fs, ok := fieldsMap["log"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_log"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("changelogs")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("logId"))

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := ChangelogQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

type ChangelogQueryFilter struct {
	Query *string
}

func (qf *ChangelogQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("changelogs"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *ChangelogQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["entity"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("entity")+" LIKE ? OR %[1]s"+dialect.Quote("entity")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["entityID"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("entityID")+" LIKE ? OR %[1]s"+dialect.Quote("entityID")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if _, ok := fieldsMap["principalID"]; ok {
		*ors = append(*ors, fmt.Sprintf("%[1]s"+dialect.Quote("principalID")+" LIKE ? OR %[1]s"+dialect.Quote("principalID")+" LIKE ?", dialect.Quote(alias)+"."))
		*values = append(*values, fmt.Sprintf("%s%%", query), fmt.Sprintf("%% %s%%", query))
	}

	if fs, ok := fieldsMap["changes"]; ok {
		_fields := []*ast.Field{}
		_alias := alias + "_changes"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote("changelog_changes")+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("logId")+" = "+dialect.Quote(alias)+".id")

		for _, f := range fs {
			for _, s := range f.SelectionSet {
				if f, ok := s.(*ast.Field); ok {
					_fields = append(_fields, f)
				}
			}
		}
		q := ChangelogChangeQueryFilter{qf.Query}
		err := q.applyQueryWithFields(dialect, _fields, query, _alias, ors, values, joins)
		if err != nil {
			return err
		}
	}

	return nil
}
