package models

type Index struct {
}

var Index_Template = `# First line is the header fields for columns and is mandatory.
# Regular expressions are supported in all fields except the first.
# Last field supports variable length command completion.
# abc[[xyz]] is expanded to abc(x(y(z)?)?)?, regexp inside [[]] is not supported
#
# Rules of Ordering:
#  - OS in alphabetical order
#  - Template name in length order (longest to shortest)
#  - When Length is the same, use alphabetical order
#  - Keep space between OS's
#
Template, Hostname, Platform, Command

alcatel_aos_show_vlan.textfsm, .*, alcatel_aos, show vlan

alcatel_sros_show_router_bgp_routes_vpn-ipv4.textfsm, .*, alcatel_sros, sh[[ow]] router bgp rou[[tes]] vpn-ipv4
alcatel_sros_show_router_bgp_summary_family.textfsm, .*, alcatel_sros, sh[[ow]] router bgp sum[[mary]] family
alcatel_sros_show_router_isis_adjacency.textfsm, .*, alcatel_sros, sh[[ow]] router isis adj[[acency]]
alcatel_sros_show_router_isis_interface.textfsm, .*, alcatel_sros, sh[[ow]] router isis int[[erface]]
alcatel_sros_show_router_mpls_interface.textfsm, .*, alcatel_sros, sh[[ow]] router mpls int[[erface]]
alcatel_sros_show_router_ospf_interface.textfsm, .*, alcatel_sros, sh[[ow]] router ospf int[[erface]]
alcatel_sros_show_router_rsvp_interface.textfsm, .*, alcatel_sros, sh[[ow]] router rsvp int[[erface]]
alcatel_sros_show_router_ldp_interface.textfsm, .*, alcatel_sros, sh[[ow]] router ldp int[[erface]]
alcatel_sros_show_router_pim_interface.textfsm, .*, alcatel_sros, sh[[ow]] router pim int[[erface]]
alcatel_sros_show_service_id_interface.textfsm, .*, alcatel_sros, sh[[ow]] se[[rvice]] id (\d+ in[[terface]]|in[[terface]])
alcatel_sros_show_service_sap-using.textfsm, .*, alcatel_sros, sh[[ow]] service sap-u[[sing]]
alcatel_sros_show_service_sdp-using.textfsm, .*, alcatel_sros, sh[[ow]] service sdp-using
alcatel_sros_show_router_interface.textfsm, .*, alcatel_sros, sh[[ow]] router (\d+ int[[erface]]|int[[erface]])
alcatel_sros_show_router_mpls_lsp.textfsm, .*, alcatel_sros, sh[[ow]] router mpls lsp
alcatel_sros_show_service_id_base.textfsm, .*, alcatel_sros, sh[[ow]] serv[[ice]] id ba[[se]]
alcatel_sros_show_service_sdp.textfsm, .*, alcatel_sros, sh[[ow]] service sdp
alcatel_sros_show_system_cpu.textfsm, .*, alcatel_sros, sh[[ow]] system cpu
alcatel_sros_oam_mac-ping.textfsm, .*, alcatel_sros, oam mac-pi[[ng]]
alcatel_sros_show_port.textfsm, .*, alcatel_sros, show port
alcatel_sros_show_lag.textfsm, .*, alcatel_sros, show lag
alcatel_sros_ping.textfsm, .*, alcatel_sros, pi[[ng]]

allied_telesis_awplus_show_lldp_neighbors_detail.textfsm, .*, allied_telesis_awplus, sh[[ow]] ll[[dp]] n[[eighbors]] d[[etail]]
allied_telesis_awplus_show_etherchannel_summary.textfsm, .*, allied_telesis_awplus, sh[[ow]] e[[therchannel]] s[[ummary]]
allied_telesis_awplus_show_interface_switchport.textfsm, .*, allied_telesis_awplus, sh[[ow]] in[[terface]] s[[witchport]]
allied_telesis_awplus_show_static-channel-group.textfsm, .*, allied_telesis_awplus, sh[[ow]] st[[atic-channel-group]]
allied_telesis_awplus_show_mac_address-table.textfsm, .*, allied_telesis_awplus, sh[[ow]] mac a[[ddress-table]]
allied_telesis_awplus_show_interface.textfsm, .*, allied_telesis_awplus, sh[[ow]] in[[terface]]
allied_telesis_awplus_show_ip_route.textfsm, .*, allied_telesis_awplus, sh[[ow]] ip ro[[ute]]
allied_telesis_awplus_show_vlan_all.textfsm, .*, allied_telesis_awplus, sh[[ow]] vl[[an]] al[[l]]
allied_telesis_awplus_show_system.textfsm, .*, allied_telesis_awplus, sh[[ow]] sy[[stem]]
allied_telesis_awplus_show_arp.textfsm, .*, allied_telesis_awplus, sh[[ow]] ar[[p]]

arista_eos_show_mac_security_participants_detail.textfsm, .*, arista_eos, sh[[ow]] ma[[c]] secu[[rity]] part[[icipants]] det[[ail]]
arista_eos_show_interfaces_transceiver_detail.textfsm, .*, arista_eos, sh[[ow]] inte[[rfaces]] tr[[ansceiver]] de[[tail]]
arista_eos_show_mac_security_mka_counters.textfsm, .*, arista_eos, sh[[ow]] ma[[c]] secu[[rity]] mk[[a]] count[[ers]]
arista_eos_show_environment_temperature.textfsm, .*, arista_eos, sh[[ow]] en[[vironment]] t[[emperature]]
arista_eos_show_ip_ospf_interface_brief.textfsm, .*, arista_eos, sh[[ow]] i[[p]] o[[spf]] int[[erface]] br[[ief]]
arista_eos_show_interfaces_description.textfsm, .*, arista_eos, sh[[ow]] int[[erfaces]] des[[cription]]
arista_eos_show_interfaces_transceiver.textfsm, .*, arista_eos, sh[[ow]] inte[[rfaces]] tr[[ansceiver]]
arista_eos_show_mac_security_interface.textfsm, .*, arista_eos, sh[[ow]] ma[[c]] secu[[rity]] int[[erface]]
arista_eos_show_lldp_neighbors_detail.textfsm, .*, arista_eos, sh[[ow]] ll[[dp]] nei[[ghbors]] d[[etail]]
arista_eos_show_port-channel_summary.textfsm, .*, arista_eos, sh[[ow]] port-c[[hannel]] s[[ummary]]
arista_eos_show_environment_cooling.textfsm, .*, arista_eos, sh[[ow]] en[[vironment]] c[[ooling]]
arista_eos_show_ip_interface_brief.textfsm, .*, arista_eos, sh[[ow]] i[[p]] int[[erface]] br[[ief]]
arista_eos_show_pim_ipv4_interface.textfsm, .*, arista_eos, sh[[ow]] pim ipv4 int[[erface]]
arista_eos_show_processes_top_once.textfsm, .*, arista_eos, sh[[ow]] pro[[cesses]] t[[op]] o[[nce]]
arista_eos_show_environment_power.textfsm, .*, arista_eos, sh[[ow]] en[[vironment]] p[[ower]]
arista_eos_show_interfaces_status.textfsm, .*, arista_eos, sh[[ow]] int[[erfaces]] st[[atus]]
arista_eos_show_ip_helper-address.textfsm, .*, arista_eos, sh[[ow]] ip he[[lper-address]]
arista_eos_show_mac_address-table.textfsm, .*, arista_eos, sh[[ow]] m[[ac]] ad[[dress-table]]
arista_eos_show_pim_ipv4_neighbor.textfsm, .*, arista_eos, sh[[ow]] pim ipv4 nei[[ghbor]]
arista_eos_show_ip_ospf_database.textfsm, .*, arista_eos, sh[[ow]] i[[p]] o[[spf]] data[[base]]
arista_eos_show_ip_ospf_neighbor.textfsm, .*, arista_eos, sh[[ow]] i[[p]] o[[spf]] nei[[ghbor]]
arista_eos_show_ipv6_bgp_summary.textfsm, .*, arista_eos, sh[[ow]] ipv6 bg[[p]] su[[mmary]]
arista_eos_show_ip_access-lists.textfsm, .*, arista_eos, sh[[ow]] i[[p]] acce[[ss-lists]]
arista_eos_show_ip_ospf_summary.textfsm, .*, arista_eos, sh[[ow]] i[[p]] o[[spf]] sum[[mary]]
arista_eos_show_ip_bgp_summary.textfsm, .*, arista_eos, sh[[ow]] (?:i[[p]] bg[[p]]|bg[[p]] ev[[pn]]) su[[mmary]]
arista_eos_show_isis_neighbors.textfsm, .*, arista_eos, sh[[ow]] isis ne[[ighbors]]
arista_eos_show_lldp_neighbors.textfsm, .*, arista_eos, sh[[ow]] ll[[dp]] nei[[ghbors]]
arista_eos_show_snmp_community.textfsm, .*, arista_eos, sh[[ow]] sn[[mp]] com[[munity]]
arista_eos_show_ip_bgp_detail.textfsm, .*, arista_eos, sh[[ow]] i[[p]] bg[[p]] d[[etail]]
arista_eos_show_reload_cause.textfsm, .*, arista_eos, sh[[ow]] relo[[ad]] ca[[use]]
arista_eos_show_boot-config.textfsm, .*, arista_eos, sh[[ow]] boot-c[[onfig]]
arista_eos_show_interfaces.textfsm, .*, arista_eos, sh[[ow]] inte[[rfaces]]
arista_eos_show_inventory.textfsm, .*, arista_eos, sh[[ow]] inv[[entory]]
arista_eos_show_hostname.textfsm, .*, arista_eos, sh[[ow]] hostn[[ame]]
arista_eos_show_ip_route.textfsm, .*, arista_eos, sh[[ow]] i[[p]] rou[[te]]
arista_eos_show_version.textfsm, .*, arista_eos, sh[[ow]] ver[[sion]]
arista_eos_show_ip_arp.textfsm, .*, arista_eos, sh[[ow]] i[[p]] ar[[p]]
arista_eos_show_ip_bgp.textfsm, .*, arista_eos, sh[[ow]] i[[p]] bg[[p]]
arista_eos_show_module.textfsm, .*, arista_eos, sh[[ow]] modu[[le]]
arista_eos_bash_df_-h.textfsm, .*, arista_eos, bas[[h]] d[[f]] [[-h]]
arista_eos_show_clock.textfsm, .*, arista_eos, sh[[ow]] clo[[ck]]
arista_eos_show_mlag.textfsm, .*, arista_eos, sh[[ow]] ml[[ag]]
arista_eos_show_vlan.textfsm, .*, arista_eos, sh[[ow]] vl[[an]]
arista_eos_show_vrf.textfsm, .*, arista_eos, sh[[ow]] vrf
arista_eos_dir.textfsm, .*, arista_eos, dir

aruba_aoscx_show_aaa_authentication_port-access_interface_all_client-status.textfsm , .*, aruba_aoscx, sh[[ow]] aa[[a]] authe[[ntication]] port-access interface all client-status
aruba_aoscx_show_lldp_neighbors-info_detail.textfsm, .*, aruba_aoscx, sh[[ow]] ll[[dp]] nei[[ghbor]]s?[[-info]] d[[etail]]
aruba_aoscx_show_bgp_all-vrfs_all_summary.textfsm, .*, aruba_aoscx, sh[[ow]] bgp all-[[vrfs]] a[[ll]] s[[ummary]]
aruba_aoscx_show_interface_dom_detail.textfsm , .*, aruba_aoscx, sh[[ow]] int[[erface]] dom d[[etail]]
aruba_aoscx_show_ip_route_all-vrfs.textfsm, .*, aruba_aoscx, sh[[ow]] ip r[[oute]] a[[ll-vrfs]]
aruba_aoscx_show_mac-address-table.textfsm, .*, aruba_aoscx, sh[[ow]] ma[[c-address-table]]
aruba_aoscx_show_ntp_associations.textfsm, .*, aruba_aoscx, sh[[ow]] ntp as[[sociations]]
aruba_aoscx_show_arp_all-vrfs.textfsm, .*, aruba_aoscx, sh[[ow]] ar[[p]] a[[ll-vrfs]]
aruba_aoscx_show_bfd_all-vrfs.textfsm, .*, aruba_aoscx, sh[[ow]] bf[[d]] a[[ll-vrfs]]
aruba_aoscx_show_vsf_detail.textfsm, .*, aruba_aoscx, sh[[ow]] vsf d[[etail]]
aruba_aoscx_show_interface.textfsm, .*, aruba_aoscx, sh[[ow]] int[[erface]]
aruba_aoscx_show_system.textfsm, .*, aruba_aoscx, sh[[ow]] sys[[tem]]
aruba_aoscx_show_vlan.textfsm, .*, aruba_aoscx, sh[[ow]] vl[[an]]

aruba_os_show_ap_bss-table_details.textfsm, .*, aruba_os, show ap bss-table details
aruba_os_show_ipv6_interface_brief.textfsm, .*, aruba_os, sh[[ow]] ipv6 in[[terface]] b[[rief]]
aruba_os_show_ip_interface_brief.textfsm, .*, aruba_os, sh[[ow]] ip in[[terface]] b[[rief]]
aruba_os_show_ap_radio-database.textfsm, .*, aruba_os, show ap radio-database
aruba_os_show_ap_database_long.textfsm, .*, aruba_os, show ap database long
aruba_os_show_ap_database.textfsm, .*, aruba_os, show ap database
aruba_os_show_arp.textfsm, .*, aruba_os, sh[[ow]] arp

avaya_ers_show_mac-address-table.textfsm, .*, avaya_ers, sh[[ow]] mac-a[[ddress-table]]
avaya_ers_show_mlt_all-members.textfsm, .*, avaya_ers, sh[[ow]] ml[[t]] a[[ll-members]]
avaya_ers_show_interface_name.textfsm, .*, avaya_ers, sh[[ow]] in[[terfaces]] n[[ames]]
avaya_ers_show_logging_config.textfsm, .*, avaya_ers, sh[[ow]] lo[[gging]] co[[nfig]]
avaya_ers_show_sys-info.textfsm, .*, avaya_ers, sh[[ow]] sys-[[info]]
avaya_ers_show_vlan.textfsm, .*, avaya_ers, sh[[ow]] vlan
avaya_ers_show_mlt.textfsm, .*, avaya_ers, sh[[ow]] ml[[t]]

avaya_vsp_show_software.textfsm, .*, avaya_vsp, sho[[w]] so[[ftware]]

broadcom_icos_show_lldp_remote-device_all.textfsm, .*, broadcom_icos, sh[[ow]] lld[[p]] r[[emote-device]] a[[ll]]
broadcom_icos_show_isdp_neighbors.textfsm, .*, broadcom_icos, sh[[ow]] is[[dp]] n[[eighbors]]
broadcom_icos_show_mac-addr-table.textfsm, .*, broadcom_icos, sh[[ow]] mac-addr-[[table]]
broadcom_icos_show_vlan_brief.textfsm, .*, broadcom_icos, sh[[ow]] vl[[an]] b[[rief]]
broadcom_icos_show_version.textfsm, .*, broadcom_icos, sh[[ow]] ver[[sion]]

brocade_fastiron_show_lldp_neighbors_detail.textfsm, .*, brocade_fastiron, sh[[ow]] ll[[dp]] n[[eighbors]] d[[etail]]
brocade_fastiron_show_running-config_vlan.textfsm, .*, brocade_fastiron, sh[[ow]] ru[[nning-config]] v[[lan]]
brocade_fastiron_show_interfaces_brief.textfsm, .*, brocade_fastiron, sh[[ow]] in[[terfaces]] b[[rief]]
brocade_fastiron_show_lldp_neighbors.textfsm, .*, brocade_fastiron, sh[[ow]] ll[[dp]] n[[eighbors]]
brocade_fastiron_show_mac-address.textfsm, .*, brocade_fastiron, sh[[ow]] ma[[c-address]]
brocade_fastiron_show_interfaces.textfsm, .*, brocade_fastiron, sh[[ow]] in[[terfaces]]
brocade_fastiron_show_lag_brief.textfsm, .*, brocade_fastiron, sh[[ow]] la[[g]] b[[rief]]
brocade_fastiron_show_monitor.textfsm, .*, brocade_fastiron, sh[[ow]] mo[[nitor]]
brocade_fastiron_show_version.textfsm, .*, brocade_fastiron, sh[[ow]] ve[[rsion]]
brocade_fastiron_show_metro.textfsm, .*, brocade_fastiron, sh[[ow]] met[[ro-ring]]
brocade_fastiron_show_trunk.textfsm, .*, brocade_fastiron, sh[[ow]] tru[[nk]]
brocade_fastiron_show_span.textfsm, .*, brocade_fastiron, sh[[ow]] sp[[an]]
brocade_fastiron_show_topo.textfsm, .*, brocade_fastiron, sh[[ow]] to[[pology-group]]
brocade_fastiron_show_arp.textfsm, .*, brocade_fastiron, sh[[ow]] a[[rp]]

brocade_netiron_show_running-config_interface.textfsm, .*, brocade_netiron, sh[[ow]] ru[[nning-config]] i[[nterface]]
brocade_netiron_show_lldp_neighbors_detail.textfsm, .*, brocade_netiron, sh[[ow]] ll[[dp]] n[[eighbors]] d[[etail]]
brocade_netiron_show_running-config_vlan.textfsm, .*, brocade_netiron, sh[[ow]] ru[[nning-config]] v[[lan]]
brocade_netiron_show_interfaces_brief.textfsm, .*, brocade_netiron, sh[[ow]] in[[terfaces]] b[[rief]]
brocade_netiron_show_monitor_actual.textfsm, .*, brocade_netiron, sh[[ow]] mon[[itor]] (?:ac|co)
brocade_netiron_show_interfaces.textfsm, .*, brocade_netiron, sh[[ow]] in[[terfaces]]
brocade_netiron_show_lag_brief.textfsm, .*, brocade_netiron, sh[[ow]] lag b[[rief]]
brocade_netiron_show_metro.textfsm, .*, brocade_netiron, sh[[ow]] met[[ro-ring]]
brocade_netiron_show_span.textfsm, .*, brocade_netiron, sh[[ow]] sp[[anning-tree]]
brocade_netiron_show_topo.textfsm, .*, brocade_netiron, sh[[ow]] to[[pology-group]]

checkpoint_gaia_show_virtual-system_all.textfsm, .*, checkpoint_gaia, show virtual-system all
checkpoint_gaia_show_arp_dynamic_all.textfsm, .*, checkpoint_gaia, show arp dynamic all
checkpoint_gaia_show_interfaces_all.textfsm, .*, checkpoint_gaia, show interfaces all
checkpoint_gaia_show_ntp_servers.textfsm, .*, checkpoint_gaia, show ntp servers
checkpoint_gaia_show_version_all.textfsm, .*, checkpoint_gaia, show version all
checkpoint_gaia_show_domainname.textfsm, .*, checkpoint_gaia, show domainname
checkpoint_gaia_show_ipv6_route.textfsm, .*, checkpoint_gaia, show ipv6 route
checkpoint_gaia_show_asset_all.textfsm, .*, checkpoint_gaia, show asset all
checkpoint_gaia_show_route.textfsm, .*, checkpoint_gaia, show route
checkpoint_gaia_show_dns.textfsm, .*, checkpoint_gaia, show dns
checkpoint_gaia_show_lom.textfsm, .*, checkpoint_gaia, show lom
checkpoint_gaia_fw_stat.textfsm, .*, checkpoint_gaia, fw stat

ciena_saos_traffic-profiling_standard-profile_show.textfsm, .*, ciena_saos, traffic-p[[rofiling]] st[[andard-profile]] sh[[ow]]
ciena_saos_port_show_ethernet-config.textfsm, .*, ciena_saos, po[[rt]] sh[[ow]] ethernet-config
ciena_saos_chassis_show_temperature.textfsm, .*, ciena_saos, ch[[assis]] sh[[ow]] te[[mperature]]
ciena_saos_port_show_network-config.textfsm, .*, ciena_saos, po[[rt]] sh[[ow]] network-config
ciena_saos_lldp_show_neighbors.textfsm, .*, ciena_saos, lldp sh[[ow]] nei[[ghbors]]
ciena_saos_ssh_server_show_key.textfsm, .*, ciena_saos, ssh se[[rver]] sh[[ow]] k[[ey]]
ciena_saos_port_show_status.textfsm, .*, ciena_saos, po[[rt]] sh[[ow]] status
ciena_saos_port_xcvr_show.textfsm, .*, ciena_saos, po[[rt]] xcvr sh[[ow]]
ciena_saos_software_show.textfsm, .*, ciena_saos, so[[ftware]] sh[[ow]]
ciena_saos_port_show.textfsm, .*, ciena_saos, po[[rt]] sh[[ow]]
ciena_saos_rstp_show.textfsm, .*, ciena_saos, rs[[tp]] sh[[ow]]
ciena_saos_vlan_show.textfsm, .*, ciena_saos, vl[[an]] sh[[ow]]

cisco_apic_fabric_show_vlan_extended.textfsm, .*, cisco_apic, fabric sh[[ow]] vlan ex[[tended]]

cisco_asa_show_module.textfsm:cisco_asa_show_module_details.textfsm:cisco_asa_show_module_status.textfsm, .*, cisco_asa, sh[[ow]] modu[[le]]
cisco_asa_show_running-config_all_crypto_map.textfsm, .*, cisco_asa, sh[[ow]] ru[[nning-config]] all cr[[ypto]] m[[ap]]
cisco_asa_show_running-config_object_network.textfsm, .*, cisco_asa, sh[[ow]] ru[[nning-config]] o[[bject]] n[[etwork]]
cisco_asa_show_asp_table_vpn-context_detail.textfsm, .*, cisco_asa, sh[[ow]] asp t[[able]] vpn-co[[ntext]] d[[etail]]
cisco_asa_show_running-config_crypto_ikev1.textfsm, .*, cisco_asa, sh[[ow]] ru[[nning-config]] cr[[ypto]] ikev1
cisco_asa_show_running-config_tunnel-group.textfsm, .*, cisco_asa, sh[[ow]] ru[[nning-config]] tu[[nnel-group]]
cisco_asa_show_running-config_crypto_map.textfsm, .*, cisco_asa, sh[[ow]] ru[[nning-config]] cr[[ypto]] m[[ap]]
cisco_asa_show_vpn-sessiondb_anyconnect.textfsm, .*, cisco_asa, sh[[ow]] vpn-[[sessiondb]] a[[nyconnect]]
cisco_asa_show_vpn-sessiondb_detail_l2l.textfsm, .*, cisco_asa, sh[[ow]] vpn-[[sessiondb]] d[[etail]] l[[2l]]
cisco_asa_show_crypto_ikev1_sa_detail.textfsm, .*, cisco_asa, sh[[ow]] cry[[pto]] ikev1 sa d[[etail]]
cisco_asa_show_object-group_network.textfsm, .*, cisco_asa, sh[[ow]] (?:ru[[nning-config]] object-[[group]]|ob[[ject-group]]) n[[etwork]]
cisco_asa_show_ospf_interface_brief.textfsm, .*, cisco_asa, sh[[ow]] ospf int[[erface]] brief
cisco_asa_show_port-channel_summary.textfsm, .*, cisco_asa, sh[[ow]] po[[rt-channel]] sum[[mary]]
cisco_asa_show_running-config_ipsec.textfsm, .*, cisco_asa, sh[[ow]] ru[[nning-config]] ips[[ec]]
cisco_asa_show_cpu_usage_detailed.textfsm, .*, cisco_asa, sh[[ow]] cpu u[[sage]] d[[etailed]]
cisco_asa_show_interface_ip_brief.textfsm, .*, cisco_asa, sh[[ow]] int[[erface]] ip br[[ief]]
cisco_asa_show_access-list_brief.textfsm, .*, cisco_asa, sh[[ow]] ac[[cess-list]] (\S+ )?br[[ief]]
cisco_asa_show_interface_detail.textfsm, .*, cisco_asa, sh[[ow]] int[[erface]] d[[etail]]
cisco_asa_show_crypto_ipsec_sa.textfsm, .*, cisco_asa, sh[[ow]] (?:cry[[pto]] ip[[sec]]|ipsec) sa
cisco_asa_show_resource_usage.textfsm, .*, cisco_asa, sh[[ow]] res[[ource]] u[[sage]]
cisco_asa_show_ospf_neighbor.textfsm, .*, cisco_asa, sh[[ow]] ospf nei[[ghbor]]
cisco_asa_show_vpn-sessiondb.textfsm, .*, cisco_asa, sh[[ow]] vpn-[[sessiondb]]
cisco_asa_show_access-list.textfsm, .*, cisco_asa, sh[[ow]] ac[[cess-list]]
cisco_asa_show_bgp_summary.textfsm, .*, cisco_asa, sh[[ow]] bg[[p]] s[[ummary]]
cisco_asa_show_license_all.textfsm, .*, cisco_asa, (?:fa[[ilover]]\s+e[[xec]]\s+)?sh[[ow]] lic[[ense]] a[[ll]]
cisco_asa_show_interface.textfsm, .*, cisco_asa, sh[[ow]] int[[erface]]
cisco_asa_show_inventory.textfsm, .*, cisco_asa, sh[[ow]] inven[[tory]]
cisco_asa_show_asp_drop.textfsm, .*, cisco_asa, sh[[ow]] asp d[[rop]]
cisco_asa_show_failover.textfsm, .*, cisco_asa, sh[[ow]] fa[[ilover]]
cisco_asa_show_logging.textfsm, .*, cisco_asa, sh[[ow]] log[[ging]]
cisco_asa_show_version.textfsm, .*, cisco_asa, sh[[ow]] ver[[sion]]
cisco_asa_show_route.textfsm, .*, cisco_asa, sh[[ow]] ro[[ute]]
cisco_asa_show_xlate.textfsm, .*, cisco_asa, sh[[ow]] x[[late]]
cisco_asa_show_name.textfsm, .*, cisco_asa, sh[[ow]] nam[[e]]
cisco_asa_show_arp.textfsm, .*, cisco_(asa|ftd), sh[[ow]] arp
cisco_asa_show_nat.textfsm, .*, cisco_asa, sh[[ow]] nat
cisco_asa_ping.textfsm, .*, cisco_(asa|ftd), ping
cisco_asa_dir.textfsm, .*, cisco_asa, dir

cisco_ftd_show_vpn-sessiondb_anyconnect.textfsm, .*, cisco_ftd, sh[[ow]] vpn-[[sessiondb]] a[[nyconnect]]

cisco_ios_show_module.textfsm:cisco_ios_show_module_status.textfsm:cisco_ios_show_module_submodule.textfsm:cisco_ios_show_module_online_diag.textfsm, .*, cisco_ios, sh[[ow]] mod[[ule]]
cisco_ios_show_switch_detail.textfsm:cisco_ios_show_switch_detail_stack_ports.textfsm, .*, cisco_ios, sh[[ow]] sw[[itch]] d[[etail]]
cisco_ios_show_authentication_sessions_method_details.textfsm, .*, cisco_ios, show authen[[tication]] ses[[sions]] met[[hod]](\s+d[[ot1x]]|\s+m[[ab]])? det[[ails]]
cisco_ios_show_running-config_partition_access-list.textfsm, .*, cisco_ios, sh[[ow]] ru[[nning-config]] p[[artition]] a[[ccess-list]]
cisco_ios_show_ip_bgp_neighbors_advertised-routes.textfsm, .*, cisco_ios, sh[[ow]] ip bgp nei[[ghbors]](\s+\d+\.\d+\.\d+\.\d+)? adv[[ertised-routes]]
cisco_ios_show_running-config_partition_route-map.textfsm, .*, cisco_ios, sh[[ow]] ru[[nning-config]] p[[artition]] route-[[map]]
cisco_ios_show_port-security_interface_interface.textfsm, .*, cisco_ios, sh[[ow]] por[[t-security]] i[[nterface]] (\S+)
cisco_ios_show_wireless_tag_policy_summary.textfsm, .*, cisco_ios, sh[[ow]] wireless tag policy sum[[mary]]
cisco_ios_show_capability_feature_routing.textfsm, .*, cisco_ios, sh[[ow]] cap[[ability]] f[[eature]] r[[outing]]
cisco_ios_show_ip_bgp_vpnv4_all_neighbors.textfsm, .*, cisco_ios, sh[[ow]] ip bgp vpnv4 all nei[[ghbors]]
cisco_ios_show_ip_eigrp_interfaces_detail.textfsm, .*, cisco_ios, sh[[ow]] ip ei[[grp]] i[[nterfaces]] de[[tail]]
cisco_ios_show_ip_eigrp_neighbors_detail.textfsm, .*, cisco_ios, sh[[ow]] ip ei[[grp]] nei[[ghbors]] de[[tail]]
cisco_ios_show_ip_dhcp_snooping_binding.textfsm, .*, cisco_ios, sh[[ow]] ip dhcp sn[[ooping]] bi[[nding]]
cisco_ios_show_ip_ospf_database_network.textfsm, .*, cisco_ios, sh[[ow]] ip ospf data[[base]] ne[[twork]]
cisco_ios_show_authentication_sessions.textfsm, .*, cisco_ios, show authen[[tication]] ses[[sions]]
cisco_ios_show_crypto_pki_certificates.textfsm, .*, cisco_ios, sh[[ow]] cry[[pto]] p[[ki]] ce[[rtificates]]
cisco_ios_show_environment_temperature.textfsm, .*, cisco_ios, sh[[ow]] envi[[ronment]] t[[emperature]]
cisco_ios_show_ip_ospf_database_router.textfsm, .*, cisco_ios, sh[[ow]] ip ospf data[[base]] r[[outer]]
cisco_ios_show_ip_ospf_interface_brief.textfsm, .*, cisco_ios, sh[[ow]] ip ospf int[[erface]]
cisco_ios_show_ip_ospf_neighbor_detail.textfsm, .*, cisco_ios, sh[[ow]] ip ospf nei[[ghbors]] d[[etail]]
cisco_ios_show_processes_memory_sorted.textfsm, .*, cisco_ios, sh[[ow]] pro[[cesses]] mem[[ory]] so[[rted]]
cisco_ios_show_crypto_ipsec_sa_detail.textfsm, .*, cisco_ios, sh[[ow]] cry[[pto]] ip[[sec]] sa d[[etail]]
cisco_ios_show_interfaces_description.textfsm, .*, cisco_ios, sh[[ow]] int[[erfaces]](?: (?:\S+))? des[[cription]]
cisco_ios_show_ip_device_tracking_all.textfsm, .*, cisco_ios, sh[[ow]] ip de[[vice]] t[[racking]] a[[ll]]
cisco_ios_show_bfd_neighbors_details.textfsm, .*, cisco_ios, sh[[ow]] bf[[d]] n[[eighbors]] (?:(?:ipv\d+|client \S+) )?de[[tails]]
cisco_ios_show_crypto_session_detail.textfsm, .*, cisco_ios, sh[[ow]] cry[[pto]] se[[ssion]] d[[etail]]
cisco_ios_show_environment_power_all.textfsm, .*, cisco_ios, sh[[ow]] envi[[ronment]] p[[ower]] a[[ll]]
cisco_ios_show_interface_transceiver.textfsm, .*, cisco_ios, sh[[ow]] int[[erface]] trans[[ceiver]]
cisco_ios_show_interfaces_switchport.textfsm, .*, cisco_ios, sh[[ow]] int[[erfaces]](?: (?:\S+))? sw[[itchport]]
cisco_ios_show_ip_http_server_status.textfsm, .*, cisco_ios, sh[[ow]] ip http ser[[ver]] statu[[s]]
cisco_ios_show_lldp_neighbors_detail.textfsm, .*, cisco_ios, sh[[ow]] lld[[p]] neig[[hbors]] det[[ail]]
cisco_ios_show_cdp_neighbors_detail.textfsm, .*, cisco_ios, sh[[ow]] c[[dp]] neig[[hbors]] det[[ail]]
cisco_ios_show_etherchannel_summary.textfsm, .*, cisco_ios, sh[[ow]] etherchann[[el]] summ[[ary]]
cisco_ios_show_ipv6_interface_brief.textfsm, .*, cisco_ios, sh[[ow]] ipv[[6]] i[[nterface]] b[[rief]]
cisco_ios_show_ip_nat_translations.textfsm, .*, cisco_ios, sh[[ow]] ip nat translation[[s]]
cisco_ios_show_mpls_l2transport_vc.textfsm, .*, cisco_ios, sh[[ow]] m[[pls]] l2[[transport]] v[[c]]
cisco_ios_show_ip_eigrp_neighbors.textfsm, .*, cisco_ios, sh[[ow]] ip ei[[grp]] nei[[ghbors]]
cisco_ios_show_ip_flow_toptalkers.textfsm, .*, cisco_ios, sh[[ow]] ip fl[[ow]] top[[-talkers]]
cisco_ios_show_ip_interface_brief.textfsm, .*, cisco_ios, sh[[ow]] ip int[[erface]] br[[ief]]
cisco_ios_show_interfaces_status.textfsm, .*, cisco_ios, sh[[ow]] int[[erfaces]](?: (?:\S+))? st[[atus]]
cisco_ios_show_ip_eigrp_topology.textfsm, .*, cisco_ios, sh[[ow]] ip eigrp top[[ology]]
cisco_ios_show_ip_source_binding.textfsm, .*, cisco_ios, sh[[ow]] ip sou[[rce]] b[[inding]]
cisco_ios_show_ip_vrf_interfaces.textfsm, .*, cisco_ios, sh[[ow]] ip vr[[f]] in[[terfaces]]
cisco_ios_show_ipv6_access-lists.textfsm, .*, cisco_ios, sh[[ow]] ipv6 acce[[ss-lists]]
cisco_ios_show_mac-address-table.textfsm, .*, cisco_ios, sh[[ow]] mac[[-address-table]]
cisco_ios_show_ap_cdp_neighbors.textfsm, .*, cisco_ios, sh[[ow]] ap c[[dp]] n[[eighbors]]
cisco_ios_show_ip_bgp_neighbors.textfsm, .*, cisco_ios, sh[[ow]] ip bgp nei[[ghbors]]
cisco_ios_show_ip_ospf_database.textfsm, .*, cisco_ios, sh[[ow]] ip ospf data[[base]]
cisco_ios_show_ip_ospf_neighbor.textfsm, .*, cisco_ios, sh[[ow]] ip ospf nei[[ghbor]]
cisco_ios_show_ip_route_summary.textfsm, .*, cisco_ios, sh[[ow]] ip ro[[ute]] sum[[mary]]
cisco_ios_show_ip_access-lists.textfsm, .*, cisco_ios, sh[[ow]] ip acce[[ss-lists]]
cisco_ios_show_ip_dhcp_binding.textfsm, .*, cisco_ios, sh[[ow]] ip dh[[cp]] b[[inding]]
cisco_ios_show_mpls_interfaces.textfsm, .*, cisco_ios, sh[[ow]] mpls interfa[[ces]]
cisco_ios_show_power_available.textfsm,  .*, cisco_ios, sh[[ow]] pow[[er]] a[[vailable]]
cisco_ios_show_access-session.textfsm, .*, cisco_ios, show access-s[[ession]]
cisco_ios_show_alert_counters.textfsm, .*, cisco_ios, sh[[ow]] alert [[counters]]
cisco_ios_show_interface_link.textfsm, .*, cisco_ios, sh[[ow]] int[[erfaces]] li[[nk]]
cisco_ios_show_ip_bgp_summary.textfsm, .*, cisco_ios, sh[[ow]] ip bgp (?:all\s+)?sum[[mary]]
cisco_ios_show_ip_prefix-list.textfsm, .*, cisco_ios, sh[[ow]] ip pre[[fix-list]]
cisco_ios_show_ipv6_neighbors.textfsm, .*, cisco_ios, sh[[ow]] ipv[[6]] ne[[ighbors]]
cisco_ios_show_isis_neighbors.textfsm, .*, cisco_ios, sh[[ow]] isis ne[[ighbors]]
cisco_ios_show_license_status.textfsm, .*, cisco_ios, sh[[ow]] lic[[ense]] st[[atus]]
cisco_ios_show_lldp_neighbors.textfsm, .*, cisco_ios, sh[[ow]] lld[[p]] neig[[hbors]]
cisco_ios_show_power_supplies.textfsm,  .*, cisco_ios, sh[[ow]] pow[[er]] su[[pplies]]
cisco_ios_show_snmp_community.textfsm, .*, cisco_ios, sh[[ow]] sn[[mp]] com[[munity]]
cisco_ios_show_cdp_neighbors.textfsm, .*, cisco_ios, sh[[ow]] c[[dp]] neig[[hbors]]
cisco_ios_show_controller_t1.textfsm, .*, cisco_ios, sh[[ow]] cont[[rollers]] t1
cisco_ios_show_hosts_summary.textfsm, .*, cisco_ios, sh[[ow]] ho[[sts]] summary
cisco_ios_show_ip_cef_detail.textfsm, .*, cisco_ios, sh[[ow]] ip ce[[f]].+?d[[etail]]
cisco_ios_show_platform_diag.textfsm, .*, cisco_ios, sh[[ow]] plat[[form]] di[[ag]]
cisco_ios_show_processes_cpu.textfsm, .*, cisco_ios, sh[[ow]] proc[[esses]] [[cpu]]
cisco_ios_show_spanning-tree.textfsm, .*, cisco_ios, sh[[ow]] sp[[anning-tree]]
cisco_ios_show_standby_brief.textfsm, .*, cisco_ios, sh[[ow]] standby(?:\s+\S+)? br[[ief]]
cisco_ios_show_file_systems.textfsm, .*, cisco_ios, sh[[ow]] fi[[le]] s[[ystems]]
cisco_ios_show_ip_interface.textfsm, .*, cisco_ios, sh[[ow]] ip int[[erface]]
cisco_ios_show_object-group.textfsm, .*, cisco_ios, sh[[ow]] ob[[ject-group]]
cisco_ios_show_power_status.textfsm, .*, cisco_ios, sh[[ow]] pow[[er]] st[[atus]]
cisco_ios_show_rep_topology.textfsm, .*, cisco_ios, sh[[ow]] rep topology
cisco_ios_show_access-list.textfsm, .*, cisco_ios, sh[[ow]] acc[[ess-list]]
cisco_ios_show_isdn_status.textfsm, .*, cisco_ios, sh[[ow]] isd[[n]] st[[atus]]
cisco_ios_show_ap_summary.textfsm, .*, cisco_ios, sh[[ow]] ap sum[[mary]]
cisco_ios_show_dhcp_lease.textfsm, .*, cisco_ios, sh[[ow]] dh[[cp]] l[[ease]]
cisco_ios_show_interfaces.textfsm, .*, cisco_ios, sh[[ow]] int[[erfaces]](?: (?:\S+))?
cisco_ios_show_ipv6_route.textfsm, .*, cisco_ios, sh[[ow]] ipv6 r[[oute]]
cisco_ios_show_redundancy.textfsm, .*, cisco_ios, sh[[ow]] redu[[ndancy]]
cisco_ios_show_snmp_group.textfsm, .*, cisco_ios, sh[[ow]] snm[[p]] g[[roup]]
cisco_ios_show_vrf_detail.textfsm, .*, cisco_ios, sh[[ow]] vrf detail
cisco_ios_show_vrrp_brief.textfsm, .*, cisco_ios, sh[[ow]] vrr[[p]] b[[rief]]
cisco_ios_show_vtp_status.textfsm, .*, cisco_ios, sh[[ow]] vtp stat[[us]]
cisco_ios_show_adjacency.textfsm, .*, cisco_ios, sh[[ow]] ad[[jacency]]
cisco_ios_show_dot1x_all.textfsm, .*, cisco_ios, sh[[ow]] dot1x a[[ll]]
cisco_ios_show_inventory.textfsm, .*, cisco_ios, sh[[ow]] inven[[tory]]
cisco_ios_show_ip_mroute.textfsm, .*, cisco_ios, sh[[ow]] ip mr[[oute]]
cisco_ios_show_nve_peers.textfsm, .*, cisco_ios, sh[[ow]] nv[[e]] p[[eers]]
cisco_ios_show_route-map.textfsm, .*, cisco_ios, sh[[ow]] route-m[[ap]]
cisco_ios_show_snmp_user.textfsm, .*, cisco_ios, sh[[ow]] sn[[mp]] u[[ser]]
cisco_ios_show_ip_route.textfsm, .*, cisco_ios, sh[[ow]] ip r[[oute]]
cisco_ios_show_vrrp_all.textfsm, .*, cisco_ios, sh[[ow]] vrr[[p]] a[[ll]]
cisco_ios_show_aliases.textfsm,  .*, cisco_ios, sh[[ow]] alia[[ses]]
cisco_ios_show_archive.textfsm,  .*, cisco_ios, sh[[ow]] arc[[hive]]
cisco_ios_show_license.textfsm, .*, cisco_ios, sh[[ow]] lic[[ense]]
cisco_ios_show_logging.textfsm, .*, cisco_ios, sh[[ow]] log[[ging]]
cisco_ios_show_nve_vni.textfsm, .*, cisco_ios, sh[[ow]] nv[[e]] v[[ni]]
cisco_ios_show_standby.textfsm, .*, cisco_ios, sh[[ow]] sta[[ndby]]
cisco_ios_show_version.textfsm, .*, cisco_ios, sh[[ow]] ver[[sion]]
cisco_ios_show_ip_arp.textfsm, .*, cisco_ios, sh[[ow]] i[[p]] a[[rp]]
cisco_ios_show_ip_bgp.textfsm, .*, cisco_ios, sh[[ow]] i[[p]] bgp
cisco_ios_show_ip_cef.textfsm, .*, cisco_ios, sh[[ow]] ip ce[[f]](?: vrf? \S+)?\s*$
cisco_ios_show_tacacs.textfsm, .*, cisco_ios, sh[[ow]] tacacs
cisco_ios_show_clock.textfsm, .*, cisco_ios, sh[[ow]] clo[[ck]]
cisco_ios_show_dmvpn.textfsm, .*, cisco_ios, sh[[ow]] dm[[vpn]]
cisco_ios_show_users.textfsm, .*, cisco_ios, sh[[ow]] users(?: (?:a[[ll]]))?
cisco_ios_show_vlans.textfsm, .*, cisco_ios, sh[[ow]] vlans
cisco_ios_traceroute.textfsm, .*, cisco_ios, tr[[aceroute]]
cisco_ios_show_boot.textfsm, .*, cisco_ios, sh[[ow]] boot
cisco_ios_show_vlan.textfsm, .*, cisco_ios, sh[[ow]] vlan
cisco_ios_show_arp.textfsm, .*, cisco_ios, sh[[ow]] arp
cisco_ios_show_vrf.textfsm, .*, cisco_ios, sh[[ow]] vrf
cisco_ios_ping.textfsm, .*, cisco_ios, ping
cisco_ios_dir.textfsm,  .*, cisco_ios, dir

cisco_nvfis_show_running-config_system_settings.textfsm, .*, cisco_nvfis, show running-config system settings
cisco_nvfis_show_running-config_system_upgrade.textfsm, .*, cisco_nvfis, show running-config system upgrade
cisco_nvfis_show_configuration_commit_changes.textfsm, .*, cisco_nvfis, show configuration commit changes
cisco_nvfis_show_running-config_snmp_enable.textfsm, .*, cisco_nvfis, show running-config snmp enable
cisco_nvfis_show_running-config_system_time.textfsm, .*, cisco_nvfis, show running-config system time
cisco_nvfis_show_running-config_snmp_group.textfsm, .*, cisco_nvfis, show running-config snmp group
cisco_nvfis_show_running-config_snmp_host.textfsm, .*, cisco_nvfis, show running-config snmp host
cisco_nvfis_show_running-config_snmp_user.textfsm, .*, cisco_nvfis, show running-config snmp user
cisco_nvfis_show_bridge-settings.textfsm, .*, cisco_nvfis, show bridge-settings
cisco_nvfis_show_version.textfsm, .*, cisco_nvfis, show version
cisco_nvfis_show_nic.textfsm, .*, cisco_nvfis, show nic

cisco_nxos_show_bgp_vrf_all_ipv4_unicast_neighbors_routes.textfsm, .*, cisco_nxos, sh[[ow]] bg[[p]] v[[rf]] (\S+) ipv4 uni[[cast]] neigh[[bors]](\s+\S+)? (ro(utes)?|a(dvertised-routes)?|re(ceived-routes)?)
cisco_nxos_show_hardware_internal_bigsur_all-ports_detail.textfsm, .*, cisco_nxos, sh[[ow]] ha[[rdware]] i[[nternal]] bi[[gsur]] al[[l-ports]] d[[etail]]
cisco_nxos_show_l2rib_internal_permanently-frozen-list.textfsm, .*, cisco_nxos, sh[[ow]] l2ri[[b]] i[[nternal]] pe[[rmanently-frozen-list]]
cisco_nxos_show_bgp_vrf_all_ipv4_unicast_detail.textfsm, .*, cisco_nxos, sh[[ow]] bg[[p]] v[[rf]] (\S+) ipv4 uni[[cast]] d[[etail]]
cisco_nxos_show_ip_pim_interface_brief_vrf_all.textfsm, .*, cisco_nxos, sh[[ow]] ip pim int[[erface]] br[[ief]] vrf all
cisco_nxos_show_nve_vni_interface_nve_1_detail.textfsm, .*, cisco_nxos, sh[[ow]] nv[[e]] vn[[i]] int[[erface]] nv[[e]] 1 det[[ail]]
cisco_nxos_show_configuration_session_summary.textfsm, .*, cisco_nxos, sh[[ow]] configu[[ration]] s[[ession]] su[[mmary]]
cisco_nxos_show_interface_transceiver_details.textfsm, .*, cisco_nxos, sh[[ow]] int[[erface]] tra[[nsceiver]] de[[tails]]
cisco_nxos_show_ip_dhcp_snooping_statistics.textfsm, .*, cisco_nxos, sh[[ow]] ip dh[[cp]] sn[[ooping]] st[[atistics]]
cisco_nxos_show_ip_pim_group-range_vrf_all.textfsm, .*, cisco_nxos, sh[[ow]] ip pim group[[-range]] vrf all
cisco_nxos_show_environment_temperature.textfsm, .*, cisco_nxos, sh[[ow]] env[[ironment]] t[[emperature]]
cisco_nxos_show_ip_msdp_summary_vrf_all.textfsm, .*, cisco_nxos, sh[[ow]] ip msdp summary vrf all
cisco_nxos_show_ip_ospf_interface_brief.textfsm, .*, cisco_nxos, sh[[ow]] ip ospf int[[erface]] b[[rief]](?: vrf \S+)?\s*$
cisco_nxos_show_ip_pim_neighbor_vrf_all.textfsm, .*, cisco_nxos, sh[[ow]] ip pim nei[[ghbor]] vrf all
cisco_nxos_show_interface_snmp-ifindex.textfsm, .*, cisco_nxos, sh[[ow]] int[[erface]] sn[[mp-ifindex]]
cisco_nxos_show_forwarding_ipv4_route.textfsm, .*, cisco_nxos, sh[[ow]] fo[[rwarding]] ipv4 ro[[ute]]
cisco_nxos_show_interface_description.textfsm, .*, cisco_nxos, sh[[ow]] int[[erface]] desc[[ription]]
cisco_nxos_show_interface_transceiver.textfsm, .*, cisco_nxos, sh[[ow]] int[[erface]] tra[[nsceiver]]
cisco_nxos_show_interfaces_switchport.textfsm, .*, cisco_nxos, sh[[ow]] int[[erfaces]] sw[[itchport]]
cisco_nxos_show_ip_arp_detail_vrf_all.textfsm, .*, cisco_nxos, sh[[ow]] i[[p]] a[[rp]] det[[ail]] vrf all
cisco_nxos_show_ip_dhcp_relay_address.textfsm, .*, cisco_nxos, sh[[ow]] ip dh[[cp]] r[[elay]] a[[ddress]]
cisco_nxos_show_lldp_neighbors_detail.textfsm, .*, cisco_nxos, sh[[ow]] ll[[dp]] nei[[ghbors]] d[[etail]]
cisco_nxos_show_cdp_neighbors_detail.textfsm, .*, cisco_nxos, sh[[ow]] c[[dp]] neig[[hbors]] det[[ail]]
cisco_nxos_show_forwarding_adjacency.textfsm, .*, cisco_nxos, sh[[ow]] fo[[rwarding]] (?:ipv4 )?ad[[jacency]]
cisco_nxos_show_ip_interface_vrf_all.textfsm, .*, cisco_nxos, sh[[ow]] ip interf[[ace]] v[[rf]] a[[ll]]
cisco_nxos_show_ipv6_interface_brief.textfsm, .*, cisco_nxos, sh[[ow]] ipv[[6]] interf[[ace]] b[[rief]]
cisco_nxos_show_port-channel_summary.textfsm, .*, cisco_nxos, sh[[ow]] po[[rt-channel]] sum[[mary]]
cisco_nxos_show_cts_interface_brief.textfsm, .*, cisco_nxos, sh[[ow]] cts inte[[rface]] br[[ief]]
cisco_nxos_show_ip_bgp_summary_vrf.textfsm, .*, cisco_nxos, sh[[ow]] ip b[[gp]] (?:all\s+)?s[[ummary]] v[[rf]]
cisco_nxos_show_ip_eigrp_neighbors.textfsm, .*, cisco_nxos, sh[[ow]] ip ei[[grp]] nei[[ghbors]]
cisco_nxos_show_ip_interface_brief.textfsm, .*, cisco_nxos, sh[[ow]] ip int[[erface]] b[[rief]](?: vrf \S+)?\s*$
cisco_nxos_show_ip_mroutes_vrf_all.textfsm, .*, cisco_nxos, sh[[ow]] ip mr[[outes]] vrf al[[l]]
cisco_nxos_show_spanning-tree_root.textfsm, .*, cisco_nxos, sh[[ow]] sp[[an]]ning-tree root
cisco_nxos_show_cts_interface_all.textfsm, .*, cisco_nxos, sh[[ow]] ct[[s]] inter[[face]] al[[l]]
cisco_nxos_show_ip_community-list.textfsm, .*, cisco_nxos, sh[[ow]] ip comm[[unity-list]]
cisco_nxos_show_ip_pim_rp_vrf_all.textfsm, .*, cisco_nxos, sh[[ow]] ip pim rp vrf all
cisco_nxos_show_mac_address-table.textfsm, .*, cisco_nxos, sh[[ow]] m[[ac]] addr[[ess-table]]
cisco_nxos_show_fabricpath_route.textfsm, .*, cisco_nxos, sh[[ow]] fabricpath route
cisco_nxos_show_interface_status.textfsm, .*, cisco_nxos, sh[[ow]] int[[erface]] st[[atus]]
cisco_nxos_show_ip_bgp_neighbors.textfsm, .*, cisco_nxos, sh[[ow]] ip bgp nei[[ghbors]]
cisco_nxos_show_ip_ospf_database.textfsm, .*, cisco_nxos, sh[[ow]] ip o[[spf]] d[[atabase]]
cisco_nxos_show_ip_ospf_neighbor.textfsm, .*, cisco_nxos, sh[[ow]] ip ospf nei[[ghbor]]
cisco_nxos_show_interface_brief.textfsm, .*, cisco_nxos, sh[[ow]] int[[erface]] br[[ief]]
cisco_nxos_show_flogi_database.textfsm, .*, cisco_nxos, sh[[ow]] fl[[ogi]] d[[atabase]]
cisco_nxos_show_ip_bgp_summary.textfsm, .*, cisco_nxos, sh[[ow]] ip b[[gp]] s[[ummary]]
cisco_nxos_show_lldp_neighbors.textfsm, .*, cisco_nxos, sh[[ow]] ll[[dp]] nei[[ghbors]]
cisco_nxos_show_switching-mode.textfsm, .*, cisco_nxos, sh[[ow]] switchi[[ng-mode]]
cisco_nxos_show_bfd_neighbors.textfsm, .*, cisco_nxos, sh[[ow]] bf[[d]] neig[[hbors]]
cisco_nxos_show_cdp_neighbors.textfsm, .*, cisco_nxos, sh[[ow]] cd[[p]] neig[[hbors]]
cisco_nxos_show_ip_arp_detail.textfsm, .*, cisco_nxos, sh[[ow]] i[[p]] a[[rp]] det[[ail]]
cisco_nxos_show_license_usage.textfsm, .*, cisco_nxos, sh[[ow]] lic[[ense]] us[[age]]
cisco_nxos_show_processes_cpu.textfsm, .*, cisco_nxos, sh[[ow]] proc[[esses]] c[[pu]]
cisco_nxos_show_vrf_interface.textfsm, .*, cisco_nxos, sh[[ow]] vrf int[[erface]]
cisco_nxos_show_access-lists.textfsm, .*, cisco_nxos, sh[[ow]] acc[[ess-lists]]
cisco_nxos_show_ip_adjacency.textfsm, .*, cisco_nxos, sh[[ow]] ip ad[[jacency]]
cisco_nxos_show_ip_interface.textfsm, .*, cisco_nxos, sh[[ow]] ip int[[erface]]
cisco_nxos_show_environment.textfsm, .*, cisco_nxos, sh[[ow]] env[[ironment]]
cisco_nxos_show_vrf_detail.textfsm, .*, cisco_nxos, sh[[ow]] vrf det[[ail]]
cisco_nxos_show_interface.textfsm, .*, cisco_nxos, sh[[ow]] inte[[rface]]
cisco_nxos_show_inventory.textfsm, .*, cisco_nxos, sh[[ow]] inv[[entory]]
cisco_nxos_show_nve_peers.textfsm, .*, cisco_nxos, sh[[ow]] nv[[e]] p[[eers]]
cisco_nxos_show_route-map.textfsm, .*, cisco_nxos, sh[[ow]] route-m[[ap]]
cisco_nxos_show_hostname.textfsm, .*, cisco_nxos, sh[[ow]] hostn[[ame]]
cisco_nxos_show_hsrp_all.textfsm, .*, cisco_nxos, sh[[ow]] hsrp all
cisco_nxos_show_ip_route.textfsm, .*, cisco_nxos, sh[[ow]] ip route
cisco_nxos_show_feature.textfsm, .*, cisco_nxos, sh[[ow]] feat[[ure]]
cisco_nxos_show_nve_vni.textfsm, .*, cisco_nxos, sh[[ow]] nv[[e]] vn[[i]]
cisco_nxos_show_version.textfsm, .*, cisco_nxos, sh[[ow]] ver[[sion]]
cisco_nxos_show_fex_id.textfsm, .*, cisco_nxos, sh[[ow]] fex (\S+)
cisco_nxos_show_ip_arp.textfsm, .*, cisco_nxos, sh[[ow]] i[[p]] a[[rp]]
cisco_nxos_show_ip_bgp.textfsm, .*, cisco_nxos, sh[[ow]] i[[p]] bgp(?: vrf \S+)?
cisco_nxos_show_module.textfsm, .*, cisco_nxos, sh[[ow]] mod[[ule]]
cisco_nxos_show_clock.textfsm, .*, cisco_nxos, sh[[ow]] clo[[ck]]
cisco_nxos_show_vlan.textfsm, .*, cisco_nxos, sh[[ow]] vl[[an]]
cisco_nxos_show_fex.textfsm, .*, cisco_nxos, sh[[ow]] fex
cisco_nxos_show_vdc.textfsm, .*, cisco_nxos, sh[[ow]] vdc
cisco_nxos_show_vpc.textfsm, .*, cisco_nxos, sh[[ow]] vpc
cisco_nxos_show_vrf.textfsm, .*, cisco_nxos, sh[[ow]] vrf
cisco_nxos_dir.textfsm,  .*, cisco_nxos, dir

cisco_s300_show_interfaces_description.textfsm, .*, cisco_s300, sh[[ow]] int[[erfaces]] description
cisco_s300_show_interfaces_switchport.textfsm, .*, cisco_s300, sh[[ow]] int[[erfaces]] switchport
cisco_s300_show_interfaces_status.textfsm, .*, cisco_s300, sh[[ow]] int[[erfaces]] st[[atus]]
cisco_s300_show_mac_address-table.textfsm, .*, cisco_s300, sh[[ow]] mac address-[[table]]
cisco_s300_show_lldp_neighbors.textfsm, .*, cisco_s300, sh[[ow]] lld[[p]] neig[[hbors]]
cisco_s300_show_ip_interface.textfsm, .*, cisco_s300, sh[[ow]] ip interface
cisco_s300_show_system_id.textfsm, .*, cisco_s300, sh[[ow]] system id
cisco_s300_show_version.textfsm, .*, cisco_s300, sh[[ow]] ver[[sion]]
cisco_s300_show_system.textfsm, .*, cisco_s300, sh[[ow]] system
cisco_s300_show_vlan.textfsm, .*, cisco_s300, sh[[ow]] vlan

cisco_viptela_show_interface.textfsm, .*, cisco_viptela, sh[[ow]] in[[terface]]
cisco_viptela_show_arp.textfsm, .*, cisco_viptela, sh[[ow]] ar[[p]]

cisco_wlc_ssh_show_flexconnect_group_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] fl[[exconnect]] gr[[oup]] s[[ummary]]
cisco_wlc_ssh_show_advanced_802.11a_channel.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ad[[vanced]] 802\.11[ab] ch[[annel]]
cisco_wlc_ssh_show_802.11a_cleanair_config.textfsm, .*, cisco_wlc_ssh, sh[[ow]] 802\.11[ab] cl[[eanair]] c[[onfig]]
cisco_wlc_ssh_show_interface_group_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] int[[erface]] gr[[oup]] s[[ummary]]
cisco_wlc_ssh_show_interface_detailed_id.textfsm, .*, cisco_wlc_ssh, sh[[ow]] int[[erface]] d[[etailed]] (\S+)
cisco_wlc_ssh_show_cdp_neighbors_detail.textfsm, .*, cisco_wlc_ssh, sh[[ow]] c[[dp]] neig[[hbors]] det[[ail]]
cisco_wlc_ssh_show_redundancy_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] red[[undancy]] su[[mmary]]
cisco_wlc_ssh_show_rf-profile_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] rf-[[profile]] s[[ummary]]
cisco_wlc_ssh_show_stats_port_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] st[[ats]] p[[ort]] s[[ummary]]
cisco_wlc_ssh_show_ap_config_general.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ap con[[fig]] ge[[neral]]
cisco_wlc_ssh_show_interface_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] int[[erface]] s[[ummary]]
cisco_wlc_ssh_show_redundancy_detail.textfsm, .*, cisco_wlc_ssh, sh[[ow]] red[[undancy]] d[[etail]]
cisco_wlc_ssh_show_mobility_anchor.textfsm, .*, cisco_wlc_ssh, sh[[ow]] mo[[bility]] an[[chor]]
cisco_wlc_ssh_show_radius_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ra[[dius]] s[[ummary]]
cisco_wlc_ssh_show_tacacs_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ta[[cacs]] s[[ummary]]
cisco_wlc_ssh_show_client_detail.textfsm, .*, cisco_wlc_ssh, sh[[ow]] cl[[ient]] det[[ail]]
cisco_wlc_ssh_show_exclusionlist.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ex[[clusionlist]]
cisco_wlc_ssh_show_ap_image_all.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ap ima[[ge]] a[[ll]]
cisco_wlc_ssh_show_mobility_sum.textfsm, .*, cisco_wlc_ssh, sh[[ow]] mo[[bility]] su[[mmary]]
cisco_wlc_ssh_show_port_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] por[[t]] s[[ummary]]
cisco_wlc_ssh_show_band-select.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ba[[nd-select]]
cisco_wlc_ssh_show_ap_summary.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ap sum[[mary]]
cisco_wlc_ssh_show_inventory.textfsm, .*, cisco_wlc_ssh, sh[[ow]] inve[[ntory]]
cisco_wlc_ssh_show_wlan_sum.textfsm, .*, cisco_wlc_ssh, sh[[ow]] wl[[an]] s[[ummary]]
cisco_wlc_ssh_show_802.11a.textfsm, .*, cisco_wlc_ssh, show 802\.11[ab]
cisco_wlc_ssh_show_sysinfo.textfsm, .*, cisco_wlc_ssh, sh[[ow]] sysi[[nfo]]
cisco_wlc_ssh_show_boot.textfsm, .*, cisco_wlc_ssh, sh[[ow]] bo[[ot]]
cisco_wlc_ssh_show_time.textfsm, .*, cisco_wlc_ssh, sh[[ow]] ti[[me]]

cisco_xr_show_controllers_fabric_fia_errors_ingress_location.textfsm, .*, cisco_xr, sh[[ow]] contr[[ollers]] fabric fi[[a]] err[[ors]] in[[gress]] loc[[ation]]
cisco_xr_show_controllers_fabric_fia_drops_ingress_location.textfsm, .*, cisco_xr, sh[[ow]] contr[[ollers]] fabric fi[[a]] dr[[ops]] in[[gress]] loc[[ation]]
cisco_xr_show_controllers_fabric_fia_errors_egress_location.textfsm, .*, cisco_xr, sh[[ow]] contr[[ollers]] fabric fi[[a]] err[[ors]] eg[[ress]] loc[[ation]]
cisco_xr_show_controllers_fabric_fia_drops_egress_location.textfsm, .*, cisco_xr, sh[[ow]] contr[[ollers]] fabric fi[[a]] dr[[ops]] eg[[ress]] loc[[ation]]
cisco_xr_show_l2vpn_mac-learning_mac_all_location.textfsm, .*, cisco_xr, sh[[ow]] l2vpn mac-l[[earning]] mac all loc[[ation]]
cisco_xr_show_lpts_pifib_hardware_police_location.textfsm, .*, cisco_xr, sh[[ow]] lpts pifib hardware police loc[[ation]]
cisco_xr_show_controllers_HundredGigabitEthernet.textfsm, .*, cisco_xr, sh[[ow]] contr[[ollers]] Hu[[ndredGigabitEthernet]]
cisco_xr_show_bgp_vrf_all_ipv4_unicast_summary.textfsm, .*, cisco_xr, sh[[ow]] bg[[p]] v[[rf]] all ip[[v4]] uni[[cast]] summ[[ary]]
cisco_xr_admin_show_controller_fabric_health.textfsm, .*, cisco_xr, adm[[in]] sh[[ow]] controller fab[[ric]] hea[[lth]]
cisco_xr_show_platform_summary_location_all.textfsm, .*, cisco_xr, sh[[ow]] plat[[form]] sum[[mary]] loc[[ation]] (\S+)
cisco_xr_show_ipv4_vrf_all_interface_brief.textfsm, .*, cisco_xr, sh[[ow]] ipv4 vrf (\S+) int[[erface]] br[[ief]]
cisco_xr_show_ospf_vrf_all_interface_brief.textfsm, .*, cisco_xr, sh[[ow]] ospf vrf (\S+) int[[erface]] brief
cisco_xr_show_controller_fabric_plane_all.textfsm, .*, cisco_xr, sh[[ow]] controller fab[[ric]] pla[[ne]] all
cisco_xr_show_controllers_hundredgige_all.textfsm, .*, cisco_xr, sh[[ow]] controllers hund[[redgige]] all
cisco_xr_show_configuration_commit_list.textfsm, .*, cisco_xr, sh[[ow]] conf[[iguration]] c[[ommit]] l[[ist]]
cisco_xr_show_asic-errors_all_location.textfsm, .*, cisco_xr, sh[[ow]] asic-er[[rors]] all loc[[ation]]
cisco_xr_show_bgp_instance_all_summary.textfsm, .*, cisco_xr, sh[[ow]] bg[[p]] ins[[tance]] (\S+) sum[[mary]]
cisco_xr_admin_show_environment_power.textfsm, .*, cisco_xr, adm[[in]] sh[[ow]] env[[ironment]] p[[ower]]
cisco_xr_show_dhcp_ipv4_proxy_binding.textfsm, .*, cisco_xr, sh[[ow]] dh[[cp]] ipv4 p[[roxy]] b[[inding]]
cisco_xr_show_mpls_ldp_neighbor_brief.textfsm, .*, cisco_xr, sh[[ow]] mp[[ls]] ld[[p]] neigh[[bor]] br[[ief]]
cisco_xr_show_running-config_hostname.textfsm, .*, cisco_xr, sh[[ow]] run[[ning-config]] host[[name]] 
cisco_xr_show_interfaces_description.textfsm, .*, cisco_xr, sh[[ow]] int[[erfaces]](?: (?:\S+))? des[[cription]]
cisco_xr_admin_show_environment_fan.textfsm, .*, cisco_xr, adm[[in]] sh[[ow]] env[[ironment]] f[[an]]
cisco_xr_show_lldp_neighbors_detail.textfsm, .*, cisco_xr, sh[[ow]] ll[[dp]] ne[[ighbors]] de[[tail]]
cisco_xr_show_ospf_vrf_all_neighbor.textfsm, .*, cisco_xr, sh[[ow]] ospf vrf (\S+) nei[[ghbor]]
cisco_xr_show_cdp_neighbors_detail.textfsm, .*, cisco_xr, sh[[ow]] c[[dp]] neig[[hbors]] det[[ail]]
cisco_xr_show_controllers_all_phy.textfsm, .*, cisco_xr, sh[[ow]] contr[[ollers]] (\S+) ph[[y]]
cisco_xr_show_cef_drops_location.textfsm, .*, cisco_xr, sh[[ow]] cef drops loc[[ation]]
cisco_xr_show_interfaces_summary.textfsm, .*, cisco_xr, sh[[ow]] int[[erfaces]] summ[[ary]]
cisco_xr_show_ip_interface_brief.textfsm, .*, cisco_xr, sh[[ow]] ip int[[erface]] br[[ief]]
cisco_xr_show_pim_ipv4_group-map.textfsm, .*, cisco_xr, sh[[ow]] pim ipv4 group-map
cisco_xr_show_pim_ipv4_interface.textfsm, .*, cisco_xr, sh[[ow]] pim ipv4 int[[erface]]
cisco_xr_show_redundancy_summary.textfsm, .*, cisco_xr, sh[[ow]] redun[[dancy]] summ[[ary]]
cisco_xr_show_pim_ipv4_neighbor.textfsm, .*, cisco_xr, sh[[ow]] pim ipv4 nei[[ghbor]]
cisco_xr_admin_show_inventory.textfsm, .*, cisco_xr, (adm[[in]] )?sh[[ow]] inven[[tory]]
cisco_xr_show_interface_brief.textfsm, .*, cisco_xr, sh[[ow]] int[[erface]] br[[ief]]
cisco_xr_admin_show_platform.textfsm, .*, cisco_xr, (adm[[in]] )?sh[[ow]] pla[[tform]]
cisco_xr_show_install_active.textfsm, .*, cisco_xr, sh[[ow]] install active
cisco_xr_show_ip_bgp_summary.textfsm, .*, cisco_xr, sh[[ow]] ip b[[gp]] s[[ummary]]
cisco_xr_show_ipv4_interface.textfsm, .*, cisco_xr, sh[[ow]] ip[[v4]](?: vrf \S+)? int[[erface]]
cisco_xr_show_ipv6_neighbors.textfsm, .*, cisco_xr, sh[[ow]] ipv6 ne[[ighbors]]
cisco_xr_show_isis_neighbors.textfsm, .*, cisco_xr, sh[[ow]] isis ne[[ighbors]]
cisco_xr_show_lldp_neighbors.textfsm, .*, cisco_xr, sh[[ow]] lld[[p]] neig[[hbors]]
cisco_xr_show_rsvp_neighbors.textfsm, .*, cisco_xr, sh[[ow]] rs[[vp]] neigh[[bors]]
cisco_xr_show_vrf_all_detail.textfsm, .*, cisco_xr, sh[[ow]] vrf (\S+) det[[ail]]
cisco_xr_show_bgp_neighbors.textfsm, .*, cisco_xr, sh[[ow]] bg[[p]] nei[[ghbors]]
cisco_xr_show_ospf_neighbor.textfsm, .*, cisco_xr, sh[[ow]] ospf nei[[ghbor]]
cisco_xr_show_processes_cpu.textfsm, .*, cisco_xr, sh[[ow]] proc[[esses]] c[[pu]]
cisco_xr_show_version_brief.textfsm, .*, cisco_xr, sh[[ow]] ver[[sion]] br[[ief]]
cisco_xr_show_bfd_sessions.textfsm, .*, cisco_xr, sh[[ow]] bf[[d]] sess[[ions]]
cisco_xr_show_drops_np_all.textfsm, .*, cisco_xr, sh[[ow]] drops np all
cisco_xr_show_pim_neighbor.textfsm, .*, cisco_xr, sh[[ow]] pi[[m]] neigh[[bor]]
cisco_xr_show_interfaces.textfsm, .*, cisco_xr, sh[[ow]] inte[[rfaces]]
cisco_xr_admin_show_vm.textfsm, .*, cisco_xr, adm[[in]] sh[[ow]] vm
cisco_xr_show_ip_route.textfsm, .*, cisco_xr, sh[[ow]] (?:ip )?ro[[ute]]
cisco_xr_show_version.textfsm, .*, cisco_xr, sh[[ow]] ver[[sion]]
cisco_xr_show_hsrp.textfsm, .*, cisco_xr, sh[[ow]] hs[[rp]]
cisco_xr_show_arp.textfsm, .*, cisco_xr, sh[[ow]] arp
cisco_xr_show_bgp.textfsm, .*, cisco_xr, sh[[ow]] bg[[p]]
cisco_xr_ping.textfsm, .*, cisco_xr, ping
cisco_xr_dir.textfsm, .*, cisco_xr, dir

dell_force10_show_ip_interface_brief.textfsm, .*, dell_force10, sh[[ow]] ip int[[erface]] br[[ief]]
dell_force10_show_vlan_brief.textfsm, .*, dell_force10, sh[[ow]] vl[[an]] br[[ief]]
dell_force10_show_version.textfsm, .*, dell_force10, sh[[ow]] ver[[sion]]
dell_force10_show_vlan.textfsm, .*, dell_force10, sh[[ow]] vl[[an]]
dell_force10_show_arp.textfsm, .*, dell_force10, sh[[ow]] ar[[p]]

dell_powerconnect_show_interfaces_description.textfsm, .*, dell_powerconnect, sh[[ow]] int[[erfaces]] des[[cription]]
dell_powerconnect_show_bridge_address_table.textfsm, .*, dell_powerconnect, sh[[ow]] br[[idge]] a[[ddress-table]]
dell_powerconnect_show_interfaces_status.textfsm, .*, dell_powerconnect, sh[[ow]] int[[erfaces]] st[[atus]]

dlink_ds_show_arpentry.textfsm, .*, dlink_ds, sh[[ow]] arpe[[ntry]]

edgecore_show_interfaces_switchport.textfsm, .*, edgecore, show interfaces switchport
edgecore_show_interfaces_status.textfsm, .*, edgecore, show interfaces status
edgecore_show_interfaces_brief.textfsm, .*, edgecore, show interfaces brief
edgecore_show_ip_interface.textfsm, .*, edgecore, show ip interface
edgecore_show_version.textfsm, .*, edgecore, show version
edgecore_show_system.textfsm, .*, edgecore, show system
edgecore_show_vlan.textfsm, .*, edgecore, show vlan

eltex_show_interfaces_description.textfsm, .*, eltex, sh[[ow]] int[[erfaces]] de[[scription]]
eltex_show_interfaces_switchport.textfsm, .*, eltex, sh[[ow]] int[[erfaces]] sw[[itchport]]
eltex_show_interfaces_status.textfsm, .*, eltex, sh[[ow]] int[[erfaces]] st[[atus]]
eltex_show_ip_interface.textfsm, .*, eltex, sh[[ow]] ip int[[erfaces]]
eltex_show_interface.textfsm, .*, eltex, sh[[ow]] int[[erfaces]]
eltex_show_system_id.textfsm, .*, eltex, sh[[ow]] syst[[em]] id
eltex_show_system.textfsm, .*, eltex, sh[[ow]] syst[[em]]
eltex_show_vlan.textfsm, .*, eltex, sh[[ow]] vl[[an]]

ericsson_ipos_show_isis_adjacency.textfsm, .*, ericsson_ipos, sh[[ow]] isis adja[[cency]]
ericsson_ipos_show_version.textfsm, .*, ericsson_ipos, sh[[ow]] ver[[sion]]
ericsson_ipos_show_arp.textfsm, .*, ericsson_ipos, sh[[ow]] ar[[p]]

extreme_exos_show_ports_information_detail.textfsm, .*, extreme_exos, show ports information detail
extreme_exos_show_ports_description.textfsm, .*, extreme_exos, show ports description
extreme_exos_show_ports_information.textfsm, .*, extreme_exos, show ports information
extreme_exos_show_vlan_description.textfsm, .*, extreme_exos, show vlan description
extreme_exos_show_ipconfig.textfsm, .*, extreme_exos, show ipconfig
extreme_exos_show_sharing.textfsm, .*, extreme_exos, show sharing
extreme_exos_show_iparp.textfsm, .*, extreme_exos, show iparp

fortinet_get_router_info_routing-table_all.textfsm, .*, fortinet, g[[et]] r[[outer]] info ro[[uting-table]] a[[ll]]
fortinet_get_router_info_bgp_neighbors.textfsm, .*, fortinet, g[[et]] r[[outer]] info bg[[p]] nei[[ghbors]]
fortinet_get_system_interface_physical.textfsm, .*, fortinet, g[[et]] sy[[stem]] in[[terface]] p[[hysical]]
fortinet_get_router_info_bgp_summary.textfsm, .*, fortinet, g[[et]] r[[outer]] info bg[[p]] su[[mmary]]
fortinet_get_router_info_ospf_status.textfsm, .*, fortinet, g[[et]] r[[outer]] info o[[spf]] s[[tatus]]
fortinet_get_hardware_nic_nic-name.textfsm, .*, fortinet, g[[et]] hard[[ware]] n[[ic]] (\S+)
fortinet_execute_dhcp_lease-list.textfsm, .*, fortinet, exec[[ute]] dhcp lease-l[[ist]]
fortinet_get_system_ha_status.textfsm, .*, fortinet, g[[et]] sy[[stem]] ha s[[tatus]]
fortinet_get_system_interface.textfsm, .*, fortinet, g[[et]] sy[[stem]] in[[terface]]
fortinet_execute_log_display.textfsm, .*, fortinet, exec[[ute]] l[[og]] di[[splay]]
fortinet_execute_traceroute.textfsm, .*, fortinet, exec[[ute]] traceroute
fortinet_fnsysctl_ifconfig.textfsm, .*, fortinet, fnsysctl ifconfig
fortinet_get_system_status.textfsm, .*, fortinet, g[[et]] sy[[stem]] stat[[us]]
fortinet_diagnose_sys_top.textfsm, .*, fortinet, d[[iagnose]] sy[[s]] top
fortinet_get_hardware_nic.textfsm, .*, fortinet, g[[et]] hard[[ware]] n[[ic]]
fortinet_get_system_arp.textfsm, .*, fortinet, g[[et]] sy[[stem]] arp
fortinet_execute_date.textfsm, .*, fortinet, exec[[ute]] da[[te]]
fortinet_execute_ping.textfsm, .*, fortinet, exec[[ute]] ping
fortinet_execute_time.textfsm, .*, fortinet, exec[[ute]] ti[[me]]

hp_comware_display_lldp_neighbor-information_verbose.textfsm, .*, hp_comware, di[[splay]] ll[[dp]] n[[eighbor-information]] v[[erbose]]
hp_comware_display_lldp_neighbor-information_list.textfsm, .*, hp_comware, di[[splay]] ll[[dp]] n[[eighbor-information]] l[[ist]]
hp_comware_display_ip_vpn-instance_instance-name.textfsm, .*, hp_comware, di[[splay]] ip vpn[[-instance]] in[[instance-name]]
hp_comware_display_counters_bound_interface.textfsm, .*, hp_comware, di[[splay]] cou[[nters]] (\S+) i[[nterface]]
hp_comware_display_link-aggregation_verbose.textfsm, .*, hp_comware, di[[splay]] link[[-aggregation]] v[[erbose]]
hp_comware_display_ip_routing-table.textfsm, .*, hp_comware, di[[splay]] ip r[[outing-table]]
hp_comware_display_device_manuinfo.textfsm, .*, hp_comware, di[[splay]] dev[[ice]] m[[anuinfo]]
hp_comware_display_interface_brief.textfsm, .*, hp_comware, di[[splay]] int[[erface]] brief
hp_comware_display_ip_vpn-instance.textfsm, .*, hp_comware, di[[splay]] ip vpn[[-instance]]
hp_comware_display_ip_interface.textfsm, .*, hp_comware, dis[[play]] ip i[[nterface]]
hp_comware_display_mac-address.textfsm, .*, hp_comware, di[[splay]] mac-ad[[dress]]
hp_comware_display_vlan_brief.textfsm, .*, hp_comware, di[[splay]] v[[lan]] b[[rief]]
hp_comware_display_interface.textfsm, .*, hp_comware, dis[[play]] int[[erface]]
hp_comware_display_vlan_all.textfsm, .*, hp_comware, di[[splay]] v[[lan]] a[[ll]]
hp_comware_display_clock.textfsm, .*, hp_comware, di[[splay]] clo[[ck]]
hp_comware_display_arp.textfsm, .*, hp_comware, di[[splay]] a[[rp]]

hp_procurve_show_lldp_info_remote-device_detail.textfsm, .*, hp_procurve, sh[[ow]] ll[[dp]] i[[nfo]] r[[emote-device]] .+
hp_procurve_show_lldp_info_remote-device.textfsm, .*, hp_procurve, sh[[ow]] ll[[dp]] i[[nfo]] r[[emote-device]]
hp_procurve_show_cdp_neighbors_detail.textfsm, .*, hp_procurve, sh[[ow]] cd[[p]] nei[[ghbors]] d[[etail]]
hp_procurve_show_interfaces_brief.textfsm, .*, hp_procurve, sh[[ow]] int[[erfaces]] b[[rief]]
hp_procurve_show_port-security.textfsm, .*, hp_procurve, sh[[ow]] port-s[[ecurity]]
hp_procurve_show_tech_buffers.textfsm, .*, hp_procurve, sh[[ow]] tec[[h]] buf[[ffers]]
hp_procurve_show_mac-address.textfsm, .*, hp_procurve, sh[[ow]] mac-a[[ddress]]
hp_procurve_show_interfaces.textfsm, .*, hp_procurve, sh[[ow]] int[[erfaces]]
hp_procurve_show_ip_route.textfsm, .*, hp_procurve, sh[[ow]] ip ro[[ute]]
hp_procurve_show_system.textfsm, .*, hp_procurve, sh[[ow]] syst[[em]]
hp_procurve_show_trunks.textfsm, .*, hp_procurve, sh[[ow]] tr[[unks]]
hp_procurve_show_vlans.textfsm, .*, hp_procurve, sh[[ow]] vl[[ans]]
hp_procurve_show_arp.textfsm, .*, hp_procurve, sh[[ow]] ar[[p]]
hp_procurve_show_ip.textfsm, .*, hp_procurve, sh[[ow]] ip

huawei_smartax_display_ont_info_summary_ont.textfsm:huawei_smartax_display_ont_info_summary_sn.textfsm, .*, huawei_smartax, di[[splay]] ont i[[nfo]] su[[mmary]] \S+ *$
huawei_smartax_display_ont_info_0_1_2.textfsm:huawei_smartax_display_ont_info_description.textfsm, .*, huawei_smartax, di[[splay]] ont i[[nfo]] (\d+\s*|\d+ \d+ \d+\s*)\s*[[all]] *$
huawei_smartax_display_ont_port_vlan_0_1_byport_eth_0.textfsm, .*, huawei_smartax, di[[splay]] ont p[[ort]] vl[[an]] \d+ \d+ byport eth \d+\s*
huawei_smartax_display_ont_port_state_0_1_eth-port.textfsm, .*, huawei_smartax, di[[splay]] on[[t]] por[[t]] st[[ate]] \d+ \d+ e[[th-port]]
huawei_smartax_display_ont_port_vlan_0_1_byvlan_0.textfsm, .*, huawei_smartax, di[[splay]] ont p[[ort]] v[[lan]] \d+ \d+ byvlan \d+
huawei_smartax_display_mac-address_ont_0_1_2_0.textfsm, .*, huawei_smartax, di[[splay]] mac-a[[ddress]] o[[nt]] \d+(\/)?\s*\d+(\/)?\s*\d+ \d+ *$
huawei_smartax_display_ont_optical-info_0_all.textfsm, .*, huawei_smartax, di[[splay]] ont o[[ptical-info]] \d+ all
huawei_smartax_display_ont_snmp-profile_0_all.textfsm, .*, huawei_smartax, di[[splay]] ont s[[nmp-profile]] \d+ all
huawei_smartax_display_ont_gemport_0_ontid_0.textfsm, .*, huawei_smartax, di[[splay]] ont g[[emport]] \d+ ontid \d+\s*
huawei_smartax_display_sysman_service_state.textfsm, .*, huawei_smartax, di[[splay]] sysman s[[ervice]] s[[tate]]
huawei_smartax_display_board_serial-number.textfsm, .*, huawei_smartax, di[[splay]] bo[[ard]] s[[erial-number]]\s*
huawei_smartax_display_service-port_all.textfsm, .*, huawei_smartax, disp[[lay]] ser[[vice-port]] a[[ll]] *$
huawei_smartax_display_ont_autofind.textfsm, .*, huawei_smartax, di[[splay]] ont a[[utofind]]\s*
huawei_smartax_display_ont_info_0_1.textfsm, .*, huawei_smartax, di[[splay]] ont i[[nfo]] (by-sn \w*\s*|\d+ \d+\s*|\d+ \d+ \d+ \d+\s*)
huawei_smartax_display_mac-address.textfsm, .*, huawei_smartax, display mac-address\s*
huawei_smartax_display_temperature.textfsm, .*, huawei_smartax, di[[splay]] tem[[perature]]\s*
huawei_smartax_display_port_info.textfsm, .*, huawei_smartax, di[[splay]] port i[[nfo]]\s*
huawei_smartax_display_sysuptime.textfsm, .*, huawei_smartax, di[[splay]] sys[[uptime]]
huawei_smartax_display_version.textfsm, .*, huawei_smartax, di[[splay]] ver[[sion]]
huawei_smartax_display_board.textfsm, .*^, huawei_smartax, di[[splay]] bo[[ard]]\s*
huawei_smartax_display_cpu.textfsm, .*, huawei_smartax, di[[splay]] cpu\s*
huawei_smartax_display_mem.textfsm, .*, huawei_smartax, di[[splay]] mem[[ory]]\s*
huawei_smartax_port_vlan.textfsm, .*, huawei_smartax, port vlan\s*
huawei_smartax_ont_add.textfsm, .*, huawei_smartax, ont add\s*

huawei_vrp_display_traffic-filter_applied-record.textfsm, .*, huawei_vrp, dis[[play]] traffic-filter applied-record
huawei_vrp_display_ip_vpn-instance_interface.textfsm, .*, huawei_vrp, di[[splay]] ip vpn[[-instance]] interface
huawei_vrp_display_snmp-agent_community_read.textfsm, .*, huawei_vrp, dis[[play]] snm[[p-agent]] c[[ommunity]] (r[[ead]]|w[[rite]])
huawei_vrp_display_ip_routing-table_verbose.textfsm, .*, huawei_vrp, dis[[play]] ip(v6)? routi[[ng-table]] ve[[rbose]]
huawei_vrp_display_interface_description.textfsm, .*, huawei_vrp, dis[[play]] inter[[face]] des[[cription]]
huawei_vrp_display_service-set_id_id.textfsm, .*, huawei_vrp, dis[[play]] service-set (id|name) \S+
huawei_vrp_display_interface_brief.textfsm, .*, huawei_vrp, dis[[play]] inter[[face]] br[[ief]]
huawei_vrp_display_ip_vpn-instance.textfsm, .*, huawei_vrp, di[[splay]] ip vpn[[-instance]]
huawei_vrp_display_service-set_all.textfsm, .*, huawei_vrp, dis[[play]] service-set all
huawei_vrp_display_ipv6_neighbors.textfsm, .*, huawei_vrp, dis[[play]] ipv6 n[[eighbors]]
huawei_vrp_display_mpls_te_tunnel.textfsm, .*, huawei_vrp, dis[[play]] mpls te tunnel
huawei_vrp_display_lldp_neighbor.textfsm, .*, huawei_vrp, dis[[play]] lldp nei[[ghbor]]
huawei_vrp_display_mac-address.textfsm, .*, huawei_vrp, disp[[lay]] mac[[-address]]
huawei_vrp_display_temperature.textfsm, .*, huawei_vrp, dis[[play]] tem[[perature]]
huawei_vrp_display_nat_server.textfsm, .*, huawei_vrp, dis[[play]] na[[t]] ser[[ver]]
huawei_vrp_display_sn_license.textfsm, .*, huawei_vrp, dis[[play]] sn l[[icence]]
huawei_vrp_display_arp_brief.textfsm, .*, huawei_vrp, dis[[play]] arp br[[ief]]
huawei_vrp_display_interface.textfsm, .*, huawei_vrp, dis[[play]] inter[[face]]\s*((?!brief|counters|description).)*$
huawei_vrp_display_isis_peer.textfsm, .*, huawei_vrp, disp[[lay]] isis p[[eer]]
huawei_vrp_display_port_vlan.textfsm, .*, huawei_vrp, dis[[play]] port vl[[an]]
huawei_vrp_display_acl_all.textfsm, .*, huawei_vrp, dis[[play]] acl(\si[[pv6]])? a[[ll]]
huawei_vrp_display_arp_all.textfsm, .*, huawei_vrp, disp[[lay]] ar[[p]] all
huawei_vrp_display_startup.textfsm, .*, huawei_vrp, dis[[play]] star[[tup]]
huawei_vrp_display_version.textfsm, .*, huawei_vrp, dis[[play]] ver[[sion]]
huawei_vrp_display_device.textfsm, .*, huawei_vrp, di[[splay]] de[[vice]]( a[[ll]]|c[[ard]])?
huawei_vrp_display_vlan.textfsm, .*, huawei_vrp, di[[splay]] v[[lan]]

ipinfusion_ocnos_show_lldp_table.textfsm, .*, ipinfusion_ocnos, show ll[[dp]] t[[able]]

juniper_junos_show_chassis_cluster_interfaces.textfsm, .*, juniper_junos, sh[[ow]] ch[[assis]] c[[luster]] i[[nterface]]
juniper_junos_show_ethernet-switching_table.textfsm, .*, juniper_junos, sh[[ow]] et[[hernet-switching]] t[[able]]
juniper_junos_show_chassis_cluster_status.textfsm, .*, juniper_junos, sh[[ow]] ch[[assis]] c[[luster]] s[[tatus]]
juniper_junos_show_chassis_firmware.textfsm, .*, juniper_junos, sh[[ow]] ch[[assis]] fi[[rmware]]
juniper_junos_show_chassis_hardware.textfsm, .*, juniper_junos, sh[[ow]] ch[[assis]] ha[[ardware]]
juniper_junos_show_lacp_interfaces.textfsm, .*, juniper_junos, sh[[ow]] la[[cp]] i[[nterfaces]]
juniper_junos_show_arp_no-resolve.textfsm, .*, juniper_junos, sh[[ow]] a[[rp]] n[[o-resolve]]
juniper_junos_show_isis_adjacency.textfsm, .*, juniper_junos, sh[[ow]] is[[is]] ad[[jacency]]
juniper_junos_show_lldp_neighbors.textfsm, .*, juniper_junos, sh[[ow]] ll[[dp]] n[[eighbors]]
juniper_junos_show_ospf_neighbor.textfsm, .*, juniper_junos, sh[[ow]] ospf n[[eighbor]]
juniper_junos_show_system_uptime.textfsm, .*, juniper_junos, sh[[ow]] sys[[tem]] up[[time]]
juniper_junos_show_bgp_summary.textfsm, .*, juniper_junos, sh[[ow]] bgp sum[[mary]]
juniper_junos_show_interfaces.textfsm, .*, juniper_junos, sh[[ow]] inte[[rfaces]]
juniper_junos_show_version.textfsm, .*, juniper_junos, sh[[ow]] ver[[sion]]
juniper_junos_show_vlans.textfsm, .*, juniper_junos, sh[[ow]] vl[[ans]]

juniper_screenos_get_route.textfsm, .*, juniper_screenos, get route

linux_ip_address_show.textfsm, .*, linux, ip a[[ddress]] [[show]]
linux_ip_route_show.textfsm, .*, linux, ip r[[oute]] [[show]]
linux_ip_link_show.textfsm, .*, linux, ip l[[ink]] [[show]]
linux_ip_vrf_show.textfsm, .*, linux, ip v[[rf]] [[show]]
linux_arp_-a.textfsm, .*, linux, arp -a

mikrotik_routeros_routing_ospf_neighbor_print_terse_without-paging.textfsm, .*, mikrotik_routeros, [[/]]r[[outing]] o[[spf]] nei[[ghbor]] p[[rint]] t[[erse]] wi[[thout-paging]]
mikrotik_routeros_ip_dhcp-server_lease_print_terse_without-paging.textfsm, .*, mikrotik_routeros, [[/]]ip dh[[cp-server]] l[[ease]] pr[[int]] t[[erse]] wi[[thout-paging]]
mikrotik_routeros_routing_bgp_peer_print_status_without-paging.textfsm, .*, mikrotik_routeros, [[/]]r[[outing]] bg[[p]] p[[eer]] p[[rint]] s[[tatus]] wi[[thout-paging]]
mikrotik_routeros_interface_ethernet_poe_print_without-paging.textfsm, .*, mikrotik_routeros, [[/]]in[[terface]] et[[hernet]] po[[e]] pr[[int]] wi[[thout-paging]]
mikrotik_routeros_ip_firewall_filter_print_all_without-paging.textfsm, .*, mikrotik_routeros, [[/]]ip(v6)? firewall filter print( all)? without-paging
mikrotik_routeros_ip_dhcp-server_lease_print_without-paging.textfsm, .*, mikrotik_routeros, [[/]]ip dhcp-server lease print without-paging
mikrotik_routeros_ip_firewall_nat_print_all_without-paging.textfsm, .*, mikrotik_routeros, [[/]]ip(v6)? firewall nat print all without-paging
mikrotik_routeros_interface_ethernet_monitor_name_once.textfsm, .*, mikrotik_routeros, [[/]]interface ethernet monitor (\S+) once
mikrotik_routeros_interface_print_terse_without-paging.textfsm, .*, mikrotik_routeros, [[/]]in[[terface]] p[[rint]] t[[erse]] wi[[thout-paging]]
mikrotik_routeros_ip_firewall_address-list_print_terse.textfsm, .*, mikrotik_routeros, [[/]]ip(v6)? firewall address-list print terse
mikrotik_routeros_ip_route_print_terse_without-paging.textfsm, .*, mikrotik_routeros, [[/]]ip r[[oute]] p[[rint]] t[[erse]] wi[[thout-paging]]
mikrotik_routeros_snmp_community_print_without-paging.textfsm, .*, mikrotik_routeros, [[/]]snmp community print without-paging
mikrotik_routeros_ipv6_neighbor_print_without-paging.textfsm, .*, mikrotik_routeros, [[/]]ipv6 neighbor print without-paging
mikrotik_routeros_ip_arp_print_terse_without-paging.textfsm, .*, mikrotik_routeros, [[/]]ip a[[rp]] p[[rint]] t[[erse]] wi[[thout-paging]]
mikrotik_routeros_log_print_detail_without-paging.textfsm, .*, mikrotik_routeros, [[/]]log p[[rint]] d[[etail]] wi[[thout-paging]]
mikrotik_routeros_routing_bgp_peer_print_status.textfsm, .*, mikrotik_routeros, [[/]]r[[outing]] bg[[p]] p[[eer]] p[[rint]] s[[tatus]]
mikrotik_routeros_ip_arp_print_without-paging.textfsm, .*, mikrotik_routeros, [[/]]ip arp print without-paging
mikrotik_routeros_routing_ospf_neighbor_print.textfsm, .*, mikrotik_routeros, [[/]]r[[outing]] o[[spf]] nei[[ghbor]] p[[rint]]
mikrotik_routeros_ip_dhcp-server_lease_print.textfsm, .*, mikrotik_routeros, [[/]]ip dh[[cp-server]] l[[ease]] p[[rint]]
mikrotik_routeros_ip_address_export_verbose.textfsm, .*, mikrotik_routeros, [[/]]ip a[[ddress]] e[[xport]] [[verbose]]
mikrotik_routeros_interface_ethernet_print.textfsm, .*, mikrotik_routeros, [[/]]in[[terface]] et[[hernet]] pr[[int]]
mikrotik_routeros_ip_neighbor_print_detail.textfsm, .*, mikrotik_routeros, [[/]]ip n[[eighbor]] p[[rint]] d[[etail]]
mikrotik_routeros_system_routerboard_print.textfsm, .*, mikrotik_routeros, [[/]]system routerboard print
mikrotik_routeros_tool_speed-test_address.textfsm, .*, mikrotik_routeros, [[/]]too[[l]] s[[peed-test]]
mikrotik_routeros_interface_print_detail.textfsm, .*, mikrotik_routeros, [[/]]in[[terface]] p[[rint]] d[[etail]]
mikrotik_routeros_interface_print_brief.textfsm, .*, mikrotik_routeros, [[/]]in[[terface]] p[[rint]] b[[rief]]
mikrotik_routeros_ip_route_print_detail.textfsm, .*, mikrotik_routeros, [[/]]ip r[[oute]] p[[rint]] d[[etail]]
mikrotik_routeros_system_identity_print.textfsm, .*, mikrotik_routeros, [[/]]sy[[stem]] i[[dentity]] p[[rint]]
mikrotik_routeros_system_resource_print.textfsm, .*, mikrotik_routeros, [[/]]sy[[stem]] re[[source]] p[[rint]]
mikrotik_routeros_ip_route_print_terse.textfsm, .*, mikrotik_routeros, [[/]]ip(v6)? route print terse
mikrotik_routeros_system_clock_print.textfsm, .*, mikrotik_routeros, [[/]]sy[[stem]] cl[[ock]] p[[rint]]
mikrotik_routeros_user_active_print.textfsm, .*, mikrotik_routeros, [[/]]us[[er]] ac[[tive]] p[[rint]]
mikrotik_routeros_ip_address_print.textfsm, .*, mikrotik_routeros, [[/]]ip ad[[dress]] p[[rint]]
mikrotik_routeros_ip_arp_print.textfsm, .*, mikrotik_routeros, [[/]]ip a[[rp]] p[[rint]]
mikrotik_routeros_tool_profile.textfsm, .*, mikrotik_routeros, [[/]]too[[l]] pr[[ofile]]
mikrotik_routeros_user_print.textfsm, .*, mikrotik_routeros, [[/]]us[[er]] p[[rint]]
mikrotik_routeros_ping.textfsm, .*, mikrotik_routeros, [[/]]pin[[g]]

oneaccess_oneos_show_voice_dial-peer_voice_voip_all.textfsm, .*, oneaccess_oneos, sh[[ow]] voice dial-peer voice voip all( al)?
oneaccess_oneos_show_voice_voice-port_pri_all_histo.textfsm, .*, oneaccess_oneos, sh[[ow]] voice voice-port pri all histo
oneaccess_oneos_show_voice_voip-call_active_all.textfsm, .*, oneaccess_oneos, sh[[ow]] voice voip-call active
oneaccess_oneos_show_cellular-radio_equipment.textfsm, .*, oneaccess_oneos, sh[[ow]] cell[[ular-radio]] equipment
oneaccess_oneos_show_voice_voice-port_pri_all.textfsm, .*, oneaccess_oneos, sh[[ow]] voice voice-port pri all
oneaccess_oneos_show_policy-interface_output.textfsm, .*, oneaccess_oneos, sh[[ow]] policy-interface output
oneaccess_oneos_show_running-config_ip_route.textfsm, .*, oneaccess_oneos, sh[[ow]] run[[ning-config]] (\| i[[nclude]] )?ip route
oneaccess_oneos_show_voice_voip-call_any_all.textfsm, .*, oneaccess_oneos, sh[[ow]] voice voip-call any all
oneaccess_oneos_show_cellular-radio_context.textfsm, .*, oneaccess_oneos, sh[[ow]] cell[[ular-radio]] context
oneaccess_oneos_show_cellular-radio_network.textfsm, .*, oneaccess_oneos, sh[[ow]] cell[[ular-radio]] network
oneaccess_oneos_show_ip_as-path-access-list.textfsm, .*, oneaccess_oneos, sh[[ow]] ip as-path-access-list
oneaccess_oneos_show_running-config_ip_dhcp.textfsm, .*, oneaccess_oneos, sh[[ow]] run[[ning-config]] (\| be[[gin]] )?ip dhcp
oneaccess_oneos_show_system_secure-crashlog.textfsm, .*, oneaccess_oneos, sh[[ow]] system secure-crashlog
oneaccess_oneos_show_transceiver_equipment.textfsm, .*, oneaccess_oneos, sh[[ow]] transceiver equipment
oneaccess_oneos_show_voice_voice-port_all.textfsm, .*, oneaccess_oneos, sh[[ow]] voice voice-port all
oneaccess_oneos_show_running-config_bind.textfsm, .*, oneaccess_oneos, sh[[ow]] run[[ning-config]] \|?bind
oneaccess_oneos_show_ip_interface_brief.textfsm, .*, oneaccess_oneos, sh[[ow]] ip int[[erface]] brief
oneaccess_oneos_show_running-config_aaa.textfsm, .*, oneaccess_oneos, sh[[ow]] run[[ning-config]] \|?aaa
oneaccess_oneos_show_product-info-area.textfsm, .*, oneaccess_oneos, sh[[ow]] product\-info\-area
oneaccess_oneos_show_voice_sip-gateway.textfsm, .*, oneaccess_oneos, sh[[ow]] voice sip-gateway
oneaccess_oneos_show_ip_access-lists.textfsm, .*, oneaccess_oneos, sh[[ow]] ip access-lists[[s]]
oneaccess_oneos_show_isdn_led-status.textfsm, .*, oneaccess_oneos, sh[[ow]] isdn led-status
oneaccess_oneos_show_isdn_status_all.textfsm, .*, oneaccess_oneos, sh[[ow]] isdn status all
oneaccess_oneos_show_reboot_counters.textfsm, .*, oneaccess_oneos, sh[[ow]] reboot counters
oneaccess_oneos_show_system_hardware.textfsm, .*, oneaccess_oneos, sh[[ow]] system hardware
oneaccess_oneos_cat_bsa_bsaboot.inf.textfsm, .*, oneaccess_oneos, cat (\/|)(BSA|bsa)(\/| )bsa(B|b)oot.inf
oneaccess_oneos_show_ip_bgp_summary.textfsm, .*, oneaccess_oneos, sh[[ow]] ip bgp summary
oneaccess_oneos_show_ip_prefix-list.textfsm, .*, oneaccess_oneos, sh[[ow]] ip prefix-list
oneaccess_oneos_show_snmp_community.textfsm, .*, oneaccess_oneos, sh[[ow]] snmp comm[[unity]]
oneaccess_oneos_show_soft-file_info.textfsm, .*, oneaccess_oneos, sh[[ow]] soft\-file info.*
oneaccess_oneos_show_software-image.textfsm, .*, oneaccess_oneos, sh[[ow]] software-image
oneaccess_oneos_show_vrrp_interface.textfsm, .*, oneaccess_oneos, sh[[ow]] vrrp int[[erface]]
oneaccess_oneos_show_system_status.textfsm, .*, oneaccess_oneos, sh[[ow]] system status
oneaccess_oneos_show_tacacs-server.textfsm, .*, oneaccess_oneos, sh[[ow]] tacacs\-server
oneaccess_oneos_show_ip_vrf_brief.textfsm, .*, oneaccess_oneos, sh[[ow]] ip vrf brief
oneaccess_oneos_show_isdn_active.textfsm, .*, oneaccess_oneos, sh[[ow]] isdn active
oneaccess_oneos_show_interfaces.textfsm, .*, oneaccess_oneos, sh[[ow]] interface[[s]]
oneaccess_oneos_show_route-map.textfsm, .*, oneaccess_oneos, sh[[ow]] route-map
oneaccess_oneos_show_track_all.textfsm, .*, oneaccess_oneos, sh[[ow]] track all
oneaccess_oneos_show_voice_mos.textfsm, .*, oneaccess_oneos, sh[[ow]] voice mos
oneaccess_oneos_show_helpers.textfsm, .*, oneaccess_oneos, sh[[ow]] helpers
oneaccess_oneos_show_ip_ssh.textfsm, .*, oneaccess_oneos, sh[[ow]] ip ssh
oneaccess_oneos_show_memory.textfsm, .*, oneaccess_oneos, sh[[ow]] memory
oneaccess_oneos_show_tacacs.textfsm, .*, oneaccess_oneos, sh[[ow]] tacacs
oneaccess_oneos_show_sntp.textfsm, .*, oneaccess_oneos, sh[[ow]] sntp
oneaccess_oneos_hostname.textfsm, .*, oneaccess_oneos, hostname
oneaccess_oneos_ls.textfsm, .*, oneaccess_oneos, ls( -[lh][lh]? )?.*

paloalto_panos_show_running_security-policy.textfsm, .*, paloalto_panos, sh[[ow]] runn[[ing]] security[[-policy]]
paloalto_panos_show_high-availability_all.textfsm, .*, paloalto_panos, sh[[ow]] high[[-availability]] all
paloalto_panos_test_security-policy-match.textfsm, .*, paloalto_panos, test security-policy-match.*
paloalto_panos_show_interface_management.textfsm, .*, paloalto_panos, sh[[ow]] int[[erface]] man[[agement]]
paloalto_panos_show_interface_hardware.textfsm, .*, paloalto_panos, sh[[ow]] int[[erface]] hard[[ware]]
paloalto_panos_show_running_nat-policy.textfsm, .*, paloalto_panos, sh[[ow]] runn[[ing]] nat[[-policy]]
paloalto_panos_show_interface_logical.textfsm, .*, paloalto_panos, sh[[ow]] int[[erface]] logi[[cal]]
paloalto_panos_request_license_info.textfsm, .*, paloalto_panos, request license info 
paloalto_panos_show_counter_global.textfsm, .*, paloalto_panos, sh[[ow]] coun[[ter]] glo[[bal]]
paloalto_panos_debug_swm_status.textfsm, .*, paloalto_panos, de[[bug]] s[[wm]] status
paloalto_panos_show_system_info.textfsm, .*, paloalto_panos, sh[[ow]] sys[[tem]] in[[fo]]
paloalto_panos_show_jobs_all.textfsm, .*, paloalto_panos, sh[[ow]] jo[[bs]] all
paloalto_panos_show_arp_all.textfsm, .*, paloalto_panos, sh[[ow]] ar[[p]] all
paloalto_panos_show_mac_all.textfsm, .*, paloalto_panos, sh[[ow]] mac all

ruckus_fastiron_show_interfaces_brief.textfsm, .*, ruckus_fastiron, sh[[ow]] int[[erfaces]] b[[rief]]
ruckus_fastiron_show_media_validation.textfsm, .*, ruckus_fastiron, sh[[ow]] med[[ia]] v[[alidation]]
ruckus_fastiron_show_mac-address.textfsm, .*, ruckus_fastiron, sh[[ow]] mac[[-address]]
ruckus_fastiron_show_interfaces.textfsm, .*, ruckus_fastiron, sh[[ow]] int[[erfaces]]
ruckus_fastiron_show_version.textfsm, .*, ruckus_fastiron, sh[[ow]] ve[[rsion]]
ruckus_fastiron_show_vlan.textfsm, .*, ruckus_fastiron, sh[[ow]] vlan
ruckus_fastiron_show_arp.textfsm, .*, ruckus_fastiron, sh[[ow]] ar[[p]]

ubiquiti_edgerouter_show_interfaces_ethernet_physical.textfsm, .*, ubiquiti_edgerouter, sh[[ow]] interfaces ethernet (\S+\s)?physical
ubiquiti_edgerouter_show_dhcpv6_server_leases.textfsm, .*, ubiquiti_edgerouter, sh[[ow]] dhcpv6 server leases
ubiquiti_edgerouter_show_ipv6_neighbors.textfsm, .*, ubiquiti_edgerouter, sh[[ow]] ipv6 neighbors
ubiquiti_edgerouter_show_dhcp_leases.textfsm, .*, ubiquiti_edgerouter, sh[[ow]] dhcp leases
ubiquiti_edgerouter_show_interfaces.textfsm, .*, ubiquiti_edgerouter, show interfaces
ubiquiti_edgerouter_show_ipv6_route.textfsm, .*, ubiquiti_edgerouter, show ipv6 route
ubiquiti_edgerouter_show_nat_rules.textfsm, .*, ubiquiti_edgerouter, sh[[ow]] nat rules
ubiquiti_edgerouter_show_ip_route.textfsm, .*, ubiquiti_edgerouter, show ip route
ubiquiti_edgerouter_show_version.textfsm, .*, ubiquiti_edgerouter, sh[[ow]] ve[[rsion]]
ubiquiti_edgerouter_show_arp.textfsm, .*, ubiquiti_edgerouter, sh[[ow]] arp

ubiquiti_edgeswitch_show_version.textfsm, .*, ubiquiti_edgeswitch, sh[[ow]] ve[[rsion]]
ubiquiti_edgeswitch_show_vlan.textfsm, .*, ubiquiti_edgeswitch, sh[[ow]] vl[[an]]
ubiquiti_edgeswitch_show_arp.textfsm, .*, ubiquiti_edgeswitch, sh[[ow]] ar[[p]]

vmware_nsxv_show_ip_bgp_neighbors.textfsm, .*, vmware_nsxv, sh[[ow]] ip b[[gp]] n[[eighbors]]
vmware_nsxv_show_ip_route.textfsm, .*, vmware_nsxv, sh[[ow]] ip r[[oute]]

vyatta_vyos_show_ip_bgp_summary.textfsm, .*, .*vyos.*, sh[[ow]] ip bgp sum[[mary]]
vyatta_vyos_show_interfaces.textfsm, .*, .*vyos.*, sh[[ow]] int[[erfaces]]
vyatta_vyos_show_arp.textfsm, .*, .*vyos.*, sh[[ow]] a[[rp]]

watchguard_firebox_show_arp.textfsm, .*, watchguard_firebox, sh[[ow]] arp

yamaha_show_environment.textfsm, .*, yamaha, sh[[ow]] en[[vironment]]
yamaha_show_ip_route.textfsm, .*, yamaha, sh[[ow]] ip ro[[ute]]

zte_zxros_show_mpls_traffic-eng_tunnels.textfsm, .*, zte_zxros, sh[[ow]] mpls traffic-eng tunnels
zte_zxros_show_interface_brief.textfsm, .*, zte_zxros, sh[[ow]] inter[[face]] br[[ief]]
zte_zxros_show_isis_adjacency.textfsm, .*, zte_zxros, sh[[ow]] is[[is]] ad[[jacency]]
zte_zxros_show_interface.textfsm, .*, zte_zxros, sh[[ow]] int[[erface]]
zte_zxros_show_version.textfsm, .*, zte_zxros, sh[[ow]] ver[[sion]]
zte_zxros_show_arp.textfsm, .*, zte_zxros, sh[[ow]] arp

zyxel_os_cfg_lan_get_--Name_name.textfsm, .*, zyxel_os, cfg lan get --Name .+
zyxel_os_zycli_Ethctl_media-type.textfsm, .*, zyxel_os, zycli Ethctl media-type(\s\w+)*
zyxel_os_cfg_firewall_acl_get.textfsm, .*, zyxel_os, cfg firewall(\s|_)acl get
zyxel_os_cfg_nat_addr_map_get.textfsm, .*, zyxel_os, cfg nat(\s|_)addr(\s|_)map get
zyxel_os_cfg_static_route_get.textfsm, .*, zyxel_os, cfg static(\s|_)route get
zyxel_os_cfg_static_dhcp_get.textfsm, .*, zyxel_os, cfg static(\s|_)dhcp get
zyxel_os_cfg_intf_group_get.textfsm, .*, zyxel_os, cfg intf(\s|_)group get
zyxel_os_cfg_lanhosts_get.textfsm, .*, zyxel_os, cfg lanhosts get
zyxel_os_cfg_ipalias_get.textfsm, .*, zyxel_os, cfg ipalias get
zyxel_os_cfg_snmp_get.textfsm, .*, zyxel_os, cfg snmp get
zyxel_os_cfg_wlan_get.textfsm, .*, zyxel_os, cfg wlan get
zyxel_os_cfg_lan_get.textfsm, .*, zyxel_os, cfg lan get
zyxel_os_cfg_nat_get.textfsm, .*, zyxel_os, cfg nat get
zyxel_os_sys_atsh.textfsm, .*, zyxel_os, sys atsh

`