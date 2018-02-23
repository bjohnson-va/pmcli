### CHANGELOG
### 2.9.0
- Add business_id, app_id and addon_id to DismissPendingActivationRequest and ResolvePendingActivationRequest

### 2.8.1
- Add missing AppID field to PendingActivation

### 2.8.0
- Add `ResolvePendingActivationRequest` and `DismissPendingActivationRequest`
- Create new `PendingActivationService`
- Create new `PendingActivation` message

### 2.7.0
- Add blame message that encapsulates meta-data about activation actions
- Extend ActivateApp, DeactivateApp, ActivateAddon, DeactivateAddon, and Delete endpoints
  to also include a Blame message

### 2.6.0
- Extend ActivateApp endpoint, add fields in ActivateAppRequest and ActivateAppResponse
- Remove Activate endpoint and ActivateRequest, ActivateResponse

### 2.5.0
- Add boolean trial field to ActivateRequest

## 2.4.1
- Fix imports to support multi package compilation

## 2.4.0
- Added Activate rpc for the Fulfillment service

## 2.3.0
- Added Fulfillment service and created protos for the Set Auto Renew Endpoint

## 2.2.0
- Added activate_on and deactivate_on date to Addon activation request

## 2.1.0
- Added order_form_submission_id to Account, so you can get that information on a Get call.

## 2.0.0
- Added order_form_submission_id to AddonActivation
- Added order_form_submission_id to ActivateAddonRequest
- Added order_form_submission_id to CreateRequest
- Added order_form_submission_id to ActivateAddonResponse
- Create CreateResponse message and added activation_id and order_form_submission_id to CreateResponse
- Modify Create to return CreateResponse instead of google.protobuf.Empty
- Removed the unused multiple addon activation objects

## 1.5.0
- Added java options to addons.proto

## 1.4.0
- Add product_id and partner_id to GetRequest. This is required to avoid circular µservice dependencies for legacy applications

## 1.3.0
- Add product_id to DeleteRequest

## 1.2.0
- Add product_id to UpdateExpiry request

## 1.1.0
- ListAddonActivations endpoint gets activations of an app's add-ons for a business

## 1.0.0
- Initial commit of accounts protos
- Only contains addon activation and deactivation at the moment
