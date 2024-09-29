package models

type CiscoXrAdminShowVm struct {
	Location	string	`json:"LOCATION"`
	Id	string	`json:"ID"`
	Status	string	`json:"STATUS"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Hb_sent	string	`json:"HB_SENT"`
	Hb_recv	string	`json:"HB_RECV"`
}

var CiscoXrAdminShowVm_Template = `Value Filldown LOCATION (\S+)
Value ID (\S+)
Value Required STATUS (\S+)
Value Required IP_ADDRESS (\S+)
Value HB_SENT (\S+)
Value HB_RECV (\S+)

Start
  ^.+UTC
  ^Location:\s${LOCATION}
  ^Id\s+Status\s+IP\sAddress\s+HB\sSent/Recv
  ^-+ -> Loc
  ^\s+$$
  ^$$
  ^.* -> Error "LINE NOT FOUND"

Loc
  ^${ID}\s+${STATUS}\s+${IP_ADDRESS}\s+${HB_SENT}\/${HB_RECV} -> Record
  ^\s+$$
  ^$$ -> Start
  ^.* -> Error "LINE NOT FOUND"

`