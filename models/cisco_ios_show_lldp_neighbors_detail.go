package models

type CiscoIosShowLldpNeighborsDetail struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_port_id	string	`json:"NEIGHBOR_PORT_ID"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Capabilities	string	`json:"CAPABILITIES"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Vlan_id	string	`json:"VLAN_ID"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Serial	string	`json:"SERIAL"`
	Manufacturer	string	`json:"MANUFACTURER"`
	Hw_revision	string	`json:"HW_REVISION"`
	Fw_revision	string	`json:"FW_REVISION"`
	Sw_revision	string	`json:"SW_REVISION"`
	Platform	string	`json:"PLATFORM"`
	Power_pair	string	`json:"POWER_PAIR"`
	Power_class	string	`json:"POWER_CLASS"`
	Power_device_type	string	`json:"POWER_DEVICE_TYPE"`
	Power_priority	string	`json:"POWER_PRIORITY"`
	Power_source	string	`json:"POWER_SOURCE"`
	Power_requested	string	`json:"POWER_REQUESTED"`
}