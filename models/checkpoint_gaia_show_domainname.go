package models

type CheckpointGaiaShowDomainname struct {
	Domainname	string	`json:"DOMAINNAME"`
}

var CheckpointGaiaShowDomainname_Template = `Value DOMAINNAME (\S+)

Start
  ^Domain name\s+${DOMAINNAME} -> Record

`