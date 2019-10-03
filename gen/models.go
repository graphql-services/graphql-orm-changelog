package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

type NotFoundError struct {
	Entity string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Entity)
}

type ChangelogChangeResultType struct {
	EntityResultType
}

type ChangelogChange struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Column    string     `json:"column" gorm:"column:column"`
	OldValue  *string    `json:"oldValue" gorm:"column:oldValue"`
	NewValue  *string    `json:"newValue" gorm:"column:newValue"`
	LogID     *string    `json:"logId" gorm:"column:logId"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Log *Changelog `json:"log"`
}

func (m *ChangelogChange) Is_Entity() {}

type ChangelogChangeChanges struct {
	ID        string
	Column    string
	OldValue  *string
	NewValue  *string
	LogID     *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string
}

type ChangelogResultType struct {
	EntityResultType
}

type Changelog struct {
	ID          string        `json:"id" gorm:"column:id;primary_key"`
	Entity      string        `json:"entity" gorm:"column:entity"`
	EntityID    string        `json:"entityID" gorm:"column:entityID"`
	PrincipalID *string       `json:"principalID" gorm:"column:principalID"`
	Type        ChangelogType `json:"type" gorm:"column:type"`
	Date        time.Time     `json:"date" gorm:"column:date"`
	UpdatedAt   *time.Time    `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time     `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string       `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string       `json:"createdBy" gorm:"column:createdBy"`

	Changes []*ChangelogChange `json:"changes" gorm:"foreignkey:LogID"`
}

func (m *Changelog) Is_Entity() {}

type ChangelogChanges struct {
	ID          string
	Entity      string
	EntityID    string
	PrincipalID *string
	Type        ChangelogType
	Date        time.Time
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string

	ChangesIDs []*string
}

// used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
