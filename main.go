package main

import (
    "BTPN/database"
    "BTPN/router"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    database.ConnectDatabase()
    router.SetupRouter(r)
    r.Run(":8080")
}
