package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    imgDb := InitDb()

    router := gin.Default()

    router.Static("/static", "./static")

    router.GET("/get_image_from_url", func(c *gin.Context) {
        imgUrl := c.Request.URL.Query().Get("img")
        resultPath := MakeMosaic(imgUrl, imgDb)
        c.Redirect(http.StatusMovedPermanently, resultPath)
    })

    router.GET("/admin/fetch_images", func(c *gin.Context) {
        fetched := FetchFromUrlFile()
        response := fmt.Sprintf("Success! Fetched %d images", fetched)
        c.String(http.StatusOK, response)
    })

    router.Run(":8080")
}