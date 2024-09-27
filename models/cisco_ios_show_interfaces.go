package models

type Cisco_ios_show_interfaces struct {
	Interface          string `json:"INTERFACE"`
	Link_status        string `json:"LINK_STATUS"`
	Protocol_status    string `json:"PROTOCOL_STATUS"`
	Hardware_type      string `json:"HARDWARE_TYPE"`
	Mac_address        string `json:"MAC_ADDRESS"`
	Bia                string `json:"BIA"`
	Description        string `json:"DESCRIPTION"`
	Ip_address         string `json:"IP_ADDRESS"`
	Prefix_length      string `json:"PREFIX_LENGTH"`
	Mtu                string `json:"MTU"`
	Duplex             string `json:"DUPLEX"`
	Speed              string `json:"SPEED"`
	Media_type         string `json:"MEDIA_TYPE"`
	Bandwidth          string `json:"BANDWIDTH"`
	Delay              string `json:"DELAY"`
	Encapsulation      string `json:"ENCAPSULATION"`
	Last_input         string `json:"LAST_INPUT"`
	Last_output        string `json:"LAST_OUTPUT"`
	Last_output_hang   string `json:"LAST_OUTPUT_HANG"`
	Queue_strategy     string `json:"QUEUE_STRATEGY"`
	Input_rate         string `json:"INPUT_RATE"`
	Output_rate        string `json:"OUTPUT_RATE"`
	Input_pps          string `json:"INPUT_PPS"`
	Output_pps         string `json:"OUTPUT_PPS"`
	Input_packets      string `json:"INPUT_PACKETS"`
	Output_packets     string `json:"OUTPUT_PACKETS"`
	Runts              string `json:"RUNTS"`
	Giants             string `json:"GIANTS"`
	Input_errors       string `json:"INPUT_ERRORS"`
	Crc                string `json:"CRC"`
	Frame              string `json:"FRAME"`
	Overrun            string `json:"OVERRUN"`
	Abort              string `json:"ABORT"`
	Output_errors      string `json:"OUTPUT_ERRORS"`
	Vlan_id            string `json:"VLAN_ID"`
	Vlan_id_inner      string `json:"VLAN_ID_INNER"`
	Vlan_id_outer      string `json:"VLAN_ID_OUTER"`
	Queue_size         string `json:"QUEUE_SIZE"`
	Queue_max          string `json:"QUEUE_MAX"`
	Queue_drops        string `json:"QUEUE_DROPS"`
	Queue_flushes      string `json:"QUEUE_FLUSHES"`
	Queue_output_drops string `json:"QUEUE_OUTPUT_DROPS"`
}
