package story

import (
	"api-web-bangarief/internal/constants"
	"api-web-bangarief/internal/errors"
	"api-web-bangarief/internal/story/request"
	"api-web-bangarief/pkg/log"
	"net/http"
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/stories", res.query)
	r.Post("/stories", res.create)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	var err error

	var ID int64
	queryID := c.Query("id")
	if queryID != "" {
		ID, err = strconv.ParseInt(queryID, 10, 64)
		if err != nil {
			r.logger.With(ctx).Info(err)
			return errors.BadRequest("")
		}
	}

	limit := constants.DefaultStoryLimit
	queryLimit := c.Query("limit")
	if queryLimit != "" {
		limit, err = strconv.ParseInt(queryLimit, 10, 64)
		if err != nil {
			r.logger.With(ctx).Info(err)
			return errors.BadRequest("")
		}
	}

	response, err := r.service.Query(ctx, ID, limit)
	if err != nil {
		return err
	}

	return c.Write(response)
}

func (r resource) create(c *routing.Context) error {
	ctx := c.Request.Context()

	var request request.Create
	if err := c.Read(&request); err != nil {
		r.logger.With(ctx).Info(err)
		return errors.BadRequest("")
	}

	if err := request.Validate(); err != nil {
		r.logger.With(ctx).Info(err)
		return err
	}

	response, err := r.service.Create(ctx, request)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(response, http.StatusCreated)
}
