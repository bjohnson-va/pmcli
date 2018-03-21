# Proto Mock Command Line Interface

*pmcli* provides mocking utilities based on protofiles

## Usage
```
pmcli serve
```

### Requirements
Before you can run `pmcli serve`, you should set up a minimal configuration 
file called mockserver.json.  Put this file in the directory from which you 
will be running *pmcli*.  This will typically be the root directory of the
project you will be a running a mockserver *in place of*.

Example:
```json
{
  "port": 28000,
  "protofiles": [
    "advertising/v1/api.proto"
  ],
  "overrides": {
  },
  "instructions": {
  },
  "exclusions": {
  }
}
```

The config file is a major component of *pmcli*'s operation. See 
[Configuration](#configuration) for more info.

### Additional options
*pmcli* can be run with command-line flags
Use `pmcli -h` for more info.

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

Each section of this config file is explained in more detail below the example.

### Example `mockserver.json`
```json
{
  "protofiles": [
    "advertising/v1/api.proto"
  ],
  "overrides": {
    "Campaigns/List": {
      "campaigns.businessId": "MY-OVERRIDE"
    },
    "Adwords/GetStatsForBusiness": {
      "conversions": 1,
      "costMicrodollars": 1000000.00
    }
  },
  "instructions": {
    "Campaigns/List": {
      "delaySeconds": 0.55,
      "statusCode": 200,
      "emptyBody": false,
      "fields": {
        "campaigns": {
          "num": 2
        }
      }
    }
  },
  "exclusions": {
    "Adwords/GetAccountInfo": [
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
overridden.  This is a limitation of pmcli and will likely be addressed later.

Also note that override was provided in `snake_case` but the response was
generated in `camelCase`.  *pmcli* will accept field keys in either format 
but will always generate its responses in `camelCase`.

---

### Instructions:

`statusCode`: The HTTP status code to return from the associated endpoint.

`emptyBody`: boolean.  Set to true to have the associated endpoint return an
empty body.

`delaySeconds`: The number of seconds to wait before the response is returned.

`fields`: A json dict containing field-specific instructions

### Instructions - Fields

These instructions apply to specific fields of an RPC's response body.

`num`: Determines how many items will be in the randomly generated list that is
output from the API.

E.g.
```
  "SomeService/SomeRPC": {
    "fields": {
      "someRepeatedField": {
        "num": 5
      }
    }
  }
```

---

### Exclusions:

NOTE: **Exclusions must be specified in snake_case.  This is temporary.**

This is a list of fields that should not be included in the response at all.
