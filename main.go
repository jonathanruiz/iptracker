package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/jonathanruiz/iptracker-app/cmd"
	_ "github.com/lib/pq"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var ips []string
	rows, err := db.Query("SELECT * FROM ip")
	fmt.Println("rows: ", rows)

	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		ips = append(ips, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": ips,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// assign environment variables to db variables
	host := os.Getenv("HOST")
	web_port := os.Getenv("PORT")
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	// Create connection string
	connStr := "postgresql://" + user + ":" + password + "@" + host + "/" + dbname + "?sslmode=disable"

	fmt.Println(connStr)

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		println("Is this working?")
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	if web_port == "" {
		web_port = "3000"
	}

	app.Static("/", "./public") // add this before starting the app

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", web_port)))

	fmt.Println("Private IP : ", cmd.GetPrivateIP())
	fmt.Println("Public IP: ", cmd.GetPublicIP())
}
