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

// Call_GetRecordingSummary forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRecordingSummaryResponse.
func Call_GetRecordingSummary(ctx context.Context, dev *onvif.Device, request search.GetRecordingSummary) (search.GetRecordingSummaryResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRecordingSummaryResponse search.GetRecordingSummaryResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRecordingSummaryResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRecordingSummary")
		return reply.Body.GetRecordingSummaryResponse, errors.Annotate(err, "reply")
	}
}
