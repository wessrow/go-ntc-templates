package mikrotik_routeros 

type IpRoutePrintTerseWithoutPaging struct {
	Index	string	`json:"INDEX"`
	Flags	string	`json:"FLAGS"`
	Comment	string	`json:"COMMENT"`
	Destination_address	string	`json:"DESTINATION_ADDRESS"`
	Destination_address_subnet	string	`json:"DESTINATION_ADDRESS_SUBNET"`
	Type	string	`json:"TYPE"`
	Routing_table	string	`json:"ROUTING_TABLE"`
	Pref_src	string	`json:"PREF_SRC"`
	Gateway	string	`json:"GATEWAY"`
	Gateway_status	string	`json:"GATEWAY_STATUS"`
	Immediate_gw	string	`json:"IMMEDIATE_GW"`
	Check_gateway	string	`json:"CHECK_GATEWAY"`
	Distance	string	`json:"DISTANCE"`
	Scope	string	`json:"SCOPE"`
	Target_scope	string	`json:"TARGET_SCOPE"`
	Vrf_interface	string	`json:"VRF_INTERFACE"`
	Suppress_hw_offload	string	`json:"SUPPRESS_HW_OFFLOAD"`
	Local_address	string	`json:"LOCAL_ADDRESS"`
	Routing_mark	string	`json:"ROUTING_MARK"`
}

var IpRoutePrintTerseWithoutPaging_Template = `Value INDEX (\d+)
Value FLAGS ([XADCSrbomBUP ]{4})
Value COMMENT (.*?(?=(?:[^=]\S+=)|\s*$))
Value DESTINATION_ADDRESS (\S+)
Value DESTINATION_ADDRESS_SUBNET (\d+)
Value TYPE (.+?(?=(?:[^=]\S+=)|\s*$))
Value ROUTING_TABLE (.+?(?=(?:[^=]\S+=)|\s*$))
Value PREF_SRC (\S+?)
Value GATEWAY (.+?(?=(?:[^=]\S+=)|\s*$))
Value GATEWAY_STATUS (.+?(?=(?:[^=]\S+=)|\s*$))
Value IMMEDIATE_GW (.+?(?=(?:[^=]\S+=)|\s*$))
Value CHECK_GATEWAY (.+?(?=(?:[^=]\S+=)|\s*$))
Value DISTANCE (\d+)
Value SCOPE (\d+)
Value TARGET_SCOPE (\d+)
Value VRF_INTERFACE (.+?(?=(?:[^=]\S+=)|\s*$))
Value SUPPRESS_HW_OFFLOAD (yes|no)
Value LOCAL_ADDRESS (.+?(?=(?:[^=]\S+=)|\s*$))
Value ROUTING_MARK (.*?(?=(?:[^=]\S+=)|\s*$))

Start
  ^\s*${INDEX}\s+${FLAGS}(?:\s+comment=${COMMENT})?\s+dst-address=${DESTINATION_ADDRESS}/${DESTINATION_ADDRESS_SUBNET}(?:\s+type=${TYPE})?(?:\s+routing-table=${ROUTING_TABLE})?(?:\s+pref-src=${PREF_SRC})?(?:\s+gateway=${GATEWAY})?(?:\s+gateway-status=${GATEWAY_STATUS})?(?:\s+immediate-gw=${IMMEDIATE_GW})?(?:\s+check-gateway=${CHECK_GATEWAY})?(?:\s+distance=${DISTANCE})?(?:\s+scope=${SCOPE})?(?:\s+target-scope=${TARGET_SCOPE})?(?:\s+vrf-interface=${VRF_INTERFACE})?(?:\s+suppress-hw-offload=${SUPPRESS_HW_OFFLOAD})?(?:\s+local-address=${LOCAL_ADDRESS})?(?:\s+routing-mark=${ROUTING_MARK})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`