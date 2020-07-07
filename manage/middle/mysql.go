package middle

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
)

func init() {
	mysqlYaml := new(MysqlYaml)
	loadMysqlConfig(mysqlYaml)
	initMysqlDB(mysqlYaml)
}

//GormDb gorm mysql db.
var GormDb *gorm.DB

//MysqlYaml mysql yaml.
type MysqlYaml struct {
	URL          string `yaml:"mysql-url"`
	MaxIdleConns int    `yaml:"max-idle-conns"`
	MaxOpenConns int    `yaml:"max-open-conns"`
	LogAble      bool   `yaml:"logable"`
}

func loadMysqlConfig(mysqlYaml *MysqlYaml) {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("base path error")
	}
	fileName := filepath.Join(basePath, "conf", "mysql.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err := yaml.Unmarshal(yamlFile, mysqlYaml); err != nil {
		log.Fatal(err)
	}
}

func initMysqlDB(mysqlYaml *MysqlYaml) {
	if gormDb, err := gorm.Open("mysql", mysqlYaml.URL); err == nil {
		GormDb = gormDb
		GormDb.LogMode(mysqlYaml.LogAble)
		GormDb.DB().SetMaxIdleConns(mysqlYaml.MaxIdleConns)
		GormDb.DB().SetMaxOpenConns(mysqlYaml.MaxOpenConns)
		// 全局禁用表名复数
		GormDb.SingularTable(true)
	} else {
		log.Fatal(err)
	}
}

//CreateTable create table.
func CreateTable(table interface{}) {
	db := GormDb
	if db.HasTable(table) {
		db.AutoMigrate(table)
	} else {
		db.CreateTable(table)
	}
}

//DropTable drop table.
func DropTable(table interface{}) {
	db := GormDb
	db.DropTableIfExists(table).Commit()
}
