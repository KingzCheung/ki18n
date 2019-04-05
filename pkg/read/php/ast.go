package php

type Ast struct {
	Type     string `json:"type"`
	Position struct {
		StartPos  int `json:"startPos"`
		EndPos    int `json:"endPos"`
		StartLine int `json:"startLine"`
		EndLine   int `json:"endLine"`
	} `json:"position"`
	Freefloating struct {
		End []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"End"`
	} `json:"freefloating"`
	Stmts []struct {
		Type     string `json:"type"`
		Position struct {
			StartPos  int `json:"startPos"`
			EndPos    int `json:"endPos"`
			StartLine int `json:"startLine"`
			EndLine   int `json:"endLine"`
		} `json:"position"`
		Freefloating struct {
			Start []struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"Start"`
			SemiColon []struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"SemiColon"`
		} `json:"freefloating"`
		Expr struct {
			Type     string `json:"type"`
			Position struct {
				StartPos  int `json:"startPos"`
				EndPos    int `json:"endPos"`
				StartLine int `json:"startLine"`
				EndLine   int `json:"endLine"`
			} `json:"position"`
			Freefloating struct {
				ArrayPairList []struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"ArrayPairList"`
				Start []struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"Start"`
			} `json:"freefloating"`
			Items []struct {
				Type     string `json:"type"`
				Position struct {
					StartPos  int `json:"startPos"`
					EndPos    int `json:"endPos"`
					StartLine int `json:"startLine"`
					EndLine   int `json:"endLine"`
				} `json:"position,omitempty"`
				Freefloating struct {
					Start []struct {
						Type  string `json:"type"`
						Value string `json:"value"`
					} `json:"Start"`
				} `json:"freefloating,omitempty"`
				Key struct {
					Type     string `json:"type"`
					Position struct {
						StartPos  int `json:"startPos"`
						EndPos    int `json:"endPos"`
						StartLine int `json:"startLine"`
						EndLine   int `json:"endLine"`
					} `json:"position"`
					Value string `json:"Value"`
				} `json:"Key,omitempty"`
				Val struct {
					Type     string `json:"type"`
					Position struct {
						StartPos  int `json:"startPos"`
						EndPos    int `json:"endPos"`
						StartLine int `json:"startLine"`
						EndLine   int `json:"endLine"`
					} `json:"position"`
					Value string `json:"Value"`
				} `json:"Val,omitempty"`
			} `json:"Items"`
		} `json:"Expr"`
	} `json:"Stmts"`
}
