package recording

import (
	"github.com/waynejared/onvif/xsd"
	"github.com/waynejared/onvif/xsd/onvif"
)

type CreateRecording struct {
	RecordingConfiguration RecordingConfiguration
}

type RecordingConfiguration struct {
	Source               RecordingSourceInformation
	Content              onvif.Description
	MaximumRetentionTime xsd.Duration
}

type RecordingSourceInformation struct {
	SourceId    xsd.AnyURI
	Name        onvif.Name
	Location    onvif.Description
	Description onvif.Description
	Address     xsd.AnyURI
}

type CreatRecordingResponse struct {
	RecordingToken xsd.Token
}
