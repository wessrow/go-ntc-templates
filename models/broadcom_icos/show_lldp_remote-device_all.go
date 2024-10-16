package broadcom_icos

type ShowLldpRemoteDeviceAll struct {
	Neighbor_name      string `json:"NEIGHBOR_NAME"`
	Local_interface    string `json:"LOCAL_INTERFACE"`
	Remote_id          string `json:"REMOTE_ID"`
	Chassis_id         string `json:"CHASSIS_ID"`
	Neighbor_interface string `json:"NEIGHBOR_INTERFACE"`
}

var ShowLldpRemoteDeviceAll_Template string = `Value LOCAL_INTERFACE (\S+)
Value REMOTE_ID (\S+)
Value CHASSIS_ID (\S+)
Value NEIGHBOR_INTERFACE (\S+)
Value NEIGHBOR_NAME (\S+)

Start
  # Captures show lldp remote-device all for:
  # Accton AS4610-54P, Accton AS5610-52X, Quanta LY2R, Quanta LB9, DNI AG3448P-R   
  # Raw data is the same in the case of all those devices
  ^LLDP
  ^Local
  ^\s*Interface\s+RemID\s+Chassis\s+ID\s+Port\s+ID\s+System\s+Name$$
  ^-+
  ^\s*${LOCAL_INTERFACE}\s+${REMOTE_ID}\s+${CHASSIS_ID}\s+${NEIGHBOR_INTERFACE}\s+${NEIGHBOR_NAME} -> Record
  # Match records that only include the interface value
  ^\s*${LOCAL_INTERFACE} -> Record
  ^\s*$$
  ^. -> Error
`
