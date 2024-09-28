package models

type CiscoAsaShowFailover struct {
	State	string	`json:"STATE"`
	Role	string	`json:"ROLE"`
	Lan_intf_name	string	`json:"LAN_INTF_NAME"`
	Lan_intf	string	`json:"LAN_INTF"`
	Lan_intf_state	string	`json:"LAN_INTF_STATE"`
	Sw_version	string	`json:"SW_VERSION"`
	Sw_version_mate	string	`json:"SW_VERSION_MATE"`
	Failover_group	[]string	`json:"FAILOVER_GROUP"`
	Last_failover_time	[]string	`json:"LAST_FAILOVER_TIME"`
	Last_failover_timezone	[]string	`json:"LAST_FAILOVER_TIMEZONE"`
	Last_failover_month	[]string	`json:"LAST_FAILOVER_MONTH"`
	Last_failover_day	[]string	`json:"LAST_FAILOVER_DAY"`
	Last_failover_year	[]string	`json:"LAST_FAILOVER_YEAR"`
	Service_state	[]string	`json:"SERVICE_STATE"`
	Service_state_mate	[]string	`json:"SERVICE_STATE_MATE"`
	Ssp_slot	[]string	`json:"SSP_SLOT"`
	Ssp_model	[]string	`json:"SSP_MODEL"`
	Ssp_status	[]string	`json:"SSP_STATUS"`
	Ssp_slot_mate	[]string	`json:"SSP_SLOT_MATE"`
	Ssp_model_mate	[]string	`json:"SSP_MODEL_MATE"`
	Ssp_status_mate	[]string	`json:"SSP_STATUS_MATE"`
	Interfaces_context	[]string	`json:"INTERFACES_CONTEXT"`
	Interfaces	[]string	`json:"INTERFACES"`
	Interfaces_status	[]string	`json:"INTERFACES_STATUS"`
	Interfaces_state	[]string	`json:"INTERFACES_STATE"`
	Interfaces_context_mate	[]string	`json:"INTERFACES_CONTEXT_MATE"`
	Interfaces_mate	[]string	`json:"INTERFACES_MATE"`
	Interfaces_status_mate	[]string	`json:"INTERFACES_STATUS_MATE"`
	Interfaces_state_mate	[]string	`json:"INTERFACES_STATE_MATE"`
}