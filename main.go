package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
)

func main() {
    imgDb := InitDb()

    router := gin.Default()
    // router.GET("/", func(c *gin.Context) {
    //     c.String(http.StatusOK, "hello world")
    // })
    // router.GET("/ping", func(c *gin.Context) {
    //     c.String(http.StatusOK, "pong")
    // })
    // router.POST("/submit", func(c *gin.Context) {
    //     c.String(http.StatusUnauthorized, "not authorized")
    // })
    // router.PUT("/error", func(c *gin.Context) {
    //     c.String(http.StatusInternalServerError, "an error happened :(")
    // })
    router.GET("/get_image_from_url", func(c *gin.Context) {

        //TODO: parse url from request
        imgUrl := c.Request.URL.Query().Get("img")
        MakeMosaic(imgUrl, imgDb)

        //TODO: handle response from MakeMosaic and return image to user
        c.String(http.StatusOK, "hello world")
    })

    router.GET("/test_parser", func(c *gin.Context){
        i := c.Request.URL.Query().Get("img")
        c.String(http.StatusOK, i)
        })


    router.GET("/fetch_images", func(c *gin.Context) {
        fetched := FetchFromUrlFile()
        response := fmt.Sprintf("Success! Fetched %d images", fetched)
        c.String(http.StatusOK, response)
    })

    router.Run(":8081")
}