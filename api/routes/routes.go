package routes

import (
	"github.com/gofiber/fiber/v2"
	"server-bof/api/controllers"
	"server-bof/stripe"
)

func Routes(app *fiber.App) {

	//payment
	app.Post("/api/order", stripe.Payment)

	//user routes
	app.Post("/api/user/register", controllers.UserCtrlRegister)
	app.Post("/api/user/login", controllers.UserCtrlLogin)
	app.Put("/api/user/update", controllers.UpdateUser)
	app.Post("/api/user/logout", controllers.Logout)
	app.Get("/api/user/get", controllers.GetUser)
	app.Post("/api/user/create_admin", controllers.CreateAdmin)

	//Review routes
	app.Post("/api/review/add", controllers.AddReview)
	app.Get("/api/review/get", controllers.GetReview)
	app.Put("/api/review/update", controllers.UpdateReview)
	app.Delete("/api/review/delete", controllers.RemoveReview)

	//Product routs
	app.Post("/api/product/add", controllers.AddProduct)
	app.Get("/api/product/get", controllers.GetProducts)
	app.Get("/api/product/:product_id", controllers.GetProduct)
	app.Put("/api/product/update/:product_id", controllers.UpdateProduct)
	app.Delete("/api/product/delete/:product_id", controllers.RemoveProduct)

}
