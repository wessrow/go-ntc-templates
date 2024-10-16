package oneaccess_oneos

type ShowPolicyInterfaceOutput struct {
	Class_parent_priority                string `json:"CLASS_PARENT_PRIORITY"`
	Dscp_remarked                        string `json:"DSCP_REMARKED"`
	Total_queue_length                   string `json:"TOTAL_QUEUE_LENGTH"`
	Qos_bytes_out                        string `json:"QOS_BYTES_OUT"`
	Qos_bytes_dropped_percent            string `json:"QOS_BYTES_DROPPED_PERCENT"`
	Qos_remaining_bw_share               string `json:"QOS_REMAINING_BW_SHARE"`
	Total_mean_output_rate_bps           string `json:"TOTAL_MEAN_OUTPUT_RATE_BPS"`
	Qos_burst_b                          string `json:"QOS_BURST_B"`
	Qos_queue_length                     string `json:"QOS_QUEUE_LENGTH"`
	Class_parent                         string `json:"CLASS_PARENT"`
	Total_mean_input_rate_bps            string `json:"TOTAL_MEAN_INPUT_RATE_BPS"`
	Qos_bytes_dropped                    string `json:"QOS_BYTES_DROPPED"`
	Qos_packets_out                      string `json:"QOS_PACKETS_OUT"`
	Qos_excess_rate_priority             string `json:"QOS_EXCESS_RATE_PRIORITY"`
	Qos_excess_packets                   string `json:"QOS_EXCESS_PACKETS"`
	Packets                              string `json:"PACKETS"`
	Total_output_packets                 string `json:"TOTAL_OUTPUT_PACKETS"`
	Total_output_bytes_dropped           string `json:"TOTAL_OUTPUT_BYTES_DROPPED"`
	Service_policy_out_parent            string `json:"SERVICE_POLICY_OUT_PARENT"`
	Service_policy_out_child             string `json:"SERVICE_POLICY_OUT_CHILD"`
	Class_child                          string `json:"CLASS_CHILD"`
	Bytes                                string `json:"BYTES"`
	Mean_intput_rate_bits_second         string `json:"MEAN_INTPUT_RATE_BITS_SECOND"`
	Total_output_packets_dropped         string `json:"TOTAL_OUTPUT_PACKETS_DROPPED"`
	Total_output_bytes                   string `json:"TOTAL_OUTPUT_BYTES"`
	Qos_packets_dropped                  string `json:"QOS_PACKETS_DROPPED"`
	Qos_packets_dropped_percent          string `json:"QOS_PACKETS_DROPPED_PERCENT"`
	Qos_bw_kb                            string `json:"QOS_BW_KB"`
	Interface                            string `json:"INTERFACE"`
	Qos_average_output_bps               string `json:"QOS_AVERAGE_OUTPUT_BPS"`
	Qos_queue_limit                      string `json:"QOS_QUEUE_LIMIT"`
	Total_bw_kbs                         string `json:"TOTAL_BW_KBS"`
	Total_output_packets_dropped_percent string `json:"TOTAL_OUTPUT_PACKETS_DROPPED_PERCENT"`
	Total_output_bytes_dropped_percent   string `json:"TOTAL_OUTPUT_BYTES_DROPPED_PERCENT"`
	Qos_average_input_bps                string `json:"QOS_AVERAGE_INPUT_BPS"`
}

var ShowPolicyInterfaceOutput_Template string = `Value Filldown INTERFACE (.*)
Value Filldown SERVICE_POLICY_OUT_PARENT (\S+)
Value Filldown CLASS_PARENT (\S+)
Value Filldown CLASS_PARENT_PRIORITY (.*\S)
Value Filldown SERVICE_POLICY_OUT_CHILD (\S+)
Value CLASS_CHILD (\S+)
Value PACKETS (\d+)
Value Required BYTES (\d+)
Value DSCP_REMARKED (\d+)
Value MEAN_INTPUT_RATE_BITS_SECOND (\d+)
Value Fillup TOTAL_BW_KBS (\d+)
Value Fillup TOTAL_QUEUE_LENGTH (\d+)
Value Fillup TOTAL_MEAN_INPUT_RATE_BPS (\d+)
Value Fillup TOTAL_MEAN_OUTPUT_RATE_BPS (\d+)
Value Fillup TOTAL_OUTPUT_PACKETS (\d+)
Value Fillup TOTAL_OUTPUT_PACKETS_DROPPED (\d+)
Value Fillup TOTAL_OUTPUT_PACKETS_DROPPED_PERCENT (\d+)
Value Fillup TOTAL_OUTPUT_BYTES (\d+)
Value Fillup TOTAL_OUTPUT_BYTES_DROPPED (\d+)
Value Fillup TOTAL_OUTPUT_BYTES_DROPPED_PERCENT (\d+)
Value QOS_BYTES_OUT (\d+)
Value QOS_BYTES_DROPPED (\d+)
Value QOS_BYTES_DROPPED_PERCENT (\d+)
Value QOS_AVERAGE_INPUT_BPS (\d+)
Value QOS_AVERAGE_OUTPUT_BPS (\d+)
Value QOS_PACKETS_OUT (\d+)
Value QOS_PACKETS_DROPPED (\d+)
Value QOS_PACKETS_DROPPED_PERCENT (\d+)
Value QOS_BW_KB (\d+)
Value QOS_BURST_B (\d+)
Value QOS_QUEUE_LENGTH (\d+)
Value QOS_QUEUE_LIMIT (\d+)
Value Fillup QOS_REMAINING_BW_SHARE (\d+)
Value Fillup QOS_EXCESS_RATE_PRIORITY (\d+)
Value Fillup QOS_EXCESS_PACKETS (\d+)

# each record is either part of the traffic stats output
#   these records do not have any QOS values filled in
# or the queueing output
#   these records will have QOS values filled in
Start
  ^${INTERFACE}:\sservice-policy\soutput\s${SERVICE_POLICY_OUT_PARENT}
  ^traffic\sstatistics: -> TrafficStats
  ^output\squeuing\sstatistics: -> OutputQueuingStats
  ^\s*$$
  ^. -> Error

TrafficStats
  ^\s{2}Class\s'${CLASS_PARENT}':
  ^\s+${PACKETS}\spackets,\s${BYTES}\sbytes,(:?\s${DSCP_REMARKED}\sdscp\sremarked,)?\smean\sinput\srate\s${MEAN_INTPUT_RATE_BITS_SECOND}\sbits/s -> Record
  ^\s+Service-policy\s${SERVICE_POLICY_OUT_CHILD}\s:
  ^\s+Class\s'${CLASS_CHILD}':
  ^\s+${PACKETS}\spackets,\s${BYTES}\sbytes,(:?\s${DSCP_REMARKED}\sdscp\sremarked,)?\smean\sinput\srate\s${MEAN_INTPUT_RATE_BITS_SECOND}\sbits/s -> Record
  ^\s+Packets\s+dropped\s+by\s+Congestion\s+Management:
  ^\s+cir\s+\d+\s+kbits/s,\s+cbs\s\d+\s+bytes
  ^\s+(conformed|exceeded)\s+\d+\s+packets,\s+\d+\s+bytes;\s+action:
  ^\s*$$ -> Start
  ^. -> Error

OutputQueuingStats
  ^\s+Class\s'${CLASS_PARENT}':\s${CLASS_PARENT_PRIORITY}
  ^\s+bandwidth\s${QOS_BW_KB}\skb/s,\sburst\s${QOS_BURST_B}\sbytes,\squeue\slength/limit\s${QOS_QUEUE_LENGTH}/${QOS_QUEUE_LIMIT} 
  ^\s+mean\sinput\srate\s${QOS_AVERAGE_INPUT_BPS}\sbits/s,\smean\soutput\srate\s${QOS_AVERAGE_OUTPUT_BPS}\sbits/s
  ^\s+packets\soutput\s${QOS_PACKETS_OUT},\spackets\sdropped\s${QOS_PACKETS_DROPPED}\s\(${QOS_PACKETS_DROPPED_PERCENT}%\)
  ^\s+remaining-bandwidth\sshare\s${QOS_REMAINING_BW_SHARE},\sexcess-rate-priority\s${QOS_EXCESS_RATE_PRIORITY},\sexcess\spackets\ssent\s${QOS_EXCESS_PACKETS}
  ^\s+bytes\soutput\s${QOS_BYTES_OUT},\sbytes\sdropped\s${QOS_BYTES_DROPPED}\s\(${QOS_BYTES_DROPPED_PERCENT}%\) -> Continue
  ^\s+bytes\soutput\s${BYTES},\sbytes\sdropped\s\d+\s\(\d+%\) -> Record
  ^\s+Service-policy\s+VDSL_QOS_4M_OUT\s+:
  ^\s+weight\s+random\s+early\s+detection\s*:
  ^\s+exponential\s+weight\s*:
  ^\s+Dscp\s+Random\s+drop\s+\s+Tail\s+drop\s+Min/Max\s+Mark\s*$$
  ^\s+pkts\s+pkts\s+threshold\s+\s+probability\s*$$
  ^\s+\d+\s+\d+\s+\d+\s+\d+/\d+\s+\d+/\d+\s*$$
  ^\s+bandwidth\s+\d+\s+kb/s,\s+burst\s+\d+\s+bytes
  ^\s+Service-policy\s+\S+\s*:
  ^\s+Interface\stotal: -> StatsInterfaceTotal
  ^\*$$
  ^. -> Error

StatsInterfaceTotal
  ^\s+bandwidth\s${TOTAL_BW_KBS}\skb/s,\squeue\slength\s${TOTAL_QUEUE_LENGTH}
  ^\s+mean\sinput\srate\s${TOTAL_MEAN_INPUT_RATE_BPS}\sbits/s,\smean\soutput\srate\s${TOTAL_MEAN_OUTPUT_RATE_BPS}\sbits/s
  ^\s+packets\soutput\s${TOTAL_OUTPUT_PACKETS},\spackets\sdropped\s${TOTAL_OUTPUT_PACKETS_DROPPED}\s\(${TOTAL_OUTPUT_PACKETS_DROPPED_PERCENT}%\)
  ^\s+bytes\soutput\s${TOTAL_OUTPUT_BYTES},\sbytes\sdropped\s${TOTAL_OUTPUT_BYTES_DROPPED}\s\(${TOTAL_OUTPUT_BYTES_DROPPED_PERCENT}%\)
  ^\s+bandwidth\s+\d+\s+kb/s
  ^\s*$$ -> Start
  ^. -> Error
`
