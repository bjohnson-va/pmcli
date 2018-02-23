# Marketplace Webhooks Changelog

## 1.1.0
- Add `GetSessionWebhookHandler` for handling for Session webhooks from marketplace

## 1.0.1
- 1.0.0 was completely broken using wrong keys for non-local and
trying to unmarshal JSON into []byte instead of json.RawMessage

## 1.0.0
- Initial Release