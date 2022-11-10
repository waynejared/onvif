package main

import (
	"testing"

	ssearch "github.com/waynejared/onvif/sdk/search"
	"github.com/waynejared/onvif/search"
)

//Preparing commands
//	systemDateAndTyme := device.GetSystemDateAndTime{}
//	getCapabilities := device.GetCapabilities{Category: "All"}

func TestGetServiceCapabilities(t *testing.T) {
	//Need to connect then call search
	//mimic the event test
	camera := New()
	SearchURI := "http://192.168.85.158/onvif/search_service"
	GSC := search.GetServiceCapabilities{
		XMLName: `xml:"tse:GetServiceCapabilities"`,
	}

	ServiceCapabilities, err := ssearch.Call_GetServiceCapabilities(ctx, camera)
	require.NoError(t, err)
}
