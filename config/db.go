package config

import "os"

type Database struct {
	Driver     string
	DBName     string
	DBPassword string
	DBHost     string
	DBUser     string
	DBFileName string
}

func LoadDatabaseConfiguration() Database {
	return Database{
		Driver:     os.Getenv("DB_DRIVER"),
		DBFileName: os.Getenv("DB_FILENAME"),
		DBName:     "",
		DBPassword: "",
		DBHost:     "",
		DBUser:     "",
	}
}
