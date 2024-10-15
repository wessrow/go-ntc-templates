package cisco_ios

type ShowInterfaces struct {
	Delay              string `json:"DELAY"`
	Encapsulation      string `json:"ENCAPSULATION"`
	Output_pps         string `json:"OUTPUT_PPS"`
	Frame              string `json:"FRAME"`
	Output_rate        string `json:"OUTPUT_RATE"`
	Input_pps          string `json:"INPUT_PPS"`
	Queue_output_drops string `json:"QUEUE_OUTPUT_DROPS"`
	Hardware_type      string `json:"HARDWARE_TYPE"`
	Bandwidth          string `json:"BANDWIDTH"`
	Input_errors       string `json:"INPUT_ERRORS"`
	Duplex             string `json:"DUPLEX"`
	Queue_strategy     string `json:"QUEUE_STRATEGY"`
	Input_rate         string `json:"INPUT_RATE"`
	Output_packets     string `json:"OUTPUT_PACKETS"`
	Bia                string `json:"BIA"`
	Description        string `json:"DESCRIPTION"`
	Ip_address         string `json:"IP_ADDRESS"`
	Prefix_length      string `json:"PREFIX_LENGTH"`
	Vlan_id_outer      string `json:"VLAN_ID_OUTER"`
	Input_packets      string `json:"INPUT_PACKETS"`
	Output_errors      string `json:"OUTPUT_ERRORS"`
	Queue_drops        string `json:"QUEUE_DROPS"`
	Queue_flushes      string `json:"QUEUE_FLUSHES"`
	Mac_address        string `json:"MAC_ADDRESS"`
	Speed              string `json:"SPEED"`
	Media_type         string `json:"MEDIA_TYPE"`
	Last_output        string `json:"LAST_OUTPUT"`
	Overrun            string `json:"OVERRUN"`
	Abort              string `json:"ABORT"`
	Interface          string `json:"INTERFACE"`
	Mtu                string `json:"MTU"`
	Last_output_hang   string `json:"LAST_OUTPUT_HANG"`
	Runts              string `json:"RUNTS"`
	Giants             string `json:"GIANTS"`
	Queue_max          string `json:"QUEUE_MAX"`
	Vlan_id            string `json:"VLAN_ID"`
	Vlan_id_inner      string `json:"VLAN_ID_INNER"`
	Queue_size         string `json:"QUEUE_SIZE"`
	Link_status        string `json:"LINK_STATUS"`
	Protocol_status    string `json:"PROTOCOL_STATUS"`
	Last_input         string `json:"LAST_INPUT"`
	Crc                string `json:"CRC"`
}

var ShowInterfaces_Template string = `Value Required INTERFACE (\S+)
Value LINK_STATUS (.+?)
Value PROTOCOL_STATUS (.+?)
Value HARDWARE_TYPE ([\w \-]+)
Value MAC_ADDRESS ([a-fA-F0-9]{4}\.[a-fA-F0-9]{4}\.[a-fA-F0-9]{4})
Value BIA ([a-fA-F0-9]{4}\.[a-fA-F0-9]{4}\.[a-fA-F0-9]{4})
Value DESCRIPTION (.+?)
Value IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value PREFIX_LENGTH (\d+)
Value MTU (\d+)
Value DUPLEX (([Ff]ull|[Aa]uto|[Hh]alf|[Aa]-).*?)
Value SPEED (.*?)
Value MEDIA_TYPE (\S+.*)
Value BANDWIDTH (\d+\s+\w+)
Value DELAY (\d+\s+\S+)
Value ENCAPSULATION (.+?)
Value LAST_INPUT (.+?)
Value LAST_OUTPUT (.+?)
Value LAST_OUTPUT_HANG (.+?)
Value QUEUE_STRATEGY (.+)
Value INPUT_RATE (\d+)
Value OUTPUT_RATE (\d+)
Value INPUT_PPS (\d+)
Value OUTPUT_PPS (\d+)
Value INPUT_PACKETS (\d+)
Value OUTPUT_PACKETS (\d+)
Value RUNTS (\d+)
Value GIANTS (\d+)
Value INPUT_ERRORS (\d+)
Value CRC (\d+)
Value FRAME (\d+)
Value OVERRUN (\d+)
Value ABORT (\d+)
Value OUTPUT_ERRORS (\d+)
Value VLAN_ID (\d+)
Value VLAN_ID_INNER (\d+)
Value VLAN_ID_OUTER (\d+)
Value QUEUE_SIZE (\d+)
Value QUEUE_MAX (\d+) 
Value QUEUE_DROPS (\d+)
Value QUEUE_FLUSHES (\d+)
Value QUEUE_OUTPUT_DROPS (\d+)

Start
  ^\S+\s+is\s+.+?,\s+line\s+protocol.*$$ -> Continue.Record
  ^${INTERFACE}\s+is\s+${LINK_STATUS},\s+line\s+protocol\s+is\s+${PROTOCOL_STATUS}\s*$$
  ^\s+Hardware\s+is\s+${HARDWARE_TYPE} -> Continue
  ^.+address\s+is\s+${MAC_ADDRESS}\s+\(bia\s+${BIA}\)\s*$$
  ^\s+Description:\s+${DESCRIPTION}\s*$$
  ^\s+Internet\s+address\s+is\s+${IP_ADDRESS}\/${PREFIX_LENGTH}\s*$$
  ^\s+MTU\s+${MTU}.*BW\s+${BANDWIDTH}.*DLY\s+${DELAY},\s*$$
  ^\s+Encapsulation\s+${ENCAPSULATION}, Vlan ID\s+${VLAN_ID}
  ^\s+Encapsulation\s+${ENCAPSULATION}, outer ID\s+${VLAN_ID_OUTER}, inner ID\s+${VLAN_ID_INNER}
  ^\s+Encapsulation\s+${ENCAPSULATION},.+$$
  ^\s+Last\s+input\s+${LAST_INPUT},\s+output\s+${LAST_OUTPUT},\s+output\s+hang\s+${LAST_OUTPUT_HANG}\s*$$
  ^\s+Input\s+queue:\s+${QUEUE_SIZE}\/${QUEUE_MAX}\/${QUEUE_DROPS}\/${QUEUE_FLUSHES}\s+\(size\/max\/drops\/flushes\);\s+Total output\s+drops:\s+${QUEUE_OUTPUT_DROPS}\s*$$
  ^\s+Queueing\s+strategy:\s+${QUEUE_STRATEGY}\s*$$
  ^\s+${DUPLEX},\s+${SPEED},.+media\stype\sis\s${MEDIA_TYPE}$$
  ^\s+${DUPLEX},\s+${SPEED},.+TX/FX$$
  ^\s+${DUPLEX},\s+${SPEED}$$
  ^.*input\s+rate\s+${INPUT_RATE}\s+\w+/sec,\s+${INPUT_PPS}\s+packets.+$$
  ^.*output\s+rate\s+${OUTPUT_RATE}\s+\w+/sec,\s+${OUTPUT_PPS}\s+packets.+$$
  ^\s+${INPUT_PACKETS}\s+packets\s+input,\s+\d+\s+bytes,\s+\d+\s+no\s+buffer\s*$$
  ^\s+${RUNTS}\s+runts,\s+${GIANTS}\s+giants,\s+\d+\s+throttles\s*$$
  ^\s+${INPUT_ERRORS}\s+input\s+errors,\s+${CRC}\s+CRC,\s+${FRAME}\s+frame,\s+${OVERRUN}\s+overrun,\s+\d+\s+ignored\s*$$
  ^\s+${INPUT_ERRORS}\s+input\s+errors,\s+${CRC}\s+CRC,\s+${FRAME}\s+frame,\s+${OVERRUN}\s+overrun,\s+\d+\s+ignored,\s+${ABORT}\s+abort\s*$$
  ^\s+${OUTPUT_PACKETS}\s+packets\s+output,\s+\d+\s+bytes,\s+\d+\s+underruns\s*$$
  ^\s+${OUTPUT_ERRORS}\s+output\s+errors,\s+\d+\s+collisions,\s+\d+\s+interface\s+resets\s*$$
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
`
