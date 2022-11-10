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

// Call_GetPTZPositionSearchResults forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPTZPositionSearchResultsResponse.
func Call_GetPTZPositionSearchResults(ctx context.Context, dev *onvif.Device, request search.GetPTZPositionSearchResults) (search.GetPTZPositionSearchResultsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPTZPositionSearchResultsResponse search.GetPTZPositionSearchResultsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPTZPositionSearchResultsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetPTZPositionSearchResults")
		return reply.Body.GetPTZPositionSearchResultsResponse, errors.Annotate(err, "reply")
	}
}
