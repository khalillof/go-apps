package main

import (
    "github.com/gin-gonic/gin"
    "github.com/khalillof/go-apps/pkg/books"
    "github.com/khalillof/go-apps/pkg/common/db"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
    h := db.Init(dbUrl)

    books.RegisterRoutes(r, h)
    // register more routes here

    r.Run(port)
}
	
