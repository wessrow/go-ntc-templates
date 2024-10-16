package arista_eos

type ShowLldpNeighborsDetail struct {
	Neighbor_count       string `json:"NEIGHBOR_COUNT"`
	Age                  string `json:"AGE"`
	Neighbor_name        string `json:"NEIGHBOR_NAME"`
	Chassis_id           string `json:"CHASSIS_ID"`
	Mgmt_address         string `json:"MGMT_ADDRESS"`
	Neighbor_description string `json:"NEIGHBOR_DESCRIPTION"`
	Neighbor_interface   string `json:"NEIGHBOR_INTERFACE"`
	Local_interface      string `json:"LOCAL_INTERFACE"`
}

var ShowLldpNeighborsDetail_Template string = `Value Required NEIGHBOR_NAME (\S+)
Value CHASSIS_ID (.+?)
Value MGMT_ADDRESS (.+?)
Value NEIGHBOR_DESCRIPTION (.+?)
Value NEIGHBOR_INTERFACE (.+?)
Value Filldown LOCAL_INTERFACE (\S+?)
Value Filldown NEIGHBOR_COUNT ([1-9]\d*)
Value AGE (.+?)

Start
  ^\S+\s+${LOCAL_INTERFACE}\s+\S+\s+${NEIGHBOR_COUNT}.*
  ^.*age\s+${AGE}$$ -> Interface

Interface
  ^\s+Chassis ID\s+:\s+${CHASSIS_ID}$$
  ^\s+-\s+Port ID type -> Port_ID
  ^\s+-\s+\S+\s+Name:\s+\"${NEIGHBOR_NAME}\"
  ^.*System Description:\s+\"${NEIGHBOR_DESCRIPTION}(\"|$$)
  ^\s+\S+\s+Address\s+:\s+${MGMT_ADDRESS}$$
  ^\s*$$ -> Record Start

Port_ID
  ^\s+Port ID\s+:(\s+\"|\s+)${NEIGHBOR_INTERFACE}(\"|$$) -> Interface
`
