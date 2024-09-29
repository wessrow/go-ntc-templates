package cisco_ios 

type ShowStandby struct {
	Interface	string	`json:"INTERFACE"`
	Group	string	`json:"GROUP"`
	Version	string	`json:"VERSION"`
	State	string	`json:"STATE"`
	State_changes	string	`json:"STATE_CHANGES"`
	State_last_change	string	`json:"STATE_LAST_CHANGE"`
	Virtual_ip	string	`json:"VIRTUAL_IP"`
	Secondary_ips	[]string	`json:"SECONDARY_IPS"`
	Active_virtual_mac	string	`json:"ACTIVE_VIRTUAL_MAC"`
	Local_virtual_mac	string	`json:"LOCAL_VIRTUAL_MAC"`
	Hello_time	string	`json:"HELLO_TIME"`
	Hold_time	string	`json:"HOLD_TIME"`
	Authentication	string	`json:"AUTHENTICATION"`
	Preemption	string	`json:"PREEMPTION"`
	Active_router	string	`json:"ACTIVE_ROUTER"`
	Active_router_priority	string	`json:"ACTIVE_ROUTER_PRIORITY"`
	Active_router_mac	string	`json:"ACTIVE_ROUTER_MAC"`
	Standby_router	string	`json:"STANDBY_ROUTER"`
	Standby_router_priority	string	`json:"STANDBY_ROUTER_PRIORITY"`
	Priority	string	`json:"PRIORITY"`
	Group_name	string	`json:"GROUP_NAME"`
	Flags	string	`json:"FLAGS"`
	Track_item	string	`json:"TRACK_ITEM"`
	Track_type	string	`json:"TRACK_TYPE"`
	Track_state	string	`json:"TRACK_STATE"`
	Track_decrement_time	string	`json:"TRACK_DECREMENT_TIME"`
}

var ShowStandby_Template = `Value INTERFACE (\S+)
Value GROUP (\d+)
Value VERSION (\d+)
Value STATE (\w+)
Value STATE_CHANGES (\d+)
Value STATE_LAST_CHANGE (\S+)
Value VIRTUAL_IP (\S+)
Value List SECONDARY_IPS (\S+)
Value ACTIVE_VIRTUAL_MAC (\S+)
Value LOCAL_VIRTUAL_MAC (\S+)
Value HELLO_TIME (\d+)
Value HOLD_TIME (\d+)
Value AUTHENTICATION (\w+)
Value PREEMPTION (\w+)
Value ACTIVE_ROUTER (\d+\.\d+\.\d+\.\d+|local|unknown)
Value ACTIVE_ROUTER_PRIORITY (\d+)
Value ACTIVE_ROUTER_MAC (\S+)
Value STANDBY_ROUTER (\d+\.\d+\.\d+\.\d+|local|unknown)
Value STANDBY_ROUTER_PRIORITY (\d+)
Value PRIORITY (\d+)
Value GROUP_NAME (\S+)
Value FLAGS (\S+)
Value TRACK_ITEM (\S+)
Value TRACK_TYPE (object|interface)
Value TRACK_STATE (\S+)
Value TRACK_DECREMENT_TIME (\d+)

Start
  ^\S+ -> Continue.Record
  ^\s*${INTERFACE}\s+-\s+Group\s+${GROUP}\s*(\(version\s+${VERSION}\))?$$
  ^.*State\s+is\s+${STATE}
  ^.*Virtual\s+IP\s+address\s+is\s+${VIRTUAL_IP}
  ^.*Secondary\s+virtual\s+IP\s+address\s+${SECONDARY_IPS}
  ^.*Active\s+virtual\s+MAC\s+address\s+is\s+${ACTIVE_VIRTUAL_MAC}
  ^.*Local\s+virtual\s+MAC\s+address\s+is\s+${LOCAL_VIRTUAL_MAC}
  ^.*Hello\s+time\s+${HELLO_TIME}.*hold\s+time\s+${HOLD_TIME}
  ^.*Next\s+hello\s+sent\s+in.*secs
  ^.*Authentication\s+${AUTHENTICATION}
  ^.*Preemption\s+${PREEMPTION}
  ^.*Active\s+router\s+is\s+${ACTIVE_ROUTER}(,\s+priority\s+${ACTIVE_ROUTER_PRIORITY})?
  ^\s+MAC\s+address\s+is\s+${ACTIVE_ROUTER_MAC}
  ^.*Standby\s+router\s+is\s+${STANDBY_ROUTER}(,\s+priority\s+${STANDBY_ROUTER_PRIORITY})?
  ^.*Priority\s+${PRIORITY}
  ^\s*Track\s+${TRACK_TYPE}(\s+${TRACK_ITEM})?(\(unknown\))?(\s+state\s+${TRACK_STATE}(\s+decrement\s+${TRACK_DECREMENT_TIME})?)?
  ^.*${STATE_CHANGES}\s+state\s+(change|changes),\s+last\s+state\s+change\s+${STATE_LAST_CHANGE}
  ^.*Group\s+name\s+is\s+"${GROUP_NAME}"
  ^.*FLAGS:\s+${FLAGS}
  ^.*IP\s+redundancy\s+name\s+is\s+\"${GROUP_NAME}\"(\s+\(default\))?
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error
`