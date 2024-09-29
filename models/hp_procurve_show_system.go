package models

type HpProcurveShowSystem struct {
	Name	string	`json:"NAME"`
	Contact	string	`json:"CONTACT"`
	Location	string	`json:"LOCATION"`
	Mac_age	string	`json:"MAC_AGE"`
	Timezone	string	`json:"TIMEZONE"`
	Daylight_rule	string	`json:"DAYLIGHT_RULE"`
	Software_version	string	`json:"SOFTWARE_VERSION"`
	Rom_version	string	`json:"ROM_VERSION"`
	Allow_mods	string	`json:"ALLOW_MODS"`
	Mac_address	string	`json:"MAC_ADDRESS"`
	Serial	string	`json:"SERIAL"`
	Uptime	string	`json:"UPTIME"`
	Cpu_util	string	`json:"CPU_UTIL"`
	Mem_tot	string	`json:"MEM_TOT"`
	Mem_free	string	`json:"MEM_FREE"`
	Packets_rx	string	`json:"PACKETS_RX"`
	Packets_tx	string	`json:"PACKETS_TX"`
	Packets_tot	string	`json:"PACKETS_TOT"`
	Buffers_free	string	`json:"BUFFERS_FREE"`
	Buffers_lowest	string	`json:"BUFFERS_LOWEST"`
	Buffers_missed	string	`json:"BUFFERS_MISSED"`
}

var HpProcurveShowSystem_Template = `Value NAME (\S+)
Value CONTACT (.+)
Value LOCATION (.+)
Value MAC_AGE (\d+)
Value TIMEZONE (\S+)
Value DAYLIGHT_RULE (\S+)
Value SOFTWARE_VERSION (\S+)
Value ROM_VERSION (\S+)
Value ALLOW_MODS ([Yy]es|[Nn]o)
Value MAC_ADDRESS (\S+)
Value SERIAL (\S+)
Value UPTIME (\d+ \w+)
Value CPU_UTIL (\d+)
Value MEM_TOT ([\d,]+)
Value MEM_FREE ([\d,]+)
Value PACKETS_RX ([\d,]+)
Value PACKETS_TX ([\d,]+)
Value PACKETS_TOT (\d+)
Value BUFFERS_FREE (\d+)
Value BUFFERS_LOWEST (\d+)
Value BUFFERS_MISSED (\d+)

Start
  ^.*Status and Counters -> INFO

INFO
  ^\s*System Name\s+:\s+${NAME}
  ^\s*System Contact\s+:\s+${CONTACT}
  ^\s*System Location\s+:\s+${LOCATION}
  ^\s*MAC Age Time[^:]*:\s+${MAC_AGE}
  ^\s*Time Zone\s+:\s+${TIMEZONE}
  ^\s*Daylight Time Rule\s+:\s+${DAYLIGHT_RULE}
  ^\s*Software revision\s+:\s+${SOFTWARE_VERSION}\s+Base MAC Addr\s+:\s+${MAC_ADDRESS}
  ^\s*ROM Version\s+:\s+${ROM_VERSION}\s+Serial Number\s+:\s+${SERIAL}
  ^\s*Allow V1 Modules\s+:\s+${ALLOW_MODS}
  ^\s*Up Time\s+:\s+${UPTIME}\s+Memory\s+- Total\s+:\s+${MEM_TOT}
  ^\s*CPU Util[^:]*:\s+${CPU_UTIL}\s+Free\s+:\s+${MEM_FREE}
  ^\s*IP Mgmt\s+- Pkts Rx\s+:\s+${PACKETS_RX}\s+Packet\s+- Total\s+:\s+${PACKETS_TOT}
  ^\s*Pkts Tx\s+:\s+${PACKETS_TX}\s+Buffers\s+Free\s+:\s+${BUFFERS_FREE}
  ^\s*Lowest\s+:\s+${BUFFERS_LOWEST}
  ^\s*Missed\s+:\s+${BUFFERS_MISSED} -> Record
  
`