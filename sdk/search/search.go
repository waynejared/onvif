package search

//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetServiceCapabilities
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetRecordingSummary
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetRecordingInformation
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetMediaAttributes
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search FindRecordings
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetRecordingSearchResults
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search FindEvents
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetEventSearchResults
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search FindPTZPosition
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetPTZPositionSearchResults
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetSearchState
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search EndSearch
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search FindMetaData
//go:generate go run github.com/waynejared/onvif/sdk/codegen search search GetMetaDataSearchResults
