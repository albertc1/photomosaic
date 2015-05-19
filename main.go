package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
)

func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "hello world")
    })
    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })
    router.POST("/submit", func(c *gin.Context) {
        c.String(http.StatusUnauthorized, "not authorized")
    })
    router.PUT("/error", func(c *gin.Context) {
        c.String(http.StatusInternalServerError, "an error happened :(")
    })

    router.GET("/fetch_images", func(c *gin.Context) {
        fetched := FetchFromUrlFile()
        response := fmt.Sprintf("Success! Fetched %d images", fetched)
        c.String(http.StatusOK, response)
    })

    router.Run(":8080")
}