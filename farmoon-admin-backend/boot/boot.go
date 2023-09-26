package boot

import (
	"context"
	"errors"
	"fmt"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Boot() {
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", global.FXConfig.System.Port),
		Handler: Router(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("【farmoon-admin】listen:%s\n", err)
		}
	}()

	go func() {
		RPC()
	}()

	logo(global.FXConfig.System.Port)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("【farmoon-admin】start to exit...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("【farmoon-admin】force exit：", err)
	}

	log.Println("【farmoon-admin】 exit complete！")
}

func logo(port int) {
	fmt.Println("Welcome to Farmoon-Admin !")
	fmt.Println("Github: https://github.com/kanyuanzhi/farmoon-admin ")
	fmt.Println("Expecting Your Star!")
	fmt.Printf("System started, listening port: %d...\n", port)
	global.FXLogger.Info(fmt.Sprintf("System started, listening port: %d", port))
}

func init() {
	utils.Reload("config", &global.FXConfig)
}
