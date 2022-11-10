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

// Call_GetMediaAttributes forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMediaAttributesResponse.
func Call_GetMediaAttributes(ctx context.Context, dev *onvif.Device, request search.GetMediaAttributes) (search.GetMediaAttributesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMediaAttributesResponse search.GetMediaAttributesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMediaAttributesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetMediaAttributes")
		return reply.Body.GetMediaAttributesResponse, errors.Annotate(err, "reply")
	}
}
