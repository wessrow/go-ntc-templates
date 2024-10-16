package aruba_os

type ShowApBssTableDetails struct {
	Tot_t           string `json:"TOT_T"`
	Acl             string `json:"ACL"`
	Cluster         string `json:"CLUSTER"`
	Standby_clients string `json:"STANDBY_CLIENTS"`
	Bss             string `json:"BSS"`
	Ip_address      string `json:"IP_ADDRESS"`
	Cur_cl          string `json:"CUR_CL"`
	Flags           string `json:"FLAGS"`
	Datazone        string `json:"DATAZONE"`
	Ess             string `json:"ESS"`
	Channel         string `json:"CHANNEL"`
	In_t            string `json:"IN_T"`
	Mtu             string `json:"MTU"`
	Acl_state       string `json:"ACL_STATE"`
	Active_clients  string `json:"ACTIVE_CLIENTS"`
	Port            string `json:"PORT"`
	Type            string `json:"TYPE"`
	Ap_name         string `json:"AP_NAME"`
	Phy             string `json:"PHY"`
	Fm              string `json:"FM"`
	Band            string `json:"BAND"`
}

var ShowApBssTableDetails_Template string = `Value BSS ([a-fA-F0-9]{2}(:[a-fA-F0-9]{2}){5})
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
