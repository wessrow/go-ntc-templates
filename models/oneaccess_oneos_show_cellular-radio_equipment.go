package models

type OneaccessOneosShowCellularRadioEquipment struct {
	Manufacturer	string	`json:"MANUFACTURER"`
	Equipment	string	`json:"EQUIPMENT"`
	Boot_revision	string	`json:"BOOT_REVISION"`
	Imei	string	`json:"IMEI"`
	Sim1_status	string	`json:"SIM1_STATUS"`
	Sim2_status	string	`json:"SIM2_STATUS"`
	Imsi	string	`json:"IMSI"`
	Card_id	string	`json:"CARD_ID"`
	Pin_status	string	`json:"PIN_STATUS"`
}

var OneaccessOneosShowCellularRadioEquipment_Template = `Value Required MANUFACTURER (.*)
Value EQUIPMENT (.*)
Value BOOT_REVISION (.*)
Value IMEI (\d+)
Value SIM1_STATUS (present|not present)
Value SIM2_STATUS (present|not present)
Value IMSI (\d+)
Value CARD_ID (\d+)
Value PIN_STATUS (.*[^\s])

Start
  ^\s*Cellular\sRadio\sModem\sInformation -> Record
  ^\s*Manufacturer\sidentification\s*:\s*${MANUFACTURER}
  ^\s*Equipment\sinformation\s*:\s*${EQUIPMENT}
  ^\s*Revision\sidentification\s*:\s*${BOOT_REVISION}
  ^\s*Equipment\sinformation\s\(IMEI\)\s*:\s*${IMEI}
  ^\s*SIM\scard\sstatus\s*:\s*SIM( card)?( is)?\s${SIM1_STATUS} -> Oneos5
  ^\s*SIM\scard\sslot\s1\s*:\s*SIM( card)?( is)?\s${SIM1_STATUS} -> Oneos6
  ^\s*Boot\s+revision\s+identification\s+:
  ^\s*SIM\s+Card\s+Information
  ^\s*$$
  ^. -> Error

Oneos5
  ^.*IMSI\s*:\s+${IMSI}
  ^\s*Integrated\sCircuit\sCard\sID\s*:\s+${CARD_ID}
  ^\s*PIN\scode\sstatus\s*:\s+${PIN_STATUS}
  ^\s*PIN\s+Information
  ^\s*Active\s+SIM\s+slot\s+:
  ^\s*SIM\s+card\s+status\s+:
  ^\s*$$
  ^. -> Error

Oneos6
  ^\s*SIM\scard\sslot\s2\s*:\s*SIM( card)?( is)?\s${SIM2_STATUS}
  ^.*IMSI\s*:\s+${IMSI}
  ^\s*Integrated\sCircuit\sCard\sID\s*:\s+${CARD_ID}
  ^\s*PIN\s+Information
  ^\s*Active\s+SIM\s+slot\s+:
  ^\s*SIM\s+card\s+status\s+:
  ^\s*$$
  ^. -> Error

`