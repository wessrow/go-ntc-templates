package models

type CiscoNxosShowNveVniInterfaceNve1Detail struct {
	Vni	string	`json:"VNI"`
	Nve_interface	string	`json:"NVE_INTERFACE"`
	Multicast_address	string	`json:"MULTICAST_ADDRESS"`
	Vni_state	string	`json:"VNI_STATE"`
	Mode	string	`json:"MODE"`
	Vni_type	string	`json:"VNI_TYPE"`
	Vni_flags	string	`json:"VNI_FLAGS"`
	Dci_multicast_address	string	`json:"DCI_MULTICAST_ADDRESS"`
	Provision_state	string	`json:"PROVISION_STATE"`
	Vlan_bd	string	`json:"VLAN_BD"`
	Svi_state	string	`json:"SVI_STATE"`
}

var CiscoNxosShowNveVniInterfaceNve1Detail_Template = `Value VNI (\d+)
Value NVE_INTERFACE (\S+)
Value MULTICAST_ADDRESS (\S+)
Value VNI_STATE (\S+)
Value MODE (\S+)
Value VNI_TYPE (\S+(?:\s+\[\S+\])?)
Value VNI_FLAGS (\S*)
Value DCI_MULTICAST_ADDRESS (\S*)
Value PROVISION_STATE (\S+)
Value VLAN_BD (\d+)
Value SVI_STATE (.*)

Start
  ^$$ -> Start
  ^VNI:\s+${VNI}
  ^\s+NVE-Interface\s+:\s+${NVE_INTERFACE}
  ^\s+Mcast-Addr\s+:\s+${MULTICAST_ADDRESS}
  ^\s+VNI State\s+:\s+${VNI_STATE}
  ^\s+Mode\s+:\s+${MODE}
  ^\s+VNI Type\s+:\s+${VNI_TYPE}
  ^\s+VNI Flags\s+:\s*${VNI_FLAGS}
  ^\s+DCI Mcast-Addr\s+:\s*${DCI_MULTICAST_ADDRESS}
  ^\s+Provision State\s+:\s+${PROVISION_STATE}
  ^\s+Vlan-BD\s+:\s+${VLAN_BD}
  ^\s+SVI State\s+:\s+${SVI_STATE} -> Record
  ^VNI:\s+\d+ -> Start
  ^.* -> Error

`