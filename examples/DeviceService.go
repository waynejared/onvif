package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	goonvif "github.com/waynejared/onvif"
	"github.com/waynejared/onvif/device"
	sdk "github.com/waynejared/onvif/sdk/device"
	"github.com/waynejared/onvif/xsd/onvif"
)

const (
	login     = "root"
	password  = "C@nadien3"
	ipAddress = "192.168.85.158:80"
)

func main() {
	ctx := context.Background()

	//Getting an camera instance
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      "192.168.85.158:80",
		Username:   login,
		Password:   password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		panic(err)
	}

	//Preparing commands
	systemDateAndTyme := device.GetSystemDateAndTime{}
	getCapabilities := device.GetCapabilities{Category: "All"}
	createUser := device.CreateUsers{
		User: onvif.User{
			Username:  "TestUser",
			Password:  "TestPassword",
			UserLevel: "User",
		},
	}

	//Commands execution
	systemDateAndTymeResponse, err := sdk.Call_GetSystemDateAndTime(ctx, dev, systemDateAndTyme)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(systemDateAndTymeResponse)
	}
	getCapabilitiesResponse, err := sdk.Call_GetCapabilities(ctx, dev, getCapabilities)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCapabilitiesResponse)
	}

	createUserResponse, err := sdk.Call_CreateUsers(ctx, dev, createUser)
	if err != nil {
		log.Println(err)
	} else {
		// You could use https://github.com/waynejared/onvif/gosoap for pretty printing response
		fmt.Println(createUserResponse)
	}

}
