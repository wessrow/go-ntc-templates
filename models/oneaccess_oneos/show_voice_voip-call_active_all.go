package oneaccess_oneos

type ShowVoiceVoipCallActiveAll struct {
	Calling_to           string   `json:"CALLING_TO"`
	Rtp_destination_ip   string   `json:"RTP_DESTINATION_IP"`
	Call_id              string   `json:"CALL_ID"`
	Call_priority        string   `json:"CALL_PRIORITY"`
	Rtp_packets_rx       string   `json:"RTP_PACKETS_RX"`
	Rtp_acom_db          string   `json:"RTP_ACOM_DB"`
	Calling_from         string   `json:"CALLING_FROM"`
	Pdd_duration         string   `json:"PDD_DURATION"`
	Rtp_source_port      string   `json:"RTP_SOURCE_PORT"`
	Rtp_play_time        string   `json:"RTP_PLAY_TIME"`
	Rtp_packets_tx       string   `json:"RTP_PACKETS_TX"`
	Rtp_mos_cq           string   `json:"RTP_MOS_CQ"`
	Setup_time           string   `json:"SETUP_TIME"`
	Rtp_source_ip        string   `json:"RTP_SOURCE_IP"`
	Rtp_excessive_jitter string   `json:"RTP_EXCESSIVE_JITTER"`
	Rtp_erl_db           string   `json:"RTP_ERL_DB"`
	Call_idx             string   `json:"CALL_IDX"`
	Rtp_packets_lost_tx  string   `json:"RTP_PACKETS_LOST_TX"`
	Rtp_mos_lq           string   `json:"RTP_MOS_LQ"`
	Calling_number       string   `json:"CALLING_NUMBER"`
	Called_number        string   `json:"CALLED_NUMBER"`
	Advice_of_charge     string   `json:"ADVICE_OF_CHARGE"`
	Rtp_destination_port string   `json:"RTP_DESTINATION_PORT"`
	Rtp_tx_coder         string   `json:"RTP_TX_CODER"`
	Call_status          string   `json:"CALL_STATUS"`
	B_channels           []string `json:"B_CHANNELS"`
	Connection_time      string   `json:"CONNECTION_TIME"`
	Rtp_rx_coder         string   `json:"RTP_RX_CODER"`
	Rtp_packets_lost_rx  string   `json:"RTP_PACKETS_LOST_RX"`
}

var ShowVoiceVoipCallActiveAll_Template string = `Value Required CALL_IDX (\d+)
Value CALLING_FROM (.*)
Value CALLING_TO (.*)
Value CALL_ID (.*\S)
Value CALL_STATUS (\S+)
Value CALLING_NUMBER (\S*)
Value CALLED_NUMBER (\S+)
Value SETUP_TIME (.*\S)
Value CONNECTION_TIME (.*\S)
Value List B_CHANNELS (\S+)
Value PDD_DURATION (.*)
Value ADVICE_OF_CHARGE (.*)
Value CALL_PRIORITY (\d+)
Value RTP_SOURCE_IP (\d+\.\d+\.\d+\.\d+)
Value RTP_SOURCE_PORT (\d+)
Value RTP_DESTINATION_IP (\d+\.\d+\.\d+\.\d+)
Value RTP_DESTINATION_PORT (\d+)
Value RTP_PLAY_TIME (\S+)
Value RTP_TX_CODER (.*\S)
Value RTP_RX_CODER (.*\S)
Value RTP_PACKETS_RX (\d+)
Value RTP_PACKETS_TX (\d+)
Value RTP_PACKETS_LOST_RX (\d+)
Value RTP_PACKETS_LOST_TX (\d+)
Value RTP_EXCESSIVE_JITTER (\d+)
Value RTP_MOS_CQ (\d+\.\d+)
Value RTP_MOS_LQ (\d+\.\d+)
Value RTP_ERL_DB (\d+|-+)
Value RTP_ACOM_DB (\d+|-+)

Start
  # start record example: 41 - Call from
  ^\s*\d+\s-\s -> Continue.Record
  ^\s*\d+\/\d+\/\d+\s\d
  ^\s*${CALL_IDX}\s-\sCall\sfrom\s${CALLING_FROM},\sto\s${CALLING_TO}\scall-id:\s${CALL_ID}\s${CALL_STATUS} ->  NewCall
  ^\s*$$
  ^. -> Error

NewCall
  ^\s*calling\s*:\s*${CALLING_NUMBER},\s*called\s*:\s*${CALLED_NUMBER}
  ^\s*setup\stime:\s*${SETUP_TIME}
  ^\s*connexion\stime:\s*${CONNECTION_TIME} -> BChannels
  ^\s*advice-of-charge:\s*${ADVICE_OF_CHARGE}
  ^\s*call\spriority:\s${CALL_PRIORITY}
  ^\s*RTP\sSource\sip\s*:${RTP_SOURCE_IP}\srtp:${RTP_SOURCE_PORT}\s*\/\s*Dest\sip\s*:${RTP_DESTINATION_IP}\srtp:${RTP_DESTINATION_PORT} -> RTP
  ^\s*$$ -> Start
  ^. -> Error

BChannels
  ^\s*B\schannel\s\(from\s\S+\)\s:\s${B_CHANNELS}
  ^\s*PDD\sduration:\s*${PDD_DURATION} -> NewCall
  ^\s*$$ -> Start
  ^. -> Error

RTP
  ^\s*Play\stime\s\(voice\)\s*:\s*${RTP_PLAY_TIME}
  ^\s*Tx\sCoder\s*:\s*${RTP_TX_CODER}\s*;\s*Rx\sCoder\s*:\s*${RTP_RX_CODER}
  ^\s*RTP\sPackets\sRX\s*\/\s*TX\s*:\s*${RTP_PACKETS_RX}\s*\/\s*${RTP_PACKETS_TX}
  ^\s*RTP\sPacket\slost.*RX\s*\/\s*TX\s\(RTCP\sreported\)\s*:\s*(${RTP_PACKETS_LOST_RX}|\-+)\s*\/\s*(${RTP_PACKETS_LOST_TX}|\-+)
  ^\s*Number\sof\sExcessive\sJitter\sevents\s*:\s*${RTP_EXCESSIVE_JITTER}
  ^\s*MOS-CQ\s*\/\s*MOS-LQ\s*:\s*${RTP_MOS_CQ}\s*\/\s*${RTP_MOS_LQ}
  ^\s*ERL\s*:\s*${RTP_ERL_DB}
  ^\s*ACOM\s*:\s*${RTP_ACOM_DB}
  ^\s*$$ -> Start
  ^. -> Error
`
