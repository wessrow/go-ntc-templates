package mikrotik_routeros

type InterfacePrintDetail struct {
	Status              string   `json:"STATUS"`
	Name                string   `json:"NAME"`
	Description         []string `json:"DESCRIPTION"`
	Max_l2mtu           string   `json:"MAX_L2MTU"`
	Mac_address         string   `json:"MAC_ADDRESS"`
	Slave               string   `json:"SLAVE"`
	Mtu                 string   `json:"MTU"`
	Actual_mtu          string   `json:"ACTUAL_MTU"`
	L2mtu               string   `json:"L2MTU"`
	Last_link_down_time string   `json:"LAST_LINK_DOWN_TIME"`
	Last_link_up_time   string   `json:"LAST_LINK_UP_TIME"`
	Link_downs          string   `json:"LINK_DOWNS"`
	Id                  string   `json:"ID"`
	Dynamic             string   `json:"DYNAMIC"`
	Default_name        string   `json:"DEFAULT_NAME"`
	Type                string   `json:"TYPE"`
}

var InterfacePrintDetail_Template string = `Value Required ID (\d+)
Value DYNAMIC (D)
Value STATUS (X|R)
Value SLAVE (S)
Value NAME (\S+)
Value List DESCRIPTION ((?!\s*$).+[^\s])
Value DEFAULT_NAME (\S+)
Value TYPE (\S+)
Value MTU (\d+|auto)
Value ACTUAL_MTU (\d+)
Value L2MTU (\d+|auto)
Value MAX_L2MTU (\d+)
Value MAC_ADDRESS ([a-zA-Z0-9]{2}(?::[a-zA-Z0-9]{2}){5})
Value LAST_LINK_DOWN_TIME ([a-z]{3}/\d+/\d+\s\d+\:\d+\:\d+)
Value LAST_LINK_UP_TIME ([a-z]{3}/\d+/\d+\s\d+\:\d+\:\d+)
Value LINK_DOWNS (\d+)

Start
  ^\s*Flags:\s+D\s+-\s+dynamic,\s+X\s+-\s+disabled,\s+R\s+-\s+running,\s+S\s+-\s+slave\s*$$ -> Interface
  ^\s*$$
  ^. -> Error

Interface
  ^\s*\d+(?:\s+D)?(?:\s*X|R)?(?:\s*S)?\s+ -> Continue.Record
  ^\s*${ID}(?:\s+${DYNAMIC})?(?:\s*${STATUS})?(?:\s*${SLAVE})?\s+name="${NAME}"(?:\s+default-name="${DEFAULT_NAME}")?\s+type="${TYPE}"(\s+mtu=${MTU})?(?:\s+actual-mtu=${ACTUAL_MTU})?(?:\s+l2mtu=${L2MTU})?(?:\s+max-l2mtu=${MAX_L2MTU})?(?:\s+mac-address=${MAC_ADDRESS})?(?:\s+last-link-down-time=${LAST_LINK_DOWN_TIME})?(?:\s+last-link-up-time=${LAST_LINK_UP_TIME})?(?:\s+link-downs=${LINK_DOWNS})?\s*$$
  ^\s*name="${NAME}"(?:\s+default-name="${DEFAULT_NAME}")?\s+type="${TYPE}"(\s+mtu=${MTU})?(?:\s+actual-mtu=${ACTUAL_MTU})?(?:\s+l2mtu=${L2MTU})?(?:\s+max-l2mtu=${MAX_L2MTU})?(?:\s+mac-address=${MAC_ADDRESS})?(?:\s+last-link-down-time=${LAST_LINK_DOWN_TIME})?(?:\s+last-link-up-time=${LAST_LINK_UP_TIME})?(?:\s+link-downs=${LINK_DOWNS})?\s*$$
  ^\s*${ID}(?:\s+${DYNAMIC})?(?:\s*${STATUS})?(?:\s*${SLAVE})?\s+;{3}\s+${DESCRIPTION}\s*$$
  ^${DESCRIPTION}\s*$$
  ^\s*$$
  ^. -> Error
`
