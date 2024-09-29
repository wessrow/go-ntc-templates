package models

type HuaweiSmartaxDisplayVersion struct {
	Olt_version	string	`json:"OLT_VERSION"`
	Patch	string	`json:"PATCH"`
	Product	string	`json:"PRODUCT"`
	Mainboard_type	string	`json:"MAINBOARD_TYPE"`
	Current_program_area	string	`json:"CURRENT_PROGRAM_AREA"`
	Current_data_area	string	`json:"CURRENT_DATA_AREA"`
	Program_area_a_version	string	`json:"PROGRAM_AREA_A_VERSION"`
	Program_area_b_version	string	`json:"PROGRAM_AREA_B_VERSION"`
	Data_area_a_version	string	`json:"DATA_AREA_A_VERSION"`
	Data_area_b_version	string	`json:"DATA_AREA_B_VERSION"`
}

var HuaweiSmartaxDisplayVersion_Template = `Value Filldown OLT_VERSION (\S+)
Value Filldown PATCH (\S+)
Value Filldown PRODUCT (\S+)
Value MAINBOARD_TYPE (\S+)
Value Required CURRENT_PROGRAM_AREA (\S+\s*\S+)
Value Required CURRENT_DATA_AREA (\S+\s*\S+)
Value PROGRAM_AREA_A_VERSION (\S+)
Value PROGRAM_AREA_B_VERSION (\S+)
Value DATA_AREA_A_VERSION (\S+)
Value DATA_AREA_B_VERSION (\S+)

Start
  ^\s*VERSION\s+:\s+${OLT_VERSION}
  ^\s*PATCH\s+:\s+${PATCH}
  ^\s*PRODUCT\s+:\s+${PRODUCT} -> BoardInfo

BoardInfo
  ^\s*${MAINBOARD_TYPE}\s+Mainboard\s+Running\s+Area\s+Information:
  ^\s*Current\s+Program\s+Area\s+:\s+${CURRENT_PROGRAM_AREA}
  ^\s*Current\s+Data\s+Area\s+:\s+${CURRENT_DATA_AREA}
  ^\s*Program\s+Area\s+A\s+Version\s+:\s+${PROGRAM_AREA_A_VERSION}
  ^\s*Program\s+Area\s+B\s+Version\s+:\s+${PROGRAM_AREA_B_VERSION}
  ^\s*Data\s+Area\s+A\s+Version\s+:\s+${DATA_AREA_A_VERSION}
  ^\s*Data\s+Area\s+B\s+Version\s+:\s+${DATA_AREA_B_VERSION} -> Record
  ^\s+-
  ^\s*Uptime\s+is\s+\d+\s+day\(s\),\s+\d+\s+hour\(s\),\s+\d+\s+minute\(s\),\s+\d+\s+second\(s\)
  ^\s*$$
  ^. -> Error
`