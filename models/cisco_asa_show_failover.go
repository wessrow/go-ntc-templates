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

var CiscoAsaShowFailover_Template = `Value STATE (\S+)
Value ROLE (\S+)
Value LAN_INTF_NAME (\S+)
Value LAN_INTF (\S+)
Value LAN_INTF_STATE (\S+)
Value SW_VERSION (\S+)
Value SW_VERSION_MATE (\S+)
Value List FAILOVER_GROUP (\d+)
Value List LAST_FAILOVER_TIME (\d+:\d+:\d+)
Value List LAST_FAILOVER_TIMEZONE (\S+)
Value List LAST_FAILOVER_MONTH (\w+)
Value List LAST_FAILOVER_DAY (\d+)
Value List LAST_FAILOVER_YEAR (\d+)
Value List SERVICE_STATE (.*?)
Value List SERVICE_STATE_MATE (.*?)
Value List SSP_SLOT (\d+)
Value List SSP_MODEL (\S+)
Value List SSP_STATUS (\S+)
Value List SSP_SLOT_MATE (\d+)
Value List SSP_MODEL_MATE (\S+)
Value List SSP_STATUS_MATE (\S+)
Value List INTERFACES_CONTEXT (\S+)
Value List INTERFACES (\S+)
Value List INTERFACES_STATUS (.+?)
Value List INTERFACES_STATE (\S+)
Value List INTERFACES_CONTEXT_MATE (\S+)
Value List INTERFACES_MATE (\S+)
Value List INTERFACES_STATUS_MATE (.+?)
Value List INTERFACES_STATE_MATE (\S+)

Start
  ^Failover\s+${STATE}\s*$$
  ^Failover\s+unit\s+${ROLE}\s*$$
  ^Failover\s+LAN\s+Interface:\s+${LAN_INTF_NAME}\s+${LAN_INTF}\s+\(${LAN_INTF_STATE}\)\s*$$
  ^Version:\s+Ours\s+${SW_VERSION},\s+Mate\s+${SW_VERSION_MATE}\s*$$
  ^(?:Group\s+${FAILOVER_GROUP}\s+|)[Ll]ast\s+[Ff]ailover\s+at:\s+${LAST_FAILOVER_TIME}\s+${LAST_FAILOVER_TIMEZONE}\s+${LAST_FAILOVER_MONTH}\s+${LAST_FAILOVER_DAY}\s+${LAST_FAILOVER_YEAR}\s*$$
  ^\s*This\s+host:.+?-\s+${SERVICE_STATE}\s*$$ -> ThisHost
  ^\s*This\s+host:\s+\S+\s*$$ -> ThisHost
  ^Reconnect\s+timeout
  ^Unit\s+Poll\s+frequency
  ^Interface\s+Poll\s+frequency
  ^Interface\s+Policy
  ^Monitored\s+Interfaces
  ^MAC\s+Address\s+Move\s+Notification\s+Interval
  ^Serial\s+Number
  ^failover\s+replication
  ^\s*$$
  ^. -> Error

ThisHost
  ^Group\s+\d+\s+State:\s*${SERVICE_STATE}\s*$$
  ^\s*Active\s+time
  ^\s*slot\s+${SSP_SLOT}:\s+${SSP_MODEL}\s+.+?status\s+\(${SSP_STATUS}.*?\)\s*$$
  ^\s*(${INTERFACES_CONTEXT}\s+|)Interface\s+${INTERFACES}.+?:\s+${INTERFACES_STATUS}(?:\s+\(${INTERFACES_STATE}\)|)\s*$$
  # Service module has different line
  ^\s*\S+,\s+\S+,\s+\S+\s*$$
  ^\s*Other\s+host:.+?-\s+${SERVICE_STATE_MATE}\s*$$ -> OtherHost
  ^\s*Other\s+host:\s+\S+\s*$$ -> OtherHost
  ^\s*slot\s+\d+:\s+empty\s*$$
  ^\s+\<omited\s+output\>
  ^\s+IPS
  ^\s*ASA\s+FirePOWER
  ^\s*$$
  ^. -> Error

OtherHost
  ^Group\s+\d+\s+State:\s*${SERVICE_STATE_MATE}\s*$$
  ^\s*Active\s+time
  ^\s*slot\s+${SSP_SLOT_MATE}:\s+${SSP_MODEL_MATE}\s+.+?status\s+\(${SSP_STATUS_MATE}.*?\)\s*$$
  ^\s*(${INTERFACES_CONTEXT_MATE}\s+|)Interface\s+${INTERFACES_MATE}.+?:\s+${INTERFACES_STATUS_MATE}(?:\s+\(${INTERFACES_STATE_MATE}\)|)\s*$$
  # Service module has different line
  ^\s*\S+,\s+\S+,\s+\S+\s*$$
  ^\s*slot\s+\d+:\s+empty\s*$$
  ^\s+\<omited\s+output\>
  ^\s+IPS
  ^\s*ASA\s+FirePOWER
  ^Stateful\s+Failover\s+Logical\s+Update\s+Statistics\s*$$ -> Stats
  ^\s*$$
  ^. -> Error

Stats
  ^Link\s*:\s+
  ^Stateful\s+Obj\s+xmit\s+xerr\s+rcv\s+rerr\s*$$
  ^.+?\d+\s+\d+\s+\d+\s+\d+\s*$$
  ^Logical\s+Update\s+Queue\s+Information\s*$$
  ^Cur\s+Max\s+Total\s*$$
  ^.+?:\s+\d+\s+\d+\s+\d+\s*$$

`