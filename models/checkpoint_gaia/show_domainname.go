package checkpoint_gaia

type ShowDomainname struct {
	Domainname string `json:"DOMAINNAME"`
}

var ShowDomainname_Template string = `Value DOMAINNAME (\S+)

Start
  ^Domain name\s+${DOMAINNAME} -> Record
`
