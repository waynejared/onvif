package unittests

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	sdk_search "github.com/waynejared/onvif/sdk/search"
	"github.com/waynejared/onvif/search"
)

func TestGetServiceCapabilities(t *testing.T) {
	//Get and check the service capabilities for search
	camera := New(&http.Client{})
	ctx := context.TODO()
	err := camera.Init(ctx)
	require.NoError(t, err)
	require.NotNil(t, camera.d)

	//SearchURI := "http://192.168.85.158/onvif/search_service"
	GSC := search.GetServiceCapabilities{
		XMLName: `xml:"tse:GetServiceCapabilities"`,
	}

	ServiceCapabilities, err := sdk_search.Call_GetServiceCapabilities(ctx, camera.d, GSC)
	require.NoError(t, err)
	require.NotNil(t, ServiceCapabilities)
	if err == nil {
		log.Printf("\nGeneralStartEvents: %v\nMetadataSearch: %v", ServiceCapabilities.Capabilities.GeneralStartEvents, ServiceCapabilities.Capabilities.MetadataSearch)
	} else {
		log.Printf("\nError: %v", err)
	}
}
