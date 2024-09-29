package hp_procurve 

type ShowLldpInfoRemoteDevice struct {
	Local_interface	string	`json:"LOCAL_INTERFACE"`
	Chassis_id	string	`json:"CHASSIS_ID"`
	Neighbor_interface	string	`json:"NEIGHBOR_INTERFACE"`
	Neighbor_description	string	`json:"NEIGHBOR_DESCRIPTION"`
	Neighbor_name	string	`json:"NEIGHBOR_NAME"`
}

var ShowLldpInfoRemoteDevice_Template = `Value Required LOCAL_INTERFACE (\S+)
Value CHASSIS_ID ((?:[a-z0-9]{2}\s){5}[a-z0-9]{2}|\S+)
Value NEIGHBOR_INTERFACE ((?:[a-z0-9]{2}\s){5}[a-z0-9]{2}|\S+)
Value NEIGHBOR_DESCRIPTION (\S+|.*?)
Value NEIGHBOR_NAME (\S+)

Start
  ^\s*LLDP\sRemote\sDevices\sInformation\s*$$
  ^\s+LocalPort\s|\sChassisId.*$$
  ^\s+-.* -> LLDP
  ^\s*$$
  ^. -> Error

LLDP
  ^\s+${LOCAL_INTERFACE}\s+\|\s${CHASSIS_ID}\s+${NEIGHBOR_INTERFACE}\s*$$ -> Record
  ^\s+${LOCAL_INTERFACE}\s+\|\s${CHASSIS_ID}\s+${NEIGHBOR_INTERFACE}\s+${NEIGHBOR_NAME}\s*$$ -> Record
  ^\s+${LOCAL_INTERFACE}\s+\|\s${CHASSIS_ID}\s+${NEIGHBOR_INTERFACE}\s+${NEIGHBOR_DESCRIPTION}\s+${NEIGHBOR_NAME}\s*$$ -> Record
  ^\S+\#\s*$$
  ^\s*$$
  ^. -> Error
`