package cisco_nvfis 

type ShowRunningConfigSnmpHost struct {
	Username	string	`json:"USERNAME"`
	Hostname	string	`json:"HOSTNAME"`
	Version	string	`json:"VERSION"`
	Port	string	`json:"PORT"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Securitylevel	string	`json:"SECURITYLEVEL"`
}

var ShowRunningConfigSnmpHost_Template = `Value USERNAME ([-\w]+)
Value HOSTNAME ([-\w]+)
Value VERSION (\d+)
Value PORT (\d+)
Value IP_ADDRESS ([0-9A-F:.]+)
Value SECURITYLEVEL (\w+)


Start
  ^snmp\shost\s${HOSTNAME}
  ^\shost-port\s+${PORT}
  ^\shost-ip-address\s+${IP_ADDRESS}
  ^\shost-version\s+${VERSION}
  ^\shost-security-level\s+${SECURITYLEVEL}
  ^\shost-user-name\s+${USERNAME} 
  ^! -> Record

EOF`