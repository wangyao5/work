package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"manage/filter"
	"manage/marketing"
	"manage/security"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

//ServerYaml 服务配置
type ServerYaml struct {
	Port    int  `yaml:"port"`
	Release bool `yaml:"release"`
}

func loadServerConf(serverConf *ServerYaml) {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	serverFile := filepath.Join(basePath, "conf", "server.yaml")
	serverYaml, _ := ioutil.ReadFile(serverFile)
	if err := yaml.Unmarshal(serverYaml, serverConf); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serverConf := new(ServerYaml)
	loadServerConf(serverConf)
	if serverConf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建一个不包含中间件的路由器
	router := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	router.Use(gin.Logger())
	// 使用 Recovery 中间件
	router.Use(gin.Recovery())
	// 配置跨域中间件
	router.Use(filter.Cors())
	public := router.Group("/public")
	{
		public.PUT("/marketing/clue", marketing.ClueAPI{}.PutClue)
		public.POST("/login", security.Security{}.Login)
	}

	security := router.Group("/security")
	security.Use(filter.Authorization())
	{
		security.GET("/marketing/clue/xlsx", marketing.ClueAPI{}.GetClue)
	}
	// router.POST("/security/valid", security.Security{}.RefreshToken)
	router.Run(fmt.Sprintf(":%d", serverConf.Port))
}
