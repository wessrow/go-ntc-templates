package cisco_s300

type ShowVersion struct {
	Sw_version   string `json:"SW_VERSION"`
	Boot_version string `json:"BOOT_VERSION"`
	Hw_version   string `json:"HW_VERSION"`
}

var ShowVersion_Template string = `Value SW_VERSION ([1-9]\.[0-9]\.[0-9].[0-9]*)
Value BOOT_VERSION ([1-9]\.[0-9]\.[0-9].[0-9]*)
Value HW_VERSION (V0[1-9])

Start
  ^SW\s+version\s+${SW_VERSION}\s.* -> FWVER1
  ^Active-image:\s.* -> FWVER2

FWVER1
  ^Boot\s+version\s+${BOOT_VERSION}\s.*
  ^HW\s+version\s+${HW_VERSION} -> Record
  ^\s*$$
  ^. -> Error

FWVER2
  ^\s+Version:\s${SW_VERSION} -> Record
  ^\s+MD5\s.*
  ^\s+Date:\s.*
  ^\s+Time:\s.* -> End
  ^\s*$$
  ^. -> Error
`
