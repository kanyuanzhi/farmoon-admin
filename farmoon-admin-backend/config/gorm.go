package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormConfig() gorm.Config {
	return gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

func PostgresConfig(config Postgres) postgres.Config {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.Database, config.Port)
	return postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}
}
