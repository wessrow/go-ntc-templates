package models

type CiscoIosShowVtpStatus struct {
	Version	string	`json:"VERSION"`
	Domain	string	`json:"DOMAIN"`
	Pruning	string	`json:"PRUNING"`
	Traps	string	`json:"TRAPS"`
	Device_id	string	`json:"DEVICE_ID"`
	Last_modified_server	string	`json:"LAST_MODIFIED_SERVER"`
	Last_modified_date	string	`json:"LAST_MODIFIED_DATE"`
	Local_updater_addr	string	`json:"LOCAL_UPDATER_ADDR"`
	Local_updater_iface	string	`json:"LOCAL_UPDATER_IFACE"`
	Mode	string	`json:"MODE"`
	Max_vlans	string	`json:"MAX_VLANS"`
	Existing_vlan_count	string	`json:"EXISTING_VLAN_COUNT"`
	Revision_number	string	`json:"REVISION_NUMBER"`
}