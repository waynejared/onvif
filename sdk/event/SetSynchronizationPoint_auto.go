// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package event

import (
	"context"
	"github.com/juju/errors"
	"github.com/waynejared/onvif"
	"github.com/waynejared/onvif/sdk"
	"github.com/waynejared/onvif/event"
)

// Call_SetSynchronizationPoint forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSynchronizationPointResponse.
func Call_SetSynchronizationPoint(ctx context.Context, dev *onvif.Device, request event.SetSynchronizationPoint) (event.SetSynchronizationPointResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSynchronizationPointResponse event.SetSynchronizationPointResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetSynchronizationPointResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetSynchronizationPoint")
		return reply.Body.SetSynchronizationPointResponse, errors.Annotate(err, "reply")
	}
}
