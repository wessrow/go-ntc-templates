package models

type CiscoNxosShowLldpNeighborsDetail struct {
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Mgmt_address	string	`json:"MGMT_ADDRESS"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Capabilities	string	`json:"CAPABILITIES"`
	Vlan_id	string	`json:"VLAN_ID"`
}

var CiscoNxosShowLldpNeighborsDetail_Template = `Value Required NEIGHBOR_NAME (\S+)
Value Required LOCAL_INTERFACE (\S+)
Value Required NEIGHBOR_INTERFACE (\S+)
Value CHASSIS_ID (\w+?\.\w+?\.\w+?)
Value MGMT_ADDRESS (\d+?\.\d+?\.\d+?\.\d+?|\w+?\.\w+?\.\w+?)
Value NEIGHBOR_DESCRIPTION (.*)
Value CAPABILITIES (.*)
Value VLAN_ID ([0-9]+)


Start
  ^[Cc]hassis\s[idID]{2}\:\s${CHASSIS_ID}$$
  ^[Pp]ort\s[idID]{2}\:\s${NEIGHBOR_INTERFACE}$$
  ^[Ll]ocal\s[Pp]ort\s[idID]{2}\:\s${LOCAL_INTERFACE}$$
  ^[Ss]ystem\s[Nn]ame\:\s${NEIGHBOR_NAME}$$
  ^[Ss]ystem\s[Dd]escription\:\s${NEIGHBOR_DESCRIPTION}$$
  ^[Ee]nabled\s[Cc]apabilities\:\s${CAPABILITIES}$$
  ^[Mm]anagement\s[Aa]ddress\:\s${MGMT_ADDRESS}$$
  ^[Vv]lan\s[idID]{2}\:\s${VLAN_ID}$$ -> Record
  ^[Vv]lan\s[idID]{2}\:\snot advertised -> Record

`