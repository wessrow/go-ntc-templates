package zyxel_os

type ZycliEthctlMediaType struct {
	Link_active       string `json:"LINK_ACTIVE"`
	Speed             string `json:"SPEED"`
	Enabled           string `json:"ENABLED"`
	Negotiated_speed  string `json:"NEGOTIATED_SPEED"`
	Negotiated_duplex string `json:"NEGOTIATED_DUPLEX"`
}

var ZycliEthctlMediaType_Template string = `Value SPEED (\w+)
Value ENABLED (\w+)
Value NEGOTIATED_SPEED (\w+)
Value NEGOTIATED_DUPLEX (\w+)
Value LINK_ACTIVE (up|down)

Start
  ^${SPEED}-negotiation\s${ENABLED}\.\s*$$
  ^The\sautonegotiated\smedia\stype\sis\s${NEGOTIATED_SPEED}\s${NEGOTIATED_DUPLEX}\sDuplex\s*$$
  ^Link\sis\s${LINK_ACTIVE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
