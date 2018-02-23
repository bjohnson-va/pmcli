## 4.6.0
- added `Restrictions` and `PermissionLists` message
- Add `restrictions` of type `Restrictions` to app and addon proto

## 4.5.0
- Add `contact_webhook_url` to integration in the app proto

## 4.4.0
- Add `reserve_id_url` string field to the app proto

## 4.3.0
- Add `allow_price_negotiation` bool field to the app proto

## 4.2.0
- Add Upload URL to OrderFormField

## 4.1.1
- Add partner_id to GetMultiOrderFormsRequest

## 4.1.0
- Add GetMultiOrderForms for getting multiple app order forms

## 4.0.1
- Fix imports to support multi package compilation

## 4.0.0
- Remove ListResellerItems and all related messages

## 3.1.0
- Add multiple activation field to addon proto and create and update endpoints

## 3.0.0
- Type change on the orderform field required attribute, from string to bool

## 2.9.0
- Update ListApprovedAddons to take in a list of additional app_ids.
- Deprecate app_id field on ListApprovedAddons

## 2.8.0
- Add SearchResellerItems endpoint
- Add ResellerItem proto

## 2.7.0
- Add discoverable and screenshots fields to Create/Update Addon Requests
- Mark description as deprecated

## 2.6.0
- Deprecate Addon description field
- Add Addon screenshots field

## 2.5.0
- Add discoverable field to the proto

## 2.4.0
- Add end_user_marketing, reseller_marketing, banner_image, icon, tagline, service_model, and billing_frequency to proto

## 2.3.0
-Add ApproveAddon, UnapproveAddon and ListApprovedAddons api protos to addons

## 2.2.1
- Add uses_order_form boolean to UpdateAddonRequest

## 2.2.0
- Add update api protos to addons

## 2.1.1
- Add uses_order_form boolean to CreateAddonRequest and the Addon itself

## 2.1.0
- Move OrderForm proto to new file
- Add OrderForm attribute to Addon proto and CreateAddonRequest rpc

## 2.0.0
- Add MarketplaceApp proto
- Remove custom REST routes for the Addon RPC's

## 1.0.0
- Initial commit of MarketplaceApps protos
