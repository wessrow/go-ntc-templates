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