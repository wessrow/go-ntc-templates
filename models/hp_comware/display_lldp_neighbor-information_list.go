package hp_comware 

type DisplayLldpNeighborInformationList struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
}

var DisplayLldpNeighborInformationList_Template = `Value Required LOCAL_INTERFACE (\S+)
Value Required CHASSIS_ID (\S+)
Value Required NEIGHBOR_INTERFACE ((\S+)|(Port\s\d+))
Value Required NEIGHBOR_NAME (\S+)

Start
  ^.*Nearest\s+nontpmr\s+bridge\s+neighbor
  ^.*Nearest\s+customer\s+bridge\s+neighbor
  ^.*Nearest\s+bridge\s+neighbor
  ^System\s+Name -> Format1
  ^Local\s+Interface -> Format2
  ^. -> Error

Format1
  ^\s*${NEIGHBOR_NAME}\s+${LOCAL_INTERFACE}\s+${CHASSIS_ID}\s+${NEIGHBOR_INTERFACE}\s*$$ -> Record Format1
  # Dropping long hostnames
  ^\s*\S+\s*$$
  ^. -> Error

Format2
  ^\s*${LOCAL_INTERFACE}\s+${CHASSIS_ID}\s+${NEIGHBOR_INTERFACE}\s+${NEIGHBOR_NAME}\s*$$ -> Record Format2
  ^. -> Error
`