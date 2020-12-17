package tufinclient

// SecureChangeNetworkObjectsResult represents one or more network objects returned from API
//type SecureChangeNetworkObjectsResult struct {
//	NetworkObjects SecureChangeNetworkObjects `json:"network_objects"`
//}

// SecureChangeCreateDeviceGroupTicket represents a ticket to create a group in SecureChange
type SecureChangeCreateDeviceGroupTicket struct {
	Ticket SecureChangeTicket `json:"ticket"`
}

// SecureChangeTicket represents a ticket within SecureChange
type SecureChangeTicket struct {
	Priority string               `json:"priority"`
	Steps    SecureChangeSteps    `json:"steps"`
	Subject  string               `json:"subject"`
	Workflow SecureChangeWorkflow `json:"workflow"`
}

// SecureChangeWorkflow represents the workflow field in a SecureChange ticket
type SecureChangeWorkflow struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// SecureChangeSteps represents the steps in a SecureChange ticket
type SecureChangeSteps struct {
	Step []SecureChangeStep `json:"step"`
}

// SecureChangeStep represents a single step in a SecureChange ticket
type SecureChangeStep struct {
	Name  string            `json:"name"`
	Tasks SecureChangeTasks `json:"tasks"`
}

// SecureChangeTasks represents tasks in a SecureChange ticket
type SecureChangeTasks struct {
	Task []SecureChangeTask `json:"task"`
}

// SecureChangeTask represents a single task in a SecureChange ticket
type SecureChangeTask struct {
	Fields SecureChangeFields `json:"fields"`
}

// SecureChangeFields represents fields within a SecureChange ticket
type SecureChangeFields struct {
	Field []SecureChangeField `json:"field"`
}

// SecureChangeField represents a single field within a SecureChange ticket
type SecureChangeField struct {
	XsiType     string                    `json:"@xsi.type"`
	GroupChange []SecureChangeGroupChange `json:"group_change"`
	Name        string                    `json:"name"`
}

// SecureChangeGroupChange represents group_change object within a SecureChange ticket
type SecureChangeGroupChange struct {
	XsiType        string                   `json:"@xsi.type"`
	ChangeAction   string                   `json:"change_action"`
	ManagementID   int64                    `json:"management_id"`
	ManagementName string                   `json:"management_name,omitempty"`
	Members        SecureChangeGroupMembers `json:"members"`
	Name           string                   `json:"name"`
}

// SecureChangeGroupMembers represents a collection of members for a group_change
type SecureChangeGroupMembers struct {
	Member []SecureChangeGroupMember `json:"member"`
}

// SecureChangeGroupMember represents a single member for a group_change
type SecureChangeGroupMember struct {
	Type                string `json:"@type"`
	XsiType             string `json:"@xsi.type"`
	ManagementID        int64  `json:"management_id,omitempty"`
	ManagementName      string `json:"management_name,omitempty"`
	Name                string `json:"name"`
	ObjectDetails       string `json:"object_details"`
	ObjectType          string `json:"object_type"`
	ObjectUpdatedStatus string `json:"object_updated_status,omitempty"`
	Status              string `json:"status"`
}
