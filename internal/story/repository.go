package story

import (
	"api-web-bangarief/internal/entity"
	"api-web-bangarief/pkg/dbcontext"
	"api-web-bangarief/pkg/log"
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository ...
type Repository interface {
	Create(ctx context.Context, story *entity.Story) error
	Query(ctx context.Context, ID, limit int64) ([]entity.Story, error)
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository ...
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Create(ctx context.Context, story *entity.Story) error {
	return r.db.With(ctx).Model(story).Insert()
}

func (r repository) Query(ctx context.Context, ID, limit int64) ([]entity.Story, error) {
	var data []entity.Story

	query := r.db.With(ctx).
		Select().
		From("story").
		OrderBy("id DESC")

	if ID > 0 {
		query.Where(dbx.NewExp("id < {:ID}", dbx.Params{"ID": ID}))
	}

	if limit > 0 {
		query.Limit(limit)
	}

	err := query.All(&data)
	return data, err
}
