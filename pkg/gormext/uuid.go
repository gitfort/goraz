package gormext

import "github.com/jinzhu/gorm"

func EnableUUIDExt(db *gorm.DB) error {
	return db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
}
