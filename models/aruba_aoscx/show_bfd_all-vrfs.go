package aruba_aoscx 

type ShowBfdAllVrfs struct {
	Admin_status	string	`json:"ADMIN_STATUS"`
	Src_ip	string	`json:"SRC_IP"`
	Session	string	`json:"SESSION"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
	Source_ip	string	`json:"SOURCE_IP"`
	Destination_ip	string	`json:"DESTINATION_IP"`
	Echo	string	`json:"ECHO"`
	State	string	`json:"STATE"`
	Application	string	`json:"APPLICATION"`
}

var ShowBfdAllVrfs_Template = `Value Filldown ADMIN_STATUS (\w+)
Value Filldown SRC_IP (\S+)
Value SESSION (\d+)
Value INTERFACE (\S+)
Value VRF (\S+)
Value SOURCE_IP (\d+\.\d+\.\d+\.\d+)
Value DESTINATION_IP (\d+\.\d+\.\d+\.\d+)
Value ECHO (\w+)
Value STATE (\w+)
Value APPLICATION (\w+)


Start
  ^Admin\s+status:\s+${ADMIN_STATUS}
  ^Echo\s+source\s+IP:\s+${SRC_IP}
  ^Statistics:.*
  ^Total.*$$
  ^Session\s+Interface\s+VRF\s+Source\s+IP\s+Destination\s+IP\s+Echo\s+State\s+Protocol
  ^\s*-+
  ^${SESSION}\s+${INTERFACE}\s+${VRF}\s+${SOURCE_IP}\s+${DESTINATION_IP}\s+${ECHO}\s+${STATE}\s+${APPLICATION} -> Record
  ^. -> Error

EOF`