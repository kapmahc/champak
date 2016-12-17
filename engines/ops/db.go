package ops

import (
	"path"

	"github.com/spf13/viper"
)

func dbMigrationsDir() string {
	return path.Join("db", viper.GetString("database.driver"), "migrations")
}
