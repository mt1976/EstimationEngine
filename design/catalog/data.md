# **Data** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Data** (data) |
|Endpoint 	    |**/Data...** [^1]|
|Endpoint Query |**DataID**|
Glyph|**fas fa-table** ()
Friendly Name|**System Properties**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Data/DataList) [Exportable]
* **View** (/Data/DataView)
* **Edit** (/Data/DataEdit)
* **Save** (/Data/DataSave)
* **New** (/Data/DataNew)








##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **dataStore**
SQL Table Key | **dataID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**DataID**|String|false|true|false|true|||||NH|dataID||false|true|false|text||
|**Class**|String|true|true|false|true|||||Y|class||true|false|false|text||
|**Field**|String|true|true|false|true|||||Y|field||true|false|false|text||
|**Value**|String|true|true|false|false|||||Y|value||false|false|false|text||
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


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/data_core.go |
| code | **adaptor** | /dao/data_adaptor.go_template |
| code | **dao** | /dao/data_core.go |
| code | **datamodel** | /datamodel/data_core.go |
| code | **menu** | /design/menu/data.json |
| html | **list** | /Data_List.html |
| html | **view** | /Data_View.html |
| html | **edit** | /Data_Edit.html |
| html | **new** | /Data_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **07/01/2023** at **23:01:28**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field