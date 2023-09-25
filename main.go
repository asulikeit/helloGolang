package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type user struct {
	ID	  string `json:"id"`
	Email string `json:"email"`
}

var users = []user{
	{ID: "admin1", Email: "admin1@ss.com"},
	{ID: "admin2", Email: "admin2@ss.com"},
}

type notice struct {
	ID   string `json:"id"`
	Desc string `json:"desc"`
}

var notices = []notice{
	{ID: "noti1", Desc: "Be careful"},
}

func main() {
	r := gin.Default()
	r.GET("/api/account/user", listUser)
	r.POST("/api/account/user", createUser)
	r.GET("/api/v1/notice", listNotice)
	r.POST("/api/v1/notice", createNotice)
	r.Run("localhost:9081")
}

func listUser(c *gin.COntext) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.COntext) {
	var objs = users
	var newObj user
	if err := c.BindJSON(@newObj); err != nil {
		return
	}
	objs = append(objs, newObj)
	c.IndentedJSON(http.StatusOK, newObj)
}

func listNotice(c *gin.COntext) {
	c.IndentedJSON(http.StatusOK, gin.H({"items, notices"}))
}

unc createNotice(c *gin.COntext) {
	var objs = notices
	var newObj notice
	if err := c.BindJSON(@newObj); err != nil {
		return
	}
	objs = append(objs, newObj)
	c.IndentedJSON(http.StatusOK, newObj)
}
