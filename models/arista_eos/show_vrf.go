package arista_eos

type ShowVrf struct {
	Interfaces []string `json:"INTERFACES"`
	Vrf        string   `json:"VRF"`
	Rd         string   `json:"RD"`
}

var ShowVrf_Template string = `Value VRF (\S+)
Value RD (\d\S+|<.+>)
Value List INTERFACES ([\w\./-]+)

Start
  ^\s+V[rR][fF]\s+RD\s+Protocols\s+State\s+Interfaces -> VRF
  ^Maximum
  ^\s*$$
  ^. -> Error

VRF
  # match a vrf with interfaces
  # key on the first line of a VRF
  ^\s+\S+\s+(\d\S+|<.+>)\s+\S+\s+(\S+:?\S+(?:\s\S+)*).*$$ -> Continue.Record
  #
  ################# First lines of VRF ###################
  #
  # match first line of vrf, first interface and it is the last interface
  ^\s+${VRF}\s+${RD}\s+\S+\s+\S+:\S+(\s\S+)*,?\s+${INTERFACES}\s*$$ -> Continue
  #
  # match first line of a vrf, first interface and it is not the last interface
  ^\s+${VRF}\s+${RD}\s+\S+\s+\S+:\S+(\s\S+)*,\s+${INTERFACES}, -> Continue
  #
  # match first line of a vrf, second interface and it is the last interface
  # two interfaces displayed per line
  ^\s+\S+\s+\S+(\s\S+)*\s+\S+\s+\S+:\S+(\s\S+)*,?\s+(?:\S+,\s){1}${INTERFACES}\s*$$ -> Continue
  #
  # match first line of a vrf, third interface and it is the last interface
  # three interfaces displayed per line
  ^\s+\S+\s+\S+(\s\S+)*\s+\S+\s+\S+:\S+(\s\S+)*,?\s+(?:\S+,\s){2}${INTERFACES}\s*$$ -> Continue
  #
  # match first line of a vrf, second interface when there are more than two interfaces
  # two or three interfaces displayed per line
  ^\s+\S+\s+\S+(\s\S+)*\s+\S+\s+\S+:\S+(\s\S+)*,?\s+(?:\S+,\s){1}${INTERFACES}, -> Continue
  #
  # match first line of a vrf, third interface when there are more than three Interfaces
  # three interfaces displayed per line
  ^\s+\S+\s+\S+(\s\S+)*\s+\S+\s+\S+:\S+(\s\S+)*,?\s+(?:\S+,\s){2}${INTERFACES}, -> Continue
  #
  ############## Lines that contains state ##################
  #
  # match state only line
  ^\s+(\S+:\S+(?:\s\S+)*)\s*$$
  #
  # match first interface of the line when it also has a state and it is the last interface
  ^\s+(\S+:\S+(?:\s\S+)*)\s+${INTERFACES}?\s*$$ -> Continue
  #
  # match first interface of the line when it also has a state and it is not the last interface
  ^\s+\S+:\S+(\s\S+)*\s+${INTERFACES}, -> Continue
  #
  # match second interface of the line when it also has a state and it is the last interface
  # two interfaces displayed per line
  ^\s+\S+:\S+(\s\S+)*\s+(?:\S+,\s){1}${INTERFACES}\s*$$ -> Continue
  #
  # match third interface of the line when it also has a state and it is the last interface
  # three interfaces displayed per line
  ^\s+\S+:\S+(\s\S+)*\s+(?:\S+,\s){2}${INTERFACES}\s*$$ -> Continue
  #
  # match second interface of the line when it also has a state and it is not the last interface
  # two or three interfaces displayed per line
  ^\s+\S+:\S+(\s\S+)*\s+(?:\S+,\s){1}${INTERFACES}, -> Continue
  #
  # match third interface of the line when it also has a state and it is not the last interface
  # three interfaces displayed per line
  ^\s+\S+:\S+(\s\S+)*\s+(?:\S+,\s){2}${INTERFACES}, -> Continue
  #
  ################# Lines that have whitespace preceding the interface list ####################
  #
  #
  # match first interface of the line when it only has whitespace and it is the last interface
  ^\s+${INTERFACES}\s*$$ -> Continue
  #
  # match first interface of the line when it only has whitespace and it is not the last interface
  ^\s+${INTERFACES}, -> Continue
  #
  # match second interface of the line when it only has whitespace and it is the last interface
  # two interfaces displayed per line
  ^\s+(?:\S+,\s){1}${INTERFACES}\s*$$ -> Continue
  #
  # match third interface of the line when it only has whitespace and it is the last interface
  # three interfaces displayed per line
  ^\s+(?:\S+,\s){2}${INTERFACES}\s*$$ -> Continue
  #
  # match second interface of the line when it only has whitespace and it is not the last interface
  # two or three interfaces displayed per line
  ^\s+(?:\S+,\s){1}${INTERFACES}, -> Continue
  #
  # match third interface of the line when it only has whitespace and it is not the last interface
  # three interfaces displayed per line
  ^\s+(?:\S+,\s){2}${INTERFACES}, -> Continue
  #
  ################## For vrfs without interfaces ###################
  #
  # match vrf line with no interfaces
  ^\s+${VRF}\s+${RD}\s+\S+\s+\S+:\S+(\s\S+)*,?\s+$$`
