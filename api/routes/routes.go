package routes

import (
	"github.com/gofiber/fiber/v2"
	"server-bof/api/controllers"
)

func Routes(app *fiber.App) {

	//Category routes
	app.Post("/api/category/add", controllers.AddCategory)
	app.Get("/api/category/categories", controllers.GetCategories)
	app.Get("/api/category/get/:category_id", controllers.GetCategory)
	app.Put("/api/category/update", controllers.UpdateCategory)
	app.Delete("/api/category/remove/:category_id", controllers.DeleteCategory)

	//Address routes
	app.Post("/api/address/add", controllers.AddAddress)
	app.Get("/api/address/get/:address_id", controllers.GetAddress)
	app.Put("/api/address/update", controllers.UpdateAddress)
	app.Delete("/api/address/remove/:address_id", controllers.DeleteAddress)

	//Order routes
	app.Post("/api/order/add", controllers.AddOrder)
	app.Get("/api/category/order", controllers.GetOrder)
	app.Get("/api/category/get/:order_id", controllers.DeleteOrder)

	//payment
	app.Post("/api/order", controllers.Stripe)

	//Customer routes
	app.Post("/api/v1/customer/register", controllers.UserCtrlRegister)
	app.Post("/api/v1/customer/login", controllers.UserCtrlLogin)
	app.Put("/api/v1/customer/update", controllers.UpdateUser)
	app.Post("/api/v1/customer/logout", controllers.Logout)
	app.Get("/api/v1/customer/get", controllers.GetUser)
	app.Post("/api/v1/customer/get/orders", controllers.GetCustomerOrderDetails)

	//Review routes
	app.Post("/api/review/add", controllers.AddReview)
	app.Put("/api/review/update", controllers.UpdateReview)
	app.Delete("/api/review/delete", controllers.RemoveReview)

	//Product routes
	app.Post("/api/product/add", controllers.AddProduct)
	app.Get("/api/product/get", controllers.GetProducts)
	app.Get("/api/product/:product_id", controllers.GetProduct)
	app.Put("/api/product/update/:product_id", controllers.UpdateProduct)
	app.Delete("/api/product/delete/:product_id", controllers.RemoveProduct)

}
