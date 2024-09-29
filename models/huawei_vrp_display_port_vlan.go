package models

type HuaweiVrpDisplayPortVlan struct {
	Interface	string	`json:"INTERFACE"`
	Link_type	string	`json:"LINK_TYPE"`
	Pvid	string	`json:"PVID"`
	Trunk_vlan_list	[]string	`json:"TRUNK_VLAN_LIST"`
}

var HuaweiVrpDisplayPortVlan_Template = `Value Required INTERFACE ([\w\.\/\-]+)
Value LINK_TYPE (trunk|access|auto|hybrid|desirable)
Value PVID (\d+)
Value List TRUNK_VLAN_LIST (\d+\-\d+|\d+)


Start
  ^\s*$$
  ^[pPoOrRtT]{4}.*$$ -> VLANS
  ^. -> Error

VLANS
  ^[\w\.\/\-]+ -> Continue.Record
  ^${INTERFACE}\s+${LINK_TYPE}\s+${PVID}\s+${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){3}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){4}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){5}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){6}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){7}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){8}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){9}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){10}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){11}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){12}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){13}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){14}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){15}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){16}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){17}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){18}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){19}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){20}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){21}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){22}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){23}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){24}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){25}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){26}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){27}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){28}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){29}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){30}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){31}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){32}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){33}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){34}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){35}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){36}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){37}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){38}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){39}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){40}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){41}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){42}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){43}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){44}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){45}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){46}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){47}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){48}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){49}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){50}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){51}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){52}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){53}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){54}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){55}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){56}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){57}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){58}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){59}${TRUNK_VLAN_LIST}\s* -> Continue
  ^[\w\.\/\-]+\s+(?:\S+\s+){60}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}\S+\s+${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){2}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){3}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){4}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){5}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){6}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){7}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){8}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){9}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){10}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){11}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){12}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){13}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){14}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){15}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){16}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){17}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){18}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){19}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){20}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){21}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){22}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){23}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){24}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){25}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){26}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){27}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){28}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){29}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){30}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){31}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){32}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){33}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){34}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){35}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){36}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){37}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){38}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){39}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){40}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){41}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){42}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){43}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){44}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){45}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){46}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){47}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){48}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){49}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){50}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){51}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){52}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){53}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){54}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){55}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){56}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){57}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){58}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){59}${TRUNK_VLAN_LIST}\s* -> Continue
  ^\s{6,}(?:\S+\s+){60}${TRUNK_VLAN_LIST}\s* -> Continue
  # Bundled interface (PVID is 0)
  ^${INTERFACE}\s+-\s+0\s+-\s*$$ -> Record
  # Access port (PVID is empty)
  ^${INTERFACE}\s+${LINK_TYPE}\s+${PVID}\s+-\s*$$ -> Record
  # Interface lists
  ^[\w\.\/\-]+\s+\S+\s+\d+\s+[0-9\- ]+$$
  ^\s{6,}
  ^-+$$
  ^. -> Error

`