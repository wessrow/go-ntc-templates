package mikrotik_routeros

type InterfacePrintTerseWithoutPaging struct {
	Slave               string `json:"SLAVE"`
	Name                string `json:"NAME"`
	Default_name        string `json:"DEFAULT_NAME"`
	Mac_address         string `json:"MAC_ADDRESS"`
	Dynamic             string `json:"DYNAMIC"`
	Mtu                 string `json:"MTU"`
	Last_link_down_time string `json:"LAST_LINK_DOWN_TIME"`
	Link_downs          string `json:"LINK_DOWNS"`
	Type                string `json:"TYPE"`
	Actual_mtu          string `json:"ACTUAL_MTU"`
	L2mtu               string `json:"L2MTU"`
	Last_link_up_time   string `json:"LAST_LINK_UP_TIME"`
	Status              string `json:"STATUS"`
	Comment             string `json:"COMMENT"`
	Max_l2mtu           string `json:"MAX_L2MTU"`
	Fast_path           string `json:"FAST_PATH"`
	Id                  string `json:"ID"`
}

var InterfacePrintTerseWithoutPaging_Template string = `Value ID (\d+)
Value DYNAMIC (D)
Value STATUS (X|R)
Value SLAVE (S)
Value COMMENT (.*)
Value NAME (.*?)
Value DEFAULT_NAME (\S+)
Value TYPE (\S+)
Value MTU (\d+|auto)
Value ACTUAL_MTU (\d+)
Value L2MTU (\d+|auto)
Value MAX_L2MTU (\d+)
Value MAC_ADDRESS ([a-zA-Z0-9]{2}(:[a-zA-Z0-9]{2}){5})
Value FAST_PATH (\S+)
Value LAST_LINK_DOWN_TIME ([a-z]{3}\/\d+\/\d+\s\d+\:\d+\:\d+)
Value LAST_LINK_UP_TIME ([a-z]{3}\/\d+\/\d+\s\d+\:\d+\:\d+)
Value LINK_DOWNS (\d+)

Start
  ^\s*${ID}\s+(?:${DYNAMIC})?(?:${STATUS}|\s)?(?:${SLAVE})?\s+(comment=${COMMENT}\s+)?name=${NAME}(\s+default-name=${DEFAULT_NAME})?\s+type=${TYPE}(\s+mtu=${MTU})?(\s+actual-mtu=${ACTUAL_MTU})?(\s+l2mtu=${L2MTU})?(\s+max-l2mtu=${MAX_L2MTU})?(\s+mac-address=${MAC_ADDRESS})?(\s+fast-path=${FAST_PATH})?(\s+last-link-down-time=${LAST_LINK_DOWN_TIME})?(\s+last-link-up-time=${LAST_LINK_UP_TIME})?(\s+link-downs=${LINK_DOWNS})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
