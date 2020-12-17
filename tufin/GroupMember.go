package tufin

type GroupMember struct {
  Groupname string
  IPAddress string
}

type GroupMemberResponse struct {
  Status  string `json:"status"`
  Reason  string `json:"reason,omitempty"`
  Date    string `json:"date"`
  Version string `json:"version"`
}
