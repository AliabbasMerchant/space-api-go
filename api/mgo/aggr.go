package mgo

import (
	"context"

	"github.com/spaceuptech/space-api-go/api/client"
	"github.com/spaceuptech/space-api-go/api/config"
	"github.com/spaceuptech/space-api-go/api/model"
)

// Aggr contains the methods for the aggregation operation
type Aggr struct {
	ctx      context.Context
	meta     *client.Meta
	op       string
	pipeline []interface{}
	config   *config.Config
}

func initAggr(ctx context.Context, db, col, op string, config *config.Config) *Aggr {
	m := &client.Meta{Col: col, DbType: db, Project: config.Project, Token: config.Token}
	p := []interface{}{}
	return &Aggr{ctx, m, op, p, config}
}

// Pipe sets the pipeline to run on the backend
func (a *Aggr) Pipe(pipeline []interface{}) *Aggr {
	a.pipeline = pipeline
	return a
}

// Apply executes the operation and returns the result
func (a *Aggr) Apply() (*model.Response, error) {
	return a.config.Transport.Aggr(a.ctx, a.meta, a.op, a.pipeline)
}
