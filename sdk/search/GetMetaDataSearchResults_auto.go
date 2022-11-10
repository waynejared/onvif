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

// Call_GetMetaDataSearchResults forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetaDataSearchResultsResponse.
func Call_GetMetaDataSearchResults(ctx context.Context, dev *onvif.Device, request search.GetMetaDataSearchResults) (search.GetMetaDataSearchResultsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetaDataSearchResultsResponse search.GetMetaDataSearchResultsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetaDataSearchResultsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetMetaDataSearchResults")
		return reply.Body.GetMetaDataSearchResultsResponse, errors.Annotate(err, "reply")
	}
}
