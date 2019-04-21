package transport

import (
	"context"
	"encoding/json"

	"github.com/spaceuptech/space-api-go/api/client"
	"github.com/spaceuptech/space-api-go/api/model"
)

// Aggr triggers the gRPC aggr function on space cloud
func (t *Transport) Aggr(ctx context.Context, meta *client.Meta, op string, pipeline interface{}) (*model.Response, error) {
	pipelineJSON, err := json.Marshal(pipeline)
	if err != nil {
		return nil, err
	}

	req := client.AggregateRequest{Pipeline: pipelineJSON, Meta: meta, Operation: op}
	res, err := t.stub.Aggregate(ctx, &req)
	if err != nil {
		return nil, err
	}

	if res.Status >= 200 || res.Status < 300 {
		return &model.Response{Status: int(res.Status), Data: res.Result}, nil
	}

	return &model.Response{Status: int(res.Status), Error: res.Error}, nil
}
