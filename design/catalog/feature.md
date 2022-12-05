# **Feature** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Feature** (feature) |
|Endpoint 	    |**/Feature...** [^1]|
|Endpoint Query |**FeatureID**|
|REST API|**/API/Feature/**|
Glyph|**fas fa-tasks** ()
Friendly Name|**Feature**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Feature/FeatureList) [Exportable]
* **View** (/Feature/FeatureView)
* **Edit** (/Feature/FeatureEdit)
* **Save** (/Feature/FeatureSave)
* **New** (/Feature/FeatureNew)








##  Provides
 * Lookup (FeatureID Name)

* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **featureStore**
SQL Table Key | **featureID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**FeatureID**|String|false|true|false|true|||||H|featureID||true|false|false|text||
|**EstimationSessionID**|String|false|true|false|false|OL|EstimationSession|EstimationSession_EstimationSessionID|EstimationSession_Name|Y|estimationSessionID||false|false|false|text||
|**ConfidenceID**|String|false|true|false|false|OL|Confidence|Confidence_ConfidenceID||Y|confidenceID||false|false|false|text||
|**Name**|String|true|true|false|false|||||Y|name||false|false|false|text||
|**DevEstimate**|String|true|true|false|false|||||Y|devEstimate||false|false|false|text||
|**DevUplift**|String|false|true|false|false|||||Y|devUplift||false|false|false|text||
|**Reqs**|String|false|true|false|false|||||Y|reqs||false|false|false|text||
|**AnalystTest**|String|false|true|false|false|||||Y|analystTest||false|false|false|text||
|**Docs**|String|false|true|false|false|||||Y|docs||false|false|false|text||
|**Mgt**|String|false|true|false|false|||||Y|mgt||false|false|false|text||
|**UatSupport**|String|false|true|false|false|||||Y|uatSupport||false|false|false|text||
|**Marketing**|String|false|true|false|false|||||Y|marketing||false|false|false|text||
|**Contingency**|String|false|true|false|false|||||Y|contingency||false|false|false|text||
|**TrackerID**|String|false|true|false|false|||||Y|trackerID||false|false|false|text||
|**AdoID**|String|false|true|false|false|||||Y|adoID||false|false|false|text||
|**FreshdeskID**|String|false|true|false|false|||||Y|freshdeskID||false|false|false|text||
|**ExtRef**|String|false|true|false|false|||||Y|extRef||false|false|false|text||
|**ExtRef2**|String|false|true|false|false|||||Y|extRef2||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**Developer**|String|false|true|false|false|OL|Resource|Resource_Code|Resource_Name|Y|developer||false|false|false|text||
|**Approver**|String|false|true|false|false|OL|Resource|Resource_Code|Resource_Name|Y|approver||false|false|false|text||
|**Notes**|String|false|true|false|true|||||Y|notes||false|false|false|textarea||
|**OffProfile**|String|false|true|false|false|||||Y|offProfile||false|false|false|text||
|**OffProfileJustification**|String|false|true|false|false|||||Y|offProfileJustification||false|false|false|text||
|**SYSActivity**|String|false|true|false|false|||||NH|_activity||false|false|true|text||
|**DfReqs**|String|false|true|false|false|||||Y|dfReqs||false|false|false|text||
|**DfAnalystTest**|String|false|true|false|false|||||Y|dfAnalystTest||false|false|false|text||
|**DfDocs**|String|false|true|false|false|||||Y|dfDocs||false|false|false|text||
|**Dfmgt**|String|false|true|false|false|||||Y|dfmgt||false|false|false|text||
|**DfuatSupport**|String|false|true|false|false|||||Y|dfuatSupport||false|false|false|text||
|**Dfmarketing**|String|false|true|false|false|||||Y|dfmarketing||false|false|false|text||
|**Dfcontingency**|String|false|true|false|false|||||Y|dfcontingency||false|false|false|text||
|**DfdevUplift**|String|false|true|false|false|||||Y|dfdevUplift||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/feature_core.go |
| code | **adaptor** | /adaptor/feature_impl.go_template |
| code | **api** | /application/feature_api.go |
| code | **dao** | /dao/feature_core.go |
| code | **datamodel** | /datamodel/feature_core.go |
| code | **menu** | /design/menu/feature.json |
| html | **list** | /Feature_List.html |
| html | **view** | /Feature_View.html |
| html | **edit** | /Feature_Edit.html |
| html | **new** | /Feature_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **05/12/2022** at **20:49:56**
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