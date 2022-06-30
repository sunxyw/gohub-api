package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type UserAuth struct {
		models.BaseModel

		UserID     uint64 `gorm:"type:int;not null;index"`
		Type       string `gorm:"type:varchar(255);not null;index"`
		Identifier string `gorm:"type:varchar(255);not null;index"`
		Credential string `gorm:"type:varchar(255);not null;index"`
		Verified   bool   `gorm:"type:tinyint(1);default:0"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&UserAuth{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&UserAuth{})
	}

	migrate.Add("2022_06_30_184700_add_user_auths_table", up, down)
}
