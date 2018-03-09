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

## Randomization Strategies (R.S.)

You can adjust the way *pmcli* generates random values with the `-r` flag.

`pmcli serve -r time`

The currently supported strategies are `breadcrumb` + `time`;

### Breadcrumb

By default, *pmcli* will use a "breadcrumb based" randomization strategy.

The purpose of the Breadcrumb R.S is to provide psuedorandom (repeatable)
values from your APIs.

*pmcli* generates a breadcrumb for each field it generates.  For example, in the
following response:
```
{
  "campaigns":[
    {
      "business": {
        "name": "My business"
      }
      "campaignId":"some campaign ID",
      "endDate":"2016-01-01",
      "orderDate":"2016-01-01",
      "startDate":"2016-01-01",
      "status":"ORDER_RECEIVED"
    }
  ]
}
```

... `campaigns.business.name` would be an example of a breadcrumb.

In the case of lists, this randomization strategy will also take the list index
into account, ensuring that each item in the generated list (
see <a href="#instructions">#Instructions</a>) has a unique value in the scope
of that list.

### Time

The time-based random value generator uses the current system time to seed its
random values.

The purpose of the Time R.S. is to generate entirely unique values on each
subsequent call to your API.

Note: In order to ensure that two (for example) integer fields in the same
message have different values, the `time` R.S. also factors in the breadcrumbs.


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
    "delaySeconds": 0.55,
    "Campaigns.List": {
      "statusCode": 200,
      "emptyBody": false,
      "campaigns": {
        "num": 1
      }
    }
  }
  "exclusions": {
    "Adwords.GetAccountInfo": [
      "account_info"
    ]
  }
}
```

---

### Overrides:
In the example above, when the endpoint `Campaigns.List` is hit, the mock
server will generate a response for that endpoint but will replace values with
the ones provided.

So, if the generated response normally looked like:
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

With the override `"campaigns.business_id": "MY-OVERRIDE"`, the response
would be:
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

### Instructions:

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

---

### Exclusions:

This is a list of fields that should not be included in the response at all.
