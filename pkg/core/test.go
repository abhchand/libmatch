package core

var memA, memB, memC, memD, memK, memL, memM Member
var plA, plB, plC, plD, plK, plL, plM PreferenceList
var pt, ptA, ptB PreferenceTable

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

func setupDoubleTable() {
	memA = Member{name: "A"}
	memB = Member{name: "B"}
	memC = Member{name: "C"}

	memK = Member{name: "K"}
	memL = Member{name: "L"}
	memM = Member{name: "M"}

	plA = PreferenceList{members: []*Member{&memK, &memL, &memM}}
	plB = PreferenceList{members: []*Member{&memL, &memM, &memK}}
	plC = PreferenceList{members: []*Member{&memM, &memL, &memK}}

	plK = PreferenceList{members: []*Member{&memB, &memC, &memA}}
	plL = PreferenceList{members: []*Member{&memA, &memC, &memB}}
	plM = PreferenceList{members: []*Member{&memA, &memB, &memC}}

	memA.SetPreferenceList(&plA)
	memB.SetPreferenceList(&plB)
	memC.SetPreferenceList(&plC)

	memK.SetPreferenceList(&plK)
	memL.SetPreferenceList(&plL)
	memM.SetPreferenceList(&plM)

	ptA = PreferenceTable{
		"A": &memA,
		"B": &memB,
		"C": &memC,
	}

	ptB = PreferenceTable{
		"K": &memK,
		"L": &memL,
		"M": &memM,
	}
}
