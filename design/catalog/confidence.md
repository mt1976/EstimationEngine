# **Confidence** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**Confidence** (confidence) |
|Endpoint 	    |**/Confidence...** [^1]|
|Endpoint Query |**ConfidenceID**|
|REST API|**/API/Confidence/**|
Glyph|**fas fa-thumbs-up** ()
Friendly Name|**Confidence**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Confidence/ConfidenceList) [Exportable]
* **View** (/Confidence/ConfidenceView)
* **Edit** (/Confidence/ConfidenceEdit)
* **Save** (/Confidence/ConfidenceSave)
* **New** (/Confidence/ConfidenceNew)
* **Delete** (/Confidence/ConfidenceDelete)







##  Provides
 * Lookup (Code Name)

* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **confidenceStore**
SQL Table Key | **confidenceID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**ConfidenceID**|String|false|true|false|true|||||H|confidenceID||true|false|false|text||
|**Code**|String|true|true|false|false|||||Y|code||false|false|false|text||
|**Name**|String|true|true|false|false|||||Y|name||false|false|false|text||
|**Perc**|String|true|true|false|false|||||Y|perc||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**Notes**|String|false|true|false|false|||||Y|notes||false|false|false|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**Comments**|String|false|true|false|false|||||Y|comments||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/confidence_core.go |
| code | **adaptor** | /dao/confidence_adaptor.go_template |
| code | **validation** | /dao/confidence_validation.go_template |
| code | **api** | /application/confidence_api.go |
| code | **dao** | /dao/confidence_core.go |
| code | **datamodel** | /datamodel/confidence_core.go |
| code | **menu** | /design/menu/confidence.json |
| html | **list** | /ConfidenceList.html |
| html | **view** | /ConfidenceView.html |
| html | **edit** | /ConfidenceEdit.html |
| html | **new** | /ConfidenceNew.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **10/03/2023** at **19:54:30**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

---
### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
    * ∀ = This lookup has a filter that can be defined in the Data Object
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field