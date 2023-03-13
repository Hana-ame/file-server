package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Post("/api/upload/binary-data", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		log.Println("Receive")
		file := c.Body()
		hash, err := hashBytes(file)
		filename := fmt.Sprintf("%x", hash)[0:8]

		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"location": "fileServer.go/hashBytes",
				"error":    err.Error(),
			})
		}

		if err := saveBytesToFile(file, filename); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"location": "fileServer.go/saveBytesToFile",
				"error":    err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"id": filename,
		})
	})

	app.Get("/0/:hash/:filename", func(c *fiber.Ctx) error {

		log.Println(c.GetReqHeaders()["Cf-Connecting-Ip"])

		hash := c.Params("hash")
		f, err := os.Open("./0/" + hash)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"location": "fileServer.go/os.Open",
				"error":    err.Error(),
			})
		}
		data, err := io.ReadAll(f)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"location": "fileServer.go/io.ReadAll",
				"error":    err.Error(),
			})
		}
		c.Set("Content-Type", "application/octet-stream")
		_, err = c.Status(fiber.StatusOK).Write(data)
		return err
	})

	app.Static("/", "./html")

	// Start server
	log.Fatal(app.Listen(":3000"))
}
