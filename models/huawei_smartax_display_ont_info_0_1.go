package models

type HuaweiSmartaxDisplayOntInfo01 struct {
	Fsp	string	`json:"FSP"`
	Ont_id	string	`json:"ONT_ID"`
	Control_flag	string	`json:"CONTROL_FLAG"`
	Run_state	string	`json:"RUN_STATE"`
	Config_state	string	`json:"CONFIG_STATE"`
	Match_state	string	`json:"MATCH_STATE"`
	Ont_llid	string	`json:"ONT_LLID"`
	Dba_type	string	`json:"DBA_TYPE"`
	Distance_m	string	`json:"DISTANCE_M"`
	Last_distance_m	string	`json:"LAST_DISTANCE_M"`
	Ont_rtt	string	`json:"ONT_RTT"`
	Battery_state	string	`json:"BATTERY_STATE"`
	Power_type	string	`json:"POWER_TYPE"`
	Memory_occupation	string	`json:"MEMORY_OCCUPATION"`
	Cpu_occupation	string	`json:"CPU_OCCUPATION"`
	Temperature	string	`json:"TEMPERATURE"`
	Authentic_type	string	`json:"AUTHENTIC_TYPE"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Management_mode	string	`json:"MANAGEMENT_MODE"`
	Software_work_mode	string	`json:"SOFTWARE_WORK_MODE"`
	Multicast_mode	string	`json:"MULTICAST_MODE"`
	Isolation_state	string	`json:"ISOLATION_STATE"`
	Ip_0_address_mask	string	`json:"IP_0_ADDRESS_MASK"`
	Ip_1_address_mask	string	`json:"IP_1_ADDRESS_MASK"`
	Description	string	`json:"DESCRIPTION"`
	Last_down_cause	string	`json:"LAST_DOWN_CAUSE"`
	Last_up_time	string	`json:"LAST_UP_TIME"`
	Last_down_time	string	`json:"LAST_DOWN_TIME"`
	Last_dying_gasp_time	string	`json:"LAST_DYING_GASP_TIME"`
	Ont_online_duration	string	`json:"ONT_ONLINE_DURATION"`
	Ont_system_up_duration	string	`json:"ONT_SYSTEM_UP_DURATION"`
	Last_restart_reason	string	`json:"LAST_RESTART_REASON"`
	Type_c_support	string	`json:"TYPE_C_SUPPORT"`
	Type_d_support	string	`json:"TYPE_D_SUPPORT"`
	Ont_nni_type	string	`json:"ONT_NNI_TYPE"`
	Ont_actual_nni_type	string	`json:"ONT_ACTUAL_NNI_TYPE"`
	Last_ont_actual_nni_type	string	`json:"LAST_ONT_ACTUAL_NNI_TYPE"`
	Interoperability_mode	string	`json:"INTEROPERABILITY_MODE"`
	Power_reduction_status	string	`json:"POWER_REDUCTION_STATUS"`
	Fec_upstream_state	string	`json:"FEC_UPSTREAM_STATE"`
	Vs_id	string	`json:"VS_ID"`
	Vs_name	string	`json:"VS_NAME"`
	Global_ont_id	string	`json:"GLOBAL_ONT_ID"`
	Voip_configure_method	string	`json:"VOIP_CONFIGURE_METHOD"`
	Line_profile_id	string	`json:"LINE_PROFILE_ID"`
	Line_profile_name	string	`json:"LINE_PROFILE_NAME"`
	Fec_upstream_switch	string	`json:"FEC_UPSTREAM_SWITCH"`
	Omcc_encrypt_switch	string	`json:"OMCC_ENCRYPT_SWITCH"`
	Qos_mode	string	`json:"QOS_MODE"`
	Mapping_mode	string	`json:"MAPPING_MODE"`
	Tr069_management	string	`json:"TR069_MANAGEMENT"`
	Tr069_ip_index	string	`json:"TR069_IP_INDEX"`
	Service_profile_id	string	`json:"SERVICE_PROFILE_ID"`
	Service_profile_name	string	`json:"SERVICE_PROFILE_NAME"`
	Tdm_port_type	string	`json:"TDM_PORT_TYPE"`
	Tdm_service_type	string	`json:"TDM_SERVICE_TYPE"`
	Mac_learning_function_switch	string	`json:"MAC_LEARNING_FUNCTION_SWITCH"`
	Ont_transparent_function_switch	string	`json:"ONT_TRANSPARENT_FUNCTION_SWITCH"`
	Ring_check_switch	string	`json:"RING_CHECK_SWITCH"`
	Ring_port_auto_shutdown	string	`json:"RING_PORT_AUTO_SHUTDOWN"`
	Ring_resume_frequency	string	`json:"RING_RESUME_FREQUENCY"`
	Ring_detect_frequency	string	`json:"RING_DETECT_FREQUENCY"`
	Ring_resume_interval	string	`json:"RING_RESUME_INTERVAL"`
	Ring_detect_period	string	`json:"RING_DETECT_PERIOD"`
	Service_profile_multicast_forward_mode	string	`json:"SERVICE_PROFILE_MULTICAST_FORWARD_MODE"`
	Service_profile_multicast_forward_vlan	string	`json:"SERVICE_PROFILE_MULTICAST_FORWARD_VLAN"`
	Service_profile_multicast_mode	string	`json:"SERVICE_PROFILE_MULTICAST_MODE"`
	Upstream_igmp_packet_forward_mode	string	`json:"UPSTREAM_IGMP_PACKET_FORWARD_MODE"`
	Upstream_igmp_packet_forward_vlan	string	`json:"UPSTREAM_IGMP_PACKET_FORWARD_VLAN"`
	Upstream_igmp_packet_priority	string	`json:"UPSTREAM_IGMP_PACKET_PRIORITY"`
	Native_vlan_option	string	`json:"NATIVE_VLAN_OPTION"`
	Upstream_pq_color_policy	string	`json:"UPSTREAM_PQ_COLOR_POLICY"`
	Downstream_pq_color_policy	string	`json:"DOWNSTREAM_PQ_COLOR_POLICY"`
	Monitor_link	string	`json:"MONITOR_LINK"`
	Mtu_byte	string	`json:"MTU_BYTE"`
	Alarm_policy_profile_id	string	`json:"ALARM_POLICY_PROFILE_ID"`
	Alarm_policy_profile_name	string	`json:"ALARM_POLICY_PROFILE_NAME"`
	Tr069_server_profile_id	string	`json:"TR069_SERVER_PROFILE_ID"`
	Tr069_server_profile_name	string	`json:"TR069_SERVER_PROFILE_NAME"`
	Fiber_route	string	`json:"FIBER_ROUTE"`
	T_cont	[]string	`json:"T_CONT"`
	Number_pots	string	`json:"NUMBER_POTS"`
	Number_eth	string	`json:"NUMBER_ETH"`
	Number_tdm	string	`json:"NUMBER_TDM"`
	Number_moca	string	`json:"NUMBER_MOCA"`
	Number_catv	string	`json:"NUMBER_CATV"`
	Number_vdsl	string	`json:"NUMBER_VDSL"`
	Number_max_adaptive_pots	string	`json:"NUMBER_MAX_ADAPTIVE_POTS"`
	Number_max_adaptive_eth	string	`json:"NUMBER_MAX_ADAPTIVE_ETH"`
	Number_max_adaptive_tdm	string	`json:"NUMBER_MAX_ADAPTIVE_TDM"`
	Number_max_adaptive_moca	string	`json:"NUMBER_MAX_ADAPTIVE_MOCA"`
	Number_max_adaptive_catv	string	`json:"NUMBER_MAX_ADAPTIVE_CATV"`
	Number_max_adaptive_vdsl	string	`json:"NUMBER_MAX_ADAPTIVE_VDSL"`
}

var HuaweiSmartaxDisplayOntInfo01_Template = `Value FSP (\d+\/\s*\d+\/\d+)
Value ONT_ID (\d+)
Value CONTROL_FLAG (\w+)
Value RUN_STATE (\w+)
Value CONFIG_STATE (\w+)
Value MATCH_STATE (\w+)
Value ONT_LLID (\w+)
Value DBA_TYPE (\w+)
Value DISTANCE_M (\d+)
Value LAST_DISTANCE_M (\d+)
Value ONT_RTT (\d+)
Value BATTERY_STATE (\w+\s*\w*)
Value POWER_TYPE (\w+)
Value MEMORY_OCCUPATION (\d+)
Value CPU_OCCUPATION (\d+)
Value TEMPERATURE (\d+(\(\w\))?)
Value AUTHENTIC_TYPE (MAC-auth|SN-auth)
Value SERIAL_NUMBER (\w+(\s*\(*\S*\))?)
Value MANAGEMENT_MODE (\w+)
Value SOFTWARE_WORK_MODE (\w+)
Value MULTICAST_MODE (IGMP-Snooping)
Value ISOLATION_STATE (\w+)
Value IP_0_ADDRESS_MASK (\S+)
Value IP_1_ADDRESS_MASK (\S+)
Value DESCRIPTION (\S+)
Value LAST_DOWN_CAUSE (\S+)
Value LAST_UP_TIME (\w+(\/|-)\w+(\/|-)\w+\s\w+:\w+:\w+\+\w+:\w+)
Value LAST_DOWN_TIME (\w+(\/|-)\w+(\/|-)\w+\s\w+:\w+:\w+\+\w+:\w+)
Value LAST_DYING_GASP_TIME (\w+(\/|-)\w+(\/|-)\w+\s\w+:\w+:\w+\+\w+:\w+)
Value ONT_ONLINE_DURATION (\d*\s*day\(s\),\s*\d*\s*hour\(s\),\s*\d*\s*minute\(s\),\s*\d*\s*second\(s\))
Value ONT_SYSTEM_UP_DURATION (\d*\s*day\(s\),\s*\d*\s*hour\(s\),\s*\d*\s*minute\(s\),\s*\d*\s*second\(s\))
Value LAST_RESTART_REASON (\w+)
Value TYPE_C_SUPPORT (\w+\s*\w*)
Value TYPE_D_SUPPORT (\w+\s*\w*)
Value ONT_NNI_TYPE (\w+)
Value ONT_ACTUAL_NNI_TYPE (\w+)
Value LAST_ONT_ACTUAL_NNI_TYPE (\w+)
Value INTEROPERABILITY_MODE (\S+)
Value POWER_REDUCTION_STATUS (\w+)
Value FEC_UPSTREAM_STATE (\S+)
Value VS_ID (\d+)
Value VS_NAME (\S+)
Value GLOBAL_ONT_ID (\d+)
Value VOIP_CONFIGURE_METHOD (\w+)
Value LINE_PROFILE_ID (\d+)
Value LINE_PROFILE_NAME (\w+)
Value FEC_UPSTREAM_SWITCH (\w+)
Value OMCC_ENCRYPT_SWITCH (\w+)
Value QOS_MODE (\w+)
Value MAPPING_MODE (\w+)
Value TR069_MANAGEMENT (\w+)
Value TR069_IP_INDEX (\d+)
Value SERVICE_PROFILE_ID (\d+)
Value SERVICE_PROFILE_NAME (\w+)
Value TDM_PORT_TYPE (\w+)
Value TDM_SERVICE_TYPE (\w+)
Value MAC_LEARNING_FUNCTION_SWITCH (\w+)
Value ONT_TRANSPARENT_FUNCTION_SWITCH (\w+)
Value RING_CHECK_SWITCH (\w+)
Value RING_PORT_AUTO_SHUTDOWN (\w+)
Value RING_RESUME_FREQUENCY (\w+\s\(\w*\))
Value RING_DETECT_FREQUENCY (\w+\s\(\w*\))
Value RING_RESUME_INTERVAL (\w+\s\(\w*\))
Value RING_DETECT_PERIOD (\w+\s\(\w*\))
Value SERVICE_PROFILE_MULTICAST_FORWARD_MODE (\w+)
Value SERVICE_PROFILE_MULTICAST_FORWARD_VLAN (\w+)
Value SERVICE_PROFILE_MULTICAST_MODE (\w+)
Value UPSTREAM_IGMP_PACKET_FORWARD_MODE (\w+)
Value UPSTREAM_IGMP_PACKET_FORWARD_VLAN (\w+)
Value UPSTREAM_IGMP_PACKET_PRIORITY (\w+)
Value NATIVE_VLAN_OPTION (\w+)
Value UPSTREAM_PQ_COLOR_POLICY (\w+)
Value DOWNSTREAM_PQ_COLOR_POLICY (\w+)
Value MONITOR_LINK (\w+)
Value MTU_BYTE (\w+)
Value ALARM_POLICY_PROFILE_ID (\d+)
Value ALARM_POLICY_PROFILE_NAME (\S+)
Value TR069_SERVER_PROFILE_ID (\d+)
Value TR069_SERVER_PROFILE_NAME (\S+)
Value FIBER_ROUTE (\S+)
Value List T_CONT (\d+)
Value NUMBER_POTS (\d+|adaptive)
Value NUMBER_ETH (\d+|adaptive)
Value NUMBER_TDM (\d+|adaptive)
Value NUMBER_MOCA (\d+|adaptive)
Value NUMBER_CATV (\d+|adaptive)
Value NUMBER_VDSL (\d+|adaptive)
Value NUMBER_MAX_ADAPTIVE_POTS (\d*)
Value NUMBER_MAX_ADAPTIVE_ETH (\d*)
Value NUMBER_MAX_ADAPTIVE_TDM (\d*)
Value NUMBER_MAX_ADAPTIVE_MOCA (\d*)
Value NUMBER_MAX_ADAPTIVE_CATV (\d*)
Value NUMBER_MAX_ADAPTIVE_VDSL (\d*)

Start
  ^\s*F/S/P\s*:\s*${FSP}
  ^\s*ONT-ID\s*:\s*${ONT_ID}
  ^\s*Control\s+flag\s*:\s*${CONTROL_FLAG}
  ^\s*Run\s+state\s*:\s*${RUN_STATE}
  ^\s*Config\s+state\s*:\s*${CONFIG_STATE}
  ^\s*Match\s+state\s*:\s*${MATCH_STATE}
  ^\s*ONT\s+LLID\s*:\s*${ONT_LLID}
  ^\s*DBA\s+type\s*:\s*${DBA_TYPE}
  ^\s*ONT\s+distance\(m\)\s*:\s*${DISTANCE_M}
  ^\s*ONT\s+last\s+distance\(m\)\s*:\s*${LAST_DISTANCE_M}
  ^\s*ONT\s+RTT\s*\(TQ\)\s*:\s*${ONT_RTT}
  ^\s*ONT\s+battery\s+state\s*:\s*${BATTERY_STATE}
  ^\s*ONT\s+power\s+type\s*:\s*${POWER_TYPE}
  ^\s*Memory\s+occupation\s*:\s*${MEMORY_OCCUPATION}(%)?
  ^\s*CPU\s+occupation\s*:\s*${CPU_OCCUPATION}(%)?
  ^\s*Temperature\s*:\s*${TEMPERATURE}
  ^\s*Authentic\s+type\s*:\s*${AUTHENTIC_TYPE}
  ^\s*SN\s*:\s*${SERIAL_NUMBER}
  ^\s*Management\s+mode\s*:\s*${MANAGEMENT_MODE}
  ^\s*Software\s+work\s+mode\s*:\s*${SOFTWARE_WORK_MODE}
  ^\s*Multicast\s*mode\s*:\s*${MULTICAST_MODE}
  ^\s*Isolation\s+state\s*:\s*${ISOLATION_STATE}
  ^\s*ONT\s+NNI\s+type\s*:\s*${ONT_NNI_TYPE}
  ^\s*ONT\s+actual\s+NNI\s+type\s*:\s*${ONT_ACTUAL_NNI_TYPE}
  ^\s*Last\s+ONT\s+actual\s+NNI\s+type\s*:\s*${LAST_ONT_ACTUAL_NNI_TYPE}
  ^\s*ONT\s+IP\s+0\s+address\/mask\s*:\s*${IP_0_ADDRESS_MASK}
  ^\s*ONT\s+IP\s+1\s+address\/mask\s*:\s*${IP_1_ADDRESS_MASK}
  ^\s*Description\s*:\s*${DESCRIPTION}
  ^\s*Last\s+down\s+cause\s*:\s*${LAST_DOWN_CAUSE}
  ^\s*Last\s+up\s+time\s*:\s*${LAST_UP_TIME}
  ^\s*Last\s+down\s+time\s*:\s*${LAST_DOWN_TIME}
  ^\s*Last\s+dying\s+gasp\s+time\s*:\s*${LAST_DYING_GASP_TIME}
  ^\s*ONT\s+online\s+duration\s*:\s*${ONT_ONLINE_DURATION}
  ^\s*ONT\s+system\s+up\s+duration\s*:\s*${ONT_SYSTEM_UP_DURATION}
  ^\s*Last\s+restart\s+reason\s*:\s*${LAST_RESTART_REASON}
  ^\s*Type\s+C\s+support\s*:\s*${TYPE_C_SUPPORT}
  ^\s*Type\s+D\s+support\s*:\s*${TYPE_D_SUPPORT}
  ^\s*Interoperability-mode\s*:\s*${INTEROPERABILITY_MODE}
  ^\s*Power\s+reduction\s+status\s*:\s*${POWER_REDUCTION_STATUS}
  ^\s*FEC\s+upstream\s+state\s*:\s*${FEC_UPSTREAM_STATE}
  ^\s*VS-ID\s*:\s*${VS_ID}
  ^\s*VS\s+name\s*:\s*${VS_NAME}
  ^\s*Global\s+ONT-ID\s*:\s*${GLOBAL_ONT_ID}
  ^\s*Fiber\s*route\s*:\s*${FIBER_ROUTE}
  ^\s*- -> Next
  ^\s*VoIP\s+configure\s+method\s*:\s*${VOIP_CONFIGURE_METHOD}
  ^\s*- -> Next
  ^\s*Line\s+profile\s+ID\s*:\s*${LINE_PROFILE_ID}
  ^\s*Line\s+profile\s+name\s*:\s*${LINE_PROFILE_NAME}
  ^\s*- -> Next
  ^\s*FEC\s+upstream\s+switch\s*:\s*${FEC_UPSTREAM_SWITCH}
  ^\s*OMCC\s+(E|e)ncrypt\s+switch\s*:\s*${OMCC_ENCRYPT_SWITCH}
  ^\s*Qos\s+mode\s*:\s*${QOS_MODE}
  ^\s*Mapping\s+mode\s*:\s*${MAPPING_MODE}
  ^\s*T(R|r)069\s+management\s*:\s*${TR069_MANAGEMENT}
  ^\s*T(R|r)069\s+IP\s+index\s*:\s*${TR069_IP_INDEX}
  ^\s*Notes:\s*\*\s*indicates\s*Discrete\s*TCONT\(TCONT\s*Unbound\) -> Tconts

Tconts
  ^\s*- -> Next
  ^\s*<T-CONT\s*${T_CONT}>\s*DBA\s*Profile-ID\s*:\s*\d+$$ -> Next
  ^\s*Service\s+profile\s+ID\s*:\s*${SERVICE_PROFILE_ID} -> Service_profile


Service_profile
  # ^\s*Service\s+profile\s+ID\s*:\s*${SERVICE_PROFILE_ID}
  ^\s*Service\s+profile\s+name\s*:\s*${SERVICE_PROFILE_NAME}
  ^\s*POTS\s*${NUMBER_POTS}\s*$$
  ^\s*ETH\s*${NUMBER_ETH}\s*$$
  ^\s*VDSL\s*${NUMBER_VDSL}\s*$$
  ^\s*TDM\s*${NUMBER_TDM}\s*$$
  ^\s*MOCA\s*${NUMBER_MOCA}\s*$$
  ^\s*CATV\s*${NUMBER_CATV}\s*$$
  ^\s*POTS\s*${NUMBER_POTS}\s*(${NUMBER_MAX_ADAPTIVE_POTS}|-)\s*$$
  ^\s*ETH\s*${NUMBER_ETH}\s*(${NUMBER_MAX_ADAPTIVE_ETH}|-)\s*$$
  ^\s*VDSL\s*${NUMBER_VDSL}\s*(${NUMBER_MAX_ADAPTIVE_VDSL}|-)\s*$$
  ^\s*TDM\s*${NUMBER_TDM}\s*(${NUMBER_MAX_ADAPTIVE_TDM}|-)\s*$$
  ^\s*MOCA\s*${NUMBER_MOCA}\s*(${NUMBER_MAX_ADAPTIVE_MOCA}|-)\s*$$
  ^\s*CATV\s*${NUMBER_CATV}\s*(${NUMBER_MAX_ADAPTIVE_CATV}|-)\s*$$
  ^\s*TDM\s+port\s+type\s*:\s*${TDM_PORT_TYPE}
  ^\s*TDM\s+service\s+type\s*:\s*${TDM_SERVICE_TYPE}
  ^\s*MAC\s+learning\s+function\s+switch\s*:\s*${MAC_LEARNING_FUNCTION_SWITCH}
  ^\s*ONT\s+transparent\s+function\s+switch\s*:\s*${ONT_TRANSPARENT_FUNCTION_SWITCH}
  ^\s*Ring\s+check\s+switch\s*:\s*${RING_CHECK_SWITCH}
  ^\s*Ring\s+port\s+auto-shutdown\s*:\s*${RING_PORT_AUTO_SHUTDOWN}
  ^\s*Ring\s+detect\s+frequency\s*:\s*${RING_RESUME_FREQUENCY}
  ^\s*Ring\s+resume\s+interval\s*:\s*${RING_RESUME_INTERVAL}
  ^\s*Ring\s*detect\s+period\s*:\s*${RING_DETECT_PERIOD}
  ^ \s+Multicast\s+forward\s+mode\s*:\s*${SERVICE_PROFILE_MULTICAST_FORWARD_MODE}
  ^\s*Multicast\s+forward\s+VLAN\s*:\s*${SERVICE_PROFILE_MULTICAST_FORWARD_VLAN}
  ^\s*Multicast\s+mode\s*:\s*${SERVICE_PROFILE_MULTICAST_MODE}
  ^\s*Upstream\s+IGMP\s+packet\s+forward\s+mode\s*:\s*${UPSTREAM_IGMP_PACKET_FORWARD_MODE}
  ^\s*Upstream\s+IGMP\s+packet\s+forward\s+VLAN\s*:\s*${UPSTREAM_IGMP_PACKET_FORWARD_VLAN}
  ^\s*Upstream\s+IGMP\s+packet\s+priority\s*:\s*${UPSTREAM_IGMP_PACKET_PRIORITY}
  ^\s*Native\s+VLAN\s+option\s*:\s*${NATIVE_VLAN_OPTION}
  ^\s*Upstream\s+PQ\s+color\s+policy\s*:\s*${UPSTREAM_PQ_COLOR_POLICY}
  ^\s*Downstream\s+PQ\s+color\s+policy\s*:\s*${DOWNSTREAM_PQ_COLOR_POLICY}
  ^\s*Monitor\s+link\s*:\s*${MONITOR_LINK}
  ^\s*MTU\(byte\)\s*:\s*${MTU_BYTE}
  ^\s*Alarm\s+policy\s+profile\s+ID\s*:\s*${ALARM_POLICY_PROFILE_ID} -> Alarm_policy

Alarm_policy
  # ^\s*Alarm\s+policy\s+profile\s+ID\s*:\s*${ALARM_POLICY_PROFILE_ID}
  ^\s*Alarm\s+policy\s+profile\s+name\s*:\s*${ALARM_POLICY_PROFILE_NAME}
  ^\s*TR069\s+server\s+profile\s+ID\s*:\s*${TR069_SERVER_PROFILE_ID}
  ^\s*TR069\s+server\s+profile\s+name\s*:\s*${TR069_SERVER_PROFILE_NAME}
`