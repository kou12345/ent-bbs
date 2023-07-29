package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/kou12345/ent-bbs/ent"
	"github.com/kou12345/ent-bbs/ent/entry"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:bbs.sqlite?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		eq := client.Entry.Query().Order(ent.Desc(entry.FieldCreatedAt)).Limit(10)
		entries := eq.AllX(context.Background())

		// entriesをJSON形式に変換して返す
		return c.JSON(http.StatusOK, entries)
	})
	e.POST("/add", func(c echo.Context) error {
		e := client.Entry.Create()
		e.SetContent(c.FormValue("content"))
		if _, err := e.Save(context.Background()); err != nil {
			log.Println(err.Error())
			return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return c.Redirect(http.StatusFound, "/")
	})
	e.DELETE("/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err.Error())
			return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		}
		err = client.Entry.DeleteOneID(id).Exec(context.Background())
		if err != nil {
			log.Println(err.Error())
			return c.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return c.Redirect(http.StatusFound, "/")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
