package models

type HuaweiSmartaxDisplayBoard struct {
	Slot_id	string	`json:"SLOT_ID"`
	Boardname	string	`json:"BOARDNAME"`
	Status	string	`json:"STATUS"`
	Subtype_0	string	`json:"SUBTYPE_0"`
	Subtype_1	string	`json:"SUBTYPE_1"`
	Online_offline	string	`json:"ONLINE_OFFLINE"`
}

var HuaweiSmartaxDisplayBoard_Template = `Value Key SLOT_ID (\d+)
Value BOARDNAME (\w*)
Value STATUS (\S*)
Value SUBTYPE_0 ([A-Z]{0,8})
Value SUBTYPE_1 ([A-Z]{0,8})
Value ONLINE_OFFLINE ((Online|Offline)?)


Start
  ^\s+SlotID\s+BoardName\s+Status\s+SubType0\s+SubType1\s+Online\/Offline
  ^\s+${SLOT_ID}\s+${BOARDNAME}\s+${STATUS}\s+${SUBTYPE_0}\s+${SUBTYPE_1}\s+${ONLINE_OFFLINE} -> Record
  ^\s+-
  ^. -> Error
`