package request

type TravelRequest struct {
	Name        *string `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
	Price       *int    `json:"price" binding:"required"`
}
