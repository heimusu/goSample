package handler

import (
    "net/http"
    "github.com/labstack/echo"
    "fmt"
    //"github.com/labstack/echo/engine/standard"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gocraft/dbr"
    //"github.com/gocraft/dbr/dialect"
    //"log"
)

type (
    accounts struct {
        ID   int     `db:"id"`
        UserName string  `db:"user_name"`
        PassWord string `db:"password"`
    }

    userinfoJSON struct {
        ID   int     `db:"id"`
        UserName string  `db:"user_name"`
        PassWord string `db:"password"`
    }

    responseData struct {
        Users        []accounts      `json:"users"`
    }
)

var (
    tablename = "accounts"
    seq   = 1
    conn, _ = dbr.Open("mysql", "root:@tcp(127.0.0.1:3306)/test", nil) // スラッシュの後ろはdatabase名
    sess = conn.NewSession(nil)
)

func MainPage() echo.HandlerFunc{
    return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
        username := c.Param("username")
        return c.String(http.StatusOK, "Hello World " + username)
    }
}

func JsonReturn() echo.HandlerFunc {
    return func(c echo.Context) error {
         jsonMap := map[string]string {
            "foo": "bar",
            "hoge":"fuga",
        }
        return c.JSON(http.StatusOK, jsonMap)
    }
}


func SelectUsers() echo.HandlerFunc {
    return func(c echo.Context) error {
        var u []accounts
        // log.Println(conn);
        sess.Select("*").From(tablename).Load(&u)
        response := new(responseData)
        response.Users = u
        return c.JSON(http.StatusOK,response)
    }
}

// func insertUser(c echo.Context) error {
func InsertUser() echo.HandlerFunc {
    return func(c echo.Context) error {
        u := new(userinfoJSON)
        if err := c.Bind(u); err != nil {
            fmt.Println(err)
            return err
        }
        sess.InsertInto(tablename).Columns("id","user_name","password").Values(u.ID,u.UserName,u.PassWord).Exec()
        return c.NoContent(http.StatusOK)
    }
}
