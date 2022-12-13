# **FeatureNew** - Object Definition
---
##  Information
|   |   |
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
|   |   |
|---|---|
SQL Table Name       | **n/a**
SQL Table Key | **n/a**
Fetch|<ul><li>**Implement in Adaptor**</li><li> func FeatureNew_GetList_impl() (int, []dm.FeatureNew, error) {return 0, nil, nil}</li><li>func FeatureNew_GetByID_impl(id string) (int, dm.FeatureNew, error) {return 0, dm.FeatureNew{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func FeatureNew_NewID_impl(rec dm.FeatureNew) (string) { return rec.ID } </li><li>func FeatureNew_Delete_impl(id string) (error) {return nil}</li><li>func FeatureNew_Update_impl(id string,rec dm.FeatureNew, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|true|||||NH|ID||false|false|false|text||
|**EstimationSession**|String|true|true|false|false|OL|EstimationSession|EstimationSessionID_EstimationSessionID|EstimationSession_Name|Y|EstimationSession||false|false|false|text||
|**Name**|String|true|true|false|false|||||Y|Name||false|false|false|text||
|**DevEstimate**|String|false|true|false|false|||||Y|DevEstimate||false|false|false|text||
|**Confidence**|String|true|true|false|false|OL|Confidence|Confidence_Code|EstimationState_Name|Y|Confidence||false|false|false|text|true|
|**Developer**|String|false|true|false|false|OL|Resource|Resource_Code|Resource_Name|Y|Developer||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/featureNew_core.go |
| code | **adaptor** | /adaptor/featureNew_impl.go_template |
| code | **api** | /application/featureNew_api.go |
| code | **dao** | /dao/featureNew_core.go |
| code | **datamodel** | /datamodel/featureNew_core.go |
| code | **menu** | /design/menu/featureNew.json |
| html | **view** | /FeatureNew_View.html |
| html | **edit** | /FeatureNew_Edit.html |
| html | **new** | /FeatureNew_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **11/12/2022** at **19:24:01**
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