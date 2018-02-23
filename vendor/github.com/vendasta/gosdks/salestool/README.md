Salestool SDK
=============
An sdk to communicate with Salestool's apis via Go.

## Versions
### 3.0.0
- BREAKING CHANGE: Contexts must be provided to client calls -- this enables tracing

### 2.0.0
- Move `GetSalesperson` out of `SnapshotClient` and into `SalespersonClient`
- Add `RoundRobin` to `SalespersonClient`

### 1.4.0
- Update required fields to include address, city, country

### 1.3.0
- Moved salesperson to its own file
- Added create account endpoint

### 1.2.0
- Can now retrieve a salesperson from ST

### 1.1.0
- Added campaignId to snapshot create

### 1.0.0
- Added snapshot client
