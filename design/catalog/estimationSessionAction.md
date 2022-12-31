# **EstimationSessionAction** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**EstimationSessionAction** (estimationsessionaction) |
|Endpoint 	    |**/EstimationSessionAction...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/EstimationSessionAction/**|
Glyph|**fas fa-angle-double-right** ()
Friendly Name|**Estimation Session Actions**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}

* **View** (/EstimationSessionAction/EstimationSessionActionView)
* **Edit** (/EstimationSessionAction/EstimationSessionActionEdit)
* **Save** (/EstimationSessionAction/EstimationSessionActionSave)
* **New** (/EstimationSessionAction/EstimationSessionActionNew)








##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **n/a**
SQL Table Key | **n/a**
Fetch|<ul><li>**Implement in Adaptor**</li><li> func EstimationSessionAction_GetList_impl() (int, []dm.EstimationSessionAction, error) {return 0, nil, nil}</li><li>func EstimationSessionAction_GetByID_impl(id string) (int, dm.EstimationSessionAction, error) {return 0, dm.EstimationSessionAction{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func EstimationSessionAction_NewID_impl(rec dm.EstimationSessionAction) (string) { return rec.ID } </li><li>func EstimationSessionAction_Delete_impl(id string) (error) {return nil}</li><li>func EstimationSessionAction_Update_impl(id string,rec dm.EstimationSessionAction, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|true|||||NH|ID||false|false|false|text||
|**EstimationSession**|String|true|true|false|false|OL|EstimationSession|EstimationSessionID_EstimationSessionID|EstimationSession_Name|Y|EstimationSession||false|false|false|text||
|**Action**|String|true|true|false|false|OL|EstimationState|EstimationState_Code|EstimationState_Name|Y|Action||false|false|false|text|true|
|**Notes**|String|false|true|false|true|||||Y|Notes||false|false|false|textarea||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/estimationSessionAction_core.go |
| code | **adaptor** | /adaptor/estimationSessionAction_impl.go_template |
| code | **api** | /application/estimationSessionAction_api.go |
| code | **dao** | /dao/estimationSessionAction_core.go |
| code | **datamodel** | /datamodel/estimationSessionAction_core.go |
| code | **menu** | /design/menu/estimationSessionAction.json |
| html | **view** | /EstimationSessionAction_View.html |
| html | **edit** | /EstimationSessionAction_Edit.html |
| html | **new** | /EstimationSessionAction_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **30/12/2022** at **12:18:20**
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