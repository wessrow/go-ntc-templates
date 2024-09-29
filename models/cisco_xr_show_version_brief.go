package models

type CiscoXrShowVersionBrief struct {
	Version	string	`json:"VERSION"`
	Hostname	string	`json:"HOSTNAME"`
	Uptime	string	`json:"UPTIME"`
	Boot_image	string	`json:"BOOT_IMAGE"`
	Family	string	`json:"FAMILY"`
	Model	string	`json:"MODEL"`
}

var CiscoXrShowVersionBrief_Template = `Value VERSION (\S[^\[]+)
Value HOSTNAME (\S+)
Value UPTIME (.+?)
Value BOOT_IMAGE (\S+)
Value FAMILY (\S+)
Value MODEL (ASR \S+)


Start
  ^\s*Cisco .+ Software, Version\s+${VERSION}
  ^\s*${HOSTNAME} uptime is ${UPTIME}\s*$$
  ^\s*System image file is "${BOOT_IMAGE}"\s*$$
  ^\s*cisco ${FAMILY}.*processor.+$$
  ^\s*${MODEL}.+$$ -> Record
`