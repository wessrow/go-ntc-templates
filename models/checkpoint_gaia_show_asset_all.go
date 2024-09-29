package models

type CheckpointGaiaShowAssetAll struct {
	Platform	string	`json:"PLATFORM"`
	Model	string	`json:"MODEL"`
	Serial	string	`json:"SERIAL"`
	Slot	string	`json:"SLOT"`
	Descr	string	`json:"DESCR"`
	Lom_status	string	`json:"LOM_STATUS"`
	Lom_rev	string	`json:"LOM_REV"`
}

var CheckpointGaiaShowAssetAll_Template = `Value PLATFORM (\S+)
Value MODEL (.*)
Value SERIAL (.*)
Value SLOT (\d)
Value DESCR (.*)
Value LOM_STATUS (\S+)
Value LOM_REV (\S+)
 
 
Start
  ^Platform:\s${PLATFORM}
  ^Model:\s${MODEL}
  ^Serial Number:\s${SERIAL} -> Record
  ^LOM\sStatus:\s${LOM_STATUS}
  ^LOM\sFirmware\sRevision:\s${LOM_REV} -> Record
  ^Line\scard\s${SLOT}\smodel:\s${MODEL}
  ^Line\scard\s${SLOT}\stype:\s${DESCR} -> Record
`