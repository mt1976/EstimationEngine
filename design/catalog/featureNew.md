# **FeatureNew** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**FeatureNew** (featurenew) |
|Endpoint 	    |**/FeatureNew...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/FeatureNew/**|
Glyph|**far fa-plus-square** ()
Friendly Name|**Create a New Feature**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}

* **View** (/FeatureNew/FeatureNewView)
* **Edit** (/FeatureNew/FeatureNewEdit)
* **Save** (/FeatureNew/FeatureNewSave)
* **New** (/FeatureNew/FeatureNewNew)








##  Provides







##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **n/a**
SQL Table Key | **n/a**
Fetch|<ul><li>**Implement in Adaptor**</li><li> func FeatureNew_GetList_impl() (int, []dm.FeatureNew, error) {return 0, nil, nil}</li><li>func FeatureNew_GetByID_impl(id string) (int, dm.FeatureNew, error) {return 0, dm.FeatureNew{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func FeatureNew_NewID_impl(rec dm.FeatureNew) (string) { return rec.ID } </li><li>func FeatureNew_Delete_impl(id string) (error) {return nil}</li><li>func FeatureNew_Update_impl(id string,rec dm.FeatureNew, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|true|||||NH|ID||false|false|false|text||
|**EstimationSessionID**|String|true|true|false|false|OL∀|EstimationSession|EstimationSessionID_EstimationSessionID|EstimationSession_Name|Y|EstimationSessionID||false|false|false|text||
|**Name**|String|true|true|false|false|||||Y|Name||false|false|false|text||
|**DeveloperEstimate**|String|true|true|false|false|||||Y|DeveloperEstimate||false|false|false|text||
|**DeveloperResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|DeveloperResource||false|false|false|text||
|**AnalystEstimate**|String|true|true|false|false|||||Y|AnalystEstimate||false|false|false|text||
|**AnalystResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|AnalystResource||false|false|false|text||
|**ConfidenceCODE**|String|true|true|false|false|OL∀|Confidence|Confidence_Code|EstimationState_Name|Y|ConfidenceCODE||false|false|false|text|true|
|**ProductManagerResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|ProductManagerResource||false|false|false|text||
|**ProjectManagerResource**|String|false|true|false|false|OL∀|Resource|Resource_Code|Resource_Name|Y|ProjectManagerResource||false|false|false|text||
|**Description**|String|false|true|false|false|||||Y|Description||false|false|false|text||
|**Comments**|String|false|true|false|false|||||Y|Comments||false|false|false|text||
|**DevOpsID**|String|false|true|false|false|||||Y|DevOpsID||false|false|false|text||
|**FreshDeskID**|String|false|true|false|false|||||Y|FreshDeskID||false|false|false|text||
|**RSCID**|String|false|true|false|false|||||Y|RSCID||false|false|false|text||
|**OtherID**|String|false|true|false|false|||||Y|OtherID||false|false|false|text||
|**OtherID2**|String|false|true|false|false|||||Y|OtherID2||false|false|false|text||
|**DefaultProfile**|String|false|true|false|false|OL∀|Profile|Profile_ProfileID|Profile_Name|N|DefaultProfile||false|false|false|text||
|**ActualProfile**|String|false|true|false|false|OL∀|Profile|Profile_ProfileID|Profile_Name|Y|ActualProfile||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/featureNew_core.go |
| code | **adaptor** | /dao/featureNew_adaptor.go_template |
| code | **validation** | /dao/featureNew_validation.go_template |
| code | **api** | /application/featureNew_api.go |
| code | **dao** | /dao/featureNew_core.go |
| code | **datamodel** | /datamodel/featureNew_core.go |
| code | **menu** | /design/menu/featureNew.json |
| html | **view** | /FeatureNewView.html |
| html | **edit** | /FeatureNewEdit.html |
| html | **new** | /FeatureNewNew.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **03/03/2023** at **13:56:51**
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