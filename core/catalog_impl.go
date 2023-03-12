package core

import (
	logs "github.com/mt1976/ebEstimates/logs"
)

var API_VERSION = "1.0.0"

//var Catalog []CatalogItem

type CatalogItem struct {
	ID     string `json:"id"`
	Path   string `json:"path"`
	Descr  string `json:"descr"`
	Query  string `json:"query"`
	Source string `json:"source"`
}

type Catalog struct {
	Items []CatalogItem
}

func (c Catalog) Generate() Catalog {
	logs.Warning("Generating Catalog")
	//items := make([]string, 0)
	API := Catalog{Items: []CatalogItem{}}
	return API
}

func (c Catalog) AddRoute(id string, path string, descr string, query string, src string) Catalog {
	//logs.Information("Adding Route: ", id+" "+path+" "+descr+" "+query+" "+src)
	var catalogItem = CatalogItem{ID: id, Path: path, Descr: descr, Query: query, Source: src}
	c.Items = append(c.Items, catalogItem)
	//	fmt.Printf("c: %v\n", c)
	//c.ListRoutes()
	//fmt.Printf("len(c.Items): %v\n", len(c.Items))

	//	fmt.Printf("c.Items[0].ID: %v\n", c.Items[0].ID)
	//	fmt.Printf("c.Items[1].ID: %v\n", c.Items[1].ID)
	return c
}

func (c Catalog) ListRoutes() {
	//xx := len(c.Items)
	//fmt.Printf("no items: %v\n", xx)
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
