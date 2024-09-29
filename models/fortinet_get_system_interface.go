package models

type FortinetGetSystemInterface struct {
	Name	string	`json:"NAME"`
	Mode	string	`json:"MODE"`
	Management_ip	string	`json:"MANAGEMENT_IP"`
	Management_netmask	string	`json:"MANAGEMENT_NETMASK"`
	Ip_address	string	`json:"IP_ADDRESS"`
	Netmask	string	`json:"NETMASK"`
	Status	string	`json:"STATUS"`
	Netbios_forward	string	`json:"NETBIOS_FORWARD"`
	Type	string	`json:"TYPE"`
	Netflow_sampler	string	`json:"NETFLOW_SAMPLER"`
	Sflow_sampler	string	`json:"SFLOW_SAMPLER"`
	Scan_botnet_connections	string	`json:"SCAN_BOTNET_CONNECTIONS"`
	Src_check	string	`json:"SRC_CHECK"`
	Explicit_web_proxy	string	`json:"EXPLICIT_WEB_PROXY"`
	Explicit_ftp_proxy	string	`json:"EXPLICIT_FTP_PROXY"`
	Proxy_captive_portal	string	`json:"PROXY_CAPTIVE_PORTAL"`
	Switch_controller_feature	string	`json:"SWITCH_CONTROLLER_FEATURE"`
	Mtu_override	string	`json:"MTU_OVERRIDE"`
	Wccp	string	`json:"WCCP"`
	Drop_overlapped_fragment	string	`json:"DROP_OVERLAPPED_FRAGMENT"`
	Drop_fragment	string	`json:"DROP_FRAGMENT"`
}

var FortinetGetSystemInterface_Template = `#
# Patrick Marc Preuss, Refried Jello, Pavlo Skliarenko
# 2020-01-13 - Initial Version
# 2021-01-08 - Update for 6.2
# 2023-09-09 - Update for 6.4.2
#
# FG Version: 5.6, 6.0.7, 6.2, 6.4.2
# HW        : varied
# HA        : AP
# VDOMS     : ENABLED
Value Required NAME (\S+)
Value MODE (\S+)
Value MANAGEMENT_IP (\d+?\.\d+?\.\d+?\.\d+?)
Value MANAGEMENT_NETMASK (\d+?\.\d+?\.\d+?\.\d+?)
Value IP_ADDRESS (\d+?\.\d+?\.\d+?\.\d+?)
Value NETMASK (\d+?\.\d+?\.\d+?\.\d+?)
Value STATUS (\S+)
Value NETBIOS_FORWARD (\S+)
Value TYPE (\S+)
Value NETFLOW_SAMPLER (\S+)
Value SFLOW_SAMPLER (\S+)
Value SCAN_BOTNET_CONNECTIONS (\S+)
Value SRC_CHECK (\S+)
Value EXPLICIT_WEB_PROXY (\S+)
Value EXPLICIT_FTP_PROXY (\S+)
Value PROXY_CAPTIVE_PORTAL (\S+)
Value SWITCH_CONTROLLER_FEATURE (\S+)
Value MTU_OVERRIDE (\S+)
Value WCCP (\S+)
Value DROP_OVERLAPPED_FRAGMENT (\S+)
Value DROP_FRAGMENT (\S+)

Start
  ^==\s+\[\s+\S+\s\]$$
  # mgmt
  # name:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:scan-botnet-connections:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:mtu-override:wccp:drop-overlapped-fragment:drop-fragment
  ^name:\s+${NAME}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+scan-botnet-connections:\s+${SCAN_BOTNET_CONNECTIONS}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # modem <=6.0
  # name:mode:management-ip:ip:netbios-forward:type:netflow-sampler:sflow-sampler:scan-botnet-connections:src-check:proxy-captive-portal:mtu-override:wccp:drop-overlapped-fragment:drop-fragment
  ^name:\s+${NAME}\s+mode:\s+${MODE}\s+management-ip:\s+${MANAGEMENT_IP}\s+${MANAGEMENT_NETMASK}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+scan-botnet-connections:\s+${SCAN_BOTNET_CONNECTIONS}\s+src-check:\s+${SRC_CHECK}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # modem =6.2
  # name:mode:management-ip:ip:netbios-forward:type:netflow-sampler:sflow-sampler:src-check:proxy-captive-portal:mtu-override:wccp:drop-overlapped-fragment:drop-fragment
  ^name:\s+${NAME}\s+mode:\s+${MODE}\s+management-ip:\s+${MANAGEMENT_IP}\s+${MANAGEMENT_NETMASK}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # type physical and vlan
  # name:mode:management-ip:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:scan-botnet-connections:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:mtu-override:wccp:drop-overlapped-fragment:drop-fragment
  ^name:\s+${NAME}\s+mode:\s+${MODE}\s+management-ip:\s+${MANAGEMENT_IP}\s+${MANAGEMENT_NETMASK}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+scan-botnet-connections:\s+${SCAN_BOTNET_CONNECTIONS}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # ha
  # name:mode:management-ip:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:mtu-override:wccp:drop-overlapped-fragment:drop-fragment
  ^name:\s+${NAME}\s+mode:\s+${MODE}\s+management-ip:\s+${MANAGEMENT_IP}\s+${MANAGEMENT_NETMASK}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # type tunnel
  # name:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:scan-botnet-connections:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:wccp
  ^name:\s+${NAME}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+scan-botnet-connections:\s+${SCAN_BOTNET_CONNECTIONS}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+wccp:\s+${WCCP}\s*$$ -> Record
  # vw1
  # name:status:type:netflow-sampler:sflow-sampler:src-check:mtu-override:wccp:drop-overlapped-fragment:drop-fragment
  ^name:\s+${NAME}\s+status:\s+${STATUS}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # ssl. that have less output - unsure why different
  # name:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:scan-botnet-connections:src-check:wccp
  ^name:\s+${NAME}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+scan-botnet-connections:\s+${SCAN_BOTNET_CONNECTIONS}\s+src-check:\s+${SRC_CHECK}\s+wccp:\s+${WCCP}\sv$$ -> Record
  # ssl. =6.2
  # name:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:wccp
  ^name:\s+${NAME}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+wccp:\s+${WCCP}\s*$$ -> Record
  # type loopback
  # name:management-ip:ip:status:type:netflow-sampler:sflow-sampler:scan-botnet-connections:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal
  ^name:\s+${NAME}\s+management-ip:\s+${MANAGEMENT_IP}\s+${MANAGEMENT_NETMASK}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+scan-botnet-connections:\s+${SCAN_BOTNET_CONNECTIONS}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s*$$ -> Record
  # loopback =6.2
  # name:management-ip:ip:status:type:netflow-sampler:sflow-sampler:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal
  ^name:\s+${NAME}\s+management-ip:\s+${MANAGEMENT_IP}\s+${MANAGEMENT_NETMASK}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s*$$ -> Record
  # npu_vlink
  # name:mode:management-ip:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:wccp:drop-overlapped-fragment:drop-fragment
  ^name:\s+${NAME}\s+mode:\s+${MODE}\s+management-ip:\s+${MANAGEMENT_IP}\s+${MANAGEMENT_NETMASK}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # name:mode:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:mtu-override:wccp:drop-overlapped-fragment:drop-fragment:
  ^name:\s+${NAME}\s+mode:\s+${MODE}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # name:mode:ip:netbios-forward:type:netflow-sampler:sflow-sampler:src-check:proxy-captive-portal:mtu-override:wccp:drop-overlapped-fragment:drop-fragment:
  ^name:\s+${NAME}\s+mode:\s+${MODE}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # name:mode:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:switch-controller-feature:mtu-override:wccp:drop-overlapped-fragment:drop-fragment:
  ^name:\s+${NAME}\s+mode:\s+${MODE}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+switch-controller-feature:\s+${SWITCH_CONTROLLER_FEATURE}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s+drop-overlapped-fragment:\s+${DROP_OVERLAPPED_FRAGMENT}\s+drop-fragment:\s+${DROP_FRAGMENT}\s*$$ -> Record
  # name:ip:status:type:netflow-sampler:sflow-sampler:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:
  ^name:\s+${NAME}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s*$$ -> Record
  # name:ip:status:netbios-forward:type:netflow-sampler:sflow-sampler:src-check:explicit-web-proxy:explicit-ftp-proxy:proxy-captive-portal:mtu-override:wccp:
  ^name:\s+${NAME}\s+ip:\s+${IP_ADDRESS}\s+${NETMASK}\s+status:\s+${STATUS}\s+netbios-forward:\s+${NETBIOS_FORWARD}\s+type:\s+${TYPE}\s+netflow-sampler:\s+${NETFLOW_SAMPLER}\s+sflow-sampler:\s+${SFLOW_SAMPLER}\s+src-check:\s+${SRC_CHECK}\s+explicit-web-proxy:\s+${EXPLICIT_WEB_PROXY}\s+explicit-ftp-proxy:\s+${EXPLICIT_FTP_PROXY}\s+proxy-captive-portal:\s+${PROXY_CAPTIVE_PORTAL}\s+mtu-override:\s+${MTU_OVERRIDE}\s+wccp:\s+${WCCP}\s*$$ -> Record
  # name:status:type
  ^name:\s+${NAME}\s+status:\s+${STATUS}\s+type:\s+${TYPE}\s*$$ -> Record
  ^\s*$$
  ^. -> Error

`