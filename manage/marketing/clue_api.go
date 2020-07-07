package marketing

import (
	"context"
	"fmt"
	"io"
	"manage/middle"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"gopkg.in/mgo.v2/bson"
)

//Extend 线索扩展数据.
type Extend struct {
	Key   string `json:"key" bson:"key"`     //填充名
	Value string `json:"value" bson:"value"` //填充值
}

//Clue 线索.
type Clue struct {
	AdID       string   `json:"ad_id" bson:"ad_id"`         //广告素材ID
	AdName     string   `json:"ad_name" bson:"ad_name"`     //广告素材名
	ClueType   string   `json:"clue_type" bson:"clue_type"` //广告线索类型
	ClueInfo   string   `json:"clue_info" bson:"clue_info"` //广告线索信息
	Extends    []Extend `json:"extends" bson:"extends"`     //扩展填写内容
	CreateTime string   `json:"_" bson:"create_time"`       //创建时间
}

//ClueAPI 线索收集API.
type ClueAPI struct {
}

//PutClue 新增线索.
func (ClueAPI) PutClue(ctx *gin.Context) {
	clue := new(Clue)
	ctx.BindJSON(clue)
	clue.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	db := middle.MongoDb
	result, err := db.Collection(middle.MongoConf.ClueCollection).InsertOne(
		context.Background(),
		clue)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": result.InsertedID})
}

//GetClue 下载线索
func (ClueAPI) GetClue(ctx *gin.Context) {
	adID := ctx.Query("ad_id")
	db := middle.MongoDb
	cursor, err := db.Collection(middle.MongoConf.ClueCollection).Find(context.Background(), bson.M{"ad_id": adID})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Writer.Header().Set("Accept-Ranges", "bytes")
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename="+fmt.Sprintf("download%s.xlsx", time.Now().Format("2006-01-02-150405"))) //文件名
	ctx.Writer.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	ctx.Writer.Header().Set("Pragma", "no-cache")
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Stream(func(w io.Writer) bool {
		// Create a xlsx file
		file := xlsx.NewFile()
		sheet, _ := file.AddSheet("clue")
		row := sheet.AddRow()
		row.AddCell().Value = "clue_id"
		row.AddCell().Value = "ad_id"
		row.AddCell().Value = "ad_name"
		row.AddCell().Value = "clue_type"
		row.AddCell().Value = "clue_info"
		row.AddCell().Value = "create_time"

		for cursor.Next(context.Background()) {
			doc := cursor.Current
			id := doc.Lookup("_id").ObjectID().Hex()
			clue := new(Clue)
			bson.Unmarshal(doc, clue)
			row := sheet.AddRow()
			row.AddCell().Value = id
			row.AddCell().Value = clue.AdID
			row.AddCell().Value = clue.AdName
			row.AddCell().Value = clue.ClueType
			row.AddCell().Value = clue.ClueInfo
			row.AddCell().Value = clue.CreateTime
			for _, extend := range clue.Extends {
				row.AddCell().Value = extend.Key
				row.AddCell().Value = extend.Value
			}
		}
		file.Write(w)

		return false
	})

}
