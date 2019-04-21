package transport

import (
	"context"
	"encoding/json"

	"github.com/spaceuptech/space-api-go/api/client"
	"github.com/spaceuptech/space-api-go/api/model"
	"github.com/spaceuptech/space-api-go/api/utils"
)

// Call triggers the gRPC call function on space cloud
func (t *Transport) Call(ctx context.Context, token, engine, function string, params utils.M, timeout int) (*model.Response, error) {
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req := client.FaaSRequest{Params: paramsJSON, Timeout: int64(timeout), Token: token, Engine: engine, Function: function}
	res, err := t.stub.Call(ctx, &req)
	if err != nil {
		return nil, err
	}

	if res.Status >= 200 || res.Status < 300 {
		return &model.Response{Status: int(res.Status), Data: res.Result}, nil
	}

	return &model.Response{Status: int(res.Status), Error: res.Error}, nil
}
