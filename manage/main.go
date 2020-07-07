package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"manage/marketing"
	"manage/middle"
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
	if serverConf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建一个不包含中间件的路由器
	engine := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	engine.Use(gin.Logger())
	// 使用 Recovery 中间件
	engine.Use(gin.Recovery())
	// 配置跨域中间件
	engine.Use(middle.Cors())

	engine.PUT("/clue", marketing.ClueAPI{}.PutClue)
	engine.GET("/clue/xlsx", marketing.ClueAPI{}.GetClue)

	serverConf := new(ServerYaml)
	loadServerConf(serverConf)
	engine.Run(fmt.Sprintf(":%d", serverConf.Port))
}
