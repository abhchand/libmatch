package validate

import "github.com/abhchand/libmatch/pkg/core"

var memA, memB, memC, memD core.Member
var plA, plB, plC, plD core.PreferenceList
var pt core.PreferenceTable

func setupMembers() {
	memA = core.NewMember("A")
	memB = core.NewMember("B")
	memC = core.NewMember("C")
	memD = core.NewMember("D")

	plA = core.NewPreferenceList([]*core.Member{&memB, &memC, &memD})
	plB = core.NewPreferenceList([]*core.Member{&memA, &memC, &memD})
	plC = core.NewPreferenceList([]*core.Member{&memA, &memB, &memD})
	plD = core.NewPreferenceList([]*core.Member{&memA, &memB, &memC})

	memA.SetPreferenceList(&plA)
	memB.SetPreferenceList(&plB)
	memC.SetPreferenceList(&plC)
	memD.SetPreferenceList(&plD)

	pt = core.PreferenceTable{
		"A": &memA,
		"B": &memB,
		"C": &memC,
		"D": &memD,
	}
}
