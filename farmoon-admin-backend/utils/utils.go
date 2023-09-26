package utils

import (
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
)

// Reload 解析配置文件configName到target
func Reload(configName string, target interface{}) {
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("无法读取配置文件:", err)
		return
	}

	err = viper.Unmarshal(target)
	if err != nil {
		log.Println("解析配置文件失败:", err)
		return
	}
}

func LoadLocalImage(path string) ([]byte, error) {
	imagePath := path
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		global.FXLogger.Error(global.FXConfig.System.LoadLocalResourceErrorMessage, zap.Any("err", err))
		return nil, err
	}
	return imageData, nil
}
