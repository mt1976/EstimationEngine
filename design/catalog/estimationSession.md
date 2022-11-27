# **EstimationSession** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**EstimationSession** (estimationsession) |
|Endpoint 	    |**/EstimationSession...** [^1]|
|Endpoint Query |**EstimationSessionID**|
|REST API|**/API/EstimationSession/**|
Glyph|**fas fa-calculator** ()
Friendly Name|**Estimation Session**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/EstimationSession/EstimationSessionList) [Exportable]
* **View** (/EstimationSession/EstimationSessionView)
* **Edit** (/EstimationSession/EstimationSessionEdit)
* **Save** (/EstimationSession/EstimationSessionSave)
* **New** (/EstimationSession/EstimationSessionNew)








##  Provides
 * Lookup (EstimationSessionID Name)

* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **estimationSessionStore**
SQL Table Key | **estimationSessionID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**EstimationSessionID**|String|false|true|false|true|||||H|estimationSessionID||true|false|false|text||
|**ProjectID**|String|false|true|false|false|OL|Project|Project_ProjectID|Project_Name|Y|projectID||false|false|false|text||
|**EstimationStateID**|String|false|true|false|false|OL|EstimationState|EstimationState_EstimationStateID||Y|estimationStateID||false|false|false|text||
|**Notes**|String|false|true|false|true|||||Y|notes||false|false|false|textarea||
|**Releases**|String|false|true|false|false|||||Y|releases||false|false|false|text||
|**Total**|String|false|true|false|false|||||Y|total||false|false|false|text||
|**Contingency**|String|false|true|false|false|||||Y|contingency||false|false|false|text||
|**ReqDays**|String|false|true|false|false|||||Y|reqDays||false|false|false|text||
|**RegCost**|String|false|true|false|false|||||Y|regCost||false|false|false|text||
|**ImpDays**|String|false|true|false|false|||||Y|impDays||false|false|false|text||
|**ImpCost**|String|false|true|false|false|||||Y|impCost||false|false|false|text||
|**UatDays**|String|false|true|false|false|||||Y|uatDays||false|false|false|text||
|**UatCost**|String|false|true|false|false|||||Y|uatCost||false|false|false|text||
|**MgtDays**|String|false|true|false|false|||||Y|mgtDays||false|false|false|text||
|**MgtCost**|String|false|true|false|false|||||Y|mgtCost||false|false|false|text||
|**RelDays**|String|false|true|false|false|||||Y|relDays||false|false|false|text||
|**RelCost**|String|false|true|false|false|||||Y|relCost||false|false|false|text||
|**SupportUplift**|String|false|true|false|false|||||Y|supportUplift||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**Name**|String|false|true|false|false|||||Y|name||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/estimationSession_core.go |
| code | **adaptor** | /adaptor/estimationSession_impl.go_template |
| code | **api** | /application/estimationSession_api.go |
| code | **dao** | /dao/estimationSession_core.go |
| code | **datamodel** | /datamodel/estimationSession_core.go |
| code | **menu** | /design/menu/estimationSession.json |
| html | **list** | /EstimationSession_List.html |
| html | **view** | /EstimationSession_View.html |
| html | **edit** | /EstimationSession_Edit.html |
| html | **new** | /EstimationSession_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **27/11/2022** at **20:46:13**
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