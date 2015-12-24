package main

import (
	"fmt"
	"github.com/bugsnag/bugsnag-go"
	"github.com/d4l3k/go-pry/pry"
	"github.com/gin-gonic/gin"
	"github.com/hoisie/mustache"
)

type M_Template struct {
	Content  string
	Template string
	View     string
}

func translateType(v_type, namespace, name string) {
	if v_type == "name" || v_type == "&" {
		return "variable"
	} else if v_type == "#" && namespace == "collection" && name != "length" {
		return "collection"
	} else if v_type == "#" && namespace == "spreadsheet" && name != "length" {
		return "spreadsheet"
	} else if v_type == "#" && namespace == "block" {
		return "block"
	} else if v_type == "#" && namespace == "helpers" {
		return "helper"
	} else if v_type == "#" || v_type == "^" {
		return "boolean"
	} else {
		return nil
	}
}

func parseAst() {
}

func findDuplicateTag() {
}

func shouldOverrideDuplicate() {
}

func filterBooleans() {
}

func index(c *gin.Context) {
	c.String(200, "")
}

func status(c *gin.Context) {
	c.String(200, "OK")
}

func parse(c *gin.Context) {
	var html M_Template
	c.Bind(&html)

	ast, _ := mustache.ParseString(html.Content)

	fmt.Println(ast)

}

func interpolate(c *gin.Context) {
	var html M_Template
	c.Bind(&html)

	template := html.Template
	view := html.View

	data := mustache.Render()
}

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:              "3a198be49a4a2ed0f1c549818d792512",
		ReleaseStage:        "production",
		NotifyReleaseStages: []string{"production", "staging"},
	})

	app := gin.Default()
	app.GET("/", index)
	app.GET("/status", status)
	app.POST("/parse", parse)
	app.POST("/interpolate", interpolate)

	app.Run(":8000")
}
