# Fmustache

Mustache template render for Forgery.

## Install

    go get github.com/ricallinson/fmustache

## Use

    package main

    import(
        "github.com/ricallinson/forgery"
        "github.com/ricallinson/fmustache"
    )

    func init() {
        app := f.CreateServer()
        app.Engine(".html", fmustache.Make())
        app.Get("/", func(req *f.Request, res *f.Response, next func()) {
            res.Render("index.html", map[string]string{"title": "Mu"})
        })
        app.Listen(3000)
    }

The above code will look for the file `./views/index.html`.