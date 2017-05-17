package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/shawguo-cn/archetype-golang/dominus/config"
	"github.com/shawguo-cn/archetype-golang/dominus/webapp/types"
	mgo "gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"encoding/json"
	"os"
	"io/ioutil"
	"github.com/urfave/cli"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	config       *config.Config
	mongoSession *mgo.Session
}

func NewHandler(config *config.Config, mongoSession *mgo.Session) *Handler {
	return &Handler{config: config, mongoSession:mongoSession}
}

func (hd *Handler) queryDoc(ctx *gin.Context) {
	result := types.Person{}
	c := hd.mongoSession.DB(hd.config.MongodbTestDb).C("people")
	err := c.Find(bson.M{"name": "shawguo"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, result)
}

func main() {
	app := cli.NewApp()
	app.Name = "golang-webapp"
	app.Usage = "golang-webapp"
	app.Action = func(ctx *cli.Context) error {

		//EE:load config file
		configFile := ctx.String("config")
		buf, err := ioutil.ReadFile(configFile)
		if err != nil {
			return err
		}
		var config config.Config
		if err := json.Unmarshal(buf, &config); err != nil {
			return err
		}
		log.Println(config)

		//EE:initialize test data
		session, err := mgo.Dial(config.MongodbUrl)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Primary, true)

		c := session.DB(config.MongodbTestDb).C("people")
		err = c.Insert(&types.Person{"shawguo", "+55 53 8116 9639"}, &types.Person{"yuki", "+55 53 8402 8510"})
		if err != nil {
			log.Fatal(err)
		}

		handler := NewHandler(&config, session)

		gin.SetMode(config.GinModel)
		r := gin.Default()
		r.GET("/ping", handler.queryDoc)
		r.Run(config.GinPort) // listen and serve on 0.0.0.0:8080

		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "./config.json",
			Usage: "config file",
		},
	}

	app.Run(os.Args)

}