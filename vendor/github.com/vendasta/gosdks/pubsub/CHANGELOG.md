# PUBSUB

## 1.0.0 (BREAKING CHANGES)
- BREAKING CHANGE: Topic.Publish now returns a `PublishResult` interface rather than a pointer to google's
    `*PublishResult`
    - fix: Assuming you've imported THIS library as vpubsub; change all occurrences of `*pubsub.PublishResult` into
        `vpubsub.PublishResult`

## (Prior Changelog missing)
