package databases

import (
	"project3/config"
	"project3/models"
)

// Fungsi untuk membuat menyewakan produk baru
func CreateProduct(product *models.Products) (*models.Products, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

// Fungsi untuk mendapatkan nama kota dari city_id
func GetCity(id int) (string, error) {
	var town models.City
	if err := config.DB.Where("id = ?", id).Find(&town); err.Error != nil {
		return "", err.Error
	}
	return town.City_Name, nil
}

// Fungsi untuk memasukkan foto product
func InsertPhoto(photo *models.Photos) (interface{}, error) {
	if err := config.DB.Create(&photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}

// Fungsi untuk mendapatkan seluruh product
func GetAllProducts() (interface{}, error) {
	var results []models.GetAllProduct
	tx := config.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, cities.city_name, products.price, products.description, products.stock, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.deleted_at IS NULL").Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}

// Fungsi untuk mendapatkan product berdasarkan id product
func GetProductByID(id uint) (*models.GetProduct, error) {
	var result models.GetProduct
	tx := config.DB.Table("products").Select("products.id, products.users_id, users.created_at, users.nama, users.phone_number, products.name, products.subcategory_id, subcategories.subcategory_name, products.city_id, cities.city_name, products.caption, products.latitude, products.longitude").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join users on products.users_id = users.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.id = ?", id).Find(&result)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return &result, nil
}

// Fungsi untuk mendapatkan product berdasarkan subcategory id
func GetProductsBySubcategoryID(id int) (interface{}, error) {
	var results []models.GetAllProduct
	tx := config.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, cities.city_name, products.description, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.deleted_at IS NULL AND products.subcategory_id = ?", id).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}

// Fungsi untuk mendapatkan product berdasarkan users_id
func GetProductsByUserID(id int) (interface{}, error) {
	var results []models.GetAllProduct
	tx := config.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, cities.city_name, products.price, products.description, products.stock, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.deleted_at IS NULL AND products.users_id = ?", id).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}

// Fungsi untuk mendapatkan seluruh url photo product tertentu
func GetUrl(id uint) ([]string, error) {
	var url []string
	tx := config.DB.Table("photos").Select("photos.url").Where("photos.products_id = ?", id).Find(&url)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return url, nil
}

//
// Fungsi untuk mendapatkan seluruh guarantee product tertentu
func GetGuarantee(id int) ([]string, error) {
	var guarantee []string
	tx := config.DB.Table("products_guarantees").Select("guarantees.guarantee_name").Joins(
		"join guarantees on products_guarantees.guarantee_id = guarantees.id").Where("products_guarantees.products_id = ?", id).Find(&guarantee)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return guarantee, nil
}

// function database untuk menghapus product  by id
func DeleteProduct(id int) {
	config.DB.Exec("DELETE from products WHERE id = ?", id)
}

// Fungsi untuk product by id
func DeleteProductByID(id int) (interface{}, error) {
	var product models.Products
	if err := config.DB.Where("id = ?", id).Delete(&product).Error; err != nil {
		return nil, err
	}
	return "deleted", nil
}

func GetProductOwner(id int) (int, error) {
	var ownerProduct int
	tx := config.DB.Raw("SELECT users_id FROM products WHERE id = ?", id).Scan(&ownerProduct)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return ownerProduct, nil
}

// Fungsi untuk mendapatkan product berdasarkan search name
func GetProductsByName(product string) (interface{}, error) {
	var results []models.GetAllProduct
	var search = "%" + product + "%"
	tx := config.DB.Table("products").Select("products.id, products.users_id, products.name, subcategories.subcategory_name, products.subcategory_id, products.city_id, cities.city_name, products.price, products.description, products.stock, photos.url").Group("products.id").Joins(
		"join subcategories on subcategories.id = products.subcategory_id").Joins(
		"join photos on photos.products_id = products.id").Joins(
		"join cities on products.city_id = cities.id").Where("products.deleted_at IS NULL AND products.name LIKE ?", search).Find(&results)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return results, nil
}
