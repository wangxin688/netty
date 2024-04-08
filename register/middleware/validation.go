package middleware

import (
	"net/http"
	"netty/utils/types"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func errorValidationBody(dto any) []types.ValidationError {
	var errors []types.ValidationError

	if err := validate.Struct(dto); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, types.ValidationError{
				Field: err.Field(),
				Value: err.Value(),
			})
		}
	}
	return errors
}

func errorValidationQuery(dto any) []types.ValidationError {
	var errors []types.ValidationError

	if err := validate.Struct(dto); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, types.ValidationError{
				Field: err.Field(),
				Value: err.Value(),
			})
		}
	}
	return errors
}

func ValidationBodyMiddleware(dto any) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyDto any // Define the type of your request body DTO

		// Bind the request body to the DTO
		if err := c.ShouldBindJSON(&bodyDto); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"code":    http.StatusUnprocessableEntity,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		// Validate the DTO
		errors := errorValidationBody(bodyDto)

		if len(errors) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Invalid request body",
				"data":    errors,
			})
			return
		}

		c.Set("validatedBody", bodyDto) // Optionally set the validated body in the context
		c.Next()
	}
}

func ValidationQueryMiddleware(dto any) gin.HandlerFunc {
	return func(c *gin.Context) {
		var queryDto any // Define the type of your query DTO

		// Bind the query parameters to the DTO
		if err := c.ShouldBindQuery(&queryDto); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"code":    http.StatusUnprocessableEntity,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		// Validate the DTO
		errors := errorValidationQuery(queryDto)

		if len(errors) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Invalid query parameters",
				"data":    errors,
			})
			return
		}

		c.Set("validatedQuery", queryDto) // Optionally set the validated query in the context
		c.Next()

	}
}
