package main

import (
	"GOandVueTest/conf"
	"GOandVueTest/handlers"
	"database/sql"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {

	db := initDB()
	//migrate(db)
	e := echo.New()

	// route - evgroup

	tasks := e.Group("/tasks")
	{
		tasks.File("/list", "public/tasks.html")
		tasks.GET("", handlers.GetTasks(db))
		tasks.PUT("", handlers.PutTask(db))
		tasks.DELETE("/:id", handlers.DeleteTask(db))
	}

	evcharger := e.Group("/evchargers")
	{
		evcharger.File("/list", "public/evcharger.html")
		evcharger.GET("", handlers.ListEvCharger)
	}

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}
	return db
}

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	t.templates.Lookup("")

	return t.templates.ExecuteTemplate(w, name, data)
}

/*
func migrate(db *sql.DB) {
	sql := `
			CREATE TABLE public."EV_USERACCOUNT"
			(
				"UserId" SERIAL PRIMARY KEY,
				"UserName" character varying(100) COLLATE pg_catalog."default" UNIQUE,
				"UserPassword" character varying(100) COLLATE pg_catalog."default" UNIQUE
			)
			WITH (
				OIDS = FALSE
			)
			TABLESPACE pg_default;

			ALTER TABLE public."EV_USERACCOUNT"
				OWNER to postgres;`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
*/