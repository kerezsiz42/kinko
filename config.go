package config

import (
	"github.com/kerezsiz42/kinko/internal/database"
	csv_database "github.com/kerezsiz42/kinko/internal/database/csv"
	"github.com/kerezsiz42/kinko/internal/utils"
)

var Db database.Database

func init() {
	csvPath := utils.GetEnv("CSV_PATH", nil)
	Db = csv_database.NewCsvDatabase(csvPath)
}
