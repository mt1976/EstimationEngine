# **Session** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Session** (session) |
|Endpoint 	    |**/Session...** [^1]|
|Endpoint Query |**SessionID**|
Glyph|**fas fa-history** (text-info)
Friendly Name|**Session**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Session/SessionList) 
* **View** (/Session/SessionView)

* **Save** (/Session/SessionSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sessionStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|false|true|false|true|||||H|Id||true|false|false|text||
|**Apptoken**|String|true|true|false|false|||||Y|Apptoken||false|false|false|text||
|**Createdate**|String|false|true|false|false|||||Y|Createdate||false|false|false|text||
|**Createtime**|String|false|true|false|false|||||Y|Createtime||false|false|false|text||
|**Uniqueid**|String|false|true|false|false|||||Y|Uniqueid||false|false|false|text||
|**Sessiontoken**|String|false|true|false|false|||||Y|Sessiontoken||false|false|false|text||
|**Username**|String|false|true|false|false|||||Y|Username||false|false|false|text||
|**Password**|String|false|true|false|false|||||Y|Password||false|false|false|text||
|**Userip**|String|false|true|false|false|||||Y|Userip||false|false|false|text||
|**Userhost**|String|false|true|false|false|||||Y|Userhost||false|false|false|text||
|**Appip**|String|false|true|false|false|||||Y|Appip||false|false|false|text||
|**Apphost**|String|false|true|false|false|||||Y|Apphost||false|false|false|text||
|**Issued**|String|false|true|false|false|||||Y|Issued||false|false|false|text||
|**Expiry**|String|false|true|false|false|||||Y|Expiry||false|false|false|text||
|**Expiryraw**|String|false|true|false|false|||||Y|Expiryraw||false|false|false|text||
|**Expires**|Time|false|true|false|false|||||Y|Expires||false|false|false|text||
|**Brand**|String|false|true|false|false|||||Y|Brand||false|false|false|text||
|**SessionRole**|String|false|true|false|false|||||Y|SessionRole||false|false|false|text||
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
| code | **application** | /application/session_core.go |
| code | **dao** | /dao/session_core.go |
| code | **datamodel** | /datamodel/session_core.go |
| code | **menu** | /design/menu/session.json |
| html | **list** | /Session_List.html |
| html | **view** | /Session_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **03/01/2023** at **19:18:11**
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