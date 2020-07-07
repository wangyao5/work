package middle

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

//MongoDb mongo db.
var MongoDb *mongo.Database

//MongoConf mongo config info.
var MongoConf *MongoYaml

//MongoYaml decode mongo config.
type MongoYaml struct {
	ApplyURI       string `yaml:"applyURI"`
	DataBase       string `yaml:"database"`
	ClueCollection string `yaml:"clueCollection"`
}

func init() {
	MongoConf = new(MongoYaml)
	loadMongoConf(MongoConf)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoConf.ApplyURI))
	if err != nil {
		log.Fatal(err)
	}

	MongoDb = client.Database(MongoConf.DataBase)
}

func loadMongoConf(mongoYaml *MongoYaml) {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("base path error")
	}
	fileName := filepath.Join(basePath, "conf", "mongo.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(yamlFile, MongoConf)
}
