package search

import (
	"github.com/waynejared/onvif/event"
	"github.com/waynejared/onvif/xsd"
	"github.com/waynejared/onvif/xsd/onvif"
)

type Capabilities struct {
	MetadataSearch     xsd.Boolean `xml:"MetadataSearch,attr"`
	GeneralStartEvents xsd.Boolean `xml:"GeneralStartEvents,attr"`
}

type GetServiceCapabilities struct {
	XMLName string `xml:"tse:GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type GetRecordingSummary struct {
	XMLName string `xml:"tse:GetRecordingSummary"`
}
type GetRecordingSummaryResponse struct {
	Summary onvif.RecordingSummary
}
type GetRecordingInformation struct {
	RecordingToken onvif.RecordingReference `xml:"RecordingReference"`
}
type GetRecordingInformationResponse struct {
	RecordingInformation onvif.RecordingInformation `xml:"RecordingInformation"`
}
type GetMediaAttributes struct {
}
type GetMediaAttributesResponse struct {
}
type FindRecordings struct {
	Scope         xsd.String //onvif.SearchScope
	MaxMatches    xsd.Int
	KeepAliveTime xsd.Duration
}
type FindRecordingsResponse struct {
	SearchToken xsd.Token
}
type GetRecordingSearchResults struct {
	SearchToken xsd.Token
	MinResults  xsd.Int
	MaxResults  xsd.Int
	WaitTime    xsd.Duration
}
type GetRecordingSearchResultsResponse struct {
	ResultList FindRecordingResultList `xml:"ResultList"`
}

type FindRecordingResultList struct {
	SearchState          xsd.String                   `xml:"SearchState"`
	RecordingInformation []onvif.RecordingInformation `xml:"RecordingInformation"`
}

// Events
type FindEvents struct {
	StartPoint       xsd.DateTime
	EndPoint         xsd.DateTime
	Scope            onvif.Scope
	SearchFilter     event.FilterType //wsdl specifies type EventFilter - I think this is right and it's an xPath expresion.
	IncludeStartDate xsd.Boolean
	MaxMatches       int
	KeepAliveTime    xsd.Duration
}
type FindEventsResponse struct {
	SearchToken xsd.Token
}
type GetEventSearchResults struct {
	SearchToken xsd.Token
	MinResults  int
	MaxResults  int
	WaitTime    xsd.Duration
}
type GetEventSearchResultsResponse struct {
	ResultList FindEventResultList
}

type FindEventResultList struct {
	SearchState State `xml:"tt:SearchState"`
	Result      []FindEventResult
}

type FindEventResult struct {
	RecordingToken  onvif.RecordingReference            `xml:"RecordingToken"`
	TrackToken      onvif.TrackReference                `xml:"TrackToken"`
	Time            xsd.DateTime                        `xml:"Time"`
	Event           event.NotificationMessageHolderType `xml:"Event"`
	StartStateEvent xsd.Boolean
}

type State xsd.String

type FindPTZPosition struct {
	StartPoint    onvif.DateTime
	EndPoint      onvif.DateTime
	Scope         onvif.Scope
	SearchFilter  event.FilterType
	MaxMatches    int
	KeepAliveTime xsd.Duration
}
type FindPTZPositionResponse struct {
	SearchToken xsd.Token
}
type GetPTZPositionSearchResults struct {
}
type GetPTZPositionSearchResultsResponse struct {
}
type GetSearchState struct {
	SearchToken xsd.Token
}
type GetSearchStateResponse struct {
}
type EndSearch struct {
	SearchToken xsd.Token
}

type EndSearchResponse struct {
	Endpoint onvif.DateTime
}
type FindMetaData struct {
	StartPoint     onvif.DateTime
	EndPoint       onvif.DateTime
	Scope          onvif.Scope
	MetaDataFilter onvif.FilterType
	MaxMatches     int
	KeepAliveTime  xsd.Duration
}
type FindMetaDataResponse struct {
	SearchToken xsd.Token
}
type GetMetaDataSearchResults struct {
}
type GetMetaDataSearchResultsResponse struct {
	SearchToken xsd.Token
}
