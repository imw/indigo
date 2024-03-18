// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package noise

// schema: app.noise.modulation.setActorGain

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// ModulationSetActorGain_Input is the input argument to a app.noise.modulation.setActorGain call.
type ModulationSetActorGain_Input struct {
	Actor string `json:"actor" cborgen:"actor"`
	Gain  int64  `json:"gain" cborgen:"gain"`
}

// ModulationSetActorGain calls the XRPC method "app.noise.modulation.setActorGain".
func ModulationSetActorGain(ctx context.Context, c *xrpc.Client, input *ModulationSetActorGain_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "app.noise.modulation.setActorGain", nil, input, nil); err != nil {
		return err
	}

	return nil
}
