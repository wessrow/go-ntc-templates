package oneaccess_oneos 

type Hostname struct {
	Hostname	string	`json:"HOSTNAME"`
}

var Hostname_Template = `Value HOSTNAME (.*)

Start
  ^${HOSTNAME}$$ -> Record
  ^. -> Error
`