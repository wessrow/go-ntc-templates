package models

type CiscoS300ShowIpInterface struct {
	Ip                          string `json:"IP"`
	Interface                   string `json:"INTERFACE"`
	Interface_status_admin_oper string `json:"INTERFACE_STATUS_ADMIN_OPER"`
	Type                        string `json:"TYPE"`
	Status                      string `json:"STATUS"`
}
