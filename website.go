package main

import (
	"html/template"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initialization
	app := fiber.New()

	// charger ce template (index.html)
	mapage := template.Must(template.ParseFiles("public/index.html"))

	// variable avec mon prenom
	prenom := "Lorenzo"
	age := 21

	// creer un contexte sous la forme d'une classe (struct)
	type ViewData struct {
		Prenom string
		Age    int
	}

	// donne le changement du dossier contenant le html
	app.Get("/", func(c *fiber.Ctx) error {
		// definir mon contexte avec sa valeur
		data := ViewData{
			Prenom: prenom,
			Age:    age,
		}

		// passer le viewdata à mon template
		err := mapage.Execute(c.Response().BodyWriter(), data)
		if err != nil {
			return err
		}

		c.Type("html")
		return nil
	})

	app.Static("/", "./public")

	// ecouter le port 3000
	log.Fatal(app.Listen(":3000"))
}
