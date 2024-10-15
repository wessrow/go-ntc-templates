package cisco_asa

type ShowVpnSessiondbDetailL2l struct {
	Login_day                       string `json:"LOGIN_DAY"`
	Login_time_zone                 string `json:"LOGIN_TIME_ZONE"`
	Idle_timeout_interval           string `json:"IDLE_TIMEOUT_INTERVAL"`
	Index                           string `json:"INDEX"`
	Ip_address                      string `json:"IP_ADDRESS"`
	Idle_timeout_interval_unit      string `json:"IDLE_TIMEOUT_INTERVAL_UNIT"`
	Reval_timeout                   string `json:"REVAL_TIMEOUT"`
	Reval_timout_unit               string `json:"REVAL_TIMOUT_UNIT"`
	Status_query_interval           string `json:"STATUS_QUERY_INTERVAL"`
	Posture_holdtime_remaining      string `json:"POSTURE_HOLDTIME_REMAINING"`
	Total_bytes_received            string `json:"TOTAL_BYTES_RECEIVED"`
	Login_month                     string `json:"LOGIN_MONTH"`
	Idle_timeout_remaining          string `json:"IDLE_TIMEOUT_REMAINING"`
	Packets_transmitted             string `json:"PACKETS_TRANSMITTED"`
	Total_ipsec_sessions            string `json:"TOTAL_IPSEC_SESSIONS"`
	Session_id                      string `json:"SESSION_ID"`
	Rekey_interval_unit             string `json:"REKEY_INTERVAL_UNIT"`
	Remote_address_network          string `json:"REMOTE_ADDRESS_NETWORK"`
	Eap_over_udp_timer              string `json:"EAP_OVER_UDP_TIMER"`
	Login_weekday                   string `json:"LOGIN_WEEKDAY"`
	Negotiaion_mode                 string `json:"NEGOTIAION_MODE"`
	Rekey_data_interval_unit        string `json:"REKEY_DATA_INTERVAL_UNIT"`
	Bytes_received                  string `json:"BYTES_RECEIVED"`
	Encryption                      string `json:"ENCRYPTION"`
	Authentication_mode             string `json:"AUTHENTICATION_MODE"`
	Total_bytes_transmitted         string `json:"TOTAL_BYTES_TRANSMITTED"`
	Reval_timeout_remaining         string `json:"REVAL_TIMEOUT_REMAINING"`
	Prf                             string `json:"PRF"`
	Redirect_url                    string `json:"REDIRECT_URL"`
	Udp_dst_port                    string `json:"UDP_DST_PORT"`
	Idle_timeout_remaining_unit     string `json:"IDLE_TIMEOUT_REMAINING_UNIT"`
	Posture_token                   string `json:"POSTURE_TOKEN"`
	Rekey_time_left_unit            string `json:"REKEY_TIME_LEFT_UNIT"`
	Eap_over_udp_timer_unit         string `json:"EAP_OVER_UDP_TIMER_UNIT"`
	Hashing                         string `json:"HASHING"`
	Login_year                      string `json:"LOGIN_YEAR"`
	Rekey_data_interval             string `json:"REKEY_DATA_INTERVAL"`
	Connection                      string `json:"CONNECTION"`
	Protocol                        string `json:"PROTOCOL"`
	Ipv6_filter_name                string `json:"IPV6_FILTER_NAME"`
	Encryption_method               string `json:"ENCRYPTION_METHOD"`
	Rekey_data_remaining_unit       string `json:"REKEY_DATA_REMAINING_UNIT"`
	Packets_received                string `json:"PACKETS_RECEIVED"`
	Duration                        string `json:"DURATION"`
	Hash_method                     string `json:"HASH_METHOD"`
	Reval_timeout_remaining_unit    string `json:"REVAL_TIMEOUT_REMAINING_UNIT"`
	Posture_holdtime_remaining_unit string `json:"POSTURE_HOLDTIME_REMAINING_UNIT"`
	Udp_src_port                    string `json:"UDP_SRC_PORT"`
	Bytes_transmitted               string `json:"BYTES_TRANSMITTED"`
	Session_type                    string `json:"SESSION_TYPE"`
	Remote_address_mask             string `json:"REMOTE_ADDRESS_MASK"`
	Rekey_time_left                 string `json:"REKEY_TIME_LEFT"`
	Rekey_data_remaining            string `json:"REKEY_DATA_REMAINING"`
	Dh_group                        string `json:"DH_GROUP"`
	Local_address_network           string `json:"LOCAL_ADDRESS_NETWORK"`
	Status_query_interval_unit      string `json:"STATUS_QUERY_INTERVAL_UNIT"`
	Total_ike_sessions              string `json:"TOTAL_IKE_SESSIONS"`
	Rekey_interval                  string `json:"REKEY_INTERVAL"`
	Connection_type                 string `json:"CONNECTION_TYPE"`
	Remote_authentication_mode      string `json:"REMOTE_AUTHENTICATION_MODE"`
	Local_authentication_mode       string `json:"LOCAL_AUTHENTICATION_MODE"`
	Local_address_mask              string `json:"LOCAL_ADDRESS_MASK"`
	Encapsulation                   string `json:"ENCAPSULATION"`
	Pfs_group                       string `json:"PFS_GROUP"`
	Login_time                      string `json:"LOGIN_TIME"`
	Filter_name                     string `json:"FILTER_NAME"`
}

var ShowVpnSessiondbDetailL2l_Template string = `Value Filldown,Required SESSION_TYPE (\S+)
Value Filldown CONNECTION (\S+)
Value Filldown INDEX (\d+)
Value Filldown IP_ADDRESS (\d+\.\d+\.\d+\.\d+)
Value Filldown PROTOCOL (.+?)
Value Filldown ENCRYPTION (.+?)
Value Filldown HASHING (.+?)
Value Filldown TOTAL_BYTES_TRANSMITTED (\d+)
Value Filldown TOTAL_BYTES_RECEIVED (\d+)
Value Filldown LOGIN_TIME (\d+:\d+:\d+)
Value Filldown LOGIN_TIME_ZONE (\S+)
Value Filldown LOGIN_WEEKDAY (\w+)
Value Filldown LOGIN_MONTH (\w+)
Value Filldown LOGIN_DAY (\d+)
Value Filldown LOGIN_YEAR (\d+)
Value Filldown DURATION (.+?)
Value Filldown FILTER_NAME (.*?)
Value Filldown TOTAL_IKE_SESSIONS (\d+)
Value Filldown TOTAL_IPSEC_SESSIONS (\d+)
Value CONNECTION_TYPE (\S+)
Value SESSION_ID (\d+)
Value UDP_SRC_PORT (\d+)
Value UDP_DST_PORT (\d+)
Value NEGOTIAION_MODE (\w+)
Value AUTHENTICATION_MODE (\w+)
Value REMOTE_AUTHENTICATION_MODE (\S+|)
Value LOCAL_AUTHENTICATION_MODE (\S+|)
Value ENCRYPTION_METHOD (\S+)
Value HASH_METHOD (\w+)
Value REKEY_INTERVAL (\d+)
Value REKEY_INTERVAL_UNIT (\S+)
Value REKEY_TIME_LEFT (\d+)
Value REKEY_TIME_LEFT_UNIT (\S+)
Value REKEY_DATA_INTERVAL (\d+)
Value REKEY_DATA_INTERVAL_UNIT (\S+)
Value REKEY_DATA_REMAINING (\d+)
Value REKEY_DATA_REMAINING_UNIT (\S+)
Value IDLE_TIMEOUT_INTERVAL (\d+)
Value IDLE_TIMEOUT_INTERVAL_UNIT (\S+)
Value IDLE_TIMEOUT_REMAINING (\d+)
Value IDLE_TIMEOUT_REMAINING_UNIT (\S+)
Value PRF (\S+)
Value DH_GROUP (\d+)
Value IPV6_FILTER_NAME (.*?)
Value LOCAL_ADDRESS_NETWORK (\d+\.\d+\.\d+\.\d+)
Value LOCAL_ADDRESS_MASK (\d+\.\d+\.\d+\.\d+)
Value REMOTE_ADDRESS_NETWORK (\d+\.\d+\.\d+\.\d+)
Value REMOTE_ADDRESS_MASK (\d+\.\d+\.\d+\.\d+)
Value ENCAPSULATION (\w+)
Value PFS_GROUP (\d+)
Value BYTES_TRANSMITTED (\d+)
Value BYTES_RECEIVED (\d+)
Value PACKETS_TRANSMITTED (\d+)
Value PACKETS_RECEIVED (\d+)
Value REVAL_TIMEOUT (\d+)
Value REVAL_TIMOUT_UNIT (\S+)
Value REVAL_TIMEOUT_REMAINING (\d+)
Value REVAL_TIMEOUT_REMAINING_UNIT (\S+)
Value STATUS_QUERY_INTERVAL (\S+)
Value STATUS_QUERY_INTERVAL_UNIT (\S+)
Value EAP_OVER_UDP_TIMER (\d+)
Value EAP_OVER_UDP_TIMER_UNIT (\S+)
Value POSTURE_HOLDTIME_REMAINING (\d+)
Value POSTURE_HOLDTIME_REMAINING_UNIT (\S+)
Value POSTURE_TOKEN (.*?)
Value REDIRECT_URL (.*?)


Start
  ^Session\s+Type:\s+${SESSION_TYPE}\s+Detailed\s*$$ -> Connection
  ^\s*$$
  ^. -> Error

Connection
  ^\s*Connection\s*:\s+${CONNECTION}\s*$$
  ^\s*Index\s*:\s+${INDEX}\s+IP\s+Addr\s*:\s+${IP_ADDRESS}\s*$$
  ^\s*Protocol\s*:\s+${PROTOCOL}(?:\s+Encryption\s*:\s+${ENCRYPTION}|)\s*$$
  ^\s*Encryption\s*:\s+${ENCRYPTION}\s+Hashing\s*:\s+${HASHING}\s*$$
  ^\s*Encryption\s*:\s+${ENCRYPTION}\s*$$
  ^\s*Hashing\s*:\s+${HASHING}\s*$$
  ^\s*Bytes\s+Tx\s*:\s+${TOTAL_BYTES_TRANSMITTED}\s+Bytes\s+Rx\s*:\s+${TOTAL_BYTES_RECEIVED}\s*$$
  ^\s*Login\s+Time\s*:\s+${LOGIN_TIME}\s+${LOGIN_TIME_ZONE}\s+${LOGIN_WEEKDAY}\s+${LOGIN_MONTH}\s+${LOGIN_DAY}\s+${LOGIN_YEAR}\s*$$
  ^\s*Duration\s*:\s+${DURATION}\s*$$
  ^\s*Filter\s+Name\s*:\s*${FILTER_NAME}\s*$$
  ^\s*IKE(?:[Vv]\d|)\s+Sessions:\s+${TOTAL_IKE_SESSIONS}\s+IPSec\s+Sessions:\s+${TOTAL_IPSEC_SESSIONS}\s*$$
  ^\s*IKE(?:[Vv]\d|)\s+Tunnels:\s*${TOTAL_IKE_SESSIONS}\s*$$
  ^\s*IP[Ss]ec\s+Tunnels:\s*${TOTAL_IPSEC_SESSIONS}\s*$$
  ^\s*IP[Ss]ecOverNatT\s+Tunnels:\s*${TOTAL_IPSEC_SESSIONS}\s*$$
  ^\s*${CONNECTION_TYPE}:\s*$$ -> Continue
  ^\s*IKE(?:[Vv]\d|): -> IKE
  ^\s*IP[Ss]ec: -> IPSec
  ^\s*NAC: -> NAC
  ^\s*Connection\s*: -> Continue.Record
  ^\s*Connection\s*:\s+${CONNECTION}\s*$$
  ^Session\s+Type -> Continue.Record
  ^Session\s+Type -> Continue.Clearall
  ^Session\s+Type:\s+${SESSION_TYPE}\s+Detailed\s*$$
  ^\s*$$
  ^. -> Error

IKE
  ^\s*(Session|Tunnel)\s+ID\s*:\s+(?:\d+\.|)${SESSION_ID}\s*$$
  ^\s*UDP\s+Src\s+Port\s*:\s+${UDP_SRC_PORT}\s+UDP\s+Dst\s+Port\s*:\s+${UDP_DST_PORT}\s*$$
  ^\s*Rem\s+Auth\s+Mode\s*:\s*${REMOTE_AUTHENTICATION_MODE}\s*$$
  ^\s*Loc\s+Auth\s+Mode\s*:\s*${LOCAL_AUTHENTICATION_MODE}\s*$$
  ^\s*IKE\s+Neg\s+Mode\s*:\s+${NEGOTIAION_MODE}\s+Auth\s+Mode\s*:\s+${AUTHENTICATION_MODE}\s*$$
  ^\s*Encryption\s*:\s+${ENCRYPTION_METHOD}\s+Hashing\s*:\s+${HASH_METHOD}\s*$$
  ^\s*Encapsulation\s+:\s*${ENCAPSULATION}\s*$$
  ^\s*Rekey\s+Int\s+\([Tt]\):\s+${REKEY_INTERVAL}\s+${REKEY_INTERVAL_UNIT}\s+Rekey\s+Left\([Tt]\):\s+${REKEY_TIME_LEFT}\s+${REKEY_TIME_LEFT_UNIT}\s*$$
  ^\s*Rekey\s+Int\s+\([Dd]\):\s+${REKEY_DATA_INTERVAL}\s+${REKEY_DATA_INTERVAL_UNIT}\s+Rekey\s+Left\([Dd]+\):\s+${REKEY_DATA_REMAINING}\s+${REKEY_DATA_REMAINING_UNIT}\s*$$
  ^\s*(?:PRF\s*:\s+${PRF}\s+|)D\/H\s+Group\s*:\s+${DH_GROUP}\s*$$
  ^\s*Filter\s+Name\s+:\s*${FILTER_NAME}\s*$$
  ^\s*IPv6\s+Filter\s+:\s*${IPV6_FILTER_NAME}\s*$$
  ^\s*\S+:\s*$$ -> Continue.Record
  ^\s*${CONNECTION_TYPE}:\s*$$ -> Continue
  ^\s*IKE(?:[Vv]\d|): -> IKE
  ^\s*IP[Ss]ec(?:OverNatT|): -> IPSec
  ^\s*NAC: -> NAC
  ^\s*Connection\s*: -> Continue.Record
  ^\s*Connection\s*:\s+${CONNECTION}\s*$$ -> Connection
  ^Session\s+Type -> Continue.Record
  ^Session\s+Type -> Continue.Clearall
  ^Session\s+Type:\s+${SESSION_TYPE}\s+Detailed\s*$$ -> Connection
  ^\s*$$
  ^. -> Error

IPSec
  ^\s*(Session|Tunnel)\s+ID\s*:\s+(?:\d+\.|)${SESSION_ID}\s*$$
  ^\s*Local\s+Addr\s*:\s+${LOCAL_ADDRESS_NETWORK}\/${LOCAL_ADDRESS_MASK}
  ^\s*Remote\s+Addr\s*:\s+${REMOTE_ADDRESS_NETWORK}\/${REMOTE_ADDRESS_MASK}
  ^\s*Encryption\s*:\s+${ENCRYPTION_METHOD}\s+Hashing\s*:\s+${HASH_METHOD}\s*$$
  ^\s*Encapsulation\s*:\s+${ENCAPSULATION}(?:\s+PFS\s+Group\s*:\s+${PFS_GROUP}|)\s*$$
  ^\s*Rekey\s+Int\s+\([Tt]\):\s+${REKEY_INTERVAL}\s+${REKEY_INTERVAL_UNIT}\s+Rekey\s+Left\([Tt]\):\s+${REKEY_TIME_LEFT}\s+${REKEY_TIME_LEFT_UNIT}\s*$$
  ^\s*Rekey\s+Int\s+\([Dd]\):\s+${REKEY_DATA_INTERVAL}\s+${REKEY_DATA_INTERVAL_UNIT}\s+Rekey\s+Left\([Dd]+\):\s+${REKEY_DATA_REMAINING}\s+${REKEY_DATA_REMAINING_UNIT}\s*$$
  ^\s*Idle\s+Time\s+Out\s*:\s+${IDLE_TIMEOUT_INTERVAL}\s+${IDLE_TIMEOUT_INTERVAL_UNIT}\s+Idle\s+TO\s+Left\s*:\s+${IDLE_TIMEOUT_REMAINING}\s+${IDLE_TIMEOUT_REMAINING_UNIT}\s*$$             
  ^\s*Bytes\s+Tx\s*:\s+${BYTES_TRANSMITTED}\s+Bytes\s+Rx\s*:\s+${BYTES_RECEIVED}\s*$$
  ^\s*Pkts\s+Tx\s*:\s+${PACKETS_TRANSMITTED}\s+Pkts\s+Rx\s*:\s+${PACKETS_RECEIVED}\s*$$
  ^\s*\S+:\s*$$ -> Continue.Record
  ^\s*${CONNECTION_TYPE}:\s*$$ -> Continue
  ^\s*IKE(?:[Vv]\d|): -> IKE
  ^\s*IP[Ss]ec(?:OverNatT|): -> IPSec
  ^\s*NAC: -> NAC
  ^\s*Connection\s*: -> Continue.Record
  ^\s*Connection\s*:\s+${CONNECTION}\s*$$ -> Connection
  ^Session\s+Type -> Continue.Record
  ^Session\s+Type -> Continue.Clearall
  ^Session\s+Type:\s+${SESSION_TYPE}\s+Detailed\s*$$ -> Connection
  ^\s*$$
  ^. -> Error

NAC
  ^\s*Reval\s+Int\s+\(\w\)\s*:\s+${REVAL_TIMEOUT}\s+${REVAL_TIMOUT_UNIT}\s+Reval\s+Left\s*\(\w\)\s*:\s+${REVAL_TIMEOUT_REMAINING}\s+${REVAL_TIMEOUT_REMAINING_UNIT}\s*$$
  ^\s*SQ\s+Int\s+\(\w\)\s*:\s+${STATUS_QUERY_INTERVAL}\s+${STATUS_QUERY_INTERVAL_UNIT}\s+EoU\s+Age\(\w\)\s*:\s+${EAP_OVER_UDP_TIMER}\s+${EAP_OVER_UDP_TIMER_UNIT}\s*$$
  ^\s*Hold\s+Left\s+\(\w\)\s*:\s+${POSTURE_HOLDTIME_REMAINING}\s+${POSTURE_HOLDTIME_REMAINING_UNIT}\s+Posture\s+Token\s*:\s*${POSTURE_TOKEN}\s*$$
  ^\s*Redirect\s+URL\s*:\s*${REDIRECT_URL}\s*$$
  ^\s*\S+:\s*$$ -> Continue.Record
  ^\s*${CONNECTION_TYPE}:\s*$$ -> Continue
  ^\s*IKE(?:[Vv]\d|): -> IKE
  ^\s*IP[Ss]ec(?:OverNatT|): -> IPSec
  ^\s*NAC: -> NAC
  ^\s*Connection\s*: -> Continue.Record
  ^\s*Connection\s*:\s+${CONNECTION}\s*$$ -> Connection
  ^Session\s+Type -> Continue.Record
  ^Session\s+Type -> Continue.Clearall
  ^Session\s+Type:\s+${SESSION_TYPE}\s+Detailed\s*$$ -> Connection
  ^\s*$$
  ^. -> Error
  `
