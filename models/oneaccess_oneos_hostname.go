package models

type OneaccessOneosHostname struct {
	Hostname	string	`json:"HOSTNAME"`
}

var OneaccessOneosHostname_Template = `Value HOSTNAME (.*)

Start
  ^${HOSTNAME}$$ -> Record
  ^. -> Error

`