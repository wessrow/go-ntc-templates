package cisco_asa

type ShowVersion struct {
	Model              string   `json:"MODEL"`
	Failover           string   `json:"FAILOVER"`
	Cluster            string   `json:"CLUSTER"`
	Version            string   `json:"VERSION"`
	Device_mgr_version string   `json:"DEVICE_MGR_VERSION"`
	Compile_date       string   `json:"COMPILE_DATE"`
	Flash              string   `json:"FLASH"`
	Interfaces         []string `json:"INTERFACES"`
	Max_intf           string   `json:"MAX_INTF"`
	Max_vlans          string   `json:"MAX_VLANS"`
	Serial             []string `json:"SERIAL"`
	Image              string   `json:"IMAGE"`
	Hostname           string   `json:"HOSTNAME"`
	Hardware           string   `json:"HARDWARE"`
	License_mode       string   `json:"LICENSE_MODE"`
	Uptime             string   `json:"UPTIME"`
	License_state      string   `json:"LICENSE_STATE"`
	Last_mod           string   `json:"LAST_MOD"`
}

var ShowVersion_Template string = `Value VERSION (\S+)
Value DEVICE_MGR_VERSION (\S+)
Value COMPILE_DATE (\d+-\w+-\d+)
Value IMAGE (\S+)
Value HOSTNAME (\S+)
Value UPTIME (.+)
Value HARDWARE (.+?)
Value MODEL (\S+)
Value FLASH (\S+)
Value List INTERFACES (\S+)
Value LICENSE_MODE (.+)
Value LICENSE_STATE (.+)
Value MAX_INTF (\S+)
Value MAX_VLANS (\d+)
Value FAILOVER (\S+)
Value CLUSTER (\S+)
Value List SERIAL (\S+)
Value LAST_MOD (.+)

Start
  ^.*Software\sVersion\s${VERSION}
  ^Device.+\s${DEVICE_MGR_VERSION}
  ^Compiled\s+on\s+\w+\s+${COMPILE_DATE}.*
  ^System image file.+"${IMAGE}"
  ^${HOSTNAME} up ${UPTIME}
  ^Hardware:\s+${HARDWARE},?\s*$$
  ^Model Id:\s+${MODEL}
  ^Internal.+Flash,\s${FLASH}
  ^\s*\d+:.\S+\s*${INTERFACES}.*
  ^License mode:\s${LICENSE_MODE}
  ^.+License State:\s${LICENSE_STATE}
  ^Maximum Physical.+:\s${MAX_INTF}
  ^Maximum VLANs.+:\s${MAX_VLANS}
  ^Failover\s+:\s${FAILOVER}
  ^Cluster\s+:\s${CLUSTER}
  ^Serial Number:\s${SERIAL}
  ^.+last modified by\s${LAST_MOD}
  ^Config\s+file\s+at\s+boot\s+was
  ^Slot\s+\d:
  ^Firepower\s+Extensible\s+Operating\s+System\s+Version
  ^BIOS\s+Flash\s+Firmware\s+Hub
  ^No\s+active\s+entitlement:
  ^[F|f]ailover\s+cluster
  ^Licensed\s+features\s+for\s+this.*:
  ^BIOS\s+Flash
  ^\*Memory\s+resource\s+allocation
  ^Licensed\s+features\s+for
  ^Encryption\s+hardware\s+device
  ^Encryption\-
  ^Inside\s+Hosts\s*:
  ^\s*Number\s+of\s+accelerators\s*:
  ^Carrier\s*:
  ^Security\s+Contexts\s*:
  ^AnyConnect\s+.*:
  ^\w+\s+VPN\s+Peers\s*:
  ^Advanced\s+Endpoint\s+Assessment\s*:
  ^Shared\s+License
  ^Failover\s+cluster\s+licensed\s+features\s+for\s+this\s+user\s+context\s*:
  ^Total\s+\w+\s+Proxy\s+Sessions\s*:
  ^Botnet\s+Traffic\s+Filter\s*:
  ^Image\s+type\s*:
  ^Cluster\s+Members\s*:
  ^Key\s+version\s*:
  ^VPN\s+Load\s+Balancing\s*:
  ^Configuration\s+register\s+is
  ^Key\s+Version\s*:
  ^This\s+platform\s+has\s+a
  ^Running\s+Permanent\s+Activation\s+Key\s*:
  ^Configuration\s+has\s+not\s+been\s+modified
  ^\s{6,}\d+\s+CPUs
  ^\s*Boot\s+microcode\s*:
  ^\s*\S+\s+microcode\s*:
  ^\s*Programmable\s+device\s*:
  ^\s*VPN\-\d*DES.*:
  ^GTP\/GPRS\s*:
  ^.*Proxy\s+Sessions\s*:
  ^Intercompany\s+Media\s+Engine\s*:
  ^\d+GE\s+I\/O\s*:
  ^SSP\s+Operating\s+System\s+Version
  ^Start\-up\s+time
  ^Active\s+entitlement\s*:
  ^Firewall\s+throughput
  ^\s+ASA:\s+\d+\s+MB\s+RAM
  ^Baseboard\s+Management\s+Controller
  ^IPS\s+Module\s*:
  ^The\s+Running\s+Activation\s+Key\s+feature\s*:
  ^FPGA\s+.*\s*:
  ^ROMMON.*:
  ^VPN.*:
  ^WebVPN.*:
  ^WARNING:
  ^Running\s+Activation\s+Key\s*:
  ^\s*$$
  ^. -> Error
`
