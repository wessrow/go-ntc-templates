package cisco_wlc 

type SshShowApConfigGeneral struct {
	Identifier	string	`json:"IDENTIFIER"`
	Name	string	`json:"NAME"`
	Country_code	string	`json:"COUNTRY_CODE"`
	Country	string	`json:"COUNTRY"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Gateway	string	`json:"GATEWAY"`
	Ap_group	string	`json:"AP_GROUP"`
	Primary_switch_name	string	`json:"PRIMARY_SWITCH_NAME"`
	Primary_switch_ip	string	`json:"PRIMARY_SWITCH_IP"`
	Secondary_switch_name	string	`json:"SECONDARY_SWITCH_NAME"`
	Secondary_switch_ip	string	`json:"SECONDARY_SWITCH_IP"`
	Tertiary_switch_name	string	`json:"TERTIARY_SWITCH_NAME"`
	Tertiary_switch_ip	string	`json:"TERTIARY_SWITCH_IP"`
	Administrative_state	string	`json:"ADMINISTRATIVE_STATE"`
	Operation_state	string	`json:"OPERATION_STATE"`
	Mode	string	`json:"MODE"`
	Model	string	`json:"MODEL"`
	Image	string	`json:"IMAGE"`
	Version	string	`json:"VERSION"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Flexconnect_vlan_mode	string	`json:"FLEXCONNECT_VLAN_MODE"`
	Uptime	string	`json:"UPTIME"`
	Lwapp_uptime	string	`json:"LWAPP_UPTIME"`
	Join_date_time	string	`json:"JOIN_DATE_TIME"`
	Join_taken_time	string	`json:"JOIN_TAKEN_TIME"`
}

var SshShowApConfigGeneral_Template = `Value IDENTIFIER (.+?)
Value NAME (.+?)
Value COUNTRY_CODE (\S+)
Value COUNTRY (.+?)
Value MAC_ADDRESS (.+?)
Value IP_ADDRESS (.+?)
Value NETMASK (.+?)
Value GATEWAY (.+?)
Value AP_GROUP (.+?)
Value PRIMARY_SWITCH_NAME (.+?)
Value PRIMARY_SWITCH_IP (.+?)
Value SECONDARY_SWITCH_NAME (.+?)
Value SECONDARY_SWITCH_IP (.+?)
Value TERTIARY_SWITCH_NAME (.+?)
Value TERTIARY_SWITCH_IP (.+?)
Value ADMINISTRATIVE_STATE (.+?)
Value OPERATION_STATE (.+?)
Value MODE (.+?)
Value MODEL (.+?)
Value IMAGE (.+?)
Value VERSION (.+?)
Value SERIAL_NUMBER (.+?)
Value FLEXCONNECT_VLAN_MODE (.+?)
Value UPTIME (.+?)
Value LWAPP_UPTIME (.+?)
Value JOIN_DATE_TIME (.+?)
Value JOIN_TAKEN_TIME (.+?)

Start
  ^Cisco\s+AP\s+Identifier\.*\s+${IDENTIFIER}\s*$$
  ^Cisco\s+AP\s+Name\.*\s+${NAME}\s*$$
  ^Country\s+code\.*\s+${COUNTRY_CODE}\s+-\s+${COUNTRY}\s*$$
  ^MAC\s+Address\.*\s+${MAC_ADDRESS}\s*$$
  ^IP\s+Address\.*\s+${IP_ADDRESS}\s*$$
  ^IP\s+NetMask\.*\s+${NETMASK}\s*$$
  ^Gateway\s+IP\s+Addr\.*\s+${GATEWAY}\s*$$
  ^Cisco\s+AP\s+Group\s+Name\.*\s+${AP_GROUP}\s*$$
  ^Primary\s+Cisco\s+Switch\s+Name\.*\s+${PRIMARY_SWITCH_NAME}\s*$$
  ^Primary\s+Cisco\s+Switch\s+IP\s+Address\.*\s+${PRIMARY_SWITCH_IP}\s*$$
  ^Secondary\s+Cisco\s+Switch\s+Name\.*\s+${SECONDARY_SWITCH_NAME}\s*$$
  ^Secondary\s+Cisco\s+Switch\s+IP\s+Address\.*\s+${SECONDARY_SWITCH_IP}\s*$$
  ^Tertiary\s+Cisco\s+Switch\s+Name\.*\s+${TERTIARY_SWITCH_NAME}\s*$$
  ^Tertiary\s+Cisco\s+Switch\s+IP\s+Address\.*\s+${TERTIARY_SWITCH_IP}\s*$$
  ^Administrative\s+State\s+\.*\s+${ADMINISTRATIVE_STATE}\s*$$
  ^Operation\s+State\s+\.*\s+${OPERATION_STATE}\s*$$
  ^AP\s+Mode\s+\.*\s+${MODE}\s*$$
  ^AP\s+Model\.*\s+${MODEL}\s*$$
  ^AP\s+Image\.*\s+${IMAGE}\s*$$
  ^IOS\s+Version\.*\s+${VERSION}\s*$$
  ^AP\s+Serial\s+Number\.*\s+${SERIAL_NUMBER}\s*$$
  ^FlexConnect\s+Vlan\s+mode\s+:\.+\s+${FLEXCONNECT_VLAN_MODE}\s*$$
  ^AP\s+Up\s+Time\.*\s+${UPTIME}\s*$$
  ^AP\s+LWAPP\s+Up\s+Time\.*\s+${LWAPP_UPTIME}\s*$$
  ^Join\s+Date\s+and\s+Time\.*\s+${JOIN_DATE_TIME}\s*$$
  ^Join\s+Taken\s+Time\.*\s+${JOIN_TAKEN_TIME}\s*$$ -> Record
  ^.+\.+
  ^-+
  ^\s*$$
  ^\S+\s+VLAN.+Mappings
  ^\s+Template\s+in\s+
  ^AP-Specific\s+FlexConnect
  ^FlexConnect\s+Local-Split
  ^WLAN\s+ID\s+PROFILE\s+NAME\s+ACL\s+TYPE\s*$$
  ^\s+Flexconnect\s+Central-Dhcp
  ^WLAN\s+ID\s+PROFILE NAME\s+Central-Dhcp\s+DNS\s+Override\s+Nat-Pat\s+Type\s*$$
  ^\s+\d+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s*$$
  ^\s+\d+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s*$$
  ^\s+\d+\s+\S+\s+\S+\s+\S+\s+\S+\s+\S+\s*$$
  ^WlanId\s+PROFILE NAME\s+Inherit-level\s+Visibility\s+Flex\s+Avc-profile\s*$$
  ^\d+\s+\S+\s+\S+\s+\S+\s+\S+\s*$$
  ^\d+\s+\S+\s+\S+\s+\S+\s+\S+\s*$$
  ^\d+\s+\S+\s+\S+\s+\S+\s+\S+\s*$$
  ^FlexConnect\s+Backup\s+Auth
  ^. -> Error
`