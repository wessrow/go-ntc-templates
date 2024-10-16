package juniper_junos

type ShowVersion struct {
	Services_crypto                                  string   `json:"SERVICES_CRYPTO"`
	Services_ssl                                     string   `json:"SERVICES_SSL"`
	Runtime_software_suite                           string   `json:"RUNTIME_SOFTWARE_SUITE"`
	Serial_number                                    string   `json:"SERIAL_NUMBER"`
	Services_rpm                                     string   `json:"SERVICES_RPM"`
	Routing_software_suite                           string   `json:"ROUTING_SOFTWARE_SUITE"`
	Other_properties_versions                        []string `json:"OTHER_PROPERTIES_VERSIONS"`
	Junos_version                                    string   `json:"JUNOS_VERSION"`
	Packet_forwarding_engine_support_mx_common       string   `json:"PACKET_FORWARDING_ENGINE_SUPPORT_MX_COMMON"`
	Appid_services                                   string   `json:"APPID_SERVICES"`
	Services_captive_portal_content_delivery_package string   `json:"SERVICES_CAPTIVE_PORTAL_CONTENT_DELIVERY_PACKAGE"`
	Hostname                                         string   `json:"HOSTNAME"`
	Border_gateway_function_package                  string   `json:"BORDER_GATEWAY_FUNCTION_PACKAGE"`
	Idp_services                                     string   `json:"IDP_SERVICES"`
	Services_jflow_container_package                 string   `json:"SERVICES_JFLOW_CONTAINER_PACKAGE"`
	Services_application_level_gateways              string   `json:"SERVICES_APPLICATION_LEVEL_GATEWAYS"`
	Py_base_i386                                     string   `json:"PY_BASE_I386"`
	Qfabric_system_id                                string   `json:"QFABRIC_SYSTEM_ID"`
	Lab_package                                      string   `json:"LAB_PACKAGE"`
	Base_os_software_suite                           string   `json:"BASE_OS_SOFTWARE_SUITE"`
	Online_documentation                             string   `json:"ONLINE_DOCUMENTATION"`
	Services_aacl_container_package                  string   `json:"SERVICES_AACL_CONTAINER_PACKAGE"`
	Services_nat                                     string   `json:"SERVICES_NAT"`
	Services_mobilenext_software_package             string   `json:"SERVICES_MOBILENEXT_SOFTWARE_PACKAGE"`
	Services_mobile_subscriber_service_package       string   `json:"SERVICES_MOBILE_SUBSCRIBER_SERVICE_PACKAGE"`
	Redis_version                                    string   `json:"REDIS_VERSION"`
	Kernel_software_suite                            string   `json:"KERNEL_SOFTWARE_SUITE"`
	Packet_forwarding_engine_support_m_t_ex_common   string   `json:"PACKET_FORWARDING_ENGINE_SUPPORT_M_T_EX_COMMON"`
	Fips_mode_utilities                              string   `json:"FIPS_MODE_UTILITIES"`
	Services_http_content_management_package         string   `json:"SERVICES_HTTP_CONTENT_MANAGEMENT_PACKAGE"`
	Services_ipsec                                   string   `json:"SERVICES_IPSEC"`
	Model                                            string   `json:"MODEL"`
	Base_os_boot                                     string   `json:"BASE_OS_BOOT"`
	Services_ll_pdf_container_package                string   `json:"SERVICES_LL_PDF_CONTAINER_PACKAGE"`
	Voice_services_container_package                 string   `json:"VOICE_SERVICES_CONTAINER_PACKAGE"`
	Other_device_properties                          []string `json:"OTHER_DEVICE_PROPERTIES"`
	Crypto_software_suite                            string   `json:"CRYPTO_SOFTWARE_SUITE"`
	Services_ptsp_container_package                  string   `json:"SERVICES_PTSP_CONTAINER_PACKAGE"`
	Services_stateful_firewall                       string   `json:"SERVICES_STATEFUL_FIREWALL"`
	Platform_software_suite                          string   `json:"PLATFORM_SOFTWARE_SUITE"`
}

var ShowVersion_Template string = `Value Required HOSTNAME (\S+)
Value MODEL (\S+)
Value JUNOS_VERSION ([\S\s]*)
Value BASE_OS_BOOT (\S+)
Value BASE_OS_SOFTWARE_SUITE (\S+)
Value KERNEL_SOFTWARE_SUITE (\S+)
Value CRYPTO_SOFTWARE_SUITE (\S+)
Value PACKET_FORWARDING_ENGINE_SUPPORT_M_T_EX_COMMON (\S+)
Value PACKET_FORWARDING_ENGINE_SUPPORT_MX_COMMON (\S+)
Value FIPS_MODE_UTILITIES (\S+)
Value ONLINE_DOCUMENTATION (\S+)
Value SERVICES_AACL_CONTAINER_PACKAGE (\S+)
Value SERVICES_APPLICATION_LEVEL_GATEWAYS (\S+)
Value APPID_SERVICES (\S+)
Value BORDER_GATEWAY_FUNCTION_PACKAGE (\S+)
Value SERVICES_CAPTIVE_PORTAL_CONTENT_DELIVERY_PACKAGE (\S+)
Value SERVICES_HTTP_CONTENT_MANAGEMENT_PACKAGE (\S+)
Value IDP_SERVICES (\S+)
Value SERVICES_JFLOW_CONTAINER_PACKAGE (\S+)
Value SERVICES_LL_PDF_CONTAINER_PACKAGE (\S+)
Value SERVICES_MOBILENEXT_SOFTWARE_PACKAGE (\S+)
Value SERVICES_MOBILE_SUBSCRIBER_SERVICE_PACKAGE (\S+)
Value SERVICES_NAT (\S+)
Value SERVICES_PTSP_CONTAINER_PACKAGE (\S+)
Value SERVICES_RPM (\S+)
Value SERVICES_STATEFUL_FIREWALL (\S+)
Value VOICE_SERVICES_CONTAINER_PACKAGE (\S+)
Value SERVICES_CRYPTO (\S+)
Value SERVICES_SSL (\S+)
Value SERVICES_IPSEC (\S+)
Value PLATFORM_SOFTWARE_SUITE (\S+)
Value RUNTIME_SOFTWARE_SUITE (\S+)
Value ROUTING_SOFTWARE_SUITE (\S+)
Value PY_BASE_I386 (\S+)
Value LAB_PACKAGE (\S+)
Value SERIAL_NUMBER (\S+)
Value QFABRIC_SYSTEM_ID (\S+)
Value List OTHER_DEVICE_PROPERTIES ([^\[]*)
Value List OTHER_PROPERTIES_VERSIONS (\S+)
Value REDIS_VERSION (\S+)

Start
  ^Hostname:\s+${HOSTNAME}
  ^Model:\s+${MODEL}
  ^Junos:\s+${JUNOS_VERSION}
  ^JUNOS\s+EX\s*Software\s+Suite\s+\[${JUNOS_VERSION}\]
  ^JUNOS\s+Base\s+OS\s+boot\s+\[${BASE_OS_BOOT}\]
  ^JUNOS\s+Base\s+OS\s+Software\s+Suite\s+\[${BASE_OS_SOFTWARE_SUITE}\]
  ^JUNOS\s+Kernel\s+Software\s+Suite\s+\[${KERNEL_SOFTWARE_SUITE}\]
  ^JUNOS\s+Crypto\s+Software\s+Suite\s+\[${CRYPTO_SOFTWARE_SUITE}\]
  ^JUNOS\s+Packet\s+Forwarding\s+Engine\s+Support\s+\(M/T/EX\s+Common\)\s+\[${PACKET_FORWARDING_ENGINE_SUPPORT_M_T_EX_COMMON}\]
  ^JUNOS\s+Packet\s+Forwarding\s+Engine\s+Support\s+\(MX\s+Common\)\s+\[${PACKET_FORWARDING_ENGINE_SUPPORT_MX_COMMON}\]
  ^JUNOS\s+FIPS\s+mode\s+utilities\s+\[${FIPS_MODE_UTILITIES}\]
  ^JUNOS\s+Online\s+Documentation\s+\[${ONLINE_DOCUMENTATION}\]
  ^JUNOS\s+Services\s+AACL\s+Container\s+package\s+\[${SERVICES_AACL_CONTAINER_PACKAGE}\]
  ^JUNOS\s+Services\s+Application\s+Level\s+Gateways\s+\[${SERVICES_APPLICATION_LEVEL_GATEWAYS}\]
  ^JUNOS\s+AppId\s+Services\s+\[${APPID_SERVICES}\]
  ^JUNOS\s+Border\s+Gateway\s+Function\s+package\s+\[${BORDER_GATEWAY_FUNCTION_PACKAGE}\]
  ^JUNOS\s+Services\s+Captive\s+Portal\s+and\s+Content\s+Delivery\s+Container\s+package\s+\[${SERVICES_CAPTIVE_PORTAL_CONTENT_DELIVERY_PACKAGE}\]
  ^JUNOS\s+Services\s+HTTP\s+Content\s+Management\s+package\s+\[${SERVICES_HTTP_CONTENT_MANAGEMENT_PACKAGE}\]
  ^JUNOS\s+IDP\s+Services\s+\[${IDP_SERVICES}\]
  ^JUNOS\s+Services\s+Jflow\s+Container\s+package\s+\[${SERVICES_JFLOW_CONTAINER_PACKAGE}\]
  ^JUNOS\s+Services\s+LL-PDF\s+Container\s+package\s+\[${SERVICES_LL_PDF_CONTAINER_PACKAGE}\]
  ^JUNOS\s+Services\s+MobileNext\s+Software\s+package\s+\[${SERVICES_MOBILENEXT_SOFTWARE_PACKAGE}\]
  ^JUNOS\s+Services\s+Mobile\s+Subscriber\s+Service\s+Container\s+package\s+\[${SERVICES_MOBILE_SUBSCRIBER_SERVICE_PACKAGE}\]
  ^JUNOS\s+Services\s+NAT\s+\[${SERVICES_NAT}\]
  ^JUNOS\s+Services\s+PTSP\s+Container\s+package\s+\[${SERVICES_PTSP_CONTAINER_PACKAGE}\]
  ^JUNOS\s+Services\s+RPM\s+\[${SERVICES_RPM}\]
  ^JUNOS\s+Services\s+Stateful\s+Firewall\s+\[${SERVICES_STATEFUL_FIREWALL}\]
  ^JUNOS\s+Voice\s+Services\s+Container\s+package\s+\[${VOICE_SERVICES_CONTAINER_PACKAGE}\]
  ^JUNOS\s+Services\s+Crypto\s+\[${SERVICES_CRYPTO}\]
  ^JUNOS\s+Services\s+SSL\s+\[${SERVICES_SSL}\]
  ^JUNOS\s+Services\s+IPSec\s+\[${SERVICES_IPSEC}\]
  ^JUNOS\s+platform\s+Software\s+Suite\s+\[${PLATFORM_SOFTWARE_SUITE}\]
  ^JUNOS\s+Runtime\s+Software\s+Suite\s+\[${RUNTIME_SOFTWARE_SUITE}\]
  ^JUNOS\s+Routing\s+Software\s+Suite\s+\[${ROUTING_SOFTWARE_SUITE}\]
  ^JUNOS\s+py-base-i386\s+\[${PY_BASE_I386}\]
  ^labpkg\s+\[${LAB_PACKAGE}\]
  ^Serial\s+Number:\s+${SERIAL_NUMBER}
  ^QFabric\s+System\s+ID:\s+${QFABRIC_SYSTEM_ID}
  ^JUNOS\s+${OTHER_DEVICE_PROPERTIES}\[${OTHER_PROPERTIES_VERSIONS}\]
  ^Redis\s+\[${REDIS_VERSION}\]
  ^fpc\d+
  ^node\d+
  ^CHEF
  ^[Jj][Uu][Nn][Oo][Ss]\s+
  ^\-+
  ^{master:\d+}
  ^{primary:\S+}
  ^\s*$$
  ^. -> Error
`
