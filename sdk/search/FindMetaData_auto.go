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

// Call_FindMetaData forwards the call to dev.CallMethod() then parses the payload of the reply as a FindMetaDataResponse.
func Call_FindMetaData(ctx context.Context, dev *onvif.Device, request search.FindMetaData) (search.FindMetaDataResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			FindMetaDataResponse search.FindMetaDataResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.FindMetaDataResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "FindMetaData")
		return reply.Body.FindMetaDataResponse, errors.Annotate(err, "reply")
	}
}
