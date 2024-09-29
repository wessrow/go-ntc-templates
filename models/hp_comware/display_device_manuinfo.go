package hp_comware 

type DisplayDeviceManuinfo struct {
	Chassis_id	string	`json:"CHASSIS_ID"`
	Slot_type	string	`json:"SLOT_TYPE"`
	Slot_id	string	`json:"SLOT_ID"`
	Device_name	string	`json:"DEVICE_NAME"`
	Device_serial_number	string	`json:"DEVICE_SERIAL_NUMBER"`
	Manufacturing_date	string	`json:"MANUFACTURING_DATE"`
	Vendor_name	string	`json:"VENDOR_NAME"`
	Mac_address	string	`json:"MAC_ADDRESS"`
}

var DisplayDeviceManuinfo_Template = `Value Filldown CHASSIS_ID (\d+)
Value SLOT_TYPE ([sS]lot|Subslot|Fan|Power|Chassis)
Value SLOT_ID (\d+|self)
Value DEVICE_NAME (.+)
Value DEVICE_SERIAL_NUMBER (\S+)
Value MANUFACTURING_DATE (\S+)
Value VENDOR_NAME (\S+)
Value MAC_ADDRESS (\S+)


Start
  ^\s*Chassis\s+${CHASSIS_ID}
  ^\s*${SLOT_TYPE}\s+${SLOT_ID}
  ^\s*DEVICE_ID
  ^\s*The\s+operation\s+is\s+not\s+supported\s+ -> Record
  ^\s*The\s+card\s+does\s+not\s+support\s+ -> Record
  ^\s*DEVICE_NAME\s*:\s*${DEVICE_NAME}
  ^\s*DEVICE_SERIAL_NUMBER\s*:\s*${DEVICE_SERIAL_NUMBER}
  ^\s*MANU\s+SERIAL\s+NUMBER\s*:\s*${DEVICE_SERIAL_NUMBER}
  ^\s*MAC_ADDRESS\s*:\s*${MAC_ADDRESS}
  ^\s*MANUFACTURING_DATE\s*:\s*${MANUFACTURING_DATE}
  ^\s*VENDOR_NAME\s*:\s*${VENDOR_NAME} -> Record
  ^\s*$$
  ^. -> Error

# necessary to avoid an extra mostly empty match at the end
EOF
`