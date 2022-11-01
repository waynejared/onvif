package event

//go:generate go run github.com/waynejared/onvif/sdk/codegen event event GetServiceCapabilities
//go:generate go run github.com/waynejared/onvif/sdk/codegen event event CreatePullPointSubscription
//go:generate go run github.com/waynejared/onvif/sdk/codegen event event GetEventProperties
//go:generate go run github.com/waynejared/onvif/sdk/codegen event event PullMessages
//go:generate go run github.com/waynejared/onvif/sdk/codegen event event Seek
//go:generate go run github.com/waynejared/onvif/sdk/codegen event event SetSynchronizationPoint
//go:generate go run github.com/waynejared/onvif/sdk/codegen event event Subscribe
//go:generate go run github.com/waynejared/onvif/sdk/codegen event event Unsubscribe
//go:generate go run github.com/waynejared/onvif/sdk/codegen event event Renew
