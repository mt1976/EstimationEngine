package core

import logs "github.com/mt1976/ebEstimates/logs"

var API_VERSION = "1.0.0"
var Catalog []CatalogItem

type CatalogItem struct {
	ID     string `json:"id"`
	Path   string `json:"path"`
	Descr  string `json:"descr"`
	Query  string `json:"query"`
	Source string `json:"source"`
}

func Catalog_Add(id string, path string, descr string, query string, src string) {
	var catalogItem = CatalogItem{ID: id, Path: path, Descr: descr, Query: query, Source: src}
	Catalog = append(Catalog, catalogItem)
}

func Catalog_List() {
	for _, k := range Catalog {
		logs.Information("Catalog Item", k.ID+" - "+k.Path+" - "+k.Descr+" - "+k.Query+" - "+k.Source)
	}
}

type ContentList struct {
	Count int               `json:"count"`
	Key   string            `json:"key"`
	Items []ContentListItem `json:"items"`
}

type ContentListItem struct {
	ID    string `json:"id"`
	Query string `json:"query"`
}
