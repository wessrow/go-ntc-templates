package models

type AristaEosShowHostname struct {
	Hostname	string	`json:"HOSTNAME"`
	Fqdn	string	`json:"FQDN"`
}

var AristaEosShowHostname_Template = `Value HOSTNAME (\S+?)
Value FQDN (\S+?)

Start
  ^Hostname:\s+${HOSTNAME}$$
  ^FQDN:\s+${FQDN}$$ -> Record

`