package main

import (
    "log"
    // "net/http"

    // "myapp/migrations"
    // "github.com/labstack/echo/v5"
    "github.com/pocketbase/pocketbase"
    // "github.com/pocketbase/pocketbase/apis"
    // "github.com/pocketbase/pocketbase/core"

    "strings"
    "os"
    "github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
    app := pocketbase.New()

    // app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
    //     // add new "GET /hello" route to the app router (echo)
    //     e.Router.AddRoute(echo.Route{
    //         Method: http.MethodGet,
    //         Path:   "/hello",
    //         Handler: func(c echo.Context) error {
    //             return c.String(200, "Hello world!")
    //         },
    //         Middlewares: []echo.MiddlewareFunc{
    //             apis.ActivityLogger(app),
    //         },
    //     })

    //     return nil
    // })

    // loosely check if it was executed using "go run"
    isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

    migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
        // enable auto creation of migration files when making collection changes in the Admin UI
        // (the isGoRun check is to enable it only during development)
        Automigrate: isGoRun,
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}