package cisco_ios

type ShowAuthenticationSessions struct {
	Method      string `json:"METHOD"`
	Domain      string `json:"DOMAIN"`
	Status      string `json:"STATUS"`
	Session     string `json:"SESSION"`
	Interface   string `json:"INTERFACE"`
	Mac_address string `json:"MAC_ADDRESS"`
}

var ShowAuthenticationSessions_Template string = `Value INTERFACE (.+?)
Value MAC_ADDRESS (.+?)
Value METHOD (.+?)
Value DOMAIN (.+?)
Value STATUS (.+?)
Value SESSION (\w+?)

Start
  ^Interface\s+(MAC Address|Identifier)\s+Method\s+Domain -> Catch
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is

Catch
  ^-+$$
  ^${INTERFACE}\s+${MAC_ADDRESS}\s+${METHOD}\s+${DOMAIN}\s+${STATUS}\s+${SESSION}$$ -> Record
  ^.* -> Start
`
