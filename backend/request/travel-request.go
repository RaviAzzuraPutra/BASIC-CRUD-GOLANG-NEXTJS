package request

type TravelRequest struct {
	Name        *string `form:"name" binding:"required"`
	Description *string `form:"description" binding:"required"`
	Price       *int    `form:"price" binding:"required"`
}
