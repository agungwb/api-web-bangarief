package request

import (
	"api-web-bangarief/internal/entity"
	"database/sql"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Create ...
type Create struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Story  string `json:"story"`
	Email  string `json:"email"`
}

// Validate ...
func (req *Create) Validate() error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Author, validation.Required),
		validation.Field(&req.Title, validation.Required),
		validation.Field(&req.Story, validation.Required),
		validation.Field(&req.Email, is.Email),
	)
}

// Populate ...
func (req *Create) Populate(data *entity.Story) {
	data.Author = sql.NullString{String: req.Author, Valid: true}
	data.Title = sql.NullString{String: req.Title, Valid: true}
	data.Story = sql.NullString{String: req.Story, Valid: true}
	data.Email = sql.NullString{String: req.Email, Valid: true}
}
