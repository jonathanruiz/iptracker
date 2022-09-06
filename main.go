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

// IndexHandler handles the index route
func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	// Initialize res and ips
	var res string
	var ips []string

	// Query database to select all ips
	rows, err := db.Query("SELECT * FROM ip")

	// Check for errors
	if err != nil {
		log.Fatal(err)
	}

	// Close query
	defer rows.Close()

	// Check for errors
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}

	// Loop through rows
	for rows.Next() {
		rows.Scan(&res)
		ips = append(ips, res)
	}

	// Return ips to index.html
	return c.Render("index", fiber.Map{
		"IPS": ips,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func main() {

	// Load .env file
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

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Create
	engine := html.New("./views", ".html")

	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Create root GET route
	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	// Create root POST route
	app.Post("/", func(c *fiber.Ctx) error {

		return postHandler(c, db)
	})

	// Assign web listening port
	if web_port == "" {
		web_port = "3000"
	}

	// Set static files
	app.Static("/", "./public")

	// Start server
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", web_port)))

	fmt.Println("Private IP : ", cmd.GetPrivateIP())
	fmt.Println("Public IP: ", cmd.GetPublicIP())
}
