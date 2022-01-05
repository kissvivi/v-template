package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"

	"v-template/config"
)

var InitServerSet = wire.NewSet(initServer)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runServer,
}

//type Application struct {
//	name string
//	version string
//	httpServer *http.Server
//}

func runServer(cmd *cobra.Command, args []string) {
	//初始化config
	//初始化application
	//启动数据库
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	//orm.DB = config.Gorm(&cfg.Database)

	//r := router.InitRouter()

	//s := initServer(cfg, r)

	s, _, _ := InitApp(cfg)

	fmt.Println(`
      ___         ___                                     ___                       ___                 
     /  /\       /  /\        ___                        /__/\        ___          /  /\          ___   
    /  /:/_     /  /:/_      /  /\                       \  \:\      /  /\        /  /::\        /  /\  
   /  /:/ /\   /  /:/ /\    /  /:/      ___     ___       \  \:\    /  /:/       /  /:/\:\      /  /:/  
  /  /:/ /:/  /  /:/ /:/_  /__/::\     /__/\   /  /\  ___  \  \:\  /__/::\      /  /:/  \:\    /  /:/   
 /__/:/ /:/  /__/:/ /:/ /\ \__\/\:\__  \  \:\ /  /:/ /__/\  \__\:\ \__\/\:\__  /__/:/ \__\:\  /  /::\   
 \  \:\/:/   \  \:\/:/ /:/    \  \:\/\  \  \:\  /:/  \  \:\ /  /:/    \  \:\/\ \  \:\ /  /:/ /__/:/\:\  
  \  \::/     \  \::/ /:/      \__\::/   \  \:\/:/    \  \:\  /:/      \__\::/  \  \:\  /:/  \__\/  \:\ 
   \  \:\      \  \:\/:/       /__/:/     \  \::/      \  \:\/:/       /__/:/    \  \:\/:/        \  \:\
    \  \:\      \  \::/        \__\/       \__\/        \  \::/        \__\/      \  \::/          \__\/
     \__\/       \__\/                                   \__\/                     \__\/`)
	fmt.Printf(`
	 	欢迎使用 飞鹿科技 feilu-iot-tempalte
	 	服务名 : %v
	 	服务版本 : %v 
	 	服务运行地址 : %v
	`, cfg.Application.Name, cfg.Application.Version, s.Addr)

	err = s.ListenAndServe()
	if err != nil {
		return
	}

}

func initServer(setting *config.Config, r *gin.Engine) *http.Server {
	s := &http.Server{
		Addr:           fmt.Sprintf("%v:%v", setting.Application.HttpServer.Host, setting.Application.HttpServer.Port),
		Handler:        r,
		ReadTimeout:    5 * time.Minute,
		WriteTimeout:   5 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

func Execute() {

	if err := serverCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
