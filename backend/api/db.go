package api

import(
        "log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
        db *gorm.DB
)

func InitDb(dsi string) error {
        var err error
        db, err = gorm.Open(postgres.Open(dsi))

        if err != nil {
                log.Fatalf("Failed to connect to DB: %v", err)
                return err
        }

        return nil
}

