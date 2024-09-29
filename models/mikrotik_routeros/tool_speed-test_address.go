package mikrotik_routeros 

type ToolSpeedTestAddress struct {
	Address	string	`json:"ADDRESS"`
	Status	string	`json:"STATUS"`
	Time_remaining	string	`json:"TIME_REMAINING"`
	Ping_min	string	`json:"PING_MIN"`
	Ping_avg	string	`json:"PING_AVG"`
	Ping_max	string	`json:"PING_MAX"`
	Jitter_min	string	`json:"JITTER_MIN"`
	Jitter_avg	string	`json:"JITTER_AVG"`
	Jitter_max	string	`json:"JITTER_MAX"`
	Loss	string	`json:"LOSS"`
	Tcp_download_bw	string	`json:"TCP_DOWNLOAD_BW"`
	Tcp_download_local_cpu	string	`json:"TCP_DOWNLOAD_LOCAL_CPU"`
	Tcp_download_remote_cpu	string	`json:"TCP_DOWNLOAD_REMOTE_CPU"`
	Tcp_upload_bw	string	`json:"TCP_UPLOAD_BW"`
	Tcp_upload_local_cpu	string	`json:"TCP_UPLOAD_LOCAL_CPU"`
	Tcp_upload_remote_cpu	string	`json:"TCP_UPLOAD_REMOTE_CPU"`
	Udp_download_bw	string	`json:"UDP_DOWNLOAD_BW"`
	Udp_download_local_cpu	string	`json:"UDP_DOWNLOAD_LOCAL_CPU"`
	Udp_download_remote_cpu	string	`json:"UDP_DOWNLOAD_REMOTE_CPU"`
	Udp_upload_bw	string	`json:"UDP_UPLOAD_BW"`
	Udp_upload_local_cpu	string	`json:"UDP_UPLOAD_LOCAL_CPU"`
	Udp_upload_remote_cpu	string	`json:"UDP_UPLOAD_REMOTE_CPU"`
}

var ToolSpeedTestAddress_Template = `Value ADDRESS (\S+)
Value STATUS (.*)
Value TIME_REMAINING (\d+\w+)
Value PING_MIN (\S+)
Value PING_AVG (\S+)
Value PING_MAX (\S+)
Value JITTER_MIN (\S+)
Value JITTER_AVG (\S+)
Value JITTER_MAX (\S+)
Value LOSS (\d+%\s*\(\d+/\d+\))
Value TCP_DOWNLOAD_BW (\S+)
Value TCP_DOWNLOAD_LOCAL_CPU (\d+)
Value TCP_DOWNLOAD_REMOTE_CPU (\d+)
Value TCP_UPLOAD_BW (\S+)
Value TCP_UPLOAD_LOCAL_CPU (\d+)
Value TCP_UPLOAD_REMOTE_CPU (\d+)
Value UDP_DOWNLOAD_BW (\S+)
Value UDP_DOWNLOAD_LOCAL_CPU (\d+)
Value UDP_DOWNLOAD_REMOTE_CPU (\d+)
Value UDP_UPLOAD_BW (\S+)
Value UDP_UPLOAD_LOCAL_CPU (\d+)
Value UDP_UPLOAD_REMOTE_CPU (\d+)

Start
  ^\s*address:\s*${ADDRESS}\s*$$
  ^\s*status:\s*${STATUS}\s*$$
  ^\s*time-remaining:\s*${TIME_REMAINING}\s*$$
  ^\s*ping-min-avg-max:\s*${PING_MIN}\s+/\s+${PING_AVG}\s+/\s+${PING_MAX}\s*$$
  ^\s*jitter-min-avg-max:\s*${JITTER_MIN}\s+/\s+${JITTER_AVG}\s+/\s+${JITTER_MAX}\s*$$
  ^\s*loss:\s*${LOSS}\s*$$
  ^\s*tcp-download:\s*${TCP_DOWNLOAD_BW}(\s+local-cpu-load:\s*${TCP_DOWNLOAD_LOCAL_CPU}%)?(\s+remote-cpu-load:\s*${TCP_DOWNLOAD_REMOTE_CPU}%)?\s*$$
  ^\s*tcp-upload:\s*${TCP_UPLOAD_BW}(\s+local-cpu-load:\s*${TCP_UPLOAD_LOCAL_CPU}%)?(\s+remote-cpu-load:\s*${TCP_UPLOAD_REMOTE_CPU}%)?\s*$$
  ^\s*udp-download:\s*${UDP_DOWNLOAD_BW}(\s+local-cpu-load:\s*${UDP_DOWNLOAD_LOCAL_CPU}%)?(\s+remote-cpu-load:\s*${UDP_DOWNLOAD_REMOTE_CPU}%)?\s*$$
  ^\s*udp-upload:\s*${UDP_UPLOAD_BW}(\s+local-cpu-load:\s*${UDP_UPLOAD_LOCAL_CPU}%)?(\s+remote-cpu-load:\s*${UDP_UPLOAD_REMOTE_CPU}%)?\s*$$
  ^\s*--\s+\[Q\s+quit\|D\s+dump\|C-z\s+pause\]\s*$$
  ^\s*--\s+\[Q\s+quit\|C-z\s+pause\]\s*$$
  ^\s*$$
  ^. -> Error
`