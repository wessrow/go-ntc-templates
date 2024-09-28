package models

type CiscoIosShowIpOspfNeighborDetail struct {
	Neighbor_id	string	`json:"NEIGHBOR_ID"`
	Address	string	`json:"ADDRESS"`
	Area	string	`json:"AREA"`
	Interface	string	`json:"INTERFACE"`
	Priority	string	`json:"PRIORITY"`
	State	string	`json:"STATE"`
	State_changes	string	`json:"STATE_CHANGES"`
	Dr	string	`json:"DR"`
	Bdr	string	`json:"BDR"`
	Options	string	`json:"OPTIONS"`
	Lls_options	string	`json:"LLS_OPTIONS"`
	Last_oob_resync	string	`json:"LAST_OOB_RESYNC"`
	Dead_time	string	`json:"DEAD_TIME"`
	Up_time	string	`json:"UP_TIME"`
	Rq_index	string	`json:"RQ_INDEX"`
	Rq_length	string	`json:"RQ_LENGTH"`
	Retransmission_number	string	`json:"RETRANSMISSION_NUMBER"`
	First_packet	string	`json:"FIRST_PACKET"`
	Next_packet	string	`json:"NEXT_PACKET"`
	Last_retransmission_scan_length	string	`json:"LAST_RETRANSMISSION_SCAN_LENGTH"`
	Max_retransmission_scan_length	string	`json:"MAX_RETRANSMISSION_SCAN_LENGTH"`
	Last_retransmission_scan_time	string	`json:"LAST_RETRANSMISSION_SCAN_TIME"`
	Max_retransmission_scan_time	string	`json:"MAX_RETRANSMISSION_SCAN_TIME"`
}