package mikrotik_routeros

type IpDhcpServerLeasePrintTerseWithoutPaging struct {
	Address_lists      string `json:"ADDRESS_LISTS"`
	Always_broadcast   string `json:"ALWAYS_BROADCAST"`
	Active_mac_address string `json:"ACTIVE_MAC_ADDRESS"`
	Host_name          string `json:"HOST_NAME"`
	Status             string `json:"STATUS"`
	Last_seen          string `json:"LAST_SEEN"`
	Index              string `json:"INDEX"`
	Flag               string `json:"FLAG"`
	Comment            string `json:"COMMENT"`
	Dhcp_option        string `json:"DHCP_OPTION"`
	Address            string `json:"ADDRESS"`
	Mac_address        string `json:"MAC_ADDRESS"`
	Server             string `json:"SERVER"`
	Active_address     string `json:"ACTIVE_ADDRESS"`
	Src_mac_address    string `json:"SRC_MAC_ADDRESS"`
	Client_id          string `json:"CLIENT_ID"`
	Expires_after      string `json:"EXPIRES_AFTER"`
	Active_client_id   string `json:"ACTIVE_CLIENT_ID"`
	Active_server      string `json:"ACTIVE_SERVER"`
}

var IpDhcpServerLeasePrintTerseWithoutPaging_Template string = `Value INDEX (\d+)
Value FLAG ([XRDB]+)
Value COMMENT (.*)
Value ADDRESS (\S+)
Value MAC_ADDRESS ((?:[0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2})
Value CLIENT_ID (\S+)
Value ADDRESS_LISTS (\S+)
Value SERVER (\S+)
Value ALWAYS_BROADCAST (yes|no)
Value DHCP_OPTION (\S+)
Value STATUS (\S+)
Value EXPIRES_AFTER (\S+)
Value LAST_SEEN (\S+)
Value ACTIVE_ADDRESS (\S+)
Value ACTIVE_MAC_ADDRESS ((?:[0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2})
Value ACTIVE_CLIENT_ID (\S+)
Value ACTIVE_SERVER (\S+)
Value HOST_NAME (\S+)
Value SRC_MAC_ADDRESS ((?:[0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2})

Start
  ^\s*${INDEX}(?:\s+${FLAG})?(?:\s+comment=${COMMENT})?\s+address=${ADDRESS}(?:\s+mac-address=${MAC_ADDRESS})?(?:\s+client-id=(?:\"?(\b${CLIENT_ID}\b)?\"?))?(?:\s+address-lists=(?:\"?(\b${ADDRESS_LISTS}\b)?\"?))?(?:\s+server=${SERVER})?(?:\s+dhcp-option=(?:\"?(\b${DHCP_OPTION}\b)?\"?))?(?:\s+status=${STATUS})?(?:\s+expires-after=${EXPIRES_AFTER})?(?:\s+last-seen=${LAST_SEEN})?(?:\s+active-address=${ACTIVE_ADDRESS})?(?:\s+active-mac-address=${ACTIVE_MAC_ADDRESS})?(?:\s+active-client-id=(?:\"?(\b${ACTIVE_CLIENT_ID}\b)?\"?))?(?:\s+active-server=${ACTIVE_SERVER})?(?:\s+host-name=(?:\"?(\b${HOST_NAME}\b)?\"?))?(?:\s+src-mac-address=${SRC_MAC_ADDRESS})?\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
