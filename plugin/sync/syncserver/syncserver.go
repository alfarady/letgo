package syncserver

import (
	"github.com/alfarady/letgo/file"
	"github.com/alfarady/letgo/lib"
	"github.com/alfarady/letgo/net/rpc"
	"github.com/alfarady/letgo/plugin/sync/api"
	"github.com/alfarady/letgo/plugin/sync/syncconfig"
	"fmt"
)
//SyncServer
type SyncServer struct{}

//Run
func (s *SyncServer)Run(values ...interface{})interface{}{
	rpc.NewServer().Register(new(api.FileSync)).Register(new(api.Command)).Run(config.IP,config.Port)
	return true
}

//sync server config
var config syncconfig.Server
//init
func init(){
	serverFile:="config/sync_server.config"
	cfgFile:=file.GetContent(serverFile)
	if cfgFile==""{
		serverConfig:=syncconfig.Server{
			IP: "127.0.0.1",
			Port: "5566",
		}
		file.PutContent(serverFile,fmt.Sprintf("%v",serverConfig))
		panic("please setting sync server config in config/sync_server.config file")
	}
	lib.StringToObject(cfgFile, &config)
}
//New
func New()*SyncServer{
	return &SyncServer{}
}