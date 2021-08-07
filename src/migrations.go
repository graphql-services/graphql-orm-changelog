package src

import (
	"github.com/graphql-services/graphql-orm-changelog/gen"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// GetMigrations ...
func GetMigrations(db *gen.DB) []*gormigrate.Migration {
	return []*gormigrate.Migration{
		&gormigrate.Migration{
			ID: "0000_INIT",
			Migrate: func(tx *gorm.DB) error {
				return db.AutoMigrate()
			},
			Rollback: func(tx *gorm.DB) error {
				// there's not much we can do if initialization/automigration fails
				return nil
			},
		},
	}
}
