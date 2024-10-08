package mikrotik_routeros 

type IpRoutePrintTerse struct {
	Index	string	`json:"INDEX"`
	Flags	string	`json:"FLAGS"`
	Comment	string	`json:"COMMENT"`
	Dst_address	string	`json:"DST_ADDRESS"`
	Routing_table	string	`json:"ROUTING_TABLE"`
	Pref_src	string	`json:"PREF_SRC"`
	Gateway	string	`json:"GATEWAY"`
	Gateway_status	string	`json:"GATEWAY_STATUS"`
	Immediate_gw	string	`json:"IMMEDIATE_GW"`
	Distance	string	`json:"DISTANCE"`
	Scope	string	`json:"SCOPE"`
	Target_scope	string	`json:"TARGET_SCOPE"`
	Suppress_hw_offload	string	`json:"SUPPRESS_HW_OFFLOAD"`
	Local_address	string	`json:"LOCAL_ADDRESS"`
}

var IpRoutePrintTerse_Template = `Value INDEX (\d+)
Value FLAGS ([XADCSrbvomcBUP\s]+(?<!\s))
Value COMMENT (.*)
Value DST_ADDRESS ([\w.:/\d]+)
Value ROUTING_TABLE (\S+)
Value PREF_SRC (\S*)
Value GATEWAY (\S+)
Value GATEWAY_STATUS (.+)
Value IMMEDIATE_GW (\S+)
Value DISTANCE (\d+)
Value SCOPE (\d+)
Value TARGET_SCOPE (\d+)
Value SUPPRESS_HW_OFFLOAD (\S+)
Value LOCAL_ADDRESS (\S+)

Start
  ^(\s?${INDEX}\s*)?${FLAGS}\s+(comment=${COMMENT}\s)?dst-address=${DST_ADDRESS}(\srouting-table=${ROUTING_TABLE})?(\spref-src=${PREF_SRC})?\sgateway=${GATEWAY}(\sgateway-status=${GATEWAY_STATUS})?(\simmediate-gw=${IMMEDIATE_GW})?\sdistance=${DISTANCE}\sscope=${SCOPE}(\starget-scope=${TARGET_SCOPE})?(\ssuppress-hw-offload=${SUPPRESS_HW_OFFLOAD})?(\slocal-address=${LOCAL_ADDRESS})?.*$$ -> Record
  ^\s*(?:\d{2}:){2}\d{2}\s+echo:\s*.*$$ -> Next
  ^. -> Error
`