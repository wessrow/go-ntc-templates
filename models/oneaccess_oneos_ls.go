package models

type OneaccessOneosLs struct {
	File	string	`json:"FILE"`
	Bytes	string	`json:"BYTES"`
}

var OneaccessOneosLs_Template = `Value Required FILE (\S+)
Value BYTES (\d+)
	   
Start
  ^Listing -> ONEOS5
  ^\S+\s+\d+\s+${BYTES}\s+.*\s+${FILE}$$ -> Record ONEOS6
  ^$$
  ^. -> Error

ONEOS5
  ^\.
  ^${FILE}\s+${BYTES} -> Record
  ^\s*$$
  ^. -> Error

ONEOS6
  ^\S+\s+\d+\s+${BYTES}\s+.*\s+${FILE}$$ -> Record
  ^\s*$$
  ^. -> Error

`