# **Schedule** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**Schedule** (schedule) |
|Endpoint 	    |**/Schedule...** [^1]|
|Endpoint Query |**Schedule**|
|REST API|**/API/Schedule/**|
Glyph|**far fa-clock** (text-info)
Friendly Name|**Scheduler**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Schedule/ScheduleList) [Exportable]
* **View** (/Schedule/ScheduleView)

* **Save** (/Schedule/ScheduleSave)









##  Provides


* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **scheduleStore**
SQL Table Key | **id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|false|true|false|true|||||H|id||true|false|false|text||
|**Name**|String|false|true|false|false|||||Y|name||false|false|false|text||
|**Description**|String|false|true|false|false|||||Y|description||false|false|false|text||
|**Schedule**|String|false|true|false|false|||||Y|schedule||false|false|false|text||
|**Started**|String|false|true|false|false|||||Y|started||false|false|false|text||
|**Lastrun**|String|false|true|false|false|||||Y|lastrun||false|false|false|text||
|**Message**|String|false|true|false|false|||||Y|message||false|false|false|text||
|**Type**|String|false|true|false|false|||||Y|type||false|false|false|text||
|**Human**|String|false|true|false|false|||||Y|human||false|false|false|text||
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
| code | **application** | /application/schedule_core.go |
| code | **api** | /application/schedule_api.go |
| code | **dao** | /dao/schedule_core.go |
| code | **datamodel** | /datamodel/schedule_core.go |
| code | **menu** | /design/menu/schedule.json |
| html | **list** | /Schedule_List.html |
| html | **view** | /Schedule_View.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **07/02/2023** at **18:52:39**
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