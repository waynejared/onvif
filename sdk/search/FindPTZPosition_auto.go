// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package search

import (
	"context"
	"github.com/juju/errors"
	"github.com/waynejared/onvif"
	"github.com/waynejared/onvif/sdk"
	"github.com/waynejared/onvif/search"
)

// Call_FindPTZPosition forwards the call to dev.CallMethod() then parses the payload of the reply as a FindPTZPositionResponse.
func Call_FindPTZPosition(ctx context.Context, dev *onvif.Device, request search.FindPTZPosition) (search.FindPTZPositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			FindPTZPositionResponse search.FindPTZPositionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.FindPTZPositionResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "FindPTZPosition")
		return reply.Body.FindPTZPositionResponse, errors.Annotate(err, "reply")
	}
}
