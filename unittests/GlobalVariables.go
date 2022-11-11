package unittests //package is main so it can be run with "go run unittests/GlobalVariables.go" from terminal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/waynejared/onvif"
	xonvif "github.com/waynejared/onvif/xsd/onvif"
)

const (
	TopicMotionAlarm = "tns1:VideoSource/MotionAlarm"
	DeviceService    = "tds:GetServices"
	//SearchServiceURI = "http://192.168.85.158:80/onvif/search_service"
)

var (
	ErrFailedNew             = errors.New("failed to set new device")
	ErrNoURIFrame            = errors.New("failed to get URI for snapshot")
	ErrNoURIStream           = errors.New("failed to get URI for stream")
	ErrSubscribe             = errors.New("failed to subscribe")
	ErrUnmarshalEventMessage = errors.New("failed to unmarshal event message")
)

var (
	EventChanged     = "Changed"
	EventInitialized = "Initialized"
)

type Onvifcam struct {
	d                  *onvif.Device
	addr               string
	username, password string
	mainProfile        xonvif.ReferenceToken
	httpClient         *http.Client
}

type EnvVars struct {
	ip   string
	user string
	pass string
	cni  string
	csc  string
	cmp  xonvif.ReferenceToken
}

func GetEnvValues() EnvVars {
	//This function just gets the variables out of the .env file so you can check them.
	//Use the ..._test.go file to run the first test that calls this.
	var MyEnvVars EnvVars

	//Force the env vars to get set
	godotenv.Load("../.env")

	//Query all env vars and store in struct
	MyEnvVars.ip = os.Getenv("CAMERA_IP_ADDRESS") //this is how you reference variables in .env
	MyEnvVars.user = os.Getenv("CAMERA_USERNAME")
	MyEnvVars.pass = os.Getenv("CAMERA_PASSWORD")
	MyEnvVars.cni = os.Getenv("CAMERA_NETWORK_INTERFACE")
	MyEnvVars.csc = os.Getenv("CAMERA_SUPPORT_COUNT")
	MyEnvVars.cmp = xonvif.ReferenceToken(os.Getenv("CAMERA_MAIN_PROFILE"))

	return MyEnvVars
}

// Connect to a camera
// New returns a new bare ONVIF device using basic authentication.
// httpClient is used also by the ONVIF device implementation. It is set to a default client if not provided.
func New(httpClient *http.Client) *Onvifcam {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	MyEnvVars := GetEnvValues()

	return &Onvifcam{
		d:           nil,
		addr:        MyEnvVars.ip,
		username:    MyEnvVars.user,
		password:    MyEnvVars.pass,
		mainProfile: MyEnvVars.cmp,
		httpClient:  httpClient,
	}
}

// Init connects to the device using basic authentication and sets it up.
// A context is currently unused (can be set to nil) but passed for future improvement of go-onvif module.
func (c *Onvifcam) Init(_ context.Context) error {
	d, err := onvif.NewDevice(onvif.DeviceParams{Xaddr: c.addr, Username: c.username, Password: c.password, HttpClient: c.httpClient})
	if err != nil {
		return fmt.Errorf("%s: %w", ErrFailedNew, err)
	}

	c.d = d

	return nil
}
