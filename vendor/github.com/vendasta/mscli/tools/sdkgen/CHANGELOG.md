# Changelog

### Don't forget to set the version in VERSION
1.1.0
- Fix to Python SDK generation - args will now be generated as keyword arguments, example:
    `def SyncData(self, gmb_insights_json=None, requested_at=None):`

1.0.1
- Fix to Typescript SDK - `_internal/index.ts` and `*.module.ts`
now include all generated services

1.0.0
- BREAKING CHANGE
 - Typescript SDK generation now generates a different file per service
 - This is to support multiple services having RPCs with the same name
 (i.e. `InvoiceService.Get` and `MerchantService.Get`)
 - It's likely that imports from your existing TS SDK will break if you
 regenerate after this point

0.16.1
- Fix to appengine import pathing. Clean up modules after failed import

0.16.0
- Fix to the python sdk generation

0.15.7
- Fix toKebabCase() bug with inputs that contain underscores
0.15.6
- Ensure we found a proto service before trying to generate the sdk
0.15.5
- int64 values are JSON encoded as strings, so use parseInt
0.15.4
- Fix quotes and semicolon in Typescript template 
0.15.3
- Fix generation for protofile paths which would contain google in the name
0.15.2
- Fix lint error in generated TS SDK
0.15.1
- Fix typescript SDK generation duplicate imports in services
0.15.0
- Update Typscript SDK-Gen api services to use the angular HttpClient instead of the deprecated Http
0.14.4 [2017-11-10]
- Add missing imports to service template file
0.14.3 [2017-11-10]
- Fix empty message generation in Typescript SDK to simply return empty object on toApiJson
0.14.2 [2017-10-03]
- Don't generate enum index file if there aren't any enum files
0.14.1 [2017-09-28]
- Fix python import path
0.14.0 [2017-09-28]
- Breaking: Python SDK Generation - from_proto will now be handled in the generated layer, and no longer needs to be done in the hand rolled layer.
This means that any SDK generations for existing SDKs will need to have from_proto removed.
0.13.2
- Remove delcare from enum template
0.13.1
 - Typescript SDK Gen no longer relies on @vendasta/core index to import session/environment
0.13.0
 - Fix a number of lint issues in the generated typescript sdk
0.12.1 [2017-09-18]
 - Extracted proto documentation parsing to internal
0.12.0 [2017-09-15]
 - Typescript SDK Generator fixed to handle maps as objects rather than list of generated entries
0.11.0 [2017-09-13]
- Change generation to assume the SDK is being added to the @vendasta/core package of frontend
0.10.1 [2017-09-12]
- Fix typescript host service generation to use secondary SSL host if available, and to scheme in urls to avoid relative calls
0.10.0 [2017-08-28]
 - Typescript SDK Generator now generates the API layer
0.7.0 [2017-08-18]
 - Add support for Wrapper WKTs (ex. BoolValue) for Python and Typescript
0.6.0 [2017-07-07]
 - Add support for Java
0.5.2 [2017-07-07]
 - Fix import strings for typescript
0.5.1 [2017-07-05]
 - Fix Python's code generator
0.5.0 [2017-07-05]
 - Added support for typescript
0.4.0 [2017-07-04]
 - Restructured file tree to support other languages
0.2.0 [2017-06-15]
 - Python
   - Fix repeated objects not having falsy values
   - Fixed repeated enums
0.1.3 [2017-05-17]
 - Python
   - Extract repeated Enums
   - Interpret empty nested messages as None rather than Zero-valued message instances
0.1.2 [2017-05-15]
 - Use FloatProperty for Doubles
0.1.1 [2017-05-15]
 - Use the extract helper functions on strings
 - Add a CLASS_VERSION and OWNER to the vobjects
 - Make nested messages appear before their parents in the generated file.
0.1.0 [2017-05-09]
 - Generates vobject representation of proto files.
0.0.7 [2017-05-01]
 - Added other user generated protos to the imports of transport.go
0.0.6 "Initial Release" [2017-05-01]
 - Fix path to Empty in transport.py

