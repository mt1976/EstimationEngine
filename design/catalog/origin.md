# **Origin** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Origin** (origin) |
|Endpoint 	    |**/Origin...** [^1]|
|Endpoint Query |**OriginID**|
|REST API|**/API/Origin/**|
Glyph|**fas fa-building** ()
Friendly Name|**Origin**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Origin/OriginList) [Exportable]
* **View** (/Origin/OriginView)
* **Edit** (/Origin/OriginEdit)
* **Save** (/Origin/OriginSave)
* **New** (/Origin/OriginNew)








##  Provides
 * Lookup (Code FullName)

* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **originStore**
SQL Table Key | **originID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**OriginID**|String|false|true|false|true|||||H|originID||true|false|false|text||
|**StateID**|String|false|true|false|false|OL|OriginState|OriginState_OriginStateID|OriginState_Name|Y|stateID||false|false|false|text||
|**DocTypeID**|String|false|true|false|false|OL|DocType|DocType_DocTypeID||Y|docTypeID||false|false|false|text||
|**Code**|String|false|true|false|true|||||Y|code||true|false|false|text||
|**FullName**|String|false|true|false|true|||||Y|fullName||false|false|false|text||
|**Rate**|String|true|true|false|false|||||Y|rate||false|false|false|text||
|**Notes**|String|false|true|false|true|||||Y|notes||false|false|false|textarea||
|**StartDate**|String|false|true|false|true|||||Y|startDate||false|false|false|date|yyyy-mm-dd|
|**EndDate**|String|false|true|false|true|||||Y|endDate||false|false|false|date|yyyy-mm-dd|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**SYSActivity**|String|false|true|false|false|||||NH|_activity||false|false|true|text||
|**Currency**|String|true|true|false|false|LL|ccy|||Y|currency||false|false|false|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**Comments**|String|false|true|false|false|||||Y|comments||false|false|false|text||
|**NoActiveProjects**|String|false|false|true|false|||||N|||false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/origin_core.go |
| code | **adaptor** | /adaptor/origin_impl.go_template |
| code | **api** | /application/origin_api.go |
| code | **dao** | /dao/origin_core.go |
| code | **datamodel** | /datamodel/origin_core.go |
| code | **job** | /jobs/origin_core.go |
| code | **menu** | /design/menu/origin.json |
| html | **list** | /Origin_List.html |
| html | **view** | /Origin_View.html |
| html | **edit** | /Origin_Edit.html |
| html | **new** | /Origin_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **11/12/2022** at **19:53:20**
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