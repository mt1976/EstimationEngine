# **Index** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**Index** (index) |
|Endpoint 	    |**/Index...** [^1]|
|Endpoint Query |**IndexID**|
|REST API|**/API/Index/**|
Glyph|**fas fa-search** ()
Friendly Name|**Index**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Index/IndexList) [Exportable]
* **View** (/Index/IndexView)

* **Save** (/Index/IndexSave)









##  Provides


* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **indexStore**
SQL Table Key | **indexID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**IndexID**|String|false|true|false|true|||||NH|indexID||false|true|false|text||
|**KeyClass**|String|false|true|false|true|||||N|keyClass||false|false|false|text||
|**KeyName**|String|false|true|false|true|||||N|keyName||false|false|false|text||
|**KeyID**|String|false|true|false|true|||||N|keyID||false|false|false|text||
|**Link**|String|false|true|false|true|||||NH|link||false|true|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**KeyValue**|String|false|true|false|false|||||Y|KeyValue||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/index_core.go |
| code | **adaptor** | /dao/index_adaptor.go_template |
| code | **api** | /application/index_api.go |
| code | **dao** | /dao/index_core.go |
| code | **datamodel** | /datamodel/index_core.go |
| code | **job** | /jobs/index_core.go |
| code | **menu** | /design/menu/index.json |
| html | **list** | /Index_List.html |
| html | **view** | /Index_View.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **07/02/2023** at **18:52:37**
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