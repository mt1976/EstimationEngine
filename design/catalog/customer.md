# **Customer** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Customer** (customer) |
|Endpoint 	    |**/Customer...** [^1]|
|Endpoint Query |**customerID**|
|REST API|**/API/Customer/**|
Glyph|**fas fa-user-cog** ()
Friendly Name|**Customer**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Customer/CustomerList) [Exportable]
* **View** (/Customer/CustomerView)
* **Edit** (/Customer/CustomerEdit)
* **Save** (/Customer/CustomerSave)
* **New** (/Customer/CustomerNew)








##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **customer**
SQL Table Key | **customerID**
Fetch|<ul><li>**Implement in Adaptor**</li><li> func Customer_GetList_impl() (int, []dm.Customer, error) {return 0, nil, nil}</li><li>func Customer_GetByID_impl(id string) (int, dm.Customer, error) {return 0, dm.Customer{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func Customer_NewID_impl(rec dm.Customer) (string) { return rec.ID } </li><li>func Customer_Delete_impl(id string) (error) {return nil}</li><li>func Customer_Update_impl(id string,rec dm.Customer, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**CustomerID**|String|true|true|false|false|||||Y|CustomerID||false|false|false|text||
|**Code**|String|true|true|false|false|||||Y|Code||false|false|false|text||
|**FullName**|String|true|true|false|false|||||Y|FullName||false|false|false|text||
|**Rate**|String|true|true|false|false|||||Y|Rate||false|false|false|text||
|**Notes**|String|true|true|false|false|||||Y|Notes||false|false|false|text||
|**StartDate**|String|true|true|false|false|||||Y|StartDate||false|false|false|text||
|**EndDate**|String|true|true|false|false|||||Y|EndDate||false|false|false|text||
|**Status**|String|true|true|false|false|||||Y|Status||false|false|false|text||
|**WorkItemType**|String|true|true|false|false|||||Y|WorkItemType||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/customer_core.go |
| code | **adaptor** | /adaptor/customer_impl.go_template |
| code | **api** | /application/customer_api.go |
| code | **dao** | /dao/customer_core.go |
| code | **datamodel** | /datamodel/customer_core.go |
| code | **menu** | /design/menu/customer.json |
| html | **list** | /Customer_List.html |
| html | **view** | /Customer_View.html |
| html | **edit** | /Customer_Edit.html |
| html | **new** | /Customer_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **27/11/2022** at **12:36:27**
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