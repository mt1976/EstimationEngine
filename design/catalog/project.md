# **Project** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Project** (project) |
|Endpoint 	    |**/Project...** [^1]|
|Endpoint Query |**ProjectID**|
|REST API|**/API/Project/**|
Glyph|**fas fa-project-diagram** ()
Friendly Name|**Project**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Project/ProjectList) [Exportable]
* **View** (/Project/ProjectView)
* **Edit** (/Project/ProjectEdit)
* **Save** (/Project/ProjectSave)
* **New** (/Project/ProjectNew)








##  Provides
 * Lookup (ProjectID Name)

* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **projectStore**
SQL Table Key | **projectID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**ProjectID**|String|false|true|false|true|||||H|projectID||true|false|false|text||
|**OriginID**|String|true|true|false|false|OL|Origin|Origin_OriginID|Origin_FullName|Y|originID||true|false|false|text||
|**ProjectStateID**|String|false|true|false|false|OL|ProjectState|ProjectState_ProjectStateID|ProjectState_Name|Y|projectStateID||false|false|false|text||
|**ProfileID**|String|false|true|false|false|OL|Profile|Profile_ProfileID|Profile_Name|Y|profileID||false|false|false|text||
|**Name**|String|true|true|false|false|||||Y|name||false|false|false|text||
|**Description**|String|false|true|false|true|||||Y|description||false|false|false|textarea||
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
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**Comments**|String|false|true|false|true|||||Y|comments||false|false|false|textarea||
|**ProjectRate**|String|false|true|false|true|||||Y|projectRate||false|false|false|number||
|**DefaultRate**|String|false|true|false|true|||||N|defaultRate||false|false|false|number||
|**ProjectAnalyst**|String|false|true|false|false|OL|Resource|Resource_Code|Resource_Name|Y|projectAnalyst||false|false|false|text||
|**ProjectEngineer**|String|false|true|false|false|OL|Resource|Resource_Code|Resource_Name|Y|projectEngineer||false|false|false|text||
|**ProjectManager**|String|false|true|false|false|OL|Resource|Resource_Code|Resource_Name|Y|projectManager||false|false|false|text||
|**Releases**|String|false|true|false|false|||||Y|releases||false|false|false|text||
|**Notes**|String|false|true|false|true|||||N|notes||false|false|false|textarea||
|**NoEstimationSessions**|String|false|false|true|false|||||N||0|false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/project_core.go |
| code | **adaptor** | /dao/project_adaptor.go_template |
| code | **api** | /application/project_api.go |
| code | **dao** | /dao/project_core.go |
| code | **datamodel** | /datamodel/project_core.go |
| code | **job** | /jobs/project_core.go |
| code | **menu** | /design/menu/project.json |
| html | **list** | /Project_List.html |
| html | **view** | /Project_View.html |
| html | **edit** | /Project_Edit.html |
| html | **new** | /Project_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **07/01/2023** at **23:01:30**
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