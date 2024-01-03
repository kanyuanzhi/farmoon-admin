package boot

import (
	"fmt"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/rpc/rpcServer"
	pb "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/rpc/rpcv1"
	"go.uber.org/zap"

	"google.golang.org/grpc"
	"net"
)

func RPC() {
	address := fmt.Sprintf("%s:%d", global.FXConfig.RPC.Host, global.FXConfig.RPC.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		global.FXLogger.Error(global.FXConfig.System.RPCErrorMessage,
			zap.Any("err", fmt.Sprintf("RPC监听端口d%d失败：%v", global.FXConfig.RPC.Port, err)))
	}

	maxSize := 100 * 1024 * 1024

	server := grpc.NewServer(grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize))
	pb.RegisterDataUpdaterServer(server, &rpcServer.DataUpdate{})
	global.FXLogger.Info("RPC服务启动")

	if err := server.Serve(listen); err != nil {
		global.FXLogger.Fatal(global.FXConfig.System.RPCErrorMessage,
			zap.Any("err", fmt.Sprintf("RPC服务启动失败%v", err)))
	}
}
