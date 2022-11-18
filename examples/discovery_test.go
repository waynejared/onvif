package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/beevik/etree"
	"github.com/waynejared/onvif"
	"github.com/waynejared/onvif/device"
	discover "github.com/waynejared/onvif/ws-discovery"
)

func TestGetAvailableDevicesAtSpecificEthernetInterface(t *testing.T) {
	s, err := onvif.GetAvailableDevicesAtSpecificEthernetInterface("Eth2")
	log.Printf("\nError: %v\nResult:\n%v", err, s)
}

func TestClient(t *testing.T) {
	dev, err := onvif.NewDevice(onvif.DeviceParams{Xaddr: "192.168.85.158", Username: "root", Password: "C@nadien3"})
	if err != nil {
		panic(err)
	}

	log.Printf("output %+v", dev.GetServices())

	res, err := dev.CallMethod(device.GetUsers{})
	bs, _ := ioutil.ReadAll(res.Body)
	log.Printf("output %+v %s", res.StatusCode, bs)
}

// Host host
type Host struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

func runDiscovery(interfaceName string) {
	var hosts []*Host
	devices, err := discover.SendProbe(interfaceName, nil, []string{"dn:NetworkVideoTransmitter"}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	if err != nil {
		log.Printf("error %s", err)
		return
	}
	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			log.Printf("error %s", err)
		} else {

			endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
			scopes := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/Scopes")

			flag := false

			host := &Host{}

			for _, xaddr := range endpoints {
				xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
				host.URL = xaddr
			}
			if flag {
				break
			}
			for _, scope := range scopes {
				re := regexp.MustCompile(`onvif:\/\/www\.onvif\.org\/name\/[A-Za-z0-9-]+`)
				match := re.FindStringSubmatch(scope.Text())
				host.Name = path.Base(match[0])
			}

			hosts = append(hosts, host)

		}

	}

	bys, _ := json.Marshal(hosts)
	log.Printf("done %s", bys)
}
