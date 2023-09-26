package boot

import (
	"database/sql"
	"fmt"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/config"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func Postgres() *gorm.DB {
	fxPostgresConfig := global.FXConfig.Postgres
	postgresConfig := config.PostgresConfig(fxPostgresConfig)
	gormConfig := config.GormConfig()
	if db, err := gorm.Open(postgres.New(postgresConfig), &gormConfig); err != nil {
		global.FXLogger.Error("Connect postgres failed: " + err.Error())
		return nil
	} else {
		global.FXLogger.Info("Connect postgres success")
		sqlDB, err := db.DB()
		if err != nil {
			global.FXLogger.Error("Get postgres instance failed: " + err.Error())
			return nil
		}
		sqlDB.SetMaxOpenConns(10) // 最大打开连接数
		sqlDB.SetMaxIdleConns(5)  // 最大闲置连接数
		return db
	}
}

var migrateList = []interface{}{
	model.SysUser{},
	model.SysDish{},
	model.SysCuisine{},
	model.SysSeasoning{},
	model.SysIngredient{},
	model.SysIngredientType{},
	model.SysIngredientShape{},
	model.SysUserDeletedDish{},
}

func InitDb() error {
	if err := createDatabase(); err != nil {
		return err
	}
	global.FXDb = Postgres()
	err := global.FXDb.AutoMigrate(migrateList...)
	if err != nil {
		global.FXLogger.Error(err.Error())
		return err
	}

	return nil
}

func createDatabase() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		global.FXConfig.Postgres.Host, global.FXConfig.Postgres.User, global.FXConfig.Postgres.Password, global.FXConfig.Postgres.Port)
	createSql := fmt.Sprintf("CREATE DATABASE %s WITH ENCODING UTF8", global.FXConfig.Postgres.Database)
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			global.FXLogger.Error("Get postgres instance failed: " + err.Error())
		}
	}(sqlDB)
	if err = sqlDB.Ping(); err != nil {
		global.FXLogger.Error("Get postgres instance failed: " + err.Error())
		return err
	}
	if _, err = sqlDB.Exec(createSql); err != nil {
		global.FXLogger.Error("Get postgres instance failed: " + err.Error())
		return err
	}
	return err
}
