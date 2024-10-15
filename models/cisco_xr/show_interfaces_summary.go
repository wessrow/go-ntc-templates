package cisco_xr

type ShowInterfacesSummary struct {
	Admin_down string `json:"ADMIN_DOWN"`
	Intf_type  string `json:"INTF_TYPE"`
	Total      string `json:"TOTAL"`
	Up         string `json:"UP"`
	Down       string `json:"DOWN"`
}

var ShowInterfacesSummary_Template string = `Value INTF_TYPE (ALL TYPES|IFT_ETHERBUNDLE|IFT_HUNDREDGE|IFT_LOOPBACK|IFT_ETHERNET|IFT_NULL|IFT_TENGETHERNET)
Value TOTAL (\d+)
Value UP (\S+)
Value DOWN (\S+)
Value ADMIN_DOWN (\S+)

Start
  ^Interface\s+Type\s+Total\s+UP\s+\Down\s+Admin\s+Down\s*$$
  ^----
  ^${INTF_TYPE}+\s+${TOTAL}\s+${UP}\s+${DOWN}\s+${ADMIN_DOWN} -> Record
  ^\s*$$
  ^. -> Error`
