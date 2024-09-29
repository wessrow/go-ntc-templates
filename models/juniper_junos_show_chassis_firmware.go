package models

type JuniperJunosShowChassisFirmware struct {
	Part	string	`json:"PART"`
	Type	string	`json:"TYPE"`
	Version	string	`json:"VERSION"`
}

var JuniperJunosShowChassisFirmware_Template = `Value Filldown PART (\w+\s\w+\s\w+|\w+\s\d+?|\w+)
Value TYPE (\S+)
Value VERSION (\d+\.+.+?|\S+\s+\S+)

Start
  ^[Pp]art
  ^${PART}\s+${TYPE}?\s+?(.+Version\s+|.+U-Boot\s.+|.+?)${VERSION}(Copy\S+$$|\s+.+$$|$$) -> Record
  ^\s+${TYPE}\s+?(.+Version\s+|.+U-Boot\s.+|.+?)${VERSION}(Copy\S+$$|\s+.+$$|$$) -> Record
  ^\s*$$
  ^{master:\d+}

EOF

`