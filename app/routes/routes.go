package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sigit-khoirun-nizam/go-fiber-onlinestore/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/products/:category", controllers.GetProductsByCategory)
	app.Post("/products", controllers.AddProduct)
	app.Get("/products", controllers.GetAllProducts)
	app.Put("/products/:id", controllers.UpdateProduct)
	app.Delete("/products/:id", controllers.DeleteProduct)

	app.Post("/cart/add", controllers.AddToCart)
	app.Get("/cart", controllers.GetCartItems)
	app.Delete("/cart/:id", controllers.RemoveFromCart)
	app.Post("/checkout", controllers.Checkout)

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
