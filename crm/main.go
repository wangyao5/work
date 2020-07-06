package main

import (
	"crm/clue"
	"crm/middle"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

//ServerYaml 服务配置
type ServerYaml struct {
	Port int `yaml:"port"`
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
	// 创建一个不包含中间件的路由器
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	engine.Use(gin.Logger())
	// 使用 Recovery 中间件
	engine.Use(gin.Recovery())
	// 配置跨域中间件
	engine.Use(middle.Cors())

	engine.PUT("/clue", clue.PutClue)
	engine.GET("/clue/xlsx", clue.GetClue)

	serverConf := new(ServerYaml)
	loadServerConf(serverConf)
	engine.Run(fmt.Sprintf(":%d", serverConf.Port))
}
