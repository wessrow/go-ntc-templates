package models

type IpinfusionOcnosShowLldpTable struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Capabilities	string	`json:"CAPABILITIES"`
	Neighbor_interface_description	string	`json:"NEIGHBOR_INTERFACE_DESCRIPTION"`
}

var IpinfusionOcnosShowLldpTable_Template = `Value LOCAL_INTERFACE (\S+)
Value NEIGHBOR_NAME (\S+)
Value NEIGHBOR_INTERFACE (\S+)
Value CAPABILITIES (R|B|O|RO|RB|BR|BO|OR|OB|RBO|ROB|BRO|BOR|ORB|OBR)
Value NEIGHBOR_INTERFACE_DESCRIPTION (\S+)

Start
  ^Capability\s+codes:\s+\(R\)\s+Router,\s+\(B\)\s+Bridge,\s+\(O\)\s+Other
  ^LocalPort\s+RemoteDevice\s+RemotePortID\s+Capability\s+RemotePortDescr
  ^[-+\s+]+ -> Row
  ^. -> Error

Row  
  ^${LOCAL_INTERFACE}\s+${NEIGHBOR_NAME}\s+${NEIGHBOR_INTERFACE}\s+${CAPABILITIES}\s+${NEIGHBOR_INTERFACE_DESCRIPTION} -> Record
  ^-+
  ^[\w+\s+]*:\s+\d+
  ^. -> Error

`