
## Usage

```golang
adClient, err := ad.NewClient(ctx, config.CurEnv())
if err != nil {
    logging.Errorf(ctx, "Error initializing advertising client %s", err.Error())
    os.Exit(-1)
}

var advertisingPackage string
if orderForm.AdvertisingPackage == nil {
    advertisingPackage = ""
} else {
    advertisingPackage = orderForm.AdvertisingPackage.Value
}

order := &ad.OrderEvent{
    BusinessID: orderForm.BusinessID,
    PartnerID: orderForm.PartnerID,
    OrderID: orderForm.OrderID, 
    PackageType: orderForm.PackageType,
    AdvertisingPackage: advertisingPackage,
    BusinessName: orderForm.BusinessName,
    BusinessAddress: orderForm.BusinessAddress,
    BusinessPhone: orderForm.BusinessPhone,
    ContactEmail: orderForm.ContactEmail,
    ContactName: orderForm.ContactName,
    ContactPhone: orderForm.ContactPhone,
    SalespersonEmail: orderForm.SalespersonEmail,
    SalespersonName: orderForm.SalespersonName,
    SalespersonPhone: orderForm.SalespersonPhone,
    CreativeLocation: orderForm.CreativeLocation,
    CustomerValue: orderForm.CustomerValue,
    FacebookPageURL: orderForm.FacebookPageURL,
    Notes: orderForm.Notes,
    Objective: orderForm.Objective,
    Promo: orderForm.Promo,
    RetailPrice: orderForm.RetailPrice,
    Specialties: orderForm.Specialties,
    Targeting: orderForm.Targeting,
    Term: orderForm.Term,
    WebsiteURL: orderForm.WebsiteURL,
    RawMarketplacePayload: orderForm.RawMarketplacePayload,
    CreatedAt: orderForm.CreatedTime,
}

_, err = adClient.CreateOrderEvent(ctx, order)
if err != nil {
    logging.Infof(ctx, "Error  to call CreateOrderEvent endpoint in advertising client for order id %s: %s", orderForm.OrderID, err.Error())
    return err
}
```
