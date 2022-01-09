package main

import (
	"Golang-blog/model"
	"Golang-blog/routers"
	"Golang-blog/utils"
)

func main() {
	utils.InitConfig()
	model.InitDB()
	routers.InitRouters()
}
