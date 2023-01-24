package dao

import (
	"errors"
	"fmt"

	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

// Indexer is the interface that wraps the basic Index method.
// Indexer_Put updates the index
func Indexer_Put(KeyClass string, KeyField string, KeyID string, KeyValue string) (dm.Index, error) {

	thisIndexID := KeyClass + "-" + KeyField + "-" + KeyID
	thisIndexID = core.EncodeString(thisIndexID)

	_, iRec, e := Index_GetByID(thisIndexID)
	if e != nil {
		logs.Error("Indexer_Put", e)
		return dm.Index{}, e
	}
	if iRec.IndexID == "" {
		iRec.IndexID = thisIndexID
		iRec.KeyClass = KeyClass
		iRec.KeyName = KeyField
	}

	iRec.KeyID = KeyID
	iRec.KeyValue = KeyValue

	prefix := ""
	iRec.Link, _ = Data_Get(iRec.KeyClass+"-Link", iRec.KeyName, "Setting")
	iRec.Link = prefix + iRec.Link
	iRec.LinkView, _ = Data_Get(iRec.KeyClass+"-View", iRec.KeyName, "Setting")
	iRec.LinkView = prefix + iRec.LinkView
	iRec.LinkEdit, _ = Data_Get(iRec.KeyClass+"-Edit", iRec.KeyName, "Setting")
	iRec.LinkEdit = prefix + iRec.LinkEdit

	if iRec.Link == "" {
		iRec.Link = "/%s"
		logs.Warning("Indexer Content Not found " + core.DQuote(iRec.KeyClass+"-Link") + " " + core.DQuote(iRec.KeyName) + " using " + core.DQuote(iRec.Link))
	}
	iRec.Link = fmt.Sprintf(iRec.Link, iRec.KeyID)
	if iRec.LinkView == "" {
		iRec.LinkView = "/%s"
		logs.Warning("Indexer Content Not found " + core.DQuote(iRec.KeyClass+"-View") + " " + core.DQuote(iRec.KeyName) + " using " + core.DQuote(iRec.LinkView))
	}
	iRec.LinkView = fmt.Sprintf(iRec.LinkView, iRec.KeyID)
	if iRec.LinkEdit == "" {
		iRec.LinkEdit = "/%s"
		logs.Warning("Indexer Content Not found " + core.DQuote(iRec.KeyClass+"-Edit") + " " + core.DQuote(iRec.KeyName) + " using " + core.DQuote(iRec.LinkEdit))
	}
	iRec.LinkEdit = fmt.Sprintf(iRec.LinkEdit, iRec.KeyID)

	err := Index_StoreSystem(iRec)
	if err != nil {
		return dm.Index{}, err
	}
	return iRec, nil
}

// Indexer_Rebuild rebuilds the index
func Indexer_Rebuild() error {
	// Get all list of index entries
	_, indexEntries, _ := Index_GetList()

	// Loop through the index entries
	for _, indexEntry := range indexEntries {
		_, err := Indexer_Put(indexEntry.KeyClass, indexEntry.KeyName, indexEntry.KeyID, indexEntry.KeyValue)
		if err != nil {
			return errors.New("unable to rebuild index for " + indexEntry.KeyClass + " " + indexEntry.KeyName + " " + indexEntry.KeyID + " " + indexEntry.KeyValue)
		}
	}
	return nil
}
