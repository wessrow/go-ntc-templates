package cisco_asa

type ShowAspDrop struct {
	Flow_drop_ctm_crypto_request_error     string `json:"FLOW_DROP_CTM_CRYPTO_REQUEST_ERROR"`
	Cluster_ccl_unknown_stub               string `json:"CLUSTER_CCL_UNKNOWN_STUB"`
	Dispatch_block_alloc                   string `json:"DISPATCH_BLOCK_ALLOC"`
	Mp_svc_no_session                      string `json:"MP_SVC_NO_SESSION"`
	Np_sp_invalid_spi                      string `json:"NP_SP_INVALID_SPI"`
	Tcp_rstfin_ooo                         string `json:"TCP_RSTFIN_OOO"`
	Tcp_bad_option_list                    string `json:"TCP_BAD_OPTION_LIST"`
	Tcp_paws_fail                          string `json:"TCP_PAWS_FAIL"`
	Dispatch_queue_limit                   string `json:"DISPATCH_QUEUE_LIMIT"`
	Flow_being_freed                       string `json:"FLOW_BEING_FREED"`
	Last_cleared                           string `json:"LAST_CLEARED"`
	Np_socket_closed                       string `json:"NP_SOCKET_CLOSED"`
	Ctm_error                              string `json:"CTM_ERROR"`
	Lu_invalid_pkt                         string `json:"LU_INVALID_PKT"`
	Rpf_violated                           string `json:"RPF_VIOLATED"`
	Unable_to_add_flow                     string `json:"UNABLE_TO_ADD_FLOW"`
	Unsupported_ip_version                 string `json:"UNSUPPORTED_IP_VERSION"`
	Fragment_reassembly_failed             string `json:"FRAGMENT_REASSEMBLY_FAILED"`
	Nat_xlate_failed                       string `json:"NAT_XLATE_FAILED"`
	Rm_inspect_rate_limit                  string `json:"RM_INSPECT_RATE_LIMIT"`
	Fo_standby                             string `json:"FO_STANDBY"`
	No_adjacency                           string `json:"NO_ADJACENCY"`
	Ttl_exceeded                           string `json:"TTL_EXCEEDED"`
	Np_socket_lock_failure                 string `json:"NP_SOCKET_LOCK_FAILURE"`
	Security_failed                        string `json:"SECURITY_FAILED"`
	Flow_drop_np_socket_data_move_failure  string `json:"FLOW_DROP_NP_SOCKET_DATA_MOVE_FAILURE"`
	Bad_ipsec_natt                         string `json:"BAD_IPSEC_NATT"`
	Host_limit                             string `json:"HOST_LIMIT"`
	Punt_rate_limit                        string `json:"PUNT_RATE_LIMIT"`
	Tcp_invalid_ack                        string `json:"TCP_INVALID_ACK"`
	Flow_drop_svc_failover                 string `json:"FLOW_DROP_SVC_FAILOVER"`
	Tcp_3whs_failed                        string `json:"TCP_3WHS_FAILED"`
	Shunned                                string `json:"SHUNNED"`
	Flow_drop_nat_rpf_failed               string `json:"FLOW_DROP_NAT_RPF_FAILED"`
	Flow_drop_out_of_memory                string `json:"FLOW_DROP_OUT_OF_MEMORY"`
	Flow_drop_ssl_handshake_failed         string `json:"FLOW_DROP_SSL_HANDSHAKE_FAILED"`
	Inspect_dns_pak_too_long               string `json:"INSPECT_DNS_PAK_TOO_LONG"`
	Inspect_icmp_seq_num_not_matched       string `json:"INSPECT_ICMP_SEQ_NUM_NOT_MATCHED"`
	Invalid_encap                          string `json:"INVALID_ENCAP"`
	Flow_drop_vpn_handle_not_found         string `json:"FLOW_DROP_VPN_HANDLE_NOT_FOUND"`
	L2_acl                                 string `json:"L2_ACL"`
	Mp_svc_flow_control                    string `json:"MP_SVC_FLOW_CONTROL"`
	Flow_drop_acl_drop                     string `json:"FLOW_DROP_ACL_DROP"`
	Intercept_unexpected                   string `json:"INTERCEPT_UNEXPECTED"`
	Interface_down                         string `json:"INTERFACE_DOWN"`
	Flow_drop_ipsec_selector_failure       string `json:"FLOW_DROP_IPSEC_SELECTOR_FAILURE"`
	Cluster_ccl_unknown                    string `json:"CLUSTER_CCL_UNKNOWN"`
	Inspect_dns_invalid_pak                string `json:"INSPECT_DNS_INVALID_PAK"`
	Inspect_dp_out_of_memory               string `json:"INSPECT_DP_OUT_OF_MEMORY"`
	Tcp_synack_ooo                         string `json:"TCP_SYNACK_OOO"`
	Inspect_icmp_bad_code                  string `json:"INSPECT_ICMP_BAD_CODE"`
	Tcp_fo_drop                            string `json:"TCP_FO_DROP"`
	Flow_drop_shunned                      string `json:"FLOW_DROP_SHUNNED"`
	Flow_drop_ssl_bad_record_detect        string `json:"FLOW_DROP_SSL_BAD_RECORD_DETECT"`
	Flow_drop_svc_replacement_conn         string `json:"FLOW_DROP_SVC_REPLACEMENT_CONN"`
	Flow_drop_ssl_record_decrypt_error     string `json:"FLOW_DROP_SSL_RECORD_DECRYPT_ERROR"`
	Tcp_conn_limit                         string `json:"TCP_CONN_LIMIT"`
	Cp_syslog_event_queue_error            string `json:"CP_SYSLOG_EVENT_QUEUE_ERROR"`
	Cluster_ccl_backup                     string `json:"CLUSTER_CCL_BACKUP"`
	Cp_event_queue_error                   string `json:"CP_EVENT_QUEUE_ERROR"`
	Inspect_icmp_invalid_pak               string `json:"INSPECT_ICMP_INVALID_PAK"`
	Ipsec_tun_down                         string `json:"IPSEC_TUN_DOWN"`
	Punt_limit                             string `json:"PUNT_LIMIT"`
	Ssl_first_record_invalid               string `json:"SSL_FIRST_RECORD_INVALID"`
	Conn_limit                             string `json:"CONN_LIMIT"`
	Punt_no_mem                            string `json:"PUNT_NO_MEM"`
	Flow_last_cleared                      string `json:"FLOW_LAST_CLEARED"`
	Invalid_ip_length                      string `json:"INVALID_IP_LENGTH"`
	Tcp_global_buffer_full                 string `json:"TCP_GLOBAL_BUFFER_FULL"`
	Tcp_not_syn                            string `json:"TCP_NOT_SYN"`
	Cluster_bad_trailer                    string `json:"CLUSTER_BAD_TRAILER"`
	Mp_svc_bad_length                      string `json:"MP_SVC_BAD_LENGTH"`
	No_route                               string `json:"NO_ROUTE"`
	Object_group_search_threshold_exceeded string `json:"OBJECT_GROUP_SEARCH_THRESHOLD_EXCEEDED"`
	Inspect_dns_id_not_matched             string `json:"INSPECT_DNS_ID_NOT_MATCHED"`
	Invalid_ip_header                      string `json:"INVALID_IP_HEADER"`
	Rm_conn_rate_limit                     string `json:"RM_CONN_RATE_LIMIT"`
	Rule_transaction_in_progress           string `json:"RULE_TRANSACTION_IN_PROGRESS"`
	Tcp_buffer_full                        string `json:"TCP_BUFFER_FULL"`
	Cluster_non_ip_pkt                     string `json:"CLUSTER_NON_IP_PKT"`
	Invalid_ip_option                      string `json:"INVALID_IP_OPTION"`
	Connection_lock                        string `json:"CONNECTION_LOCK"`
	Inspect_dns_invalid_domain_label       string `json:"INSPECT_DNS_INVALID_DOMAIN_LABEL"`
	Unable_to_create_flow                  string `json:"UNABLE_TO_CREATE_FLOW"`
	Bad_tcp_cksum                          string `json:"BAD_TCP_CKSUM"`
	Bad_tcp_flags                          string `json:"BAD_TCP_FLAGS"`
	Flow_expired                           string `json:"FLOW_EXPIRED"`
	Natt_keepalive                         string `json:"NATT_KEEPALIVE"`
	Tcp_ack_syn_diff                       string `json:"TCP_ACK_SYN_DIFF"`
	Flow_drop_svc_selector_failure         string `json:"FLOW_DROP_SVC_SELECTOR_FAILURE"`
	Invalid_tcp_hdr_length                 string `json:"INVALID_TCP_HDR_LENGTH"`
	No_inspect                             string `json:"NO_INSPECT"`
	Punt_queue_limit                       string `json:"PUNT_QUEUE_LIMIT"`
	Tcp_data_past_fin                      string `json:"TCP_DATA_PAST_FIN"`
	Flow_drop_cluster_redirect             string `json:"FLOW_DROP_CLUSTER_REDIRECT"`
	Flow_drop_dtls_hello_close             string `json:"FLOW_DROP_DTLS_HELLO_CLOSE"`
	Acl_drop                               string `json:"ACL_DROP"`
	Inspect_icmp_error_no_existing_conn    string `json:"INSPECT_ICMP_ERROR_NO_EXISTING_CONN"`
	Mcast_in_nonactive_device              string `json:"MCAST_IN_NONACTIVE_DEVICE"`
	Nat_failed                             string `json:"NAT_FAILED"`
	Tcp_dup_in_queue                       string `json:"TCP_DUP_IN_QUEUE"`
	Tcp_seq_syn_diff                       string `json:"TCP_SEQ_SYN_DIFF"`
	Mp_svc_no_fragment                     string `json:"MP_SVC_NO_FRAGMENT"`
	Tcp_rst_syn_in_win                     string `json:"TCP_RST_SYN_IN_WIN"`
	Tcp_seq_past_win                       string `json:"TCP_SEQ_PAST_WIN"`
	Tcp_buffer_timeout                     string `json:"TCP_BUFFER_TIMEOUT"`
	Tcp_dual_open                          string `json:"TCP_DUAL_OPEN"`
	Flow_drop_inspect_fail                 string `json:"FLOW_DROP_INSPECT_FAIL"`
	Flow_drop_vpn_overlap_conflict         string `json:"FLOW_DROP_VPN_OVERLAP_CONFLICT"`
	Async_lock_queue_limit                 string `json:"ASYNC_LOCK_QUEUE_LIMIT"`
	Ifc_classify                           string `json:"IFC_CLASSIFY"`
	Invalid_udp_length                     string `json:"INVALID_UDP_LENGTH"`
	Sp_security_failed                     string `json:"SP_SECURITY_FAILED"`
	Flow_drop_vpn_missing_decrypt          string `json:"FLOW_DROP_VPN_MISSING_DECRYPT"`
	Mp_svc_no_channel                      string `json:"MP_SVC_NO_CHANNEL"`
	Nat_no_xlate_to_pat_pool               string `json:"NAT_NO_XLATE_TO_PAT_POOL"`
}

var ShowAspDrop_Template string = `Value FLOW_DROP_ACL_DROP (\d+)
Value FLOW_DROP_SHUNNED (\d+)
Value FLOW_DROP_CLUSTER_REDIRECT (\d+)
Value FLOW_DROP_INSPECT_FAIL (\d+)
Value FLOW_DROP_IPSEC_SELECTOR_FAILURE (\d+)
Value FLOW_DROP_NAT_RPF_FAILED (\d+)
Value FLOW_DROP_OUT_OF_MEMORY (\d+)
Value FLOW_DROP_SSL_BAD_RECORD_DETECT (\d+)
Value FLOW_DROP_SSL_HANDSHAKE_FAILED (\d+)
Value FLOW_DROP_VPN_OVERLAP_CONFLICT (\d+)
Value FLOW_DROP_VPN_MISSING_DECRYPT (\d+)
Value FLOW_DROP_SVC_REPLACEMENT_CONN (\d+)
Value FLOW_DROP_SVC_SELECTOR_FAILURE (\d+)
Value FLOW_DROP_SVC_FAILOVER (\d+)
Value FLOW_DROP_SSL_RECORD_DECRYPT_ERROR (\d+)
Value FLOW_DROP_CTM_CRYPTO_REQUEST_ERROR (\d+)
Value FLOW_DROP_VPN_HANDLE_NOT_FOUND (\d+)
Value FLOW_DROP_NP_SOCKET_DATA_MOVE_FAILURE (\d+)
Value FLOW_DROP_DTLS_HELLO_CLOSE (\d+)
Value FLOW_LAST_CLEARED (.+?)
Value ACL_DROP (\d+)
Value ASYNC_LOCK_QUEUE_LIMIT (\d+)
Value BAD_IPSEC_NATT (\d+)
Value BAD_TCP_CKSUM (\d+)
Value BAD_TCP_FLAGS (\d+)
Value CLUSTER_BAD_TRAILER (\d+)
Value CLUSTER_CCL_BACKUP (\d+)
Value CLUSTER_CCL_UNKNOWN (\d+)
Value CLUSTER_CCL_UNKNOWN_STUB (\d+)
Value CLUSTER_NON_IP_PKT (\d+)
Value CONN_LIMIT (\d+)
Value CONNECTION_LOCK (\d+)
Value CP_EVENT_QUEUE_ERROR (\d+)
Value CP_SYSLOG_EVENT_QUEUE_ERROR (\d+)
Value CTM_ERROR (\d+)
Value DISPATCH_QUEUE_LIMIT (\d+)
Value DISPATCH_BLOCK_ALLOC (\d+)
Value FLOW_BEING_FREED (\d+)
Value FLOW_EXPIRED (\d+)
Value FO_STANDBY (\d+)
Value FRAGMENT_REASSEMBLY_FAILED (\d+)
Value HOST_LIMIT (\d+)
Value IFC_CLASSIFY (\d+)
Value INSPECT_DNS_ID_NOT_MATCHED (\d+)
Value INSPECT_DNS_INVALID_DOMAIN_LABEL (\d+)
Value INSPECT_DNS_INVALID_PAK (\d+)
Value INSPECT_DNS_PAK_TOO_LONG (\d+)
Value INSPECT_DP_OUT_OF_MEMORY (\d+)
Value INSPECT_ICMP_BAD_CODE (\d+)
Value INSPECT_ICMP_ERROR_NO_EXISTING_CONN (\d+)
Value INSPECT_ICMP_INVALID_PAK (\d+)
Value INSPECT_ICMP_SEQ_NUM_NOT_MATCHED (\d+)
Value INTERCEPT_UNEXPECTED (\d+)
Value INTERFACE_DOWN (\d+)
Value INVALID_ENCAP (\d+)
Value INVALID_IP_HEADER (\d+)
Value INVALID_IP_LENGTH (\d+)
Value INVALID_IP_OPTION (\d+)
Value INVALID_TCP_HDR_LENGTH (\d+)
Value INVALID_UDP_LENGTH (\d+)
Value IPSEC_TUN_DOWN (\d+)
Value L2_ACL (\d+)
Value LAST_CLEARED (.+?)
Value LU_INVALID_PKT (\d+)
Value MCAST_IN_NONACTIVE_DEVICE (\d+)
Value MP_SVC_BAD_LENGTH (\d+)
Value MP_SVC_FLOW_CONTROL (\d+)
Value MP_SVC_NO_CHANNEL (\d+)
Value MP_SVC_NO_FRAGMENT (\d+)
Value MP_SVC_NO_SESSION (\d+)
Value NAT_FAILED (\d+)
Value NAT_NO_XLATE_TO_PAT_POOL (\d+)
Value NAT_XLATE_FAILED (\d+)
Value NATT_KEEPALIVE (\d+)
Value NO_ADJACENCY (\d+)
Value NO_INSPECT (\d+)
Value NO_ROUTE (\d+)
Value NP_SOCKET_CLOSED (\d+)
Value NP_SOCKET_LOCK_FAILURE (\d+)
Value NP_SP_INVALID_SPI (\d+)
Value OBJECT_GROUP_SEARCH_THRESHOLD_EXCEEDED (\d+)
Value PUNT_QUEUE_LIMIT (\d+)
Value PUNT_LIMIT (\d+)
Value PUNT_NO_MEM (\d+)
Value PUNT_RATE_LIMIT (\d+)
Value RM_CONN_RATE_LIMIT (\d+)
Value RM_INSPECT_RATE_LIMIT (\d+)
Value RPF_VIOLATED (\d+)
Value RULE_TRANSACTION_IN_PROGRESS (\d+)
Value SECURITY_FAILED (\d+)
Value SHUNNED (\d+)
Value SP_SECURITY_FAILED (\d+)
Value SSL_FIRST_RECORD_INVALID (\d+)
Value TCP_3WHS_FAILED (\d+)
Value TCP_ACK_SYN_DIFF (\d+)
Value TCP_BAD_OPTION_LIST (\d+)
Value TCP_BUFFER_FULL (\d+)
Value TCP_BUFFER_TIMEOUT (\d+)
Value TCP_CONN_LIMIT (\d+)
Value TCP_DATA_PAST_FIN (\d+)
Value TCP_DUAL_OPEN (\d+)
Value TCP_DUP_IN_QUEUE (\d+)
Value TCP_FO_DROP (\d+)
Value TCP_GLOBAL_BUFFER_FULL (\d+)
Value TCP_INVALID_ACK (\d+)
Value TCP_NOT_SYN (\d+)
Value TCP_PAWS_FAIL (\d+)
Value TCP_RST_SYN_IN_WIN (\d+)
Value TCP_RSTFIN_OOO (\d+)
Value TCP_SEQ_PAST_WIN (\d+)
Value TCP_SEQ_SYN_DIFF (\d+)
Value TCP_SYNACK_OOO (\d+)
Value TTL_EXCEEDED (\d+)
Value UNABLE_TO_ADD_FLOW (\d+)
Value UNABLE_TO_CREATE_FLOW (\d+)
Value UNSUPPORTED_IP_VERSION (\d+)

Start
  ^Frame\s+drop:
  ^\s+$$
  ^$$
  ^Last\s+clearing:\s+${LAST_CLEARED}\s*$$
  ^.+acl-drop\)\s+${ACL_DROP}
  ^.+async-lock-queue-limit\)\s+${ASYNC_LOCK_QUEUE_LIMIT}
  ^.+bad-ipsec-natt\)\s+${BAD_IPSEC_NATT}
  ^.+bad-tcp-cksum\)\s+${BAD_TCP_CKSUM}
  ^.+bad-tcp-flags\)\s+${BAD_TCP_FLAGS}
  ^.+cluster-bad-trailer\)\s+${CLUSTER_BAD_TRAILER}
  ^.+cluster-ccl-backup\)\s+${CLUSTER_CCL_BACKUP}
  ^.+cluster-ccl-unknown\)\s+${CLUSTER_CCL_UNKNOWN}
  ^.+cluster-ccl-unknown-stub\)\s+${CLUSTER_CCL_UNKNOWN_STUB}
  ^.+cluster-non-ip-pkt\)\s+${CLUSTER_NON_IP_PKT}
  ^.+conn-limit\)\s+${CONN_LIMIT}
  ^.+connection-lock\)\s+${CONNECTION_LOCK}
  ^.+cp-event-queue-error\)\s+${CP_EVENT_QUEUE_ERROR}
  ^.+cp-syslog-event-queue-error\)\s+${CP_SYSLOG_EVENT_QUEUE_ERROR}
  ^.+ctm-error\)\s+${CTM_ERROR}
  ^.+dispatch-queue-limit\)\s+${DISPATCH_QUEUE_LIMIT}
  ^.+dispatch-block-alloc\)\s+${DISPATCH_BLOCK_ALLOC}
  ^.+flow-being-freed\)\s+${FLOW_BEING_FREED}
  ^.+flow-expired\)\s+${FLOW_EXPIRED}
  ^.+fo-standby\)\s+${FO_STANDBY}
  ^.+fragment-reassembly-failed\)\s+${FRAGMENT_REASSEMBLY_FAILED}
  ^.+host-limit\)\s+${HOST_LIMIT}
  ^.+ifc-classify\)\s+${IFC_CLASSIFY}
  ^.+inspect-dns-id-not-matched\)\s+${INSPECT_DNS_ID_NOT_MATCHED}
  ^.+inspect-dns-invalid-domain-label\)\s+${INSPECT_DNS_INVALID_DOMAIN_LABEL}
  ^.+inspect-dns-invalid-pak\)\s+${INSPECT_DNS_INVALID_PAK}
  ^.+inspect-dns-pak-too-long\)\s+${INSPECT_DNS_PAK_TOO_LONG}
  ^.+inspect-dp-out-of-memory\)\s+${INSPECT_DP_OUT_OF_MEMORY}
  ^.+inspect-icmp-bad-code\)\s+${INSPECT_ICMP_BAD_CODE}
  ^.+inspect-icmp-error-no-existing-conn\)\s+${INSPECT_ICMP_ERROR_NO_EXISTING_CONN}
  ^.+inspect-icmp-invalid-pak\)\s+${INSPECT_ICMP_INVALID_PAK}
  ^.+inspect-icmp-seq-num-not-matched\)\s+${INSPECT_ICMP_SEQ_NUM_NOT_MATCHED}
  ^.+intercept-unexpected\)\s+${INTERCEPT_UNEXPECTED}
  ^.+interface-down\)\s+${INTERFACE_DOWN}
  ^.+invalid-encap\)\s+${INVALID_ENCAP}
  ^.+invalid-ip-header\)\s+${INVALID_IP_HEADER}
  ^.+invalid-ip-length\)\s+${INVALID_IP_LENGTH}
  ^.+invalid-ip-option\)\s+${INVALID_IP_OPTION}
  ^.+invalid-tcp-hdr-length\)\s+${INVALID_TCP_HDR_LENGTH}
  ^.+invalid-udp-length\)\s+${INVALID_UDP_LENGTH}
  ^.+ipsec-tun-down\)\s+${IPSEC_TUN_DOWN}
  ^.+l2_acl\)\s+${L2_ACL}
  ^.+lu-invalid-pkt\)\s+${LU_INVALID_PKT}
  ^.+mcast-in-nonactive-device\)\s+${MCAST_IN_NONACTIVE_DEVICE}
  ^.+mp-svc-bad-length\)\s+${MP_SVC_BAD_LENGTH}
  ^.+mp-svc-flow-control\)\s+${MP_SVC_FLOW_CONTROL}
  ^.+mp-svc-no-channel\)\s+${MP_SVC_NO_CHANNEL}
  ^.+mp-svc-no-fragment\)\s+${MP_SVC_NO_FRAGMENT}
  ^.+mp-svc-no-session\)\s+${MP_SVC_NO_SESSION}
  ^.+nat-failed\)\s+${NAT_FAILED}
  ^.+nat-no-xlate-to-pat-pool\)\s+${NAT_NO_XLATE_TO_PAT_POOL}
  ^.+natt-keepalive\)\s+${NATT_KEEPALIVE}
  ^.+nat-xlate-failed\)\s+${NAT_XLATE_FAILED}
  ^.+no-adjacency\)\s+${NO_ADJACENCY}
  ^.+no-inspect\)\s+${NO_INSPECT}
  ^.+no-route\)\s+${NO_ROUTE}
  ^.+np-socket-closed\)\s+${NP_SOCKET_CLOSED}
  ^.+np-socket-lock-failu\s+${NP_SOCKET_LOCK_FAILURE}
  ^.+np-sp-invalid-spi\)\s+${NP_SP_INVALID_SPI}
  ^.+object-group-search-threshold-exceeded\)\s+${OBJECT_GROUP_SEARCH_THRESHOLD_EXCEEDED}
  ^.+punt-queue-limit\)\s+${PUNT_QUEUE_LIMIT}
  ^.+punt-limit\)\s+${PUNT_LIMIT}
  ^.+punt-no-mem\)\s+${PUNT_NO_MEM}
  ^.+punt-rate-limit\)\s+${PUNT_RATE_LIMIT}
  ^.+rm-conn-rate-limit\)\s+${RM_CONN_RATE_LIMIT}
  ^.+rm-inspect-rate-limit\)\s+${RM_INSPECT_RATE_LIMIT}
  ^.+rpf-violated\)\s+${RPF_VIOLATED}
  ^.+rule-transaction-in-progress\)\s+${RULE_TRANSACTION_IN_PROGRESS}
  ^.+security-failed\)\s+${SECURITY_FAILED}
  ^.+shunned\)\s+${SHUNNED}
  ^.+sp-security-failed\)\s+${SP_SECURITY_FAILED}
  ^.+ssl-first-record-invalid\)\s+${SSL_FIRST_RECORD_INVALID}
  ^.+tcp-3whs-failed\)\s+${TCP_3WHS_FAILED}
  ^.+tcp-ack-syn-diff\)\s+${TCP_ACK_SYN_DIFF}
  ^.+tcp-bad-option-list\)\s+${TCP_BAD_OPTION_LIST}
  ^.+tcp-buffer-full\)\s+${TCP_BUFFER_FULL}
  ^.+tcp-buffer-timeout\)\s+${TCP_BUFFER_TIMEOUT}
  ^.+tcp-conn-limit\)\s+${TCP_CONN_LIMIT}
  ^.+tcp-data-past-fin\)\s+${TCP_DATA_PAST_FIN}
  ^.+tcp-dual-open\)\s+${TCP_DUAL_OPEN}
  ^.+tcp-dup-in-queue\)\s+${TCP_DUP_IN_QUEUE}
  ^.+tcp-fo-drop\)\s+${TCP_FO_DROP}
  ^.+tcp-global-buffer-full\)\s+${TCP_GLOBAL_BUFFER_FULL}
  ^.+tcp-invalid-ack\)\s+${TCP_INVALID_ACK}
  ^.+tcp-not-syn\)\s+${TCP_NOT_SYN}
  ^.+tcp-paws-fail\)\s+${TCP_PAWS_FAIL}
  ^.+tcp-rst-syn-in-win\)\s+${TCP_RST_SYN_IN_WIN}
  ^.+tcp-rstfin-ooo\)\s+${TCP_RSTFIN_OOO}
  ^.+tcp-seq-past-win\)\s+${TCP_SEQ_PAST_WIN}
  ^.+tcp-seq-syn-diff\)\s+${TCP_SEQ_SYN_DIFF}
  ^.+tcp-synack-ooo\)\s+${TCP_SYNACK_OOO}
  ^.+ttl-exceeded\)\s+${TTL_EXCEEDED}
  ^.+unable-to-add-flow\)\s+${UNABLE_TO_ADD_FLOW}
  ^.+unable-to-create-flow\)\s+${UNABLE_TO_CREATE_FLOW}
  ^.+unsupported-ip-version\)\s+${UNSUPPORTED_IP_VERSION}
  ^Flow\s+drop: -> FLOW
  ^.* -> Error "LINE NOT FOUND"

FLOW
  ^\s+$$
  ^$$
  ^Last\s+clearing:\s+${FLOW_LAST_CLEARED}\s*$$
  ^.+Flow\s+is\s+denied\s+by\s+access\s+rule\s+\(acl-drop\)\s+${FLOW_DROP_ACL_DROP}
  ^.+cluster-redirect\)\s+${FLOW_DROP_CLUSTER_REDIRECT}
  ^.+Flow\s+shunned \(shunned\)\s+${FLOW_DROP_SHUNNED}
  ^.+inspect-fail\)\s+${FLOW_DROP_INSPECT_FAIL}
  ^.+ipsec-selector-failure\)\s+${FLOW_DROP_IPSEC_SELECTOR_FAILURE}
  ^.+nat-rpf-failed\)\s+${FLOW_DROP_NAT_RPF_FAILED}
  ^.+out-of-memory\)\s+${FLOW_DROP_OUT_OF_MEMORY}
  ^.+ssl-bad-record-detect\)\s+${FLOW_DROP_SSL_BAD_RECORD_DETECT}
  ^.+ssl-handshake-failed\)\s+${FLOW_DROP_SSL_HANDSHAKE_FAILED}
  ^.+vpn-overlap-conflict\)\s+${FLOW_DROP_VPN_OVERLAP_CONFLICT}
  ^.+vpn-missing-decrypt\)\s+${FLOW_DROP_VPN_MISSING_DECRYPT}
  ^.+svc-replacement-conn\)\s+${FLOW_DROP_SVC_REPLACEMENT_CONN}
  ^.+svc-selector-failure\)\s+${FLOW_DROP_SVC_SELECTOR_FAILURE}
  ^.+svc-failover\)\s+${FLOW_DROP_SVC_FAILOVER}
  ^.+ssl-record-decrypt-error\)\s+${FLOW_DROP_SSL_RECORD_DECRYPT_ERROR}
  ^.+ctm-crypto-request-error\)\s+${FLOW_DROP_CTM_CRYPTO_REQUEST_ERROR}
  ^.+vpn-handle-not-found\)\s+${FLOW_DROP_VPN_HANDLE_NOT_FOUND}
  ^.+np-socket-data-move-failure\)\s+${FLOW_DROP_NP_SOCKET_DATA_MOVE_FAILURE}
  ^.+dtls-hello-close\)\s+${FLOW_DROP_DTLS_HELLO_CLOSE}
  ^.* -> Error "LINE NOT FOUND"
`
