package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB puntero global de conexión a base de datos
var DB *gorm.DB

//DBConfig estructura de configuración
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

//BuildDBConfig tiene la configuración de conexión
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     "root",
		DBName:   "books",
		Password: "root",
	}
	return &dbConfig
}

//DbURL construye la ruta de conexión
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
