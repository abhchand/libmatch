package core

var memA, memB, memC, memD Member
var plA, plB, plC, plD PreferenceList
var pt PreferenceTable

func setupMembers() {
	memA = Member{name: "A"}
	memB = Member{name: "B"}
	memC = Member{name: "C"}
	memD = Member{name: "D"}

	plA = PreferenceList{members: []*Member{&memB, &memC, &memD}}
	plB = PreferenceList{members: []*Member{&memA, &memC, &memD}}
	plC = PreferenceList{members: []*Member{&memA, &memB, &memD}}
	plD = PreferenceList{members: []*Member{&memA, &memB, &memC}}

	memA.SetPreferenceList(&plA)
	memB.SetPreferenceList(&plB)
	memC.SetPreferenceList(&plC)
	memD.SetPreferenceList(&plD)

	pt = PreferenceTable{
		"A": &memA,
		"B": &memB,
		"C": &memC,
		"D": &memD,
	}
}
