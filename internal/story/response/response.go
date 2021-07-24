package response

import "api-web-bangarief/internal/entity"

type story struct {
	ID        int64  `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Story     string `json:"story"`
	Status    int32  `json:"status"`
	Email     string `json:"email"`
	CreatedOn string `json:"created_on"`
}

// Query ...
type Query struct {
	Message string `json:"message"`
	Data    struct {
		Stories []story `json:"stories"`
	} `json:"data"`
}

// Set ...
func (res Query) Set(data []entity.Story) {
	res.Data.Stories = make([]story, 0)

	for _, item := range data {
		res.Data.Stories = append(res.Data.Stories, story{
			ID:        item.ID,
			Author:    item.Author.String,
			Title:     item.Title.String,
			Story:     item.Story.String,
			Status:    item.Status.Int32,
			Email:     item.Email.String,
			CreatedOn: item.CreatedOn.Time.String(),
		})
	}
}

// Create ...
type Create struct {
	Message string `json:"message"`
	Data    struct {
		Story story `json:"story"`
	} `json:"data"`
}

// Set ...
func (res Create) Set(data *entity.Story) {
	res.Data.Story.ID = data.ID
	res.Data.Story.Author = data.Author.String
	res.Data.Story.Title = data.Title.String
	res.Data.Story.Story = data.Story.String
	res.Data.Story.Status = data.Status.Int32
	res.Data.Story.Email = data.Email.String
	res.Data.Story.CreatedOn = data.CreatedOn.Time.String()
}
