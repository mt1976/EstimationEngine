package core

import logs "github.com/mt1976/ebEstimates/logs"

var API_VERSION = "1.0.0"

//var Catalog []CatalogItem

type Catalog struct {
	Items []CatalogItem
}

type CatalogItem struct {
	ID     string `json:"id"`
	Path   string `json:"path"`
	Descr  string `json:"descr"`
	Query  string `json:"query"`
	Source string `json:"source"`
}

type CatalogInterface interface {
	Add(id string, path string, descr string, query string, src string)
	List()
}

func test() {
	var c Catalog
	c.AddRoute("1", "2", "3", "4", "5")
	c.ListRoutes()
}

func (c Catalog) AddRoute(id string, path string, descr string, query string, src string) {
	var catalogItem = CatalogItem{ID: id, Path: path, Descr: descr, Query: query, Source: src}
	c.Items = append(c.Items, catalogItem)
}

func (c Catalog) ListRoutes() {
	for _, k := range c.Items {
		logs.Catalog(k.ID, k.Path, k.Query, k.Source)
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
