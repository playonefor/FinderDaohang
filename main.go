package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"               // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	_ "github.com/GoAdminGroup/themes/sword"                       // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"

	"FinderDaohang/models"
	"FinderDaohang/pages"
	"FinderDaohang/tables"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()
	// 设置模板路径
	r.LoadHTMLGlob("html/index.html")
	r.Static("/public", "./html/public")

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromJSON("./config.json").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)

	models.Init(eng.SqliteConnection())

	db := models.GetDB()
	var urlGroups []models.UrlGroup
	// 查询 UrlGroup 表中的 id 和 name 字段
	db.Select("id, group_name").Preload("UrlInfos").Find(&urlGroups)

	//前台页面
	r.GET("/", frontPageHandler())

	srv := &http.Server{
		Addr:    ":9115",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Print("closing database connection")
	eng.SqliteConnection().Close()

	log.Println("Server exiting")
}

// 定义你的前台页面处理函数
// frontPageHandler 返回一个 gin.HandlerFunc
func frontPageHandler() gin.HandlerFunc {
	// 返回一个闭包
	return func(c *gin.Context) {
		db := models.GetDB()
		var urlGroups []models.UrlGroup
		db.Select("id, group_name").Preload("UrlInfos").Find(&urlGroups)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"UrlGroups": urlGroups,
		})
	}
}
