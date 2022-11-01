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

// Call_Seek forwards the call to dev.CallMethod() then parses the payload of the reply as a SeekResponse.
func Call_Seek(ctx context.Context, dev *onvif.Device, request event.Seek) (event.SeekResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SeekResponse event.SeekResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SeekResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Seek")
		return reply.Body.SeekResponse, errors.Annotate(err, "reply")
	}
}
