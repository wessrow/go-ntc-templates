package cisco_ios 

type ShowIpOspfNeighborDetail struct {
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

var ShowIpOspfNeighborDetail_Template = `Value NEIGHBOR_ID (\S+)
Value ADDRESS (\S+)
Value AREA (\S+)
Value INTERFACE (\S+)
Value PRIORITY (\d+)
Value STATE (\S+/\s+-|\S+)
Value STATE_CHANGES (\d+)
Value DR (\S+)
Value BDR (\S+)
Value OPTIONS (0[xX][0-9a-fA-F]+)
Value LLS_OPTIONS (.+?(?=,))
Value LAST_OOB_RESYNC (\d+:\d+:\d+)
Value DEAD_TIME (\d+:\d+:\d+)
Value UP_TIME (\d+:\d+:\d+)
Value RQ_INDEX (\d+/\d+)
Value RQ_LENGTH (\d+)
Value RETRANSMISSION_NUMBER (\d+)
Value FIRST_PACKET (\S+)
Value NEXT_PACKET (\S+)
Value LAST_RETRANSMISSION_SCAN_LENGTH (\d+)
Value MAX_RETRANSMISSION_SCAN_LENGTH (\d+)
Value LAST_RETRANSMISSION_SCAN_TIME (\d+)
Value MAX_RETRANSMISSION_SCAN_TIME (\d+)

Start
  ^\s*Neighbor\s\S+,\sinterface\saddress\s\S+\s*$$ -> Continue.Record
  ^\s*Neighbor\s${NEIGHBOR_ID},\sinterface\saddress\s${ADDRESS}\s*$$
  ^\s*In\sthe\sarea\s${AREA}\svia\sinterface\s${INTERFACE}\s*$$
  ^\s*Neighbor\spriority\sis\s${PRIORITY},\sState\sis\s${STATE},\s+${STATE_CHANGES}\sstate\schanges\s*$$
  ^\s*DR\sis\s${DR}\sBDR\sis\s${BDR}\s*$$
  ^\s*Options\sis\s${OPTIONS}\s*$$
  ^\s*LLS\sOptions\sis\s${LLS_OPTIONS},\slast\sOOB-Resync\s${LAST_OOB_RESYNC}\sago\s*$$
  ^\s*Dead\stimer\sdue\sin\s${DEAD_TIME}\s*$$
  ^\s*Neighbor\sis\sup\sfor\s${UP_TIME}\s*$$
  ^\s*Index\s${RQ_INDEX},\sretransmission\squeue\slength\s${RQ_LENGTH},\snumber\sof\sretransmission\s${RETRANSMISSION_NUMBER}\s*$$
  ^\s*First\s${FIRST_PACKET}\sNext\s${NEXT_PACKET}\s*$$
  ^\s*Last\sretransmission\sscan\slength\sis\s${LAST_RETRANSMISSION_SCAN_LENGTH},\smaximum\sis\s${MAX_RETRANSMISSION_SCAN_LENGTH}\s*$$
  ^\s*Last\sretransmission\sscan\stime\sis\s${LAST_RETRANSMISSION_SCAN_TIME}\smsec,\smaximum\sis\s${MAX_RETRANSMISSION_SCAN_TIME}\smsec\s*$$
  ^\s*$$
  ^. -> Error
`