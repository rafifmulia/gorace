package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/utils"
)

var (
	f1 *os.File
)

func init() {
	var err error
	f1, err = os.OpenFile("f1.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 666)
	if err != nil {
		panic(err)
	}
}

// Why must be copied? String is immutable, not like []byte.
// result := c.Params("foo")
// Make a copy
// buffer := make([]byte, len(result))
// copy(buffer, result)
// resultCopy := string(buffer)
// https://docs.gofiber.io/#zero-allocation

// The answer is, because fiber uses [utils.UnsafeString] in [fiber.New].

func main() {
	defer f1.Close()
	app := fiber.New()
	app.Get("/race", func(c *fiber.Ctx) error {
		key := c.Query("k")
		go aToF1(key)
		return nil
	})
	err := app.Listen("127.0.0.1:8080")
	if err != nil {
		log.Fatal("Fiber closed:", err)
	}
}

// Append to a file.
func aToF1(s string) {
	time.Sleep(time.Second * 1)
	f1.WriteString(s)
}
