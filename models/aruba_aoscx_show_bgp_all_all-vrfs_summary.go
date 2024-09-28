package models

type ArubaAoscxShowBgpAllAllVrfsSummary struct {
	Vrf	string	`json:"VRF"`
	Af	string	`json:"AF"`
	Neighbour	string	`json:"NEIGHBOUR"`
	Remote_as	string	`json:"REMOTE_AS"`
	State	string	`json:"STATE"`
	Admin_status	string	`json:"ADMIN_STATUS"`
}