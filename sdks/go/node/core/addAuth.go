package core

import (
	"context"
	"time"

	"github.com/opctl/opctl/sdks/go/model"
)

func (this core) AddAuth(
	ctx context.Context,
	req model.AddAuthReq,
) error {
	// killing an op is async
	this.pubSub.Publish(
		model.Event{
			AuthAdded: &model.AuthAdded{
				Auth: model.Auth{
					Creds:     req.Creds,
					Resources: req.Resources,
				},
			},
			Timestamp: time.Now().UTC(),
		},
	)
	return nil
}
