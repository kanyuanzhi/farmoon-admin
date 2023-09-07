package main

import (
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/boot"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"log/slog"
)

func main() {
	global.FXDb = boot.Postgres()
	if global.FXDb == nil {
		err := boot.InitDb()
		if err != nil {
			slog.Error("初始化数据库失败")
			return
		}
	}

	boot.Boot()
}
