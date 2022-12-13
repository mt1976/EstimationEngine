# **UserRole** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**UserRole** (userrole) |
|Endpoint 	    |**/UserRole...** [^1]|
|Endpoint Query |**Id**|
|REST API|**/API/UserRole/**|
Glyph|**fas fa-users-cog** ()
Friendly Name|**User Roles**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/UserRole/UserRoleList) [Exportable]
* **View** (/UserRole/UserRoleView)
* **Edit** (/UserRole/UserRoleEdit)
* **Save** (/UserRole/UserRoleSave)
* **New** (/UserRole/UserRoleNew)
* **Delete** (/UserRole/UserRoleDelete)







##  Provides
 * Lookup (Id Name)

* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **roleStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|true|true|false|true|||||Y|Id||true|false|false|text||
|**Name**|String|true|true|false|false|||||Y|Name||false|false|false|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/userRole_core.go |
| code | **api** | /application/userRole_api.go |
| code | **dao** | /dao/userRole_core.go |
| code | **datamodel** | /datamodel/userRole_core.go |
| code | **menu** | /design/menu/userRole.json |
| html | **list** | /UserRole_List.html |
| html | **view** | /UserRole_View.html |
| html | **edit** | /UserRole_Edit.html |
| html | **new** | /UserRole_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **11/12/2022** at **19:24:03**
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