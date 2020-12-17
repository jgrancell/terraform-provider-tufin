package tufinclient

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

// TufinClient represents Tufin API client
type TufinClient struct {
	SecureTrack  SecureTrackClient
	SecureChange SecureChangeClient
}

//SecureTrackClient represents a client for the SecureTrack API
type SecureTrackClient struct {
	*resty.Client
}

// SecureChangeClient represents a client for the SecureChange API
type SecureChangeClient struct {
	*resty.Client
}

// NewTufinClient instantiates REST client for Tufin
func NewTufinClient(secureChangeHost string, secureTrackHost string, username string, password string, insecure bool, debug bool) *TufinClient {
	secureChangeClient := resty.New()
	secureTrackClient := resty.New()

	if debug {
		secureChangeClient.SetDebug(true)
		secureTrackClient.SetDebug(true)
	}

	secureChangeClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: insecure})
	secureTrackClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: insecure})

	secureChangeClient.SetHostURL(fmt.Sprintf("https://%s/securechangeworkflow/api", secureChangeHost))
	secureTrackClient.SetHostURL(fmt.Sprintf("https://%s/securetrack/api", secureChangeHost))

	secureTrackClient.SetBasicAuth(username, password)
	secureChangeClient.SetBasicAuth(username, password)

	return &TufinClient{
		SecureTrack:  SecureTrackClient{secureTrackClient},
		SecureChange: SecureChangeClient{secureChangeClient},
	}
}

// AddIPToGroup adds an IP to a group
func (t *TufinClient) AddIPToGroup(ip string, group string) (bool, error) {
	objs, err := t.SecureTrack.GetNetworkObjectsByName(group)
	added := false
	if err != nil {
		return added, err
	}
	// Avoid nil objs as this means the FW group does not exist
	if objs == nil {
		return added, nil
	}
	for _, obj := range *objs {
		// DisplayName check accounts for incorrectly cased results coming back from exact_match object search
		if obj.DisplayName != group {
			continue
		}
		deviceID := strconv.FormatInt(obj.DeviceID, 10)
		for _, member := range obj.Member {
			if member.Name == ip {
				continue
			}
			obj, err := t.SecureTrack.GetDeviceNetworkObjectByName(ip, deviceID, true)
			if err != nil {
				return added, err
			}
			member := SecureChangeGroupMember{
				Name:          ip,
				XsiType:       "groupMemberNetworkObjectDTO",
				ObjectDetails: ip + "/255.255.255.255",
				ObjectType:    "Host",
				Status:        "ADDED",
			}
			if obj == nil {
				member.Type = "Host"
				member.ObjectUpdatedStatus = "NEW"
			} else {
				member.Type = "Object"
				member.ObjectUpdatedStatus = "EXISTING_NOT_EDITED"
				member.ManagementID = obj.DeviceID
			}
			err = t.SecureChange.AddMemberToDeviceGroup(&member, group, deviceID)
			if err != nil {
				return added, err
			}
			added = true
		}
	}
	return added, nil
}

// RemoveIPFromGroup removes an IP from a group
func (t *TufinClient) RemoveIPFromGroup(ip string, group string) (bool, error) {
	objs, err := t.SecureTrack.GetNetworkObjectsByName(group)
	removed := false
	if err != nil {
		return removed, err
	}
	// Avoid nil objs as this means the FW group does not exist
	if objs == nil {
		return removed, nil
	}
	for _, obj := range *objs {
		// DisplayName check accounts for incorrectly cased results coming back from exact_match object search
		if obj.DisplayName != group {
			continue
		}
		for _, member := range obj.Member {
			if member.DisplayName == ip {
				memberToDelete := SecureChangeGroupMember{
					XsiType:      "groupMemberNetworkObjectDTO",
					Type:         "Object",
					ManagementID: obj.DeviceID,
					Status:       "DELETED",
					Name:         ip,
					ObjectType:   "Host",
				}
				err := t.SecureChange.RemoveMemberFromDeviceGroup(&memberToDelete, group, strconv.FormatInt(obj.DeviceID, 10))
				if err != nil {
					return removed, err
				}
				removed = true
			}
		}
	}
	return removed, nil
}
