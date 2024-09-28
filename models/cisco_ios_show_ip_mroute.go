package models

type CiscoIosShowIpMroute struct {
	Multicast_source_ip	string	`json:"MULTICAST_SOURCE_IP"`
	Multicast_group_ip	string	`json:"MULTICAST_GROUP_IP"`
	Up_time	string	`json:"UP_TIME"`
	Expiration_time	string	`json:"EXPIRATION_TIME"`
	Rendezvous_point	string	`json:"RENDEZVOUS_POINT"`
	Flags	string	`json:"FLAGS"`
	Incoming_interface	string	`json:"INCOMING_INTERFACE"`
	Reverse_path_forwarding_neighbour_ip	string	`json:"REVERSE_PATH_FORWARDING_NEIGHBOUR_IP"`
	Registering	string	`json:"REGISTERING"`
	Outgoing_interface	[]string	`json:"OUTGOING_INTERFACE"`
	Forward_mode	[]string	`json:"FORWARD_MODE"`
	Outgoing_multicast_up_time	[]string	`json:"OUTGOING_MULTICAST_UP_TIME"`
	Outgoing_multicast_expiration_time	[]string	`json:"OUTGOING_MULTICAST_EXPIRATION_TIME"`
}