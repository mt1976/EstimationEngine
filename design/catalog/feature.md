# **Feature** - Object Definition
##  Information
| Information  | Value  |
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
* **Delete** (/Feature/FeatureDelete)







##  Provides
 * Lookup (FeatureID Name)

* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **featureStore**
SQL Table Key | **featureID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**FeatureID**|String|false|true|false|true|||||H|featureID||true|true|false|text||
|**EstimationSessionID**|String|true|true|false|false|OL|EstimationSession|EstimationSession_EstimationSessionID|EstimationSession_Name|N|estimationSessionID||false|true|false|text||
|**ConfidenceCODE**|String|true|true|false|false|OL∀|Confidence|Confidence_ConfidenceID||Y|confidenceCODE||false|false|false|text||
|**Name**|String|false|true|false|true|||||Y|name||false|true|false|text||
|**DevelopmentEstimate**|String|false|true|false|true|||||N|developmentEstimate||false|true|false|number||
|**DevelopmentFactored**|String|false|true|false|true|||||N|developmentFactored||false|true|false|number||
|**Requirements**|String|false|true|false|true|||||Y|requirements||false|true|false|number||
|**AnalystTest**|String|false|true|false|true|||||Y|analystTest||false|true|false|number||
|**Documentation**|String|false|true|false|true|||||Y|documentation||false|true|false|number||
|**Management**|String|false|true|false|true|||||Y|management||false|true|false|number||
|**UatSupport**|String|false|true|false|false|||||Y|uatSupport||false|false|false|text||
|**Marketing**|String|false|true|false|true|||||Y|marketing||false|true|false|number||
|**Contingency**|String|false|true|false|true|||||Y|contingency||false|true|false|number||
|**TrackerID**|String|false|true|false|true|||||Y|trackerID||false|true|false|text||
|**AdoID**|String|false|true|false|true|||||Y|adoID||false|true|false|text||
|**FreshdeskID**|String|false|true|false|true|||||Y|freshdeskID||false|true|false|text||
|**ExtRef**|String|false|true|false|true|||||Y|extRef||false|true|false|text||
|**ExtRef2**|String|false|true|false|true|||||Y|extRef2||false|true|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**DeveloperResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|developerResource||false|true|false|text||
|**ApproverResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|approverResource||false|true|false|text||
|**Activity**|String|false|true|false|true|||||N|activity||false|false|false|textarea||
|**OffProfile**|String|false|true|false|true|LL|tf|||N|offProfile||false|false|false|text||
|**OffProfileJustification**|String|false|true|false|true|||||Y|offProfileJustification||false|true|false|textarea||
|**SYSActivity**|String|false|true|false|false|||||NH|_activity||false|false|true|text||
|**RequirementsDefault**|String|false|true|false|true|||||N|requirementsDefault||false|false|false|text||
|**AnalystTestDefault**|String|false|true|false|true|||||N|analystTestDefault||false|false|false|text||
|**DocumentationDefault**|String|false|true|false|true|||||N|documentationDefault||false|false|false|text||
|**ManagementDefault**|String|false|true|false|true|||||N|managementDefault||false|false|false|text||
|**UatSupportDefault**|String|false|true|false|true|||||N|uatSupportDefault||false|false|false|text||
|**MarketingDefault**|String|false|true|false|true|||||N|marketingDefault||false|false|false|text||
|**ContingencyDefault**|String|false|true|false|true|||||N|contingencyDefault||false|false|false|text||
|**DevelopmentFactoredDefault**|String|false|true|false|true|||||N|developmentFactoredDefault||false|false|false|text||
|**Total**|String|false|true|false|true|||||N|total||false|true|false|number||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**Comments**|String|false|true|false|true|||||Y|comments||false|false|false|textarea||
|**Description**|String|false|true|false|true|||||Y|description||false|false|false|textarea||
|**AnalystResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|analystResource||false|false|false|text||
|**ProductManagerResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|productManagerResource||false|false|false|text||
|**ProjectManagerResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|projectManagerResource||false|false|false|text||
|**Training**|String|false|true|false|true|||||Y|training||false|true|false|number||
|**TrainingDefault**|String|false|true|false|true|||||N|trainingDefault||false|false|false|text||
|**ProfileDefault**|String|false|true|false|true|OL|Profile|Profile_ProfileID|Profile_Name|N|profileDefault||false|false|false|text||
|**ProfileSelected**|String|false|true|false|false|OL∀|Profile|Profile_ProfileID|Profile_Name|Y|profileSelected||false|false|false|text||
|**EstimateEffort**|String|false|true|false|true|||||Y|estimateEffort||false|true|false|number||
|**EstimateEffortDefault**|String|false|true|false|true|||||NH|estimateEffortDefault||false|false|false|text||
|**AnalystEstimate**|String|false|true|false|true|||||N|analystEstimate||false|true|false|number||
|**TotalDefault**|String|false|true|false|true|||||N|totalDefault||false|false|false|number||
|**RequirementsPerc**|String|false|true|false|true|||||NH|requirementsPerc||false|true|false|text||
|**AnalystTestPerc**|String|false|true|false|true|||||NH|analystTestPerc||false|true|false|text||
|**DocumentationPerc**|String|false|true|false|true|||||NH|documentationPerc||false|true|false|text||
|**ManagementPerc**|String|false|true|false|true|||||NH|managementPerc||false|true|false|text||
|**UatSupportPerc**|String|false|true|false|true|||||NH|uatSupportPerc||false|true|false|text||
|**MarketingPerc**|String|false|true|false|true|||||NH|marketingPerc||false|true|false|text||
|**ContingencyPerc**|String|false|true|false|true|||||NH|contingencyPerc||false|true|false|text||
|**TrainingPerc**|String|false|true|false|true|||||NH|trainingPerc||false|true|false|text||
|**RoundedTo**|String|false|true|false|true|||||NH|roundedTo||false|true|false|text||
|**OriginName**|String|false|false|true|true|||||N|||false|true|false|text||
|**ProjectName**|String|false|false|true|true|||||N|||false|true|false|text||
|**OriginID**|String|false|false|true|true|||||N|||false|true|false|text||
|**ProjectID**|String|false|false|true|true|||||N|||false|true|false|text||
|**ProfileSelectedOld**|String|false|false|true|false|||||N|||false|true|false|text||
|**CCY**|String|false|false|true|false|||||N|||false|true|false|text||
|**OriginCode**|String|false|false|true|true|||||N|||false|true|false|text||
|**EstimationSessionName**|String|false|false|true|true|||||N|||false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/feature_core.go |
| code | **adaptor** | /dao/feature_adaptor.go_template |
| code | **validation** | /dao/feature_validation.go_template |
| code | **api** | /application/feature_api.go |
| code | **dao** | /dao/feature_core.go |
| code | **datamodel** | /datamodel/feature_core.go |
| code | **menu** | /design/menu/feature.json |
| html | **list** | /FeatureList.html |
| html | **view** | /FeatureView.html |
| html | **edit** | /FeatureEdit.html |
| html | **new** | /FeatureNew.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **15/03/2023** at **19:24:48**
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