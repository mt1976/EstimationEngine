# **Profile** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**Profile** (profile) |
|Endpoint 	    |**/Profile...** [^1]|
|Endpoint Query |**ProfileID**|
|REST API|**/API/Profile/**|
Glyph|**fas fa-sliders-h** ()
Friendly Name|**Profile**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Profile/ProfileList) [Exportable]
* **View** (/Profile/ProfileView)
* **Edit** (/Profile/ProfileEdit)
* **Save** (/Profile/ProfileSave)
* **New** (/Profile/ProfileNew)
* **Delete** (/Profile/ProfileDelete)







##  Provides
 * Lookup (Code Name)

* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **profileStore**
SQL Table Key | **profileID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**ProfileID**|String|false|true|false|true|||||H|profileID||true|false|false|text||
|**Code**|String|true|true|false|true|||||Y|code||true|false|false|text||
|**Name**|String|true|true|false|false|||||Y|name||false|false|false|text||
|**StartDate**|String|false|true|false|true|||||Y|startDate||false|false|false|date|yyyy-mm-dd|
|**EndDate**|String|false|true|false|true|||||Y|endDate||false|false|false|date|yyyy-mm-dd|
|**DefaultReleases**|String|false|true|false|true|||||Y|DefaultReleases||false|false|false|number||
|**DefaultReleaseHours**|String|false|true|false|true|||||Y|DefaultReleaseHours||false|false|false|number||
|**BlendedRate**|String|true|true|false|true|||||Y|BlendedRate||false|false|false|number||
|**Rounding**|String|false|true|false|true|||||Y|Rounding||false|false|false|number||
|**HoursPerDay**|String|true|true|false|true|||||Y|HoursPerDay||false|false|false|number||
|**REQPerc**|String|false|true|false|true|||||Y|REQPerc||false|false|false|number||
|**ANAPerc**|String|false|true|false|true|||||Y|ANAPerc||false|false|false|number||
|**DOCPerc**|String|false|true|false|true|||||Y|DOCPerc||false|false|false|number||
|**PMPerc**|String|false|true|false|true|||||Y|PMPerc||false|false|false|number||
|**UATPerc**|String|false|true|false|true|||||Y|UATPerc||false|false|false|number||
|**GTMPerc**|String|false|true|false|true|||||Y|GTMPerc||false|false|false|number||
|**SupportUplift**|String|false|true|false|true|||||Y|SupportUplift||false|false|false|number||
|**ContigencyPerc**|String|false|true|false|false|||||Y|ContigencyPerc||false|false|false|text||
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
|**Notes**|String|false|true|false|false|||||Y|notes||false|false|false|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**Comments**|String|false|true|false|false|||||Y|comments||false|false|false|text||
|**TrainingPerc**|String|false|true|false|true|||||Y|TrainingPerc||false|false|false|number||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/profile_core.go |
| code | **adaptor** | /dao/profile_adaptor.go_template |
| code | **validation** | /dao/profile_validation.go_template |
| code | **api** | /application/profile_api.go |
| code | **dao** | /dao/profile_core.go |
| code | **datamodel** | /datamodel/profile_core.go |
| code | **menu** | /design/menu/profile.json |
| html | **list** | /ProfileList.html |
| html | **view** | /ProfileView.html |
| html | **edit** | /ProfileEdit.html |
| html | **new** | /ProfileNew.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **15/03/2023** at **19:24:49**
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