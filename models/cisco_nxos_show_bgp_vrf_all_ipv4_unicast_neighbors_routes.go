package models

type CiscoNxosShowBgpVrfAllIpv4UnicastNeighborsRoutes struct {
	Status	string	`json:"STATUS"`
	Path_selection	string	`json:"PATH_SELECTION"`
	Route_source	string	`json:"ROUTE_SOURCE"`
	Prefix	string	`json:"PREFIX"`
	Prefix_length	string	`json:"PREFIX_LENGTH"`
	Next_hop	string	`json:"NEXT_HOP"`
	Metric	string	`json:"METRIC"`
	Local_pref	string	`json:"LOCAL_PREF"`
	Weight	string	`json:"WEIGHT"`
	As_path	string	`json:"AS_PATH"`
	Origin	string	`json:"ORIGIN"`
}