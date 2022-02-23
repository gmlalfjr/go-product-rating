package rating_web

type RatingCreateRequest struct {
	Star      int
	Comment   string
	ProductId uint
}

type RatingCreateResponse struct {
	Star      int    `json:"star"`
	Comment   string `json:"comment"`
	ProductId uint   `json:"product_id"`
}
