package mikrotik_routeros

type IpRoutePrintTerse struct {
	Suppress_hw_offload string `json:"SUPPRESS_HW_OFFLOAD"`
	Index               string `json:"INDEX"`
	Flags               string `json:"FLAGS"`
	Dst_address         string `json:"DST_ADDRESS"`
	Pref_src            string `json:"PREF_SRC"`
	Immediate_gw        string `json:"IMMEDIATE_GW"`
	Distance            string `json:"DISTANCE"`
	Comment             string `json:"COMMENT"`
	Gateway_status      string `json:"GATEWAY_STATUS"`
	Local_address       string `json:"LOCAL_ADDRESS"`
	Target_scope        string `json:"TARGET_SCOPE"`
	Routing_table       string `json:"ROUTING_TABLE"`
	Gateway             string `json:"GATEWAY"`
	Scope               string `json:"SCOPE"`
}

var IpRoutePrintTerse_Template string = `Value INDEX (\d+)
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
