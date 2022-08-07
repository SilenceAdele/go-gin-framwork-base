package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `form:"name" binding:"required,len=6"`
	Age  int    `form:"age" binding:"numeric,min=18,max=100"`
}

func main() {
	engine := gin.Default()
	engine.GET("/hello", test)
	engine.GET("/json", jsonTest)
	engine.GET("/xml", xmlTest)
	htmlTest(engine)
	requestParam(engine)
	uploadFile(engine)
	routeGroup(engine)
	engine.Run()
}

// route group
func routeGroup(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.GET("/get_user", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "api get_user")
		})

		api.GET("/get_info", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "api get_info")
		})
	}

	admin := engine.Group("/admin")
	{
		admin.GET("/login", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "login")
		})

		admin.GET("/hello", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "admin")
		})
	}
}

// upload file
func uploadFile(engine *gin.Engine) {
	engine.LoadHTMLGlob("templates/*")

	// through get request show upload.html
	engine.GET("/upload", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "upload.html", nil)
	})

	// through post request realization file upload
	engine.POST("/upload", func(ctx *gin.Context) {
		// get signal file

		// f, err := ctx.FormFile("file")
		// if err != nil {
		// 	log.Println(err)
		// }
		// err = ctx.SaveUploadedFile(f, fmt.Sprintf("uploads/%s", f.Filename))

		// if err != nil {
		// 	ctx.String(http.StatusOK, "上传文件失败-> %v", err)
		// } else {
		// 	ctx.String(http.StatusOK, "上传文件成功")
		// }

		// reserve mult files
		from, err := ctx.MultipartForm()
		if err != nil {
			log.Println(err)
		}
		files := from.File["file"]
		for _, file := range files {
			err := ctx.SaveUploadedFile(file, fmt.Sprintf("uploads/%s", file.Filename))
			if err != nil {
				ctx.String(http.StatusOK, "上传文件失败-> %v", err)
				return
			}
		}
		ctx.String(http.StatusOK, "上传文件成功")
	})
}

// request param
func requestParam(engine *gin.Engine) {
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/get", func(ctx *gin.Context) {
		fmt.Println(ctx.Query("name"))
		fmt.Println(ctx.Query("age"))

		ctx.String(http.StatusOK, "get")
	})

	engine.POST("/post", func(ctx *gin.Context) {
		fmt.Println(ctx.PostForm("name"))
		fmt.Println(ctx.PostForm("age"))

		ctx.String(http.StatusOK, "post")
	})

	engine.GET("/get_user/:id", func(ctx *gin.Context) {
		fmt.Println(ctx.Param("id"))

		ctx.String(http.StatusOK, "url param")
	})

	engine.GET("/get2", func(ctx *gin.Context) {
		var user User
		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
		} else {
			ctx.String(http.StatusOK, "name-> %s,age-> %d", user.Name, user.Age)
		}

	})
}

// template
func htmlTest(engine *gin.Engine) {
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/static", "./static")
	engine.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"name":  "admin",
			"age":   "24",
			"users": []string{"周晓晓", "王小五", "成小小"},
		})
	})
}

// test
func test(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello easystack")
}

// json
func jsonTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "admin",
		"age":  24,
	})
}

// xml
func xmlTest(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"name": "admin",
		"age":  24,
	})
}
