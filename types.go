package twentythreeandme

type TwentyThreeAndMe struct {
	Token string   `json:"token"`
	Scope []string `json:"scope"`
}

type GeneMarker struct {
	ID           string     `json:"id"`
	AlternateIds []string   `json:"alternate_ids"`
	GeneNames    []string   `json:"gene_names"`
	AccessionID  string     `json:"accession_id"`
	Start        int        `json:"start"`
	End          int        `json:"end"`
	IsGenotyped  bool       `json:"is_genotyped"`
	IsAssayed    bool       `json:"is_assayed"`
	IsNoCall     bool       `json:"is_no_call"`
	Variants     *[]Variant `json:"variants"`
}

type Variant struct {
	AccessionID    string   `json:"accession_id"`
	Start          int      `json:"start"`
	End            int      `json:"end"`
	Allele         string   `json:"allele"`
	PlatformLabels []string `json:"platform_labels"`
	Dosage         int      `json:"dosage"`
	IsAssayed      bool     `json:"is_assayed"`
	IsNoCall       bool     `json:"is_no_call"`
}
