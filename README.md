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
By default, *pmcli* will attempt to read the file `mockserver.json` in the same
directory from which it was run.

### Example `mockserver.json`
```json
{
  "protofiles": [
    "advertising/v1/api.proto"
  ],
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
    "delaySeconds": 0.55,
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

`statusCode`: The HTTP status code to return from the associated endpoint.

`emptyBody`: boolean.  Set to true to have the associated endpoint return an
empty body.

`delaySeconds`: The number of seconds to wait before the response is returned.

`num`: Specified on a specific field of an RPC.  Determines how many items will
be in the randomly generated list that is output from the API.

E.g.
```
  "Some.RPC": {
    "some_repeated_field": {
      "num": 5
    }
  }
```
