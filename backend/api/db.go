package api

import(
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
        db *gorm.DB
)

func InitDb(dsi string) error {
        var err error
        db, err = gorm.Open(postgres.Open(dsi))

        return err
}

