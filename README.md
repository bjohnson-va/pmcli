# Proto Mock Command Line Interface

*pmcli* provides mocking utilities based on protofiles

## Usage
```
pmcli serve
```

### Additional options
```
-p: port on which to serve generated responsed

    E.g. pmcli serve -p 28003
```

## Configuration
By default, *pmcli* will attempt to read file `config.json` in the same
directory from which it was run.

### Example config.json
```json
{
  "overrides": {
    "Campaigns.List": {
      "campaigns.business_id": "MY-OVERRIDE"
    },
    "Adwords.GetStatsForBusiness": {
      "conversions": 1,
      "cost_microdollars": 1000000.00
    }
  },
  "instructions": {
    "statusCode": 200,
    "emptyBody": false,
    "Campaigns.List": {
      "campaigns": {
        "num": 1
      }
    }
  }
}
```

---

Overrides:
> In the example above, when the endpoint `Campaigns.List` is hit, the mock
server will generate a response for that endpoint but will replace values with
the ones provided.

> So, if the generated response normally looked like:
```json
{
  "campaigns": [
    {
      "businessId": "B123"
    },
    {
      "businessId": "B456"
    }
  ]
}
```

> With the override, the response would be
```json
{
  "campaigns": [
    {
      "businessId": "MY-OVERRIDE"
    },
    {
      "businessId": "MY-OVERRIDE"
    },
  ]
}
```

Note that, since `campaigns` is a list field, *every* businessId field will be
overridden.  This is a limitation of pmcli and will likely be address later.

Also note that override was provided in `snake_case` but the response was
generated in `camelCase`.

---

Instructions:


TODO: Document this
Basically, you can specify how many list items will be generated by changing `num`.
