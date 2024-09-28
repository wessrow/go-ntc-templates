package models

type AristaEosShowIpBgpDetail struct {
	Vrf	string	`json:"VRF"`
	Local_pref	string	`json:"LOCAL_PREF"`
	Weight	string	`json:"WEIGHT"`
	Origin	string	`json:"ORIGIN"`
	Prefix	string	`json:"PREFIX"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	As_path	string	`json:"AS_PATH"`
	Valid	string	`json:"VALID"`
	Path_type	string	`json:"PATH_TYPE"`
	Active	string	`json:"ACTIVE"`
	Community	string	`json:"COMMUNITY"`
	Next_hop	string	`json:"NEXT_HOP"`
	Neighbor	string	`json:"NEIGHBOR"`
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Router_id	string	`json:"ROUTER_ID"`
	Local_as	string	`json:"LOCAL_AS"`
	Metric	string	`json:"METRIC"`
	Backup	string	`json:"BACKUP"`
}