# **Resource** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Resource** (resource) |
|Endpoint 	    |**/Resource...** [^1]|
|Endpoint Query |**ResourceID**|
|REST API|**/API/Resource/**|
Glyph|**fas fa-cogs** ()
Friendly Name|**Resource**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Resource/ResourceList) [Exportable]
* **View** (/Resource/ResourceView)
* **Edit** (/Resource/ResourceEdit)
* **Save** (/Resource/ResourceSave)
* **New** (/Resource/ResourceNew)








##  Provides
 * Lookup (Code Name)

* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **resourceStore**
SQL Table Key | **resourceID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**ResourceID**|String|false|true|false|true|||||NH|resourceID||false|false|false|text||
|**Code**|String|true|true|false|false|||||Y|code||false|false|false|text||
|**Name**|String|true|true|false|false|||||Y|name||false|false|false|text||
|**Email**|String|true|true|false|false|||||Y|email||false|false|false|text||
|**Class**|String|false|true|false|false|LL|resourcetype|||Y|class||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**UserActive**|String|false|true|false|false|LL|tf|||Y|userActive||false|false|false|text||
|**SYSActivity**|String|false|true|false|false|||||NH|_activity||false|false|true|text||
|**Notes**|String|false|true|false|false|||||Y|notes||false|false|false|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**Comments**|String|false|true|false|false|||||Y|comments||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/resource_core.go |
| code | **adaptor** | /adaptor/resource_impl.go_template |
| code | **api** | /application/resource_api.go |
| code | **dao** | /dao/resource_core.go |
| code | **datamodel** | /datamodel/resource_core.go |
| code | **menu** | /design/menu/resource.json |
| html | **list** | /Resource_List.html |
| html | **view** | /Resource_View.html |
| html | **edit** | /Resource_Edit.html |
| html | **new** | /Resource_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **11/12/2022** at **19:24:02**
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