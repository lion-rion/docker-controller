package controller

import (
	"context"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", getContaier)
	router.GET("/start/:id", StartContainer)
	router.GET("/stop/:id", StopContainer)
	return router
}

func getContaier(c *gin.Context) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	// for _, container := range containers {
	// 	fmt.Printf("%s %s %s\n", container.ID[:10], container.Image, container.State)
	// }

	c.HTML(http.StatusOK, "index.html", gin.H{
		"containers": containers,
	})
}

func StartContainer(c *gin.Context) {
	//パラメータからidの取得
	id := c.Param("id")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//docker container start
	cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})

	//リダイレクト
	c.Redirect(http.StatusSeeOther, "/")
}

func StopContainer(c *gin.Context) {
	//パラメータからidの取得
	id := c.Param("id")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//docker container stop
	cli.ContainerStop(context.Background(), id, nil)

	//リダイレクト
	c.Redirect(http.StatusSeeOther, "/")
}
