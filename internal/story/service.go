package story

import (
	"api-web-bangarief/internal/constants"
	"api-web-bangarief/internal/entity"
	"api-web-bangarief/internal/story/request"
	"api-web-bangarief/internal/story/response"
	"api-web-bangarief/pkg/log"
	"context"
	"database/sql"
	"time"
)

// Service ...
type Service interface {
	Query(ctx context.Context, ID, limit int64) (*response.Query, error)
	Create(ctx context.Context, request request.Create) (*response.Create, error)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService ...
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Query(ctx context.Context, ID, limit int64) (*response.Query, error) {
	stories, err := s.repo.Query(ctx, ID, limit)
	if err != nil {
		return nil, err
	}

	var response response.Query
	response.Message = "success.query_stories"
	response.Set(stories)

	return &response, nil
}

func (s service) Create(ctx context.Context, request request.Create) (*response.Create, error) {
	var story entity.Story
	request.Populate(&story)

	story.Status = sql.NullString{String: constants.StorySubmitted, Valid: true}
	story.CreatedOn = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	if err := s.repo.Create(ctx, &story); err != nil {
		return nil, err
	}

	var response response.Create
	response.Message = "success.post_stories"
	response.Set(&story)

	return &response, nil
}
