package huawei_smartax 

type DisplayOntAutofind struct {
	Number	string	`json:"NUMBER"`
	Fsp	string	`json:"FSP"`
	Serial_number	string	`json:"SERIAL_NUMBER"`
	Password	string	`json:"PASSWORD"`
	Loid	string	`json:"LOID"`
	Checkcode	string	`json:"CHECKCODE"`
	Vendor_id	string	`json:"VENDOR_ID"`
	Ont_version	string	`json:"ONT_VERSION"`
	Ont_software_version	string	`json:"ONT_SOFTWARE_VERSION"`
	Ont_equipment_id	string	`json:"ONT_EQUIPMENT_ID"`
	Ont_customized_info	string	`json:"ONT_CUSTOMIZED_INFO"`
	Ont_mac	string	`json:"ONT_MAC"`
	Ont_equipment_serial_number	string	`json:"ONT_EQUIPMENT_SERIAL_NUMBER"`
	Ont_autofind_time	string	`json:"ONT_AUTOFIND_TIME"`
	Multi_channel	string	`json:"MULTI_CHANNEL"`
}

var DisplayOntAutofind_Template = `Value NUMBER (\d+)
Value FSP (\d+\/\d+\/\d+)
Value SERIAL_NUMBER (\w+\s*\(.*\))
Value PASSWORD (\S+)
Value LOID (\S*)
Value CHECKCODE (\S*)
Value VENDOR_ID (\w+)
Value ONT_VERSION (\S+)
Value ONT_SOFTWARE_VERSION (\S+)
Value ONT_EQUIPMENT_ID (\S+)
Value ONT_CUSTOMIZED_INFO (\S+)
Value ONT_MAC (\S+)
Value ONT_EQUIPMENT_SERIAL_NUMBER (\S+)
Value ONT_AUTOFIND_TIME (\S+\s*\S+)
Value MULTI_CHANNEL (\w*)

Start
  ^\s*Number\s+:\s+${NUMBER}
  ^\s*F\/S\/P\s+:\s+${FSP}
  ^\s*Ont\s+SN\s+:\s+${SERIAL_NUMBER}
  ^\s*Password\s+:\s+${PASSWORD}
  ^\s*Loid\s+:\s+${LOID}
  ^\s*Checkcode\s+:\s+${CHECKCODE}
  ^\s*VendorID\s+:\s+${VENDOR_ID}
  ^\s*Ont\s+Version\s+:\s+${ONT_VERSION}
  ^\s*Ont\s+SoftwareVersion\s+:\s+${ONT_SOFTWARE_VERSION}
  ^\s*Ont\s+EquipmentID\s+:\s+${ONT_EQUIPMENT_ID}
  ^\s*Ont\s+Customized\s+Info\s+:\s+${ONT_CUSTOMIZED_INFO}
  ^\s*Ont\s+MAC\s+:\s+${ONT_MAC}
  ^\s*Ont\s+Equipment\s+SN\s+:\s+${ONT_EQUIPMENT_SERIAL_NUMBER}
  ^\s*Ont\s+autofind\s+time\s+:\s+${ONT_AUTOFIND_TIME}
  ^\s*Multi\s+channel\s+:\s+${MULTI_CHANNEL} -> Record
  ^\s*
  ^. -> Error`