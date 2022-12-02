# **Profile** - Object Definition
---
##  Information
|   |   |
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








##  Provides
 * Lookup (Code Name)

* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **profileStore**
SQL Table Key | **profileID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**ProfileID**|String|false|true|false|true|||||H|profileID||true|false|false|text||
|**Code**|String|true|true|false|false|||||Y|code||false|false|false|text||
|**Name**|String|true|true|false|false|||||Y|name||false|false|false|text||
|**StartDate**|String|false|true|false|true|||||Y|startDate||false|false|false|date|yyyy-mm-dd|
|**EndDate**|String|false|true|false|true|||||Y|endDate||false|false|false|date|yyyy-mm-dd|
|**DefaultReleases**|String|true|true|false|false|||||Y|DefaultReleases||false|false|false|text||
|**DefaultReleaseHours**|String|true|true|false|false|||||Y|DefaultReleaseHours||false|false|false|text||
|**BlendedRate**|String|true|true|false|false|||||Y|BlendedRate||false|false|false|text||
|**Rounding**|String|true|true|false|false|||||Y|Rounding||false|false|false|text||
|**HoursPerDay**|String|true|true|false|false|||||Y|HoursPerDay||false|false|false|text||
|**REQPerc**|String|false|true|false|false|||||Y|REQPerc||false|false|false|text||
|**ANAPerc**|String|false|true|false|false|||||Y|ANAPerc||false|false|false|text||
|**DOCPerc**|String|false|true|false|false|||||Y|DOCPerc||false|false|false|text||
|**PMPerc**|String|false|true|false|false|||||Y|PMPerc||false|false|false|text||
|**UATPerc**|String|false|true|false|false|||||Y|UATPerc||false|false|false|text||
|**GTMPerc**|String|false|true|false|false|||||Y|GTMPerc||false|false|false|text||
|**SupportUplift**|String|false|true|false|false|||||Y|SupportUplift||false|false|false|text||
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


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/profile_core.go |
| code | **adaptor** | /adaptor/profile_impl.go_template |
| code | **api** | /application/profile_api.go |
| code | **dao** | /dao/profile_core.go |
| code | **datamodel** | /datamodel/profile_core.go |
| code | **menu** | /design/menu/profile.json |
| html | **list** | /Profile_List.html |
| html | **view** | /Profile_View.html |
| html | **edit** | /Profile_Edit.html |
| html | **new** | /Profile_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **01/12/2022** at **09:40:02**
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