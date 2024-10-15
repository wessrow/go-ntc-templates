package cisco_nxos

type ShowHostname struct {
	Hostname string `json:"HOSTNAME"`
}

var ShowHostname_Template string = `Value HOSTNAME (.+?)

Start
  ^${HOSTNAME}\s*$$ -> Record`
