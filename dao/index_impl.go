package dao

import (
	"errors"
	"fmt"

	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
	"github.com/tj/go-spin"
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
	iRec.Link, _ = Data_Get(iRec.KeyClass, iRec.KeyName, "IndexLink")
	iRec.Link = prefix + iRec.Link

	if iRec.Link == "" {
		iRec.Link = "/{{VALUE}}"
		logs.Warning("Indexer Content Not found " + core.DQuote(iRec.KeyClass) + " " + core.DQuote(iRec.KeyName) + " using " + core.DQuote(iRec.Link))
		Data_Put(iRec.KeyClass, iRec.KeyName, "IndexLink", iRec.Link)
	}
	//iRec.Link = fmt.Sprintf(iRec.Link, iRec.KeyID)

	iRec.Link = core.ReplaceWildcard(iRec.Link, "ID", iRec.KeyID)
	iRec.Link = core.ReplaceWildcard(iRec.Link, "VALUE", iRec.KeyValue)

	if KeyValue == "" {
		// Delete the index entry
		Index_Delete(iRec.IndexID)
		return iRec, nil
	}

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
	s := spin.New()

	// Loop through the index entries
	for _, indexEntry := range indexEntries {
		fmt.Printf("\r  \033[36mProcessing\033[m %s ", s.Next())

		_, err := Indexer_Put(indexEntry.KeyClass, indexEntry.KeyName, indexEntry.KeyID, indexEntry.KeyValue)
		if err != nil {
			return errors.New("unable to rebuild index for " + indexEntry.KeyClass + " " + indexEntry.KeyName + " " + indexEntry.KeyID + " " + indexEntry.KeyValue)
		}
	}
	return nil
}
