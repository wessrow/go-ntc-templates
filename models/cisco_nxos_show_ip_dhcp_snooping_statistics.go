package models

type CiscoNxosShowIpDhcpSnoopingStatistics struct {
	Dhcp_discover_rx	string	`json:"DHCP_DISCOVER_RX"`
	Dhcp_discover_tx	string	`json:"DHCP_DISCOVER_TX"`
	Dhcp_discover_drops	string	`json:"DHCP_DISCOVER_DROPS"`
	Dhcp_offer_rx	string	`json:"DHCP_OFFER_RX"`
	Dhcp_offer_tx	string	`json:"DHCP_OFFER_TX"`
	Dhcp_offer_drops	string	`json:"DHCP_OFFER_DROPS"`
	Dhcp_request_rx	string	`json:"DHCP_REQUEST_RX"`
	Dhcp_request_tx	string	`json:"DHCP_REQUEST_TX"`
	Dhcp_request_drops	string	`json:"DHCP_REQUEST_DROPS"`
	Dhcp_ack_rx	string	`json:"DHCP_ACK_RX"`
	Dhcp_ack_tx	string	`json:"DHCP_ACK_TX"`
	Dhcp_ack_drops	string	`json:"DHCP_ACK_DROPS"`
	Dhcp_release_rx	string	`json:"DHCP_RELEASE_RX"`
	Dhcp_release_tx	string	`json:"DHCP_RELEASE_TX"`
	Dhcp_release_drops	string	`json:"DHCP_RELEASE_DROPS"`
	Dhcp_decline_rx	string	`json:"DHCP_DECLINE_RX"`
	Dhcp_decline_tx	string	`json:"DHCP_DECLINE_TX"`
	Dhcp_decline_drops	string	`json:"DHCP_DECLINE_DROPS"`
	Dhcp_inform_rx	string	`json:"DHCP_INFORM_RX"`
	Dhcp_inform_tx	string	`json:"DHCP_INFORM_TX"`
	Dhcp_inform_drops	string	`json:"DHCP_INFORM_DROPS"`
	Dhcp_nack_rx	string	`json:"DHCP_NACK_RX"`
	Dhcp_nack_tx	string	`json:"DHCP_NACK_TX"`
	Dhcp_nack_drops	string	`json:"DHCP_NACK_DROPS"`
	Dhcp_total_messages_rx	string	`json:"DHCP_TOTAL_MESSAGES_RX"`
	Dhcp_total_messages_tx	string	`json:"DHCP_TOTAL_MESSAGES_TX"`
	Dhcp_total_messages_drops	string	`json:"DHCP_TOTAL_MESSAGES_DROPS"`
	Dhcp_l2_forwarding_total_packets_tx	string	`json:"DHCP_L2_FORWARDING_TOTAL_PACKETS_TX"`
	Dhcp_l2_forwarding_total_packets_rx	string	`json:"DHCP_L2_FORWARDING_TOTAL_PACKETS_RX"`
	Dhcp_l2_forwarding_total_packets_dropped	string	`json:"DHCP_L2_FORWARDING_TOTAL_PACKETS_DROPPED"`
	Non_dhcp_total_packets_rx	string	`json:"NON_DHCP_TOTAL_PACKETS_RX"`
	Non_dhcp_total_packets_tx	string	`json:"NON_DHCP_TOTAL_PACKETS_TX"`
	Non_dhcp_total_packets_dropped	string	`json:"NON_DHCP_TOTAL_PACKETS_DROPPED"`
	Dhcp_drop_received_on_untrusted_port	string	`json:"DHCP_DROP_RECEIVED_ON_UNTRUSTED_PORT"`
	Dhcp_drop_unknown_failure	string	`json:"DHCP_DROP_UNKNOWN_FAILURE"`
	Dhcp_drop_source_mac_validation_failed	string	`json:"DHCP_DROP_SOURCE_MAC_VALIDATION_FAILED"`
	Dhcp_drop_binding_entry_validation_failed	string	`json:"DHCP_DROP_BINDING_ENTRY_VALIDATION_FAILED"`
	Dhcp_drop_invalid_dhcp_message_type	string	`json:"DHCP_DROP_INVALID_DHCP_MESSAGE_TYPE"`
	Dhcp_drop_interface_error	string	`json:"DHCP_DROP_INTERFACE_ERROR"`
	Dhcp_drop_tx_over_trusted_port_failed	string	`json:"DHCP_DROP_TX_OVER_TRUSTED_PORT_FAILED"`
	Dhcp_drop_trust_port_not_configured	string	`json:"DHCP_DROP_TRUST_PORT_NOT_CONFIGURED"`
	Dhcp_drop_vlan_validation_failure	string	`json:"DHCP_DROP_VLAN_VALIDATION_FAILURE"`
	Dhcp_drop_insertion_of_option_82_failed	string	`json:"DHCP_DROP_INSERTION_OF_OPTION_82_FAILED"`
	Dhcp_drop_packet_malformed	string	`json:"DHCP_DROP_PACKET_MALFORMED"`
}