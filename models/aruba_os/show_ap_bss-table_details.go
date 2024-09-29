package aruba_os 

type ShowApBssTableDetails struct {
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

var ShowApBssTableDetails_Template = `Value BSS ([a-fA-F0-9]{2}(:[a-fA-F0-9]{2}){5})
Value ESS (\S+(\s\S+){0,})
Value PORT (\S+)
Value IP_ADDRESS ((([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))
Value BAND (((2.4|5|6).*z)|N/A)
Value CHANNEL (\S+)
Value TYPE (\S+)
Value CUR_CL (\S+)
Value AP_NAME (\S+)
Value IN_T (\S+)
Value TOT_T (\S+)
Value MTU ((\d+|N/A|.*?))
Value PHY (\S+)
Value ACL_STATE (\S+)
Value ACL (\S+)
Value FM (\S+)
Value FLAGS ((\S+|.*?))
Value CLUSTER ((A|sA|sU|U|N/A).*?)
Value ACTIVE_CLIENTS (\S+)
Value STANDBY_CLIENTS (\S+)
Value DATAZONE (\S+)


Start
  ^bss\s+ess\s+port\s+ip\s+phy\s* -> Before_ArubaOS_8_9
  ^bss\s+ess\s+port\s+ip\s+band\s* -> After_ArubaOS_8_9
  ^Aruba AP
  ^.+\.+
  ^bss
  ^fm
  ^\s*$$
  ^\s+
  ^cluster
  ^Channel
  ^"Spectrum"
  ^Num
  ^Flags
  ^-+
  ^. -> Error


Before_ArubaOS_8_9
  ^-+\s+-+
  ^${BSS}\s+${ESS}\s+${PORT}\s+${IP_ADDRESS}\s+${PHY}\s+${TYPE}\s+${CHANNEL}\s+${CUR_CL}\s+${AP_NAME}\s+${IN_T}\s+${TOT_T}\s+${MTU}\s+${ACL_STATE}\s+${ACL}\s+${FM}\s+${FLAGS}\s+${CLUSTER}\s+ -> Record
  ^Aruba AP
  ^-\s*
  ^bss
  ^Channel -> Start
  ^Num -> End
  ^. -> Error

After_ArubaOS_8_9
  ^-+\s+-+
  ^${BSS}\s+${ESS}\s+${PORT}\s+${IP_ADDRESS}\s+${BAND}\s+${CHANNEL}\s+${TYPE}\s+${CUR_CL}\s+${AP_NAME}\s+${IN_T}\s+${TOT_T}\s+${MTU}\s+${ACL_STATE}\s+${ACL}\s+${FM}\s+${FLAGS}\s+${CLUSTER}\s+${ACTIVE_CLIENTS}\s+${STANDBY_CLIENTS}\s+${DATAZONE}\s*$$ -> Record
  ^Channel -> Start
  ^Num -> End
  ^. -> Error
`