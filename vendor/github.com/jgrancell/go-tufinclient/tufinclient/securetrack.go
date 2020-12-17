package tufinclient

import (
	"fmt"
	"log"
	"net"
)

// GetNetworkObjectsByName searches SecureTrack a network_object by a specified name
func (c *SecureTrackClient) GetNetworkObjectsByName(name string) (*[]SecureTrackNetworkObject, error) {
	response, err := c.R().
		SetResult(&SecureTrackNetworkObjectsResult{}).
		SetQueryParams(map[string]string{
			"filter":      "text",
			"exact_match": "true",
			"name":        name,
		}).
		SetHeader("Accept", "application/json").
		Get("/network_objects/search.json")
	if err != nil {
		log.Fatal(err)
	}

	switch response.StatusCode() {
	case 401:
		return nil, fmt.Errorf("Unauthorized request. Tufin API returned: %s", response.String())
	case 200:
		objs := response.Result().(*SecureTrackNetworkObjectsResult)
		if objs.NetworkObjects.Count == 0 {
			return nil, nil
		}
		return &objs.NetworkObjects.NetworkObject, nil
	default:
		return nil, fmt.Errorf("%s", response.String())
	}
}

// GetDeviceNetworkObjectByName searches a SecureTrack device for a network_object by a specified name
func (c *SecureTrackClient) GetDeviceNetworkObjectByName(name string, deviceID string, caseSensitive bool) (*SecureTrackNetworkObject, error) {
	response, err := c.R().
		SetResult(&SecureTrackNetworkObjectsResult{}).
		SetQueryParams(map[string]string{
			"filter":      "text",
			"exact_match": "true",
			"name":        name,
			"device_id":   deviceID,
		}).
		SetHeader("Accept", "application/json").
		Get("/network_objects/search.json")
	if err != nil {
		log.Fatal(err)
	}

	switch response.StatusCode() {
	case 401:
		return nil, fmt.Errorf("Unauthorized request. Tufin API returned: %s", response.String())
	case 200:
		objs := response.Result().(*SecureTrackNetworkObjectsResult)
		if objs.NetworkObjects.Count == 0 {
			return nil, nil
		}
		if objs.NetworkObjects.Count == 1 {
			if caseSensitive {
				if objs.NetworkObjects.NetworkObject[0].DisplayName == name {
					return &objs.NetworkObjects.NetworkObject[0], nil
				}
				return nil, nil
			}
			return &objs.NetworkObjects.NetworkObject[0], nil
		}
		return nil, fmt.Errorf("Multiple network objects found for device %s", name)
	default:
		return nil, fmt.Errorf("%s", response.String())
	}
}

// GetDevices retrieves all SecureTrack devices
func (c *SecureTrackClient) GetDevices() (*SecureTrackDevicesResult, error) {
	response, err := c.R().
		SetResult(&SecureTrackDevicesResult{}).
		SetQueryParams(map[string]string{
			"start": "0",
			"name":  "999",
		}).
		SetHeader("Accept", "application/json").
		Get("/devices.json")
	if err != nil {
		log.Fatal(err)
	}

	switch response.StatusCode() {
	case 401:
		return new(SecureTrackDevicesResult), fmt.Errorf("Unauthorized request. Tufin API returned: %s", response.String())
	case 200:
		return response.Result().(*SecureTrackDevicesResult), nil
	default:
		return nil, fmt.Errorf("%s", response.String())
	}
}

// GetDevice retrieves a SecureTrack device by ip or name
func (c *SecureTrackClient) GetDevice(str string) (*SecureTrackDevice, error) {
	var query string
	addr := net.ParseIP(str)
	if addr != nil {
		query = "ip"
	} else {
		query = "name"
	}
	response, err := c.R().
		SetResult(&SecureTrackDevicesResult{}).
		SetQueryString(fmt.Sprintf("%s=%s", query, str)).
		SetHeader("Accept", "application/json").
		Get("/devices.json")
	if err != nil {
		log.Fatal(err)
	}

	switch response.StatusCode() {
	case 401:
		return new(SecureTrackDevice), fmt.Errorf("Unauthorized request. Tufin API returned: %s", response.String())
	case 200:
		devices := response.Result().(*SecureTrackDevicesResult)
		if devices.Devices.Count == 0 {
			return nil, nil
		}
		if devices.Devices.Count == 1 {
			return &devices.Devices.Device[0], nil
		}
		return nil, fmt.Errorf("Multiple device objects found for device %s", str)
	default:
		return nil, fmt.Errorf("%s", response.String())
	}
}
