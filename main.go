package main

import (
    "./handler"
    "./intercepter"
    //"fmt"
    //"net/http"
    "github.com/labstack/echo"
    //"github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo/middleware"

    _ "github.com/go-sql-driver/mysql"
    //"github.com/gocraft/dbr/dialect"
)



func main() {
    // echoのインスタンス
    e := echo.New()

    // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Basic Auth
    e.Use(interceptor.BasicAuth())

    // CORS
    e.Use(middleware.CORS())

    // ルーティング
    e.GET("/hello/:username", handler.MainPage(), interceptor.BasicAuth())

    e.GET("/json", handler.JsonReturn())

    e.POST("/users/", handler.InsertUser())

    e.GET("/users",handler.SelectUsers())


    // サーバー起動
    // e.Start(":1323")
    //  e.Run(standard.New(":1323"))
    e.Logger.Fatal(e.Start(":3000"))
}
