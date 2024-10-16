package fortinet

type FnsysctlIfconfig struct {
	Nic                   string `json:"NIC"`
	Link_local6_prefixlen string `json:"LINK_LOCAL6_PREFIXLEN"`
	Rx_frame              string `json:"RX_FRAME"`
	Tx_errors             string `json:"TX_ERRORS"`
	Collisions            string `json:"COLLISIONS"`
	Link_encap            string `json:"LINK_ENCAP"`
	Link_local6           string `json:"LINK_LOCAL6"`
	Rx_overruns           string `json:"RX_OVERRUNS"`
	Tx_packets            string `json:"TX_PACKETS"`
	Tx_carrier            string `json:"TX_CARRIER"`
	Rx_bytes              string `json:"RX_BYTES"`
	Bcast                 string `json:"BCAST"`
	Mask                  string `json:"MASK"`
	Multicast             string `json:"MULTICAST"`
	Metric                string `json:"METRIC"`
	Rx_packets            string `json:"RX_PACKETS"`
	Tx_overruns           string `json:"TX_OVERRUNS"`
	Hardware_address      string `json:"HARDWARE_ADDRESS"`
	Ip_address            string `json:"IP_ADDRESS"`
	Mtu                   string `json:"MTU"`
	Rx_errors             string `json:"RX_ERRORS"`
	Rx_dropped            string `json:"RX_DROPPED"`
	Tx_dropped            string `json:"TX_DROPPED"`
	Tx_queue_len          string `json:"TX_QUEUE_LEN"`
	Tx_bytes              string `json:"TX_BYTES"`
}

var FnsysctlIfconfig_Template string = `Value NIC (.*?)
Value LINK_ENCAP (.*?)
Value HARDWARE_ADDRESS ([a-zA-Z0-9]{2}(:[a-zA-Z0-9]{2}){5})
Value IP_ADDRESS (\S+)
Value BCAST (\S+)
Value MASK (\S+)
Value LINK_LOCAL6 (\S+)
Value LINK_LOCAL6_PREFIXLEN (\d+)
Value MULTICAST (.*(?<!\s))
Value MTU (\d+)
Value METRIC (\d+)
Value RX_PACKETS (\d+)
Value RX_ERRORS (\d+)
Value RX_DROPPED (\d+)
Value RX_OVERRUNS (\d+)
Value RX_FRAME (\d+)
Value TX_PACKETS (\d+)
Value TX_ERRORS (\d+)
Value TX_DROPPED (\d+)
Value TX_OVERRUNS (\d+)
Value TX_CARRIER (\d+)
Value COLLISIONS (\d+)
Value TX_QUEUE_LEN (\d+)
Value RX_BYTES (\d+)
Value TX_BYTES (\d+)

Start
  ^\s*${NIC}\s+Link\s+encap:${LINK_ENCAP}(?:\s+HWaddr\s+${HARDWARE_ADDRESS})?\s*$$
  ^\s*inet\s+addr:${IP_ADDRESS}(?:\s+Bcast:${BCAST})?\s+Mask:${MASK}\s*$$
  ^\s*link-local6:\s+${LINK_LOCAL6}\s+prefixlen\s+${LINK_LOCAL6_PREFIXLEN}\s*$$
  ^\s*${MULTICAST}\s+MTU:${MTU}\s+Metric:${METRIC}\s*$$
  ^\s*RX\s+packets:${RX_PACKETS}\s+errors:${RX_ERRORS}\s+dropped:${RX_DROPPED}\s+overruns:${RX_OVERRUNS}\s+frame:${RX_FRAME}\s*$$
  ^\s*TX\s+packets:${TX_PACKETS}\s+errors:${TX_ERRORS}\s+dropped:${TX_DROPPED}\s+overruns:${TX_OVERRUNS}\s+carrier:${TX_CARRIER}\s*$$
  ^\s*collisions:${COLLISIONS}\s+txqueuelen:${TX_QUEUE_LEN}\s*$$
  ^\s*RX\s+bytes:${RX_BYTES}\s+\(.*\)\s+TX\s+bytes:${TX_BYTES}\s+\(.*\)\s*$$ -> Record
  ^\s*$$
  ^. -> Error
`
