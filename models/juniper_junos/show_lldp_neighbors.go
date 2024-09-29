package juniper_junos 

type ShowLldpNeighbors struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Parent_interface	string	`json:"PARENT_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
}

var ShowLldpNeighbors_Template = `Value Required LOCAL_INTERFACE (\S+)
Value Required PARENT_INTERFACE (\S+)
Value Required CHASSIS_ID (\S+)
Value Required NEIGHBOR_INTERFACE (\S+)
Value Required NEIGHBOR_NAME (\S+)

Start
  ^Local\s+Interface\s+Parent\s+Interface\s+Chassis\s+Id\s+Port\s+info\s+System\s+Name
  ^${LOCAL_INTERFACE}\s*${PARENT_INTERFACE}\s+${CHASSIS_ID}\s+${NEIGHBOR_INTERFACE}\s+${NEIGHBOR_NAME} -> Record
  ^\s*$$
  ^. -> Error
`