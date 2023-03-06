package main

import (
	"bupt.nxy/ginessential/common"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// db := common.GetDB()
	// defer db.Close()

	db := common.InitDB()
	log.Println(db)

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
