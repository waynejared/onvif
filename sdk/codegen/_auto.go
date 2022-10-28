// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package 

import (
	"context"
	"github.com/juju/errors"
	"github.com/waynejared/onvif"
	"github.com/waynejared/onvif/sdk"
	"github.com/waynejared/onvif/"
)

// Call_ forwards the call to dev.CallMethod() then parses the payload of the reply as a Response.
func Call_(ctx context.Context, dev *onvif.Device, request .) (.Response, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			Response .Response
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.Response, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "")
		return reply.Body.Response, errors.Annotate(err, "reply")
	}
}
