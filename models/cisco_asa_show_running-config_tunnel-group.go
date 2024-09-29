package models

type CiscoAsaShowRunningConfigTunnelGroup struct {
	Name	string	`json:"NAME"`
	Type	string	`json:"TYPE"`
	Dhcp_selection	[]string	`json:"DHCP_SELECTION"`
	Dhcp_server	[]string	`json:"DHCP_SERVER"`
	Authorization	string	`json:"AUTHORIZATION"`
	Authorization_intfc	string	`json:"AUTHORIZATION_INTFC"`
	Authorization_grp	string	`json:"AUTHORIZATION_GRP"`
	Accounting_grp	string	`json:"ACCOUNTING_GRP"`
	Ipv4_pool_intfc	string	`json:"IPV4_POOL_INTFC"`
	Ipv4_pool_name	string	`json:"IPV4_POOL_NAME"`
	Authentication_intfc	string	`json:"AUTHENTICATION_INTFC"`
	Authentication_pri_grp	string	`json:"AUTHENTICATION_PRI_GRP"`
	Authentication_sec_grp	string	`json:"AUTHENTICATION_SEC_GRP"`
	Authentication_svr	string	`json:"AUTHENTICATION_SVR"`
	Authenticated_user	string	`json:"AUTHENTICATED_USER"`
	Group_policy	string	`json:"GROUP_POLICY"`
	Ipv6_pool_intfc	string	`json:"IPV6_POOL_INTFC"`
	Ipv6_pool_name	string	`json:"IPV6_POOL_NAME"`
	Nat_intfc	string	`json:"NAT_INTFC"`
	Scep_enroll	string	`json:"SCEP_ENROLL"`
	Sec_authentication_grp_intfc	string	`json:"SEC_AUTHENTICATION_GRP_INTFC"`
	Sec_authentication_pri_grp	string	`json:"SEC_AUTHENTICATION_PRI_GRP"`
	Sec_authentication_sec_grp	string	`json:"SEC_AUTHENTICATION_SEC_GRP"`
	Sec_authentication_use_pri_user	string	`json:"SEC_AUTHENTICATION_USE_PRI_USER"`
	Sec_cert_codes	string	`json:"SEC_CERT_CODES"`
	Sec_cert_fallback	string	`json:"SEC_CERT_FALLBACK"`
	Pri_cert_codes	string	`json:"PRI_CERT_CODES"`
	Pri_cert_fallback	string	`json:"PRI_CERT_FALLBACK"`
	Strip_grp	string	`json:"STRIP_GRP"`
	Strip_realm	string	`json:"STRIP_REALM"`
	Ikev1_psk	string	`json:"IKEV1_PSK"`
	Chain	string	`json:"CHAIN"`
	Client_os	[]string	`json:"CLIENT_OS"`
	Update_url	[]string	`json:"UPDATE_URL"`
	Revs	[]string	`json:"REVS"`
	Trust_point	string	`json:"TRUST_POINT"`
	Ikev1_user_auth_intfc	[]string	`json:"IKEV1_USER_AUTH_INTFC"`
	Ikev1_user_auth_type	[]string	`json:"IKEV1_USER_AUTH_TYPE"`
	Ikev2_local_auth_cert	string	`json:"IKEV2_LOCAL_AUTH_CERT"`
	Ikev2_remote_auth_type	[]string	`json:"IKEV2_REMOTE_AUTH_TYPE"`
	Peer_id_validate	string	`json:"PEER_ID_VALIDATE"`
	Radius_with_expiry	string	`json:"RADIUS_WITH_EXPIRY"`
	Ppp_auth_type	[]string	`json:"PPP_AUTH_TYPE"`
	Webvpn_auth	string	`json:"WEBVPN_AUTH"`
	Webvpn_custom	string	`json:"WEBVPN_CUSTOM"`
	Dns_grp	string	`json:"DNS_GRP"`
	Webvpn_grp_alias	[]string	`json:"WEBVPN_GRP_ALIAS"`
	Webvpn_grp_alias_state	[]string	`json:"WEBVPN_GRP_ALIAS_STATE"`
	Webvpn_grp_url	[]string	`json:"WEBVPN_GRP_URL"`
	Webvpn_grp_url_state	[]string	`json:"WEBVPN_GRP_URL_STATE"`
	Nbns_svr	[]string	`json:"NBNS_SVR"`
	Nbns_master	[]string	`json:"NBNS_MASTER"`
	Nbns_timeout	[]string	`json:"NBNS_TIMEOUT"`
	Nbns_retry	[]string	`json:"NBNS_RETRY"`
	Override_svc_download	string	`json:"OVERRIDE_SVC_DOWNLOAD"`
	Pri_prefill_user_type	[]string	`json:"PRI_PREFILL_USER_TYPE"`
	Sec_prefill_user_type	[]string	`json:"SEC_PREFILL_USER_TYPE"`
	Proxy_auth	string	`json:"PROXY_AUTH"`
	Radius_reject_msg	string	`json:"RADIUS_REJECT_MSG"`
	Saml_id_provider	string	`json:"SAML_ID_PROVIDER"`
	Without_csd	string	`json:"WITHOUT_CSD"`
}

var CiscoAsaShowRunningConfigTunnelGroup_Template = `Value Required NAME (\S+|\d+\.\d+\.\d+\.\d+)
Value Required TYPE (\S+)
Value List DHCP_SELECTION ([linksubet\-co]+)
Value List DHCP_SERVER (\d+\.\d+\.\d+\.\d+)
Value AUTHORIZATION (\S+)
Value AUTHORIZATION_INTFC (\S+)
Value AUTHORIZATION_GRP (\S+)
Value ACCOUNTING_GRP (\S+)
Value IPV4_POOL_INTFC (\S+)
Value IPV4_POOL_NAME (\S+)
Value AUTHENTICATION_INTFC (\S+)
Value AUTHENTICATION_PRI_GRP (\S+)
Value AUTHENTICATION_SEC_GRP (\S+)
Value AUTHENTICATION_SVR (\S+)
Value AUTHENTICATED_USER (\S+)
Value GROUP_POLICY (\S+)
Value IPV6_POOL_INTFC (\S+)
Value IPV6_POOL_NAME (\S+)
Value NAT_INTFC (\S+)
Value SCEP_ENROLL (\S+)
Value SEC_AUTHENTICATION_GRP_INTFC (\S+)
Value SEC_AUTHENTICATION_PRI_GRP (\S+)
Value SEC_AUTHENTICATION_SEC_GRP (\S+)
Value SEC_AUTHENTICATION_USE_PRI_USER (\S+)
Value SEC_CERT_CODES (\S+\s+\S+|\S+)
Value SEC_CERT_FALLBACK (\S+)
Value PRI_CERT_CODES (\S+\s+\S+|\S+)
Value PRI_CERT_FALLBACK (\S+)
Value STRIP_GRP (strip\-group)
Value STRIP_REALM (strip\-realm)
Value IKEV1_PSK (\*+)
Value CHAIN (chain)
Value List CLIENT_OS (\S+)
Value List UPDATE_URL (\S+)
Value List REVS ((\d+,){0,10}\d+)
Value TRUST_POINT (\S+)
Value List IKEV1_USER_AUTH_INTFC (\S+)
Value List IKEV1_USER_AUTH_TYPE (\S+)
Value IKEV2_LOCAL_AUTH_CERT (\S+)
Value List IKEV2_REMOTE_AUTH_TYPE (\S+\s+\S+|\S+)
Value PEER_ID_VALIDATE (\S+)
Value RADIUS_WITH_EXPIRY (radius\-with\-expiry)
Value List PPP_AUTH_TYPE (pap|ms-chap-v2|eap-proxy)
Value WEBVPN_AUTH (\S+\s+\S+|\S+)
Value WEBVPN_CUSTOM (\S+)
Value DNS_GRP (\S+)
Value List WEBVPN_GRP_ALIAS (\S+)
Value List WEBVPN_GRP_ALIAS_STATE (enable|disable)
Value List WEBVPN_GRP_URL (\S+)
Value List WEBVPN_GRP_URL_STATE (enable|disable)
Value List NBNS_SVR (\d+\.\d+\.\d+\.\d+)
Value List NBNS_MASTER (master)
Value List NBNS_TIMEOUT (\d+)
Value List NBNS_RETRY (\d+)
Value OVERRIDE_SVC_DOWNLOAD (override\-svc\-download)
Value List PRI_PREFILL_USER_TYPE (\S+\s+\S+|\S+)
Value List SEC_PREFILL_USER_TYPE (\S+\s+\S+|\S+)
Value PROXY_AUTH (\S+)
Value RADIUS_REJECT_MSG (radius\-reject\-message)
Value SAML_ID_PROVIDER (\S+)
Value WITHOUT_CSD (\S+\s+\S+|\S+)

Start
  ^tunnel\-group\s+(\S+|\d+\.\d+\.\d+\.\d+)\s+type -> Continue.Record
  ^tunnel\-group\s+${NAME}\s+type\s+${TYPE}\s*
  # GENERAL ATTRIBUTES
  ^\s+dhcp\-server(\s+){0,1}(${DHCP_SELECTION}{0,1})\s+${DHCP_SERVER}\s*
  ^\s+authorization\-server\-group(\s+\(${AUTHORIZATION_INTFC}\)){0,1}\s+${AUTHORIZATION_GRP}\s*
  ^\s+authorization\-${AUTHORIZATION}\s*
  ^\s+accounting\-server\-group\s+${ACCOUNTING_GRP}\s*
  ^\s+address\-pool(\s+\(${IPV4_POOL_INTFC}\)){0,1}\s+${IPV4_POOL_NAME}\s*
  ^\s+authentication\-server\-group(\s+\(${AUTHENTICATION_INTFC}\)){0,1}\s+${AUTHENTICATION_PRI_GRP}(\s+${AUTHENTICATION_SEC_GRP}){0,1}\s*
  ^\s+authentication\-attr\-from\-server\s+${AUTHENTICATION_SVR}\s*
  ^\s+authenticated\-session\-username\s+${AUTHENTICATED_USER}\s*
  ^\s+default\-group\-policy\s+${GROUP_POLICY}\s*
  ^\s+ipv6\-address\-pool(\s+\(${IPV6_POOL_INTFC}\)){0,1}\s+${IPV6_POOL_NAME}\s*
  ^\s+nat\-assigned\-to\-public\-ip\s+${NAT_INTFC}\s*
  ^\s+scep\-enrollment\s+${SCEP_ENROLL}\s*
  ^\s+secondary\-authentication\-server\-group(\s+\(${SEC_AUTHENTICATION_GRP_INTFC}\)){0,1}\s+${SEC_AUTHENTICATION_PRI_GRP}(\s+${SEC_AUTHENTICATION_SEC_GRP}){0,1}(\s+${SEC_AUTHENTICATION_USE_PRI_USER}){0,1}\s*
  ^\s+secondary\-username\-from\-certificate\s+${SEC_CERT_CODES}(\s+no-certificate-fallback\s+${SEC_CERT_FALLBACK}){0,1}\s*
  ^\s+username\-from\-certificate\s+${PRI_CERT_CODES}(\s+no-certificate-fallback\s+${PRI_CERT_FALLBACK}){0,1}\s*
  ^\s+${STRIP_GRP}\s*
  ^\s+${STRIP_REALM}\s*
  # PPP ATTRIBUTES
  ^\s+authentication\s+${PPP_AUTH_TYPE}\s*
  # IPSEC ATTRIBUTES
  ^\s+ikev1\s+pre\-shared\-key\s+${IKEV1_PSK}\s*
  ^\s+${CHAIN}\s*
  ^\s+ikev1\s+client\-update\s+type\s+${CLIENT_OS}\s+url\s+${UPDATE_URL}\s+rev\-nums\s+${REVS}\s*
  ^\s+ikev1\s+trust\-point\s+${TRUST_POINT}\s+
  ^\s+ikev1\s+user\-authentication(\s+\(${IKEV1_USER_AUTH_INTFC}\)){0,1}\s+${IKEV1_USER_AUTH_TYPE}\s*
  ^\s+ikev2\s+local\-authentication\s+certificate\s+${IKEV2_LOCAL_AUTH_CERT}\s*
  ^\s+ikev2\s+remote\-authentication\s+${IKEV2_REMOTE_AUTH_TYPE}\s*
  ^\s+peer\-id\-validate\s+${PEER_ID_VALIDATE}\s*
  ^\s+${RADIUS_WITH_EXPIRY}\s*
  # WEBVPN ATTRIBUTES
  ^\s+authentication\s+${WEBVPN_AUTH}\s*
  ^\s+customization\s+${WEBVPN_CUSTOM}\s*
  ^\s+dns\-group\s+${DNS_GRP}\s*
  ^\s+group\-alias\s+${WEBVPN_GRP_ALIAS}\s+${WEBVPN_GRP_ALIAS_STATE}\s*
  ^\s+group\-url\s+${WEBVPN_GRP_URL}\s+${WEBVPN_GRP_URL_STATE}\s*
  ^\s+nbns\-server\s+${NBNS_SVR}(\s+${NBNS_MASTER}){0,1}\s+timeout\s+${NBNS_TIMEOUT}\s+retry\s+${NBNS_RETRY}\s*
  ^\s+${OVERRIDE_SVC_DOWNLOAD}\s*
  ^\s+pre\-fill\-username\s+${PRI_PREFILL_USER_TYPE}\s*
  ^\s+proxy\-auth\s+${PROXY_AUTH}\s*
  ^\s+${RADIUS_REJECT_MSG}\s*
  ^\s+saml\s+identity\-provider${SAML_ID_PROVIDER}\s*
  ^\s+secondary\s+pre\-fill\-username\s+${SEC_PREFILL_USER_TYPE}\S*
  ^\s+without\-csd\s+${WITHOUT_CSD}\s*

`