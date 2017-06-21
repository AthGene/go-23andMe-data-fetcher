package twentythreeandme

import (
	"fmt"
	"sync"
	"testing"
)

func Test_getGeneMarker(t *testing.T) {
	var wg sync.WaitGroup
	geneMarker := &GeneMarker{}
	err := getGeneMarker("rs1815739", "demo_oauth_token", geneMarker, &wg)

	if err != nil {
		t.Fatal("Request failed")
	}

	fmt.Printf("%#v", geneMarker)
}

func Test_GetTwentyThreeAndMeData(t *testing.T) {

	ttam := &TwentyThreeAndMe{
		Token: "demo_oauth_token",
		Scope: []string{
			"rs7781370", "rs10048146", "rs1430742", "rs1054627", "rs3755955", "rs1007738",
			"rs2941740", "rs1038304", "rs1999805", "rs2504063", "rs4870044", "rs7776725",
			"rs4820268", "rs4263839", "rs6478108", "rs6478109", "rs1800629", "rs9533090",
			"rs6495122", "rs762551", "rs4410790", "rs1801725", "rs1799971", "rs4680",
			"rs2395182", "rs7775228", "rs2187668", "rs970547", "rs1800562", "rs1799945",
			"rs2073711", "rs143383", "rs2253206", "rs6822844", "rs13119723", "rs7454108",
			"rs182549", "rs4988235", "rs2349775", "rs242924", "rs7209436", "rs806378",
			"rs17070145", "rs4908449", "rs35753505", "rs950809", "rs10884402", "rs1799990",
			"rs1801260", "rs1800169", "rs7832552", "rs2267668", "rs10887741", "rs12612420",
			"rs6258", "rs1799941", "rs727428", "rs722208", "rs5934505", "rs5400", "rs17822931",
			"rs4654748", "rs1805087", "rs1801394", "rs1801131", "rs1801133", "rs12150660",
			"rs11091046", "rs731236", "rs1544410", "rs7041", "rs2282679", "rs602662",
			"rs9362667", "rs15705", "rs11549465", "rs1870377", "rs2070744", "rs8192678",
			"rs12970134", "rs2472297", "rs1800497", "rs3135506", "rs17300539", "rs3736228",
			"rs12934922", "rs7501331", "rs1229984", "rs671", "rs17151919", "rs17782313",
			"rs780094", "rs7034200", "rs11708067", "rs174550", "rs7944584", "rs599083",
			"rs560887", "rs4607517", "rs10885122", "rs2191349", "rs13266634", "rs11605924",
			"rs7903146", "rs12255372", "rs7221412", "rs33972313", "rs11950646", "rs10830963",
			"rs1042714", "rs4994", "rs699", "rs5082", "rs1800012", "rs9939609", "rs1801282",
			"rs1815739", "rs1042713", "rs7117858", "rs884205", "rs7521902", "rs1049434",
			"rs2908004", "rs6696981", "rs7524102", "rs10166942", "rs11172113", "rs1801222",
			"rs838133", "rs3923809", "rs1026732", "rs1975197", "rs2300478", "rs9357271",
			"rs9296249", "rs1800588", "rs6420424", "rs12272004", "rs964184", "rs2108622",
			"rs11057830", "rs662799", "rs11898505", "rs10876432", "rs9466056", "rs9594759",
			"rs1021188", "rs3018362", "rs6993813", "rs6469804", "rs4355801", "rs1366594",
		},
	}

	geneMarkers := GetTwentyThreeAndMeData(ttam)

	// for _, geneMarker := range *geneMarkers {
	// 	fmt.Printf("%#v \n", geneMarker)
	// }

	fmt.Println("Number of gene markers downloaded: ", len(*geneMarkers))
}
