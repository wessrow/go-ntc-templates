package models

type ArubaOsShowApBssTableDetails struct {
	Bss	string	`json:"BSS"`
	Ess	string	`json:"ESS"`
	Port	string	`json:"PORT"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Band	string	`json:"BAND"`
	Channel	string	`json:"CHANNEL"`
	Type	string	`json:"TYPE"`
	Cur_cl	string	`json:"CUR_CL"`
	Ap_name	string	`json:"AP_NAME"`
	In_t	string	`json:"IN_T"`
	Tot_t	string	`json:"TOT_T"`
	Mtu	string	`json:"MTU"`
	Phy	string	`json:"PHY"`
	Acl_state	string	`json:"ACL_STATE"`
	Acl	string	`json:"ACL"`
	Fm	string	`json:"FM"`
	Flags	string	`json:"FLAGS"`
	Cluster	string	`json:"CLUSTER"`
	Active_clients	string	`json:"ACTIVE_CLIENTS"`
	Standby_clients	string	`json:"STANDBY_CLIENTS"`
	Datazone	string	`json:"DATAZONE"`
}