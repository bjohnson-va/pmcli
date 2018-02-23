### 1.11.0
- Added AddonActivation message
- Added `addon_activations` to Order message

### 1.10.0
- Updated CommonField, CustomField, and Package to support Addons

### 1.9.0
- Add `fulfilled`, `declined` and `approved` dates to `Order` proto message

### 1.8.0
- Added filter for business_id to ListSalesOrders

### 1.7.0
- Added response for approve, decline and activate endpoints

### 1.6.1
- Update imports to use full path

### 1.6.0
- Update Field to contain a label, description, and field_type
- Add FieldType enum

### 1.5.0
- Add 'ACTIVATION_ERRORS' as an order status

### 1.4.0
- Update activate products request to contain business_id, common_fields, and custom_fields

### 1.3.0
- Add activate products endpoint

### 1.2.0
- Added approve and decline endpoints

### 1.1.0
- Added list endpoint for listing sales orders for partners

### 1.0.0
- Initial commit of sales_orders protos
