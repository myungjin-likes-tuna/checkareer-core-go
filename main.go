package main

import (
	"checkareer-core/cmd"
	_ "checkareer-core/docs"
)

// @title todolist API
// @version 21.0.0
// @description todolist API

// @contact.name API Support
// @contact.url https://github.com/myungjin-likes-tuna/checkareer-core-go/issues
// @contact.email nexters@kakao.com

// @host
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cmd.Execute()
}
