package interceptor

import (
    "github.com/labstack/echo/middleware"
    "github.com/labstack/echo"
)


func BasicAuth() echo.MiddlewareFunc {
    return middleware.BasicAuth(func(username, password string, c echo.Context) (error, bool) {
        if username == "joe" && password == "secret" {
            return nil, true
        }
        return nil, false
    })
}
