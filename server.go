package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:email?", func(c *fiber.Ctx) error {
		if c.Params("email") != "" {
			email := c.Params("email")
			emailMD5 := md5.Sum([]byte(email))
			emailHash := hex.EncodeToString(emailMD5[:])
			c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
			return c.SendString("<img style=\"height: 512px; width: 512px; background: grey;\" src=\"https://www.libravatar.org/avatar/" + emailHash + "?s=512\">")
		}
		return c.SendString("To use please open in this format https://grabatar.co/you@email.com where 'you@email.com' is your email id")
	})

	app.Listen(":8080")
}
