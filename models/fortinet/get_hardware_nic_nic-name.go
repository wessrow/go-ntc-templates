package fortinet

type GetHardwareNicNicName struct {
	Netdev_oid       string   `json:"NETDEV_OID"`
	Admin            string   `json:"ADMIN"`
	Rx_packets       string   `json:"RX_PACKETS"`
	Host_rx_bytes    string   `json:"HOST_RX_BYTES"`
	Member_ports     []string `json:"MEMBER_PORTS"`
	Driver_name      string   `json:"DRIVER_NAME"`
	Permanent_hwaddr string   `json:"PERMANENT_HWADDR"`
	Host_tx_dropped  string   `json:"HOST_TX_DROPPED"`
	Current_hwaddr   string   `json:"CURRENT_HWADDR"`
	Tx_packets       string   `json:"TX_PACKETS"`
	Rx_bytes         string   `json:"RX_BYTES"`
	Duplex_setting   string   `json:"DUPLEX_SETTING"`
	Speed_setting    string   `json:"SPEED_SETTING"`
	Host_tx_bytes    string   `json:"HOST_TX_BYTES"`
	Lif_id           string   `json:"LIF_ID"`
	Lif_oid          string   `json:"LIF_OID"`
	Duplex           string   `json:"DUPLEX"`
	Link_status      string   `json:"LINK_STATUS"`
	Tx_bytes         string   `json:"TX_BYTES"`
	Host_tx_packets  string   `json:"HOST_TX_PACKETS"`
	Fragtxcreate     string   `json:"FRAGTXCREATE"`
	Description      string   `json:"DESCRIPTION"`
	Link_setting     string   `json:"LINK_SETTING"`
	Fragtxok         string   `json:"FRAGTXOK"`
	Fragtxdrop       string   `json:"FRAGTXDROP"`
	Netdev_status    string   `json:"NETDEV_STATUS"`
	Autonego_setting string   `json:"AUTONEGO_SETTING"`
	Speed            string   `json:"SPEED"`
	Host_rx_packets  string   `json:"HOST_RX_PACKETS"`
	Board            string   `json:"BOARD"`
}

var GetHardwareNicNicName_Template string = `Value DESCRIPTION (.+)
Value DRIVER_NAME (.+)
Value BOARD (\w+)
Value LIF_ID (\d+)
Value LIF_OID (\d+)
Value NETDEV_OID (\d+)
Value CURRENT_HWADDR ([a-zA-Z0-9]{2}(:[a-zA-Z0-9]{2}){5})
Value PERMANENT_HWADDR ([a-zA-Z0-9]{2}(:[a-zA-Z0-9]{2}){5})
Value ADMIN (\S+)
Value NETDEV_STATUS (\S+)
Value AUTONEGO_SETTING (\d+)
Value LINK_SETTING (\d+)
Value SPEED_SETTING (\d+)
Value DUPLEX_SETTING (\d+)
Value SPEED (\d+)
Value DUPLEX (\S+)
Value LINK_STATUS (\S+)
Value RX_PACKETS (\d+)
Value RX_BYTES (\d+)
Value TX_PACKETS (\d+)
Value TX_BYTES (\d+)
Value HOST_RX_PACKETS (\d+)
Value HOST_RX_BYTES (\d+)
Value HOST_TX_PACKETS (\d+)
Value HOST_TX_BYTES (\d+)
Value HOST_TX_DROPPED (\d+)
Value FRAGTXCREATE (\d+)
Value FRAGTXOK (\d+)
Value FRAGTXDROP (\d+)
Value List MEMBER_PORTS (\[\d+\]:\s+\S+)

Start
  ^Description\s*:${DESCRIPTION}\s*$$
  ^Driver\s+Name\s*:${DRIVER_NAME}\s*$$
  ^Board\s*:${BOARD}\s*$$
  ^lif\s+id\s*:${LIF_ID}\s*$$
  ^lif\s+oid\s*:${LIF_OID}\s*$$
  ^netdev\s+oid\s*:${NETDEV_OID}\s*$$
  ^Current_HWaddr\s*${CURRENT_HWADDR}\s*$$
  ^Permanent_HWaddr\s*${PERMANENT_HWADDR}\s*$$
  ^=+\s+Link\s+Status\s+=+ -> LinkStatus
  ^=+\s+Counters\s+=+ -> Counters
  ^\s*$$
  ^. -> Error

LinkStatus
  ^Admin\s*:${ADMIN}\s*$$
  ^netdev\s+status\s*:${NETDEV_STATUS}\s*$$
  ^autonego_setting\s*:${AUTONEGO_SETTING}\s*$$
  ^link_setting\s*:${LINK_SETTING}\s*$$
  ^speed_setting\s*:${SPEED_SETTING}\s*$$
  ^duplex_setting\s*:${DUPLEX_SETTING}\s*$$
  ^Speed\s*:${SPEED}\s*$$
  ^Duplex\s*:${DUPLEX}\s*$$
  ^link_status\s*:${LINK_STATUS}\s*$$ -> Start
  ^\s*$$
  ^. -> Error

Counters
  ^Rx\s+Pkts\s*:${RX_PACKETS}\s*$$
  ^Rx\s+Bytes\s*:${RX_BYTES}\s*$$
  ^Tx\s+Pkts\s*:${TX_PACKETS}\s*$$
  ^Tx\s+Bytes\s*:${TX_BYTES}\s*$$
  ^Host\s+Rx\s+Pkts\s*:${HOST_RX_PACKETS}\s*$$
  ^Host\s+Rx\s+Bytes\s*:${HOST_RX_BYTES}\s*$$
  ^Host\s+Tx\s+Pkts\s*:${HOST_TX_PACKETS}\s*$$
  ^Host\s+Tx\s+Bytes\s*:${HOST_TX_BYTES}\s*$$
  ^Host\s+Tx\s+dropped\s*:${HOST_TX_DROPPED}\s*$$
  ^FragTxCreate\s*:${FRAGTXCREATE}\s*$$
  ^FragTxOk\s*:${FRAGTXOK}\s*$$
  ^FragTxDrop\s*:${FRAGTXDROP}\s*$$
  # No 'Member Ports' parsing
  ^Member\s+Ports\s*\:\s*$$
  ^\s*${MEMBER_PORTS}\s*$$
  ^\s*$$
  ^. -> Error
`
