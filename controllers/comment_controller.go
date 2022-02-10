package controllers

import (
	"net/http"
	"project3/lib/databases"
	"project3/middlewares"
	"project3/models"
	"project3/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddCommentController(c echo.Context) error {
	product_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.FalseParamResponse())
	}
	var comment models.Comment
	c.Bind(&comment)
	comment.ProductsID = uint(product_id)
	middlewares.ExtractTokenUserId(c)

	if comment.Rating <= 0 || comment.Rating > 5 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "please choose between 1 - 5",
		})
	}

	databases.AddRatingToProduct(int(comment.ProductsID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse())
	}
	return c.JSON(http.StatusOK, response.SuccessResponseNonData())
}
