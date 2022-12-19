package unittests

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	sdk_search "github.com/waynejared/onvif/sdk/search"
	"github.com/waynejared/onvif/search"
	"github.com/waynejared/onvif/xsd"
)

func TestGetSearchServiceCapabilities(t *testing.T) {
	//Get and check the service capabilities for search
	ctx := context.TODO()
	camera := New(&http.Client{})
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

func TestGetRecordingSummary(t *testing.T) {
	ctx := context.TODO()
	camera := New(&http.Client{})
	err := camera.Init(ctx)

	GRS := search.GetRecordingSummary{
		XMLName: `xml:"tse:GetRecordingSummary"`,
	}

	RecordingSummary, err := sdk_search.Call_GetRecordingSummary(ctx, camera.d, GRS)
	if err == nil {
		log.Printf("\nRecordingSummary: %v", RecordingSummary)
	} else {
		log.Printf("\nError: %v", err)
	}

}

func TestFindRecordings(t *testing.T) {
	ctx := context.TODO()
	camera := New(&http.Client{})
	err := camera.Init(ctx)
	require.NoError(t, err)

	SearchFilter := search.FindRecordings{

		KeepAliveTime: "PT5M",
		MaxMatches:    100,
	}

	SessionToken, err := sdk_search.Call_FindRecordings(ctx, camera.d, SearchFilter)
	require.NotNil(t, SessionToken)
	require.NoError(t, err)

	GetRecordingSearchResults := search.GetRecordingSearchResults{
		SearchToken: SessionToken.SearchToken,
		MinResults:  50,
		MaxResults:  50,
		WaitTime:    "PT10M",
	}

	RecordingSearchResults, err := sdk_search.Call_GetRecordingSearchResults(ctx, camera.d, GetRecordingSearchResults)
	require.NoError(t, err)
	log.Println(RecordingSearchResults)

	RIToken := search.GetRecordingInformation{
		RecordingToken: RecordingSearchResults.ResultList.RecordingInformation[0].RecordingToken,
	}
	if RIToken.RecordingToken != "" {
		OneRecording, err := sdk_search.Call_GetRecordingInformation(ctx, camera.d, RIToken)
		log.Println(OneRecording)
		require.NotEmpty(t, OneRecording)
		require.NoError(t, err)
	}

	MyToken := search.GetSearchState{
		SearchToken: SessionToken.SearchToken,
	}

	SearchState, err := sdk_search.Call_GetSearchState(ctx, camera.d, MyToken)
	log.Println(SearchState)
	require.NoError(t, err)
	require.NotEmpty(t, SearchState)

	EndSearchResponse, err := sdk_search.Call_EndSearch(ctx, camera.d, search.EndSearch(SessionToken))
	require.NotNil(t, EndSearchResponse)
	require.NoError(t, err)
}

func TestFindEvents(t *testing.T) {
	/*
	This returns recording events (not actual events)
	Looks like a set of 3 events for 
	*/
	ctx := context.TODO()
	camera := New(&http.Client{})
	err := camera.Init(ctx)
	require.NoError(t, err)
	CurrentDT := time.Now().UTC()
	StartDT := CurrentDT.AddDate(0, 0, -1)

	FindEventsFilter := search.FindEvents{
		StartPoint:    xsd.DateTime(CurrentDT.String()),
		EndPoint:      xsd.DateTime(xsd.DateTime(StartDT.String())),
		MaxMatches:    100,
		KeepAliveTime: "PT10M",
	}
	log.Printf(string(FindEventsFilter.StartPoint))
	log.Printf(string(FindEventsFilter.EndPoint))

	SearchToken, err := sdk_search.Call_FindEvents(ctx, camera.d, FindEventsFilter)
	require.NotNil(t, SearchToken)
	require.NoError(t, err)
	log.Printf("Search Token\n")
	log.Println(SearchToken)

	GetEventSearchResults := search.GetEventSearchResults{
		SearchToken: SearchToken.SearchToken,
		MinResults:  0,
		MaxResults:  100,
		WaitTime:    "PT10M",
	}

	SearchResults, err := sdk_search.Call_GetEventSearchResults(ctx, camera.d, GetEventSearchResults)
	require.NoError(t, err)
	log.Println(SearchResults)

	EndSearchResponse, err := sdk_search.Call_EndSearch(ctx, camera.d, search.EndSearch(SearchToken))
	require.NotNil(t, EndSearchResponse)
	require.NoError(t, err)
}
