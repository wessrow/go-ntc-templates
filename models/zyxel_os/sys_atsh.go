package zyxel_os

type SysAtsh struct {
	Boot_version   string `json:"BOOT_VERSION"`
	Vendor         string `json:"VENDOR"`
	Hardware_model string `json:"HARDWARE_MODEL"`
	Serial_number  string `json:"SERIAL_NUMBER"`
	Mac_address    string `json:"MAC_ADDRESS"`
	Version        string `json:"VERSION"`
}

var SysAtsh_Template string = `Value VERSION (\S+)
Value BOOT_VERSION (.+)
Value VENDOR (.+)
Value HARDWARE_MODEL (\S+)
Value SERIAL_NUMBER (\S+)
Value MAC_ADDRESS ([0-9A-F]{12})

Start
  ^\d{4}-\d{2}-\d{2}.* -> Next
  ^Firmware\sVersion\s*:\s${VERSION}
  ^Bootbase\sVersion\s*:\s${BOOT_VERSION}
  ^Vendor\sName\s*:\s${VENDOR}
  ^Product\sModel\s*:\s${HARDWARE_MODEL}
  ^Serial\sNumber\s*:\s${SERIAL_NUMBER}
  ^First\sMAC\sAddress\s*:\s${MAC_ADDRESS}
  ^Last\sMAC\sAddress\s*:.+
  ^MAC\sAddress\sQuantity\s*:.+
  ^Default\sCountry\sCode\s*:.+
  ^Boot\sModule\sDebug Flag\s*:.+
  ^Kernel\sChecksum\s*:.+
  ^RootFS\sChecksum\s*:.+
  ^Romfile\sChecksum\s*:.+
  ^Main\sFeature\sBits\s*:.+
  ^Other\sFeature\sBits\s*:.*
  ^[0-9a-f]{1,8}:\s[0-9a-f]{1,8}\s[0-9a-f]{1,8}\s[0-9a-f]{1,8}\s[0-9a-f]{1,8}
  ^[0-9a-f]{1,8}:\s[0-9a-f]{1,8}\s[0-9a-f]{1,8}\s[0-9a-f]{1,8}(\s[0-9a-f]{1,8})?
  ^. -> Error
`
