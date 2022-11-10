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

// Call_GetSearchState forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSearchStateResponse.
func Call_GetSearchState(ctx context.Context, dev *onvif.Device, request search.GetSearchState) (search.GetSearchStateResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSearchStateResponse search.GetSearchStateResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSearchStateResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetSearchState")
		return reply.Body.GetSearchStateResponse, errors.Annotate(err, "reply")
	}
}
