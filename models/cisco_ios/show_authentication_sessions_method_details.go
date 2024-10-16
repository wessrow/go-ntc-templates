package cisco_ios

type ShowAuthenticationSessionsMethodDetails struct {
	Current_policy                      string   `json:"CURRENT_POLICY"`
	Server_policy_vlan_group            string   `json:"SERVER_POLICY_VLAN_GROUP"`
	Server_policy_vn                    string   `json:"SERVER_POLICY_VN"`
	Method_state_list                   []string `json:"METHOD_STATE_LIST"`
	Interface                           string   `json:"INTERFACE"`
	Mac_address                         string   `json:"MAC_ADDRESS"`
	Status                              string   `json:"STATUS"`
	Session_timeout                     string   `json:"SESSION_TIMEOUT"`
	Ipv6_address                        string   `json:"IPV6_ADDRESS"`
	Server_policy_sgt                   []string `json:"SERVER_POLICY_SGT"`
	Resultant_policy_vlan_group         string   `json:"RESULTANT_POLICY_VLAN_GROUP"`
	Operational_control_dir             string   `json:"OPERATIONAL_CONTROL_DIR"`
	Accounting_update_remaining_seconds string   `json:"ACCOUNTING_UPDATE_REMAINING_SECONDS"`
	Server_template                     string   `json:"SERVER_TEMPLATE"`
	Method_type_list                    []string `json:"METHOD_TYPE_LIST"`
	Ipv4_address                        string   `json:"IPV4_ADDRESS"`
	Device_name                         string   `json:"DEVICE_NAME"`
	Domain                              string   `json:"DOMAIN"`
	Operational_host_mode               string   `json:"OPERATIONAL_HOST_MODE"`
	Server_session_timeout              string   `json:"SERVER_SESSION_TIMEOUT"`
	Resultant_policy_sgt                []string `json:"RESULTANT_POLICY_SGT"`
	Resultant_policy_vn                 string   `json:"RESULTANT_POLICY_VN"`
	Username                            string   `json:"USERNAME"`
	Device_type                         string   `json:"DEVICE_TYPE"`
	Timeout_action                      string   `json:"TIMEOUT_ACTION"`
	Accounting_update_seconds           string   `json:"ACCOUNTING_UPDATE_SECONDS"`
}

var ShowAuthenticationSessionsMethodDetails_Template string = `Value INTERFACE (\S+)
Value MAC_ADDRESS (\S+)
Value IPV4_ADDRESS (\S+)
Value IPV6_ADDRESS (\S+)
Value USERNAME (\S+)
Value DEVICE_TYPE (\S+)
Value DEVICE_NAME (\S+)
Value STATUS (\S+)
Value DOMAIN (\S+)
Value OPERATIONAL_HOST_MODE (\S+)
Value OPERATIONAL_CONTROL_DIR (\S+)
Value SESSION_TIMEOUT (\S+)
Value TIMEOUT_ACTION (\S+)
Value ACCOUNTING_UPDATE_SECONDS (\d+)
Value ACCOUNTING_UPDATE_REMAINING_SECONDS (\d+)
Value CURRENT_POLICY (\S+)
Value SERVER_POLICY_VLAN_GROUP (\d+)
Value List SERVER_POLICY_SGT (\d+)
Value SERVER_POLICY_VN (\S+)
Value SERVER_SESSION_TIMEOUT (\d+)
Value SERVER_TEMPLATE (\S+)
Value RESULTANT_POLICY_VLAN_GROUP (\d+)
Value List RESULTANT_POLICY_SGT (\d+)
Value RESULTANT_POLICY_VN (\S+)
Value List METHOD_TYPE_LIST (\S+)
Value List METHOD_STATE_LIST (.*)

Start
  ^\s*Interface: -> Continue.Record
  ^\s*Interface:\s+${INTERFACE}
  ^\s*IIF-ID:\s+.*
  ^\s*MAC Address:\s+${MAC_ADDRESS}
  ^\s*IPv6 Address:\s+${IPV6_ADDRESS}
  ^\s*IPv4 Address:\s+${IPV4_ADDRESS}
  ^\s*User-Name:\s+${USERNAME}
  ^\s*Device-type:\s+${DEVICE_TYPE}
  ^\s*Device-name:\s+${DEVICE_NAME}
  ^\s*Status:\s+${STATUS}
  ^\s*Domain:\s+${DOMAIN}
  ^\s*Oper\s+host\s+mode:\s+${OPERATIONAL_HOST_MODE}
  ^\s*Oper\s+control\s+dir:\s+${OPERATIONAL_CONTROL_DIR}
  ^\s*Session\s+timeout:\s+${SESSION_TIMEOUT}
  ^\s*Timeout\s+action:\s+${TIMEOUT_ACTION}
  ^\s*Acct\s+update\s+timeout:\s+${ACCOUNTING_UPDATE_SECONDS}s\s+\(local\),\s+Remaining:\s+${ACCOUNTING_UPDATE_REMAINING_SECONDS}s
  ^\s*Common\s+Session\s+ID:\s+.*
  ^\s*Acct\s+Session\s+ID:\s+.*
  ^\s*Handle:\s+.*
  ^\s*Current\s+Policy:\s+${CURRENT_POLICY}
  ^\s*Local\s+Policies: -> LOCAL_POLICIES
  ^Server\s+Policies: -> SERVER_POLICIES
  ^Resultant\s+Policies: -> RESULTANT_POLICIES
  ^-+\s*$$
  ^\s*$$
  ^. -> Error

LOCAL_POLICIES
  ^\s*$$
  ^Server Policies: -> SERVER_POLICIES
  ^Resultant\s+Policies: -> RESULTANT_POLICIES
  ^. -> Error LocalPolicy

SERVER_POLICIES
  ^\s*Vlan\s+Group:\s+Vlan:\s+${SERVER_POLICY_VLAN_GROUP}
  ^\s*SGT\s+Value:\s+${SERVER_POLICY_SGT}
  ^\s*VN\s+Value:\s+${SERVER_POLICY_VN}
  ^Resultant\s+Policies: -> RESULTANT_POLICIES
  ^\s*Method\s+status\s+list: -> METHOD_STATUS
  ^\s*Session-Timeout:\s+${SERVER_SESSION_TIMEOUT} sec
  ^\s*Interface\s+Template:\s+${SERVER_TEMPLATE}
  ^-+\s*$$ -> Start
  ^\s*$$
  ^. -> Error ServerPolicy

RESULTANT_POLICIES
  ^\s*Vlan\s+Group:\s+Vlan:\s+${RESULTANT_POLICY_VLAN_GROUP}
  ^\s*SGT\s+Value:\s+${RESULTANT_POLICY_SGT}
  ^\s*VN\s+Value:\s+${RESULTANT_POLICY_VN}
  ^-+\s*$$ -> Start
  ^\s*Method\s+status\s+list: -> METHOD_STATUS
  ^\s*$$
  ^. -> Error ResultantPolicy

METHOD_STATUS
  ^\s+Method\s+State
  ^\s+${METHOD_TYPE_LIST}\s+${METHOD_STATE_LIST}\s*$$
  ^-+\s*$$ -> Start
  ^\s*$$
  ^. -> Error MethodStatus
`
