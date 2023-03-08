# **ssloader** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**ssloader** (ssloader) |
|Endpoint 	    |**/Ssloader...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-cubes** (text-info)
Friendly Name|**Spreadsheet Loader**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}













##  Provides







##  Data Source 
| Information  | Value  |
|---|---|




##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|false|||||Y|ID||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **job** | /jobs/ssloader_core.go |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **08/03/2023** at **18:42:26**
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