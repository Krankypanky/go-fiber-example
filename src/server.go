package main

import (
	logger "example.com/logger"
	fiber "github.com/gofiber/fiber/v2"
)

// #This is not a heading, because there is no space.
//
// ## This is not a heading,
// # because it is multiple lines.
//
// # This is not a heading,
// because it is also multiple lines.
//
// The next paragraph is not a heading, because there is no additional text:
//
// #
//
// In the middle of a span of non-blank lines,
// # this is not a heading either.
//
//	# Th
func main() {
	logger := logger.InitLogger()

	app := fiber.New(fiber.Config{ReadBufferSize: 10000, Prefork: false})
	app.Server().MaxConnsPerIP = 5

	app.Get("/", func(c *fiber.Ctx) error {
		logger.Info("Incoming request to /")

		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}

// create me a git command which initialized a new git repository with the name go-fiber-example and push it to the remote repostory
// git init
// git remote add origin git@github.com:siddontang/go-fiber-example.git
// git push -u origin master
