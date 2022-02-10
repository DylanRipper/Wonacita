package databases

import (
	"project3/config"
	"project3/models"
)

func AddComment(review *models.Comment) (interface{}, error) {
	if err := config.DB.Create(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func AddRatingToProduct(id int) {
	config.DB.Exec("UPDATE products SET rating = (SELECT AVG(rating) FROM reviews WHERE products_id = ?) WHERE id = ?", id, id)
}
