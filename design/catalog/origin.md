# **Origin** - Object Definition
##  Information
| Information  | Value  |
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
* **Delete** (/Origin/OriginDelete)







##  Provides
 * Lookup (Code FullName)

* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **originStore**
SQL Table Key | **originID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**OriginID**|String|false|true|false|true|||||H|originID||true|false|false|text||
|**StateID**|String|false|true|false|false|OL∀|OriginState|OriginState_OriginStateID|OriginState_Name|Y|stateID||false|true|false|text||
|**DocTypeID**|String|false|true|false|false|OL∀|DocType|DocType_DocTypeID||Y|docTypeID||false|false|false|text||
|**Code**|String|false|true|false|true|||||H|code||true|true|false|text||
|**FullName**|String|false|true|false|true|||||Y|fullName||false|true|false|text||
|**Rate**|String|false|true|false|true|||||Y|rate||false|true|false|text||
|**Notes**|String|false|true|false|true|||||Y|notes||false|false|false|textarea||
|**StartDate**|String|false|true|false|true|||||Y|startDate||false|true|false|date|yyyy-mm-dd|
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
|**ProjectManager**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|projectManager||false|false|false|text||
|**AccountManager**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|accountManager||false|false|false|text||
|**NoActiveProjects**|String|false|false|true|true|||||N|||false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/origin_core.go |
| code | **adaptor** | /dao/origin_adaptor.go_template |
| code | **validation** | /dao/origin_validation.go_template |
| code | **api** | /application/origin_api.go |
| code | **dao** | /dao/origin_core.go |
| code | **datamodel** | /datamodel/origin_core.go |
| code | **job** | /jobs/origin_core.go |
| code | **menu** | /design/menu/origin.json |
| html | **list** | /OriginList.html |
| html | **view** | /OriginView.html |
| html | **edit** | /OriginEdit.html |
| html | **new** | /OriginNew.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **08/03/2023** at **18:42:24**
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