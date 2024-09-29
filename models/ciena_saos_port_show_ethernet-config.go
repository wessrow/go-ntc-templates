package models

type CienaSaosPortShowEthernetConfig struct {
	Name	string	`json:"NAME"`
	Type	string	`json:"TYPE"`
	Admin_status	string	`json:"ADMIN_STATUS"`
	Speed	string	`json:"SPEED"`
	Duplex	string	`json:"DUPLEX"`
	Flow_control	string	`json:"FLOW_CONTROL"`
	Flow_control_adv	string	`json:"FLOW_CONTROL_ADV"`
	Auto_negotiation	string	`json:"AUTO_NEGOTIATION"`
	Mtu	string	`json:"MTU"`
	Mirror_state	string	`json:"MIRROR_STATE"`
	Mirror_eg	string	`json:"MIRROR_EG"`
	Mirror_ig	string	`json:"MIRROR_IG"`
}

var CienaSaosPortShowEthernetConfig_Template = `Value NAME (\S+)
Value TYPE (\S+)
Value ADMIN_STATUS (\S+)
Value SPEED (\S+)
Value DUPLEX (\S+)
Value FLOW_CONTROL (\S+)
Value FLOW_CONTROL_ADV (\S+)
Value AUTO_NEGOTIATION (\S+)
Value MTU (\d+)
Value MIRROR_STATE (\S+)
Value MIRROR_EG (\d+)
Value MIRROR_IG (\d+)

Start
  ^\+-+\s*PORT\s*ETHERNET\s*CONFIGURATION\s*-+\+ 
  ^(\|\s+){10}Mirror\s+\|\s*$$
  ^\|\s*Port\s*\|\s*Port\s*\|\s*Admin\s*\|(?:\s+\|){3}\s*FC\s*\|\s*Auto\s*\|\s*MTU\s*\|\s*Status\s*\|\s*$$
  ^\|\s*Name\s*\|\s*Type\s*\|\s*Status\|\s*Speed\s*\|\s*Dplx\s*\|\s*FC\s*\|\s*Adv\s*\|\s*Neg\s*\|\s*Size\s*\|\s*State\|\s*Eg\s*\|\s*Ig\s*\|\s*$$
  ^\+-+
  ^\|\s*${NAME}\s*\|\s*${TYPE}\s*\|\s*${ADMIN_STATUS}\s*\|\s*${SPEED}\s*\|\s*(${DUPLEX}|\s+)\s*\|\s*${FLOW_CONTROL}\s*\|\s*${FLOW_CONTROL_ADV}\s*\|\s*${AUTO_NEGOTIATION}\s*\|\s*${MTU}\s*\|\s*${MIRROR_STATE}\s*\|\s*${MIRROR_EG}\s*\|\s*${MIRROR_IG}\s*\| -> Record
  ^\s*$$
  ^. -> Error

`