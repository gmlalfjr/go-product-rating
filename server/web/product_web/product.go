package product_web

type ProductCreateRequest struct {
	Title       string
	Description string
}

type ProductCreateResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProductResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
