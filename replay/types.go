package replay

import (
	"github.com/waynejared/onvif/xsd"
	"github.com/waynejared/onvif/xsd/onvif"
)

type GetReplayConfiguration struct {
	XMLName xsd.String `xml:"trp:GetReplayConfiguration"`
}

type GetReplayConfigurationResponse struct {
	Configuration ReplayConfiguration
}

type ReplayConfiguration struct {
	SessionTimeout xsd.Duration
}

type GetReplayUri struct {
	StreamSetup    onvif.StreamSetup
	RecordingToken xsd.Token
}

type GetReplayUriResponse struct {
	Uri xsd.AnyURI
}

type GetServiceCapabilities struct {
	XMLName xsd.String `xml:"trp:GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type Capabilities struct {
	ReversePlayback     xsd.Boolean `xml:"ReversePlayback,attr"`
	SessionTimeoutRange xsd.String  `xml:"SessionTimeoutRange,attr"`
	RTP_RTSP_TCP        xsd.Boolean `xml:"RTP_RTSP_TCP,attr"`
	RTSPWebSocketUri    xsd.AnyURI  `xml:"RTSPWebSocketURI,attr"`
}

type SetReplayConfiguration struct {
	Configuration ReplayConfiguration
}
