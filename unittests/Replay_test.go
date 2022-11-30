package unittests

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/waynejared/onvif/replay"
	"github.com/waynejared/onvif/xsd/onvif"
	"golift.io/ffmpeg"

	sdk_replay "github.com/waynejared/onvif/sdk/replay"
)

func TestGetReplayServiceCapabilities(t *testing.T) {
	//Get and check the service capabilities for search
	ctx := context.TODO()
	camera := New(&http.Client{})
	err := camera.Init(ctx)
	require.NoError(t, err)
	require.NotNil(t, camera.d)

	GSC := replay.GetServiceCapabilities{
		XMLName: `xml:"trp:GetServiceCapabilities"`,
	}

	ServiceCapabilities, err := sdk_replay.Call_GetServiceCapabilities(ctx, camera.d, GSC)
	require.NoError(t, err)
	require.NotNil(t, ServiceCapabilities)
	if err == nil {
		log.Printf("\nReversePlayback: %v\nSessionTimeoutRange: %v\nRTP_RTSP_TCP: %v\nRTSPWebSocketURI: %v\n", ServiceCapabilities.Capabilities.ReversePlayback, ServiceCapabilities.Capabilities.SessionTimeoutRange, ServiceCapabilities.Capabilities.RTP_RTSP_TCP, ServiceCapabilities.Capabilities.RTSPWebSocketUri)
	} else {
		log.Printf("\nError: %v", err)
	}
}

func TestGetReplayURI(t *testing.T) {
	ctx := context.TODO()
	camera := New(&http.Client{})
	err := camera.Init(ctx)
	require.NoError(t, err)
	require.NotNil(t, camera.d)

	ReplayURI := replay.GetReplayUri{
		StreamSetup: onvif.StreamSetup{
			Stream: "RTP-Unicast",
			Transport: onvif.Transport{
				Protocol: "RTSP",
			},
		},
		RecordingToken: "21393060360015590418",
	}

	ReplayURIResp, err := sdk_replay.Call_GetReplayUri(ctx, camera.d, ReplayURI)
	if err == nil {
		log.Print(ReplayURIResp)
	} else {
		log.Printf("\nError: %v", err)
	}
	MyEnvVars := GetEnvValues()

	ffmpeg.DefaultFFmpegPath = "C:/Program Files/ffmpeg/ffmpeg.exe"
	dahua := string(ReplayURIResp.Uri)[0:7] + MyEnvVars.user + ":" + MyEnvVars.pass + "@" + string(ReplayURIResp.Uri)[7:]
	output := "c:/temp/RecordedVideo.m4v"
	encode := ffmpeg.Get(&ffmpeg.Config{
		Audio:  true, // retain audio stream
		Time:   10,   // 10 seconds
		Width:  1920,
		Height: 1080,
		CRF:    23,
		Level:  "4.0",
		Rate:   5,
		Prof:   "baseline", // or main or high
	})

	cmd, out, err := encode.SaveVideo(dahua, output, "DahuaVideoTitle")

	log.Println("Command Used:", cmd)
	log.Println("Command Output:", out)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Saved file from", dahua, "to", output)

}
