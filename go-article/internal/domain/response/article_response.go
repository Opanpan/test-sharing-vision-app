package response

// article response
type ArticleResponse struct {
	ID    	 int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

// articles response
type ArticlesResponse struct {
	Data []ArticleResponse `json:"data"`
	Total int64`json:"total"`
}

// respose error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
