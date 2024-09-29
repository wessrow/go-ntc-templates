package huawei_smartax 

type DisplayPortInfo struct {
	Fsp	string	`json:"FSP"`
	Min_distance_km	string	`json:"MIN_DISTANCE_KM"`
	Max_distance_km	string	`json:"MAX_DISTANCE_KM"`
	Max_guaranteed_bandwidth_kbps	string	`json:"MAX_GUARANTEED_BANDWIDTH_KBPS"`
	Left_guaranteed_bandwidth_kbps	string	`json:"LEFT_GUARANTEED_BANDWIDTH_KBPS"`
	Number_of_t_conts	string	`json:"NUMBER_OF_T_CONTS"`
	Autofind	string	`json:"AUTOFIND"`
	Fec_check	string	`json:"FEC_CHECK"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Ont_encryption_key_switching_interval_m	string	`json:"ONT_ENCRYPTION_KEY_SWITCHING_INTERVAL_M"`
	Pon_id_switch	string	`json:"PON_ID_SWITCH"`
	Pon_id_identifier	string	`json:"PON_ID_IDENTIFIER"`
	Jumbo_frame_switch	string	`json:"JUMBO_FRAME_SWITCH"`
	Port_mtu_bytes	string	`json:"PORT_MTU_BYTES"`
	Surplus_bandwidth_assignment	string	`json:"SURPLUS_BANDWIDTH_ASSIGNMENT"`
	Best_effort_bandwidth_assignment	string	`json:"BEST_EFFORT_BANDWIDTH_ASSIGNMENT"`
	Traffic_alarm_profile_id	string	`json:"TRAFFIC_ALARM_PROFILE_ID"`
	Ont_online_power_threshold_dbm	string	`json:"ONT_ONLINE_POWER_THRESHOLD_DBM"`
	Low_latency	string	`json:"LOW_LATENCY"`
	Multichannel_low_latency	string	`json:"MULTICHANNEL_LOW_LATENCY"`
	Optical_module_work_mode	string	`json:"OPTICAL_MODULE_WORK_MODE"`
	Channel_type	string	`json:"CHANNEL_TYPE"`
	Online_ont_number_threshold	string	`json:"ONLINE_ONT_NUMBER_THRESHOLD"`
}

var DisplayPortInfo_Template = `Value FSP (\d+\/\d+\/\d+)
Value MIN_DISTANCE_KM (\d+)
Value MAX_DISTANCE_KM (\d+)
Value MAX_GUARANTEED_BANDWIDTH_KBPS (\d*)
Value LEFT_GUARANTEED_BANDWIDTH_KBPS (\d*)
Value NUMBER_OF_T_CONTS (\d+)
Value AUTOFIND (Enable|Disable)
Value FEC_CHECK (Enable|Disable)
Value ADMIN_STATE (\w+)
Value ONT_ENCRYPTION_KEY_SWITCHING_INTERVAL_M (\d+)
Value PON_ID_SWITCH (Disable|Enable)
Value PON_ID_IDENTIFIER (\d+)
Value JUMBO_FRAME_SWITCH (Disable|Enable)
Value PORT_MTU_BYTES (\d+)
Value SURPLUS_BANDWIDTH_ASSIGNMENT (Enable|Disable)
Value BEST_EFFORT_BANDWIDTH_ASSIGNMENT (\w*)
Value TRAFFIC_ALARM_PROFILE_ID (\w*)
Value ONT_ONLINE_POWER_THRESHOLD_DBM (\w*)
Value LOW_LATENCY (yes|no)
Value MULTICHANNEL_LOW_LATENCY (Enable|Disable)
Value OPTICAL_MODULE_WORK_MODE (\w*)
Value CHANNEL_TYPE (\w*)
Value ONLINE_ONT_NUMBER_THRESHOLD (Enable|Disable)

Start
  ^\s*F\/S\/P\s*${FSP}
  ^\s*Min\s*distance\(km\)\s*${MIN_DISTANCE_KM}
  ^\s*Max\s*distance\(km\)\s*${MAX_DISTANCE_KM}
  ^\s*Max\s*guaranteed\s*bandwidth\(kbps\)\s*${MAX_GUARANTEED_BANDWIDTH_KBPS}
  ^\s*Left\s*guaranteed\s*bandwidth\(kbps\)\s*${LEFT_GUARANTEED_BANDWIDTH_KBPS}
  ^\s*Number\s*of\s*T-CONTs\s*${NUMBER_OF_T_CONTS}
  ^\s*Autofind\s*${AUTOFIND}
  ^\s*FEC\s*check\s*${FEC_CHECK}
  ^\s*Admin\s*State\s*${ADMIN_STATE}
  ^\s*ONT\s*encryption\s*key\s*switching\s*interval\(m\)\s*${ONT_ENCRYPTION_KEY_SWITCHING_INTERVAL_M}
  ^\s*PON-ID\s*switch\s*${PON_ID_SWITCH}
  ^\s*PON-ID\s*identifier\s*${PON_ID_IDENTIFIER}
  ^\s*Jumbo\s*frame\s*switch\s*${JUMBO_FRAME_SWITCH}
  ^\s*Port\s*MTU\(bytes\)\s*${PORT_MTU_BYTES}
  ^\s*Surplus\s*bandwidth\s*assignment\s*${SURPLUS_BANDWIDTH_ASSIGNMENT}
  ^\s*Best-effort\s*bandwidth\s*assignment\s*${BEST_EFFORT_BANDWIDTH_ASSIGNMENT}
  ^\s*Traffic\s*alarm-profile\s*ID\s*${TRAFFIC_ALARM_PROFILE_ID}
  ^\s*ONT\s*online\s*power\s*threshold\(dBm\)\s*${ONT_ONLINE_POWER_THRESHOLD_DBM}
  ^\s*Low-latency\s*${LOW_LATENCY}
  ^\s*Multichannel\s*low\s*latency\s*${MULTICHANNEL_LOW_LATENCY}
  ^\s*Optical\s*module\s*work\s*mode\s*${OPTICAL_MODULE_WORK_MODE}
  ^\s*Channel\w*\s*Information\s*
  ^\s*Channel\s*Type\s*${CHANNEL_TYPE}
  ^\s*Online\s*ONT\s*number\s*threshold\s*${ONLINE_ONT_NUMBER_THRESHOLD}
  ^\s*
  ^. -> Error`