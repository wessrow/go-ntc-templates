package oneaccess_oneos 

type ShowSntp struct {
	Server	string	`json:"SERVER"`
	Source	string	`json:"SOURCE"`
	Stratum	string	`json:"STRATUM"`
	Version	string	`json:"VERSION"`
	Last	string	`json:"LAST"`
	Receive	string	`json:"RECEIVE"`
	Learn_from	string	`json:"LEARN_FROM"`
	Interface	string	`json:"INTERFACE"`
	Vrf	string	`json:"VRF"`
	Failed_auth_pkts	string	`json:"FAILED_AUTH_PKTS"`
	Broadcast	string	`json:"BROADCAST"`
	Authentication	string	`json:"AUTHENTICATION"`
}

var ShowSntp_Template = `Value Required SERVER (\S+)
Value SOURCE (\S+(\s\S+)?)
Value STRATUM (\d+)
Value VERSION (\d+)
Value LAST (\d+:\d+:\d+)
Value RECEIVE (\S+)
Value LEARN_FROM (\S+)
Value INTERFACE (\S+)
Value VRF (\S+)
Value FAILED_AUTH_PKTS (\d+)
Value Fillup BROADCAST (enabled|not enabled)
Value Fillup AUTHENTICATION (enabled|not enabled)

Start
  ^server\s+source\s+stratum\s+version\s+last\s+receive\s+learn -> ONEOS6
  ^server\s+source\s+stratum\s+version\s+last\s+receive -> ONEOS5
  ^\s*$$
  ^. -> Error

ONEOS5
  ^${SERVER}\s+${SOURCE}\s+(${STRATUM}|)\s+(${VERSION}|)\s+(${LAST}|)\s+(${RECEIVE}|) -> Record
  ^\s*broadcast\sclient\smode\sis\s${BROADCAST}
  ^\s*SNTP\sAuthentication\sis\s${AUTHENTICATION}
  ^\s*$$
  ^. -> Error

ONEOS6
  ^${SERVER}\s+${SOURCE}\s+(${STRATUM}|-)\s+(${VERSION}|-)\s+(${LAST}|)\s+(${RECEIVE}|-)\s+(${LEARN_FROM}|-)\s+(${INTERFACE}|-)\s+(${VRF}|-)\s+(${FAILED_AUTH_PKTS}) -> Record
  ^\s*broadcast\sclient\smode\sis\s${BROADCAST}
  ^\s*SNTP\sAuthentication\sis\s${AUTHENTICATION}
  ^\s*$$
  ^. -> Error
`