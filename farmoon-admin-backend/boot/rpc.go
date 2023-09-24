package boot

import (
	"fmt"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/rpc/rpcServer"
	pb "github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/rpc/rpcv1"

	"google.golang.org/grpc"
	"log"
	"net"
)

func RPC() {
	address := fmt.Sprintf("%s:%d", global.FXConfig.RPC.Host, global.FXConfig.RPC.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("无法监听端口: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterDataUpdaterServer(server, &rpcServer.DataUpdate{})
	log.Println("RPC服务启动")

	if err := server.Serve(listen); err != nil {
		log.Fatalf("无法启动RPC服务: %v", err)
	}
}
