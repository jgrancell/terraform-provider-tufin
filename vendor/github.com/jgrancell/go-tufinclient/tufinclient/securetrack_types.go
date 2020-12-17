package tufinclient

// SecureTrackNetworkObjectsResult represents one or more network objects returned from the API
type SecureTrackNetworkObjectsResult struct {
	NetworkObjects SecureTrackNetworkObjects `json:"network_objects"`
}

// SecureTrackNetworkObjects represents a collection of network objects in SecureTrack
type SecureTrackNetworkObjects struct {
	Count         int64                      `json:"count"`
	NetworkObject []SecureTrackNetworkObject `json:"network_object"`
	Total         int64                      `json:"total"`
}

// SecureTrackNetworkObject represents a single SecureTrack network object
type SecureTrackNetworkObject struct {
	XsiType     string                           `json:"@xsi.type"`
	ClassName   string                           `json:"class_name"`
	Comment     string                           `json:"comment"`
	DeviceID    int64                            `json:"device_id"`
	DisplayName string                           `json:"display_name"`
	Exclusion   []interface{}                    `json:"exclusion,omitempty"`
	Global      bool                             `json:"global,omitempty"`
	ID          string                           `json:"id"`
	Implicit    bool                             `json:"implicit"`
	IP          string                           `json:"ip,omitempty"`
	IPType      string                           `json:"ip_type"`
	Name        string                           `json:"name"`
	Member      []SecureTrackNetworkObjectMember `json:"member,omitempty"`
	Netmask     string                           `json:"netmask,omitempty"`
	Overrides   bool                             `json:"overrides"`
	Type        string                           `json:"type"`
	UID         string                           `json:"uid"`
}

// SecureTrackNetworkObjectMember represents a member of a SecureTrack network object
type SecureTrackNetworkObjectMember struct {
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	UID         string `json:"uid"`
}

// SecureTrackDevicesResult represents one of more devices returned from the API
type SecureTrackDevicesResult struct {
	Devices SecureTrackDevices `json:"devices"`
}

// SecureTrackDevice represents data about a device in SecureTrack
type SecureTrackDevice struct {
	ContextName    string `json:"context_name"`
	DomainID       string `json:"domain_id"`
	DomainName     string `json:"domain_name"`
	ID             string `json:"id"`
	IP             string `json:"ip"`
	LatestRevision string `json:"latest_revision"`
	Model          string `json:"model"`
	ModuleUID      string `json:"module_uid"`
	Name           string `json:"name"`
	Offline        bool   `json:"offline"`
	Topology       bool   `json:"topology"`
	Vendor         string `json:"vendor"`
	VirtualType    string `json:"virtual_type"`
}

// SecureTrackDevices represents a collection of devices in SecureTrack
type SecureTrackDevices struct {
	Count  int64               `json:"count"`
	Device []SecureTrackDevice `json:"device"`
	Total  int64               `json:"total"`
}
