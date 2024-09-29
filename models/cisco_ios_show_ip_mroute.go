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

var CiscoIosShowIpMroute_Template = `Value MULTICAST_SOURCE_IP (\*|\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value MULTICAST_GROUP_IP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value UP_TIME (\S+?)
Value EXPIRATION_TIME (\S+?)
Value RENDEZVOUS_POINT (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value FLAGS (\w*)
Value INCOMING_INTERFACE (\S+)
Value REVERSE_PATH_FORWARDING_NEIGHBOUR_IP (\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})
Value REGISTERING (Registering)
Value List OUTGOING_INTERFACE (\S+)
Value List FORWARD_MODE (Forward\/Sparse|Forward\/Dense)
Value List OUTGOING_MULTICAST_UP_TIME (\S+)
Value List OUTGOING_MULTICAST_EXPIRATION_TIME (\S+?)

Start
  ^\((\*|\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}),\s(\*|\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\) -> Continue.Record
  ^\(${MULTICAST_SOURCE_IP},\s${MULTICAST_GROUP_IP}\),\s${UP_TIME}\/${EXPIRATION_TIME}(,\sRP\s${RENDEZVOUS_POINT})?,\sflags:\s${FLAGS}
  ^\s+Incoming\sinterface:\s${INCOMING_INTERFACE},\sRPF\snbr\s${REVERSE_PATH_FORWARDING_NEIGHBOUR_IP}(,\s${REGISTERING})?
  ^\s+Outgoing\s+interface\s+list:(?:\s+Null|)\s*$$
  ^\s+${OUTGOING_INTERFACE},\s${FORWARD_MODE},\s${OUTGOING_MULTICAST_UP_TIME}\/${OUTGOING_MULTICAST_EXPIRATION_TIME}\s*$$
  ^\s*$$
  ^IP\s+Multicast\s+(?:Forwarding|Routing)
  ^.*[Ff]lags
  ^\s+\S+\s+-\s+
  ^\s+(?:Timers|Interface\s+state):
  # Capture time-stamp if vty line has command time-stamping turned on
  ^Load\s+for\s+
  ^Time\s+source\s+is
  ^. -> Error

`