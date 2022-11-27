package datamodel

//AppMenuItem holds the menu data for a specfic login instance
type AppMenuItem struct {
	MenuHeaderText string        `json:"MenuHeaderText"`
	SectionTrees   []SectionTree `json:"SectionTrees"`
}
type SectionTree struct {
	SectionID    string        `json:"SectionID"`
	SectionName  string        `json:"SectionName"`
	SectionNodes []SectionNode `json:"SectionNodes"`
}

type SectionNode struct {
	MenuGlyph     string `json:"MenuGlyph"`
	MenuHREF      string `json:"MenuHREF"`
	MenuID        string `json:"MenuID"`
	MenuOnClick   string `json:"MenuOnClick"`
	MenuText      string `json:"MenuText"`
	MenuTextClass string `json:"MenuTextClass"`
	MenuBreak     string `json:"MenuBreak"`
}
