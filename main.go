package main

import (
	"fmt"
	"myapp/app"
	"myapp/marketplace/amazon/mws"
	"myapp/pkg/database"
	"myapp/pkg/debug"
	"myapp/pkg/http/auth"
	"myapp/pkg/times"
	"myapp/pkg/utils"
	//"myapp/controller"
	//"myapp/controller/admin"
	//"myapp/controller/blog"
	//"myapp/controller/forum"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./app/config.json"); err != nil {
		panic(fmt.Errorf("Failed to read the configuration file: %s", err))
	}

	//debug.Print(app.Config)
	fmt.Println(app.Config.Database.Host)
	fmt.Println(app.Version)

	//mws.initMarketplaces()
	debug.Print(mws.GetMarketplace("CA"))
	debug.Print(mws.GetMarketplace("A2EUQ1WTGCTBG2"))
	fmt.Println(mws.GetEndpoint("CA"))
}

func test1() {
	fmt.Println(times.Format("Y-m-d H:i:s"))
	fmt.Println(auth.CreateToken(123))

	config := map[string]string{"dbname": "bte"}
	database.Open(config)

	if utils.FileExists("c:/Windows/notepad.exe") {
		fmt.Println("File Exists")
	} else {
		fmt.Println("File Not Exists")
	}
	if utils.DirExists("c:/Windows/notepad.exe") {
		fmt.Println("Dir Exists")
	} else {
		fmt.Println("Dir Not Exists")
	}
}

func testRoute() {
	/*
		g := e.Group("")
		controller.RegisterRoutes(g)

		adminGroup := e.Group("/admin")
		admin.RegisterRoutes(adminGroup)

		blogGroup := e.Group("/blog")
		blog.RegisterRoutes(blogGroup)

		forumGroup := e.Group("/forum")
		forum.RegisterRoutes(forumGroup)
	*/
}
