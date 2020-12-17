package tufinclient

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// CreateDeviceGroup creates a new network object group via a SecureChange ticket
func (c *SecureChangeClient) CreateDeviceGroup(name string, managementID string) error {
	mgmtID, err := strconv.ParseInt(managementID, 10, 64)
	if err != nil {
		return fmt.Errorf("Could not convert management_id %s to integer", managementID)
	}
	ticket := SecureChangeCreateDeviceGroupTicket{
		Ticket: SecureChangeTicket{
			Priority: "Normal",
			Subject:  "Create Group " + name,
			Workflow: SecureChangeWorkflow{
				ID:   34,
				Name: "Group Change Template",
			},
			Steps: SecureChangeSteps{
				Step: []SecureChangeStep{
					{
						Name: "Submit network object group request",
						Tasks: SecureChangeTasks{
							Task: []SecureChangeTask{
								{
									Fields: SecureChangeFields{
										Field: []SecureChangeField{
											{
												XsiType: "multi_group_change",
												Name:    "Modify network object group",
												GroupChange: []SecureChangeGroupChange{
													{
														XsiType:      "group_change",
														Name:         name,
														ChangeAction: "CREATE",
														ManagementID: mgmtID,
														Members: SecureChangeGroupMembers{
															Member: []SecureChangeGroupMember{
																{
																	Name:          "169.254.255.255",
																	XsiType:       "groupMemberNetworkObjectDTO",
																	Type:          "Host",
																	ObjectDetails: "169.254.255.255/255.255.255.255",
																	ObjectType:    "Host",
																	Status:        "ADDED",
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	response, err := c.R().
		SetBody(ticket).
		Post("/securechange/tickets.json")
	if err != nil {
		log.Fatal(err)
	}

	switch response.StatusCode() {
	case 201:
		return nil
	case 400:
		matched, _ := regexp.MatchString(fmt.Sprintf(`object with name %s already exists`, name), response.String())
		if matched {
			return nil
		}
		return fmt.Errorf("%s", response.String())
	default:
		return fmt.Errorf("%s", response.String())
	}
}

// AddMemberToDeviceGroup adds a SecureChangeGroupMember to an existing group on a device
func (c *SecureChangeClient) AddMemberToDeviceGroup(member *SecureChangeGroupMember, group string, managementID string) error {
	mgmtID, err := strconv.ParseInt(managementID, 10, 64)
	if err != nil {
		return fmt.Errorf("Could not convert management_id %s to integer", managementID)
	}
	ticket := SecureChangeCreateDeviceGroupTicket{
		Ticket: SecureChangeTicket{
			Priority: "Normal",
			Subject:  "Add member " + member.Name + " to Group " + group,
			Workflow: SecureChangeWorkflow{
				ID:   34,
				Name: "Group Change Template",
			},
			Steps: SecureChangeSteps{
				Step: []SecureChangeStep{
					{
						Name: "Submit network object group request",
						Tasks: SecureChangeTasks{
							Task: []SecureChangeTask{
								{
									Fields: SecureChangeFields{
										Field: []SecureChangeField{
											{
												XsiType: "multi_group_change",
												Name:    "Modify network object group",
												GroupChange: []SecureChangeGroupChange{
													{
														XsiType:      "group_change",
														Name:         group,
														ChangeAction: "UPDATE",
														ManagementID: mgmtID,
														Members: SecureChangeGroupMembers{
															Member: []SecureChangeGroupMember{
																*member,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	response, err := c.R().
		SetBody(ticket).
		Post("/securechange/tickets.json")
	if err != nil {
		log.Fatal(err)
	}

	switch response.StatusCode() {
	case 201:
		return nil
	case 400:
		matched, _ := regexp.MatchString(fmt.Sprintf(`object with name %s in group %s already exists`, member.Name, group), response.String())
		if matched {
			return nil
		}
		return fmt.Errorf("%s", response.String())
	default:
		return fmt.Errorf("%s", response.String())
	}
}

// RemoveMemberFromDeviceGroup removes a SecureChangeGroupMember from an existing group on a device
func (c *SecureChangeClient) RemoveMemberFromDeviceGroup(member *SecureChangeGroupMember, group string, managementID string) error {
	mgmtID, err := strconv.ParseInt(managementID, 10, 64)
	if err != nil {
		return fmt.Errorf("Could not convert management_id %s to integer", managementID)
	}
	ticket := SecureChangeCreateDeviceGroupTicket{
		Ticket: SecureChangeTicket{
			Priority: "Normal",
			Subject:  "Remove Member " + member.Name + " from Group " + group,
			Workflow: SecureChangeWorkflow{
				ID:   34,
				Name: "Group Change Template",
			},
			Steps: SecureChangeSteps{
				Step: []SecureChangeStep{
					{
						Name: "Submit network object group request",
						Tasks: SecureChangeTasks{
							Task: []SecureChangeTask{
								{
									Fields: SecureChangeFields{
										Field: []SecureChangeField{
											{
												XsiType: "multi_group_change",
												Name:    "Modify network object group",
												GroupChange: []SecureChangeGroupChange{
													{
														XsiType:      "group_change",
														Name:         group,
														ChangeAction: "UPDATE",
														ManagementID: mgmtID,
														Members: SecureChangeGroupMembers{
															Member: []SecureChangeGroupMember{
																*member,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	response, err := c.R().
		SetBody(ticket).
		Post("/securechange/tickets.json")
	if err != nil {
		log.Fatal(err)
	}

	switch response.StatusCode() {
	case 201:
		return nil
	case 400:
		matched, _ := regexp.MatchString(fmt.Sprintf(`in group %s does not exist`, group), response.String())
		if matched {
			return nil
		}
		return fmt.Errorf("%s", response.String())
	default:
		return fmt.Errorf("%s", response.String())
	}
}
