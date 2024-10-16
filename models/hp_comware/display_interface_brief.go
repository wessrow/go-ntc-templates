package hp_comware

type DisplayInterfaceBrief struct {
	Type        string `json:"TYPE"`
	Pvid        string `json:"PVID"`
	Interface   string `json:"INTERFACE"`
	Link        string `json:"LINK"`
	Speed       string `json:"SPEED"`
	Duplex      string `json:"DUPLEX"`
	Protocol    string `json:"PROTOCOL"`
	Main_ip     string `json:"MAIN_IP"`
	Description string `json:"DESCRIPTION"`
}

var DisplayInterfaceBrief_Template string = `Value INTERFACE (\S+)
Value LINK (\S+)
Value PROTOCOL ((?:UP|DOWN)(?:\(s\))?)
Value MAIN_IP (\S+)
Value SPEED (\S+)
Value DUPLEX (A|H|F)
Value TYPE (A|T|H)
Value PVID (\d+)
Value DESCRIPTION (.*?)

Start
  ^\s*The\s+brief\s+information\s+of\s+interface\(s\)\s+under\s+route\s+mode:\s*$$
  ^\s*Link:\s+ADM\s+-\s+administratively\s+down;\s+Stby\s+-\s+standby\s*$$
  ^\s*Protocol:\s+\(s\)\s+-\s+spoofing\s*$$
  ^\s*Interface\s+Link\s+Protocol\s+Main\s+IP\s+Description\s*$$ -> RouteModeIfaces
  ^\s*The\s+brief\s+information\s+of\s+interface\(s\)\s+under\s+bridge\s+mode:\s*$$
  ^\s*Link:\s+ADM\s+-\s+administratively\s+down;\s+Stby\s+-\s+standby\s*$$
  ^\s*Speed\s+or\s+Duplex:\s+\(a\)/A\s+-\s+auto;\s+H\s+-\s+half;\s+F\s+-\s+full\s*$$
  ^\s*Type:\s+A\s+-\s+access;\s+T\s+-\s+trunk;\s+H\s+-\s+hybrid\s*$$
  ^\s*Interface\s+Link\s+Speed\s+Duplex\s+Type\s+PVID\s+Description\s*$$ -> BridgeModeIfaces
  ^\s*$$
  ^. -> Error

RouteModeIfaces
  ^\s*${INTERFACE}\s+${LINK}\s+${PROTOCOL}\s+${MAIN_IP}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$ -> Start
  ^. -> Error

BridgeModeIfaces
  ^${INTERFACE}\s+${LINK}\s+${SPEED}\s+${DUPLEX}\s+${TYPE}\s+${PVID}(?:\s+${DESCRIPTION})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
