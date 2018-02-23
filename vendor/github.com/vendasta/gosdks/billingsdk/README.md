Billing SDK
=============
An sdk to communicate with Billing's apis via Go.


## Versions
### 2.1.0
- Billing expects prices to be cents (int64) now

### 2.0.3
- Rename 'expiry' to 'expiryDateTime' as expected by vapi

### 2.0.2
- Format expiry using RFC3339 standard

### 2.0.1
- Add context to billing client mock functions

### 2.0.0
- Make PurchaseOptions public

### 1.0.0
- Added billing client
- Added ProductCreate and ProductUpdate
- Added ActivateProduct and DeactivateProduct
