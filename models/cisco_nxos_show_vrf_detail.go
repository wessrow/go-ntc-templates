package models

type CiscoNxosShowVrfDetail struct {
	Name	string	`json:"NAME"`
	Id	string	`json:"ID"`
	Vpn_id	string	`json:"VPN_ID"`
	State	string	`json:"STATE"`
	Default_rd	string	`json:"DEFAULT_RD"`
	Max_routes	string	`json:"MAX_ROUTES"`
	Mid_threshold	string	`json:"MID_THRESHOLD"`
	Table_id	[]string	`json:"TABLE_ID"`
	Address_family	[]string	`json:"ADDRESS_FAMILY"`
	Fwd_id	[]string	`json:"FWD_ID"`
	Table_status	[]string	`json:"TABLE_STATUS"`
}