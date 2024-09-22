package model

type CardRequest struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Username    string `json:"username"`
	Status      string `json:"status"`
}

type CardResponse struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   string `json:"createdAt"`
}

type CardWithCommentsResponse struct {
	Id          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      string            `json:"status"`
	Comments    []CommentResponse `json:"comments"`
	CreatedBy   string            `json:"created_by"`
	CreatedAt   string            `json:"createdAt"`
}

type PaginationRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
