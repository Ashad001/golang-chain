package api

type ChatInput struct {
	UserQuery string `json:"user_query" binding:"required"`
}

type ChatOutput struct {
	Score       int                      `json:"score" binding:"required"`
	Description string                   `json:"description" binding:"required"`
	Messages    []map[string]interface{} `json:"messages" binding:"required"`
}
