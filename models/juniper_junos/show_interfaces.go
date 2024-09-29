package juniper_junos 

type ShowInterfaces struct {
	Interface	string	`json:"INTERFACE"`
	Link_status	string	`json:"LINK_STATUS"`
	Admin_state	string	`json:"ADMIN_STATE"`
	Hardware_type	string	`json:"HARDWARE_TYPE"`
	Description	string	`json:"DESCRIPTION"`
	Destination	string	`json:"DESTINATION"`
	Local	string	`json:"LOCAL"`
	Mtu	string	`json:"MTU"`
}

var ShowInterfaces_Template = `Value Required INTERFACE (\S+)
Value LINK_STATUS (\w+)
Value ADMIN_STATE (\S+)
Value HARDWARE_TYPE (\S+)
Value DESCRIPTION (.*)
Value DESTINATION (\S+)
Value LOCAL (\S+)
Value MTU (\d+|Unlimited)

Start
  ^\s*\S+\s+[Ii]nterface -> Continue.Record
  ^Physical\s+interface:\s+${INTERFACE},\s+${ADMIN_STATE},\s+Physical\s+link\s+is\s+${LINK_STATUS}
  ^\s+Interface index:
  ^\s+Description:\s+${DESCRIPTION}$$
  # match "Type: xx, Link-level type:" and "  Link-level type:"
  ^.*\s+Link-level\s+type:\s+${HARDWARE_TYPE},\s+MTU:\s+${MTU}(,.+)?$$
  ^\s+Pad to minimum frame size:
  ^\s+Link-mode
  ^\s+MAC-REWRITE
  ^\s+Flow control
  ^\s+Device [Ff]lags
  ^\s+Interface flags
  ^\s+Link flags
  ^\s+Physical info
  ^\s+CoS queues
  ^\s+Schedulers
  ^\s+Current address
  ^\s+Last flapped
  ^\s+Input rate
  ^\s+Output rate
  ^\s+Active alarms
  ^\s+Active defects
  ^\s+PCS statistics
  ^\s+Bit errors
  ^\s+Errored blocks
  ^\s+Interface transmit statistics
  #
  ^\s+Logical\s+interface\s+${INTERFACE}
  ^\s+Flags:\s+${LINK_STATUS}\s+\S+\s+\w+\s+(\S+\s+)*Encapsulation:
  ^\s+Protocol\s+inet.*,\s+MTU:\s+${MTU}
  ^\s+Flags
  ^\s+Input packets
  ^\s+Output packets
  ^\s+Security
  ^\s+Protocol
  ^\s+Flags
  ^\s+Address
  ^\s+Bandwidth
  ^\s+Routing
  ^\s+Max
  ^\s+Destination:\s+${DESTINATION},\s+Local:\s+${LOCAL},.*
  ^\s+MAC:
  ^{master:\d+}
  ^\s*$$
  ^. -> Error
`