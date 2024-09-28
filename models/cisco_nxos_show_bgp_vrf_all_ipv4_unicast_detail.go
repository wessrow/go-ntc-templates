package models

type CiscoNxosShowBgpVrfAllIpv4UnicastDetail struct {
	Active	string	`json:"ACTIVE"`
	As_path	string	`json:"AS_PATH"`
	Community	string	`json:"COMMUNITY"`
	Filtered	string	`json:"FILTERED"`
	Local_pref	string	`json:"LOCAL_PREF"`
	Metric	string	`json:"METRIC"`
	Neighbor	string	`json:"NEIGHBOR"`
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Next_hop	string	`json:"NEXT_HOP"`
	Origin	string	`json:"ORIGIN"`
	Path_type	string	`json:"PATH_TYPE"`
	Prefix	string	`json:"PREFIX"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Valid	string	`json:"VALID"`
	Vrf	string	`json:"VRF"`
	Weight	string	`json:"WEIGHT"`
}