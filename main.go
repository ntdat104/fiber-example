package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Post("/test/post-http", httpPost)

	app.Get("/test/http", httpGet)

	app.Get("/:name", handleName)

	app.Get("/:name/:age", handleNameAndAge)

	app.Get("/api/*", handleWildcard)

    app.Listen(":3000")
}

func handleName(c *fiber.Ctx) error {
	name := c.Params("name")
	return c.SendString("Hello, " + name + "👋!")
}

func handleNameAndAge(c *fiber.Ctx) error {
	name := c.Params("name")
	age := c.Params("age")
	return c.SendString("Hello, " + name + "👋! You are " + age + " years old.")
}

func handleWildcard(c *fiber.Ctx) error {
	param := c.Params("*")
	return c.SendString("You are on the api route. " + param)
}

func httpGet(c*fiber.Ctx) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://ga-mobile-api.loklok.tv/cms/app/homePage/getHome?page=8", nil)

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"lang": []string{"en"},
		"versioncode": []string{"11"},
		"clienttype": []string{"ios_jike_default"},
		"deviceid": []string{""},
	}
	res, err := client.Do(req)
	
	if err != nil {
		return c.SendString(err.Error())
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(string(body))
}

func httpPost(c *fiber.Ctx) error {
	// payload := struct {
    //     SearchKeyWord  string `json:"searchKeyWord"`
    //     Size int `json:"size"`
    // }{}

	// err := c.BodyParser(&payload);
	
	// if err != nil {
    //     return err
    // }

	// return c.JSON(payload)

	payload := struct {
        SearchKeyWord  string `json:"searchKeyWord"`
        Size int `json:"size"`
    }{}

	var buf bytes.Buffer
    err := json.NewEncoder(&buf).Encode(payload)
    if err != nil {
		return c.SendString(err.Error())
    }

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://ga-mobile-api.loklok.tv/cms/app/search/searchLenovo", &buf)

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"lang": []string{"en"},
		"versioncode": []string{"11"},
		"clienttype": []string{"ios_jike_default"},
		"deviceid": []string{""},
	}
	res, err := client.Do(req)

	if err != nil {
		return c.SendString(err.Error())
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString(string(body))
}