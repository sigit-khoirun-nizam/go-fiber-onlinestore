package controllers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sigit-khoirun-nizam/go-fiber-onlinestore/app/models"
	"github.com/sigit-khoirun-nizam/go-fiber-onlinestore/app/utils"
)

// GetProductsByCategory mengambil daftar produk berdasarkan kategori dari database
func GetProductsByCategory(c *fiber.Ctx) error {
	category := c.Params("category")
	var products []models.Product
	utils.DB.Where("category = ?", category).Find(&products)
	return c.JSON(products)
}

// AddProduct menambahkan produk ke dalam database
func AddProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	utils.DB.Create(&product)
	return c.JSON(product)
}

// GetAllProducts mengambil semua produk dari database
func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	utils.DB.Find(&products)
	return c.JSON(products)
}

// UpdateProduct mengupdate informasi produk berdasarkan ID
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	utils.DB.First(&product, id)
	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Product tidak ditemukan")
	}
	if err := c.BodyParser(&product); err != nil {
		return err
	}
	utils.DB.Save(&product)
	return c.JSON(product)
}

// DeleteProduct menghapus produk berdasarkan ID
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	utils.DB.First(&product, id)
	if product.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Product tidak ditemukan")
	}
	utils.DB.Delete(&product)
	return c.SendStatus(fiber.StatusNoContent)
}

// AddToCart menambahkan produk ke keranjang belanja
func AddToCart(c *fiber.Ctx) error {
	var item models.CartItem
	if err := c.BodyParser(&item); err != nil {
		return err
	}
	utils.DB.Create(&item)
	return c.JSON(item)
}

// GetCartItems mengambil daftar item dalam keranjang belanja
func GetCartItems(c *fiber.Ctx) error {
	var items []models.CartItem
	utils.DB.Find(&items)
	return c.JSON(items)
}

// RemoveFromCart menghapus item dari keranjang belanja berdasarkan ID
func RemoveFromCart(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.CartItem
	utils.DB.First(&item, id)
	if item.ID == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Item tidak ditemukan")
	}
	utils.DB.Delete(&item)
	return c.SendStatus(fiber.StatusNoContent)
}

// Checkout menangani proses pembayaran dan checkout
func Checkout(c *fiber.Ctx) error {
	return c.SendString("Pembayaran Berhasil!")
}

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	utils.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var existingUser models.User
	utils.DB.Where("username = ?", user.Username).First(&existingUser)

	if existingUser.ID == 0 || existingUser.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(existingUser)
}
