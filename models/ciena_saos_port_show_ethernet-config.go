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