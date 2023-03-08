# **EstimationSession** - Object Definition
##  Information
| Information  | Value  |
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
* **Delete** (/EstimationSession/EstimationSessionDelete)







##  Provides
 * Lookup (EstimationSessionID Name)

* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **estimationSessionStore**
SQL Table Key | **estimationSessionID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**EstimationSessionID**|String|false|true|false|true|||||H|estimationSessionID||true|true|false|text||
|**ProjectID**|String|false|true|false|false|OL|Project|Project_ProjectID|Project_Name|Y|projectID||false|false|false|text||
|**EstimationStateID**|String|false|true|false|true|OL|EstimationState|EstimationState_EstimationStateID||Y|estimationStateID||false|true|false|text||
|**Activity**|String|true|true|false|false|||||Y|activity||false|false|false|text||
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
|**Name**|String|true|true|false|true|||||Y|name||true|false|false|text||
|**AdoID**|String|false|true|false|true|||||Y|adoID||false|true|false|text||
|**FreshdeskID**|String|false|true|false|true|||||Y|freshdeskID||false|true|false|text||
|**TrackerID**|String|false|true|false|true|||||Y|trackerID||false|true|false|text||
|**EstRef**|String|false|true|false|false|||||Y|estRef||false|false|false|text||
|**ExtRef**|String|false|true|false|false|||||Y|extRef||false|false|false|text||
|**SYSActivity**|String|false|true|false|false|||||NH|_activity||false|false|true|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**Comments**|String|false|true|false|true|||||Y|comments||false|false|false|textarea||
|**ProjectManager**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|projectManager||false|false|false|text||
|**ProductManager**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|productManager||false|false|false|text||
|**Approver**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|approver||false|true|false|text||
|**IssueDate**|String|false|true|false|true|||||Y|IssueDate||false|true|false|date|yyyy-mm-dd|
|**ExpiryDate**|String|false|true|false|true|||||Y|ExpiryDate||false|true|false|date|yyyy-mm-dd|
|**Origin**|String|false|false|true|false|||||N|||false|true|false|text||
|**OriginStateID**|String|false|false|true|false|||||N|||false|true|false|text||
|**OriginState**|String|false|false|true|false|||||N|||false|true|false|text||
|**OriginDocTypeID**|String|false|false|true|false|||||N|||false|true|false|text||
|**OriginDocType**|String|false|false|true|false|||||N|||false|true|false|text||
|**OriginCode**|String|false|false|true|false|||||N|||false|true|false|text||
|**OriginName**|String|false|false|true|false|||||N|||false|true|false|text||
|**OriginRate**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectProfileID**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectProfile**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectDefaultReleases**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectDefaultReleaseHours**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectBlendedRate**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectStateID**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectState**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectName**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectStartDate**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProjectEndDate**|String|false|false|true|false|||||N|||false|true|false|text||
|**ProfileSupportUpliftPerc**|String|false|false|true|false|||||N|||false|true|false|text||
|**CCY**|String|false|false|true|false|||||N|||false|true|false|text||
|**CCYCode**|String|false|false|true|false|||||N|||false|true|false|text||
|**EffortTotal**|String|false|false|true|false|||||N|||false|true|false|text||
|**FreshDeskURI**|String|false|false|true|false|||||N|||false|true|false|text||
|**ADOURI**|String|false|false|true|false|||||N|||false|true|false|text||
|**NoActiveFeatures**|String|false|false|true|false|||||N||0|false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/estimationSession_core.go |
| code | **adaptor** | /dao/estimationSession_adaptor.go_template |
| code | **validation** | /dao/estimationSession_validation.go_template |
| code | **api** | /application/estimationSession_api.go |
| code | **dao** | /dao/estimationSession_core.go |
| code | **datamodel** | /datamodel/estimationSession_core.go |
| code | **job** | /jobs/estimationSession_core.go |
| code | **menu** | /design/menu/estimationSession.json |
| html | **list** | /EstimationSessionList.html |
| html | **view** | /EstimationSessionView.html |
| html | **edit** | /EstimationSessionEdit.html |
| html | **new** | /EstimationSessionNew.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **06/03/2023** at **17:13:45**
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