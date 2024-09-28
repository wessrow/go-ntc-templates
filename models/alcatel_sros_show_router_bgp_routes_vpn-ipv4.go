package models

type AlcatelSrosShowRouterBgpRoutesVpnIpv4 struct {
	In_out_use	string	`json:"IN_OUT_USE"`
	Rd	string	`json:"RD"`
	Network	string	`json:"NETWORK"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Local_pref	string	`json:"LOCAL_PREF"`
	Med	string	`json:"MED"`
	Next_hop	string	`json:"NEXT_HOP"`
	Path_id	string	`json:"PATH_ID"`
	Label	string	`json:"LABEL"`
	As_path	string	`json:"AS_PATH"`
}