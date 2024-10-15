package arista_eos

type ShowHostname struct {
	Hostname string `json:"HOSTNAME"`
	Fqdn     string `json:"FQDN"`
}

var ShowHostname_Template string = `Value HOSTNAME (\S+?)
Value FQDN (\S+?)

Start
  ^Hostname:\s+${HOSTNAME}$$
  ^FQDN:\s+${FQDN}$$ -> Record
`
