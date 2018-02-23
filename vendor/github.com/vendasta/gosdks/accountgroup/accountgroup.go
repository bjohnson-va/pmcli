package accountgroup

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/vendasta/gosdks/pb/account_group/v1"
)

var paymentToProtoMap map[PaymentMethod]accountgroup_v1.RichData_PaymentMethods
var genderToProtoMap map[Gender]accountgroup_v1.HealthCareProfessionalInformation_Gender

func init() {
	paymentToProtoMap = map[PaymentMethod]accountgroup_v1.RichData_PaymentMethods{
		AmericanExpress: accountgroup_v1.RichData_AMERICAN_EXPRESS,
		AndroidPay:      accountgroup_v1.RichData_ANDROID_PAY,
		ApplePay:        accountgroup_v1.RichData_APPLE_PAY,
		Cash:            accountgroup_v1.RichData_CASH,
		Check:           accountgroup_v1.RichData_CHECK,
		Debit:           accountgroup_v1.RichData_DEBIT,
		DinersClub:      accountgroup_v1.RichData_DINERS_CLUB,
		Discover:        accountgroup_v1.RichData_DISCOVER,
		Mastercard:      accountgroup_v1.RichData_MASTERCARD,
		Paypal:          accountgroup_v1.RichData_PAYPAL,
		SamsungPay:      accountgroup_v1.RichData_SAMSUNG_PAY,
		StoreCard:       accountgroup_v1.RichData_STORE_CARD,
		TravelersCheck:  accountgroup_v1.RichData_TRAVELERS_CHECK,
		VISA:            accountgroup_v1.RichData_VISA,
	}

	genderToProtoMap = map[Gender]accountgroup_v1.HealthCareProfessionalInformation_Gender{
		NotSpecified: accountgroup_v1.HealthCareProfessionalInformation_NotSpecified,
		Female:       accountgroup_v1.HealthCareProfessionalInformation_Female,
		Male:         accountgroup_v1.HealthCareProfessionalInformation_Male,
		Other:        accountgroup_v1.HealthCareProfessionalInformation_Other,
	}
}

// FromPB returns an *AccountGroup from an account group proto
func FromPB(ag *accountgroup_v1.AccountGroup, pf *accountgroup_v1.ProjectionFilter) (*AccountGroup, error) {
	if ag == nil {
		return nil, nil
	}
	accountGroup := &AccountGroup{
		AccountGroupID: ag.AccountGroupId,
		Version:        ag.Version,
	}
	if ag.Created != nil {
		var err error
		accountGroup.Created, err = ptypes.Timestamp(ag.Created)
		if err != nil {
			return nil, err
		}
	}
	if ag.Updated != nil {
		var err error
		accountGroup.Updated, err = ptypes.Timestamp(ag.Updated)
		if err != nil {
			return nil, err
		}
	}
	if ag.Deleted != nil {
		var err error
		accountGroup.Deleted, err = ptypes.Timestamp(ag.Deleted)
		if err != nil {
			return nil, err
		}
	}

	if pf.GetNapData() && ag.NapData != nil {
		accountGroup.NAPData = &NAPData{
			CompanyName:        ag.NapData.GetCompanyName(),
			Address:            ag.NapData.GetAddress(),
			City:               ag.NapData.GetCity(),
			State:              ag.NapData.GetState(),
			Zip:                ag.NapData.GetZip(),
			Country:            ag.NapData.GetCountry(),
			Website:            ag.NapData.GetWebsite(),
			WorkNumber:         ag.NapData.GetWorkNumber(),
			CallTrackingNumber: ag.NapData.GetCallTrackingNumber(),
			Location:           GeoFromPB(ag.NapData.GetLocation()),
			Timezone:           ag.NapData.GetTimezone(),
		}
	}
	if pf.GetContactDetails() && ag.ContactDetails != nil {
		accountGroup.ContactDetails = &ContactDetails{
			FirstName:   ag.ContactDetails.FirstName,
			LastName:    ag.ContactDetails.LastName,
			Email:       ag.ContactDetails.Email,
			PhoneNumber: ag.ContactDetails.PhoneNumber,
		}
	}
	if pf.GetAccounts() && ag.Accounts != nil && len(ag.Accounts.Accounts) > 0 {
		for _, account := range ag.Accounts.Accounts {
			var expiry time.Time
			var err error
			if account.Expiry != nil {
				expiry, err = ptypes.Timestamp(account.Expiry)
				if err != nil {
					return nil, err
				}
			}

			accountGroup.Accounts = append(accountGroup.Accounts, &Account{
				IsTrial:          account.IsTrial,
				Tags:             account.Tags,
				MarketplaceAppID: account.MarketplaceAppId,
				AccountID:        account.AccountId,
				Expiry:           expiry,
			})
		}
	}
	if pf.GetListingDistribution() && ag.ListingDistribution != nil {
		var fromDate, thruDate time.Time
		var err error
		if ag.ListingDistribution.FromDate != nil {
			fromDate, err = ptypes.Timestamp(ag.ListingDistribution.FromDate)
			if err != nil {
				return nil, err
			}
		}
		if ag.ListingDistribution.ThruDate != nil {
			thruDate, err = ptypes.Timestamp(ag.ListingDistribution.ThruDate)
			if err != nil {
				return nil, err
			}
		}

		accountGroup.ListingDistribution = &ListingDistribution{
			OrderID:    ag.ListingDistribution.OrderId,
			PurchaseID: ag.ListingDistribution.PurchaseId,
			FromDate:   fromDate,
			ThruDate:   thruDate,
			AutoRenew:  ag.ListingDistribution.AutoRenew,
		}
	}

	if pf.GetListingSyncPro() && ag.ListingSyncPro != nil {
		var purchaseDate, expiryDate time.Time
		var err error
		if ag.ListingSyncPro.PurchaseDate != nil {
			purchaseDate, err = ptypes.Timestamp(ag.ListingSyncPro.PurchaseDate)
			if err != nil {
				return nil, err
			}
		}
		if ag.ListingSyncPro.ExpiryDate != nil {
			expiryDate, err = ptypes.Timestamp(ag.ListingSyncPro.ExpiryDate)
			if err != nil {
				return nil, err
			}
		}
		accountGroup.ListingSyncPro = &ListingSyncPro{
			PurchaseDate:     purchaseDate,
			BillingFrequency: BillingFrequency(ag.ListingSyncPro.BillingFrequency),
			ExpiryDate:       expiryDate,
			Country:          ag.ListingSyncPro.Country,
			DiscountFlag:     ag.ListingSyncPro.DiscountFlag,
			ServiceProvider:  ServiceProviders(ag.ListingSyncPro.ServiceProvider),
		}
	}

	if pf.GetAssociations() && len(ag.Associations) > 0 {
		for _, association := range ag.Associations {
			accountGroup.Associations = append(accountGroup.Associations, &Association{
				Label:           association.Label,
				ProductID:       association.ProductId,
				ProductUserID:   association.ProductUserId,
				VBCUserID:       association.VbcUserId,
				DefaultLocation: association.DefaultLocation,
			})
		}
	}

	if pf.GetAccountGroupExternalIdentifiers() && ag.AccountGroupExternalIdentifiers != nil {
		accountGroup.ExternalIdentifiers = &ExternalIdentifiers{
			Origin:                   ag.AccountGroupExternalIdentifiers.Origin,
			JobID:                    ag.AccountGroupExternalIdentifiers.JobId,
			CustomerIdentifier:       ag.AccountGroupExternalIdentifiers.CustomerIdentifier,
			Tags:                     ag.AccountGroupExternalIdentifiers.Tags,
			ActionLists:              ag.AccountGroupExternalIdentifiers.ActionLists,
			SocialProfileID:          ag.AccountGroupExternalIdentifiers.SocialProfileId,
			PartnerID:                ag.AccountGroupExternalIdentifiers.PartnerId,
			MarketID:                 ag.AccountGroupExternalIdentifiers.MarketId,
			TaxIDs:                   ag.AccountGroupExternalIdentifiers.TaxIds,
			SalesPersonID:            ag.AccountGroupExternalIdentifiers.SalesPersonId,
			AdditionalSalesPersonIDs: ag.AccountGroupExternalIdentifiers.AdditionalSalesPersonIds,
			SalesforceID:             ag.AccountGroupExternalIdentifiers.SalesforceId,
		}
	}

	if pf.GetSocialUrls() && ag.SocialUrls != nil {
		accountGroup.SocialURLs = &SocialURLs{
			GoogleplusURL: ag.SocialUrls.GoogleplusUrl,
			LinkedinURL:   ag.SocialUrls.LinkedinUrl,
			FoursquareURL: ag.SocialUrls.FoursquareUrl,
			TwitterURL:    ag.SocialUrls.TwitterUrl,
			FacebookURL:   ag.SocialUrls.FacebookUrl,
			RssURL:        ag.SocialUrls.RssUrl,
			YoutubeURL:    ag.SocialUrls.YoutubeUrl,
			InstagramURL:  ag.SocialUrls.InstagramUrl,
			PinterestURL:  ag.SocialUrls.PinterestUrl,
		}
	}

	if pf.GetHoursOfOperation() && ag.HoursOfOperation != nil && len(ag.HoursOfOperation.HoursOfOperation) > 0 {
		spans := make([]*Span, len(ag.HoursOfOperation.HoursOfOperation))
		for x, span := range ag.HoursOfOperation.HoursOfOperation {
			spans[x] = (*Span)(span)
		}
		accountGroup.HoursOfOperation = &HoursOfOperation{
			HoursOfOperation: spans,
		}
	}

	if pf.GetSnapshotReports() && ag.SnapshotReports != nil && len(ag.SnapshotReports.Snapshots) > 0 {
		reports := make([]*Snapshot, len(ag.SnapshotReports.Snapshots))
		for n, report := range ag.SnapshotReports.Snapshots {
			var created, expiry time.Time
			var err error

			if report.Created != nil {
				created, err = ptypes.Timestamp(report.Created)
				if err != nil {
					return nil, err
				}
			}
			if report.Expiry != nil {
				expiry, err = ptypes.Timestamp(report.Expiry)
				if err != nil {
					return nil, err
				}
			}

			reports[n] = &Snapshot{
				Created:    created,
				Expiry:     expiry,
				SnapshotID: report.SnapshotId,
			}
		}
		accountGroup.Snapshots = &Snapshots{
			Snapshots: reports,
		}
	}

	if pf.GetLegacyProductDetails() && ag.LegacyProductDetails != nil {
		accountGroup.LegacyProductDetails = &LegacyProductDetails{
			KeyPerson:             ag.LegacyProductDetails.KeyPerson,
			ShareOfVoiceService:   ag.LegacyProductDetails.ShareOfVoiceService,
			FaxNumber:             ag.LegacyProductDetails.FaxNumber,
			CommonName:            ag.LegacyProductDetails.CommonName,
			CellNumber:            ag.LegacyProductDetails.CellNumber,
			Competitor:            ag.LegacyProductDetails.Competitor,
			AdminNotes:            ag.LegacyProductDetails.AdminNotes,
			SeoCategory:           ag.LegacyProductDetails.SeoCategory,
			Email:                 ag.LegacyProductDetails.Email,
			Place:                 ag.LegacyProductDetails.Place,
			Tagline:               ag.LegacyProductDetails.Tagline,
			SubscribedToCampaigns: ag.LegacyProductDetails.SubscribedToCampaigns,
		}
	}

	if pf.GetStatus() && ag.Status != nil {
		accountGroup.Status = (*Status)(ag.Status)
	}

	if pf.GetRichData() && ag.RichData != nil {
		var healthCareInfo *HealthCareProfessionalInformation
		var customFields []*CustomField
		var paymentMethods []PaymentMethod

		if ag.RichData.HealthCareProfessionalInformation != nil {
			hcpi := ag.RichData.HealthCareProfessionalInformation
			var dateOfBirth time.Time
			var isTakingPatients *bool
			var err error
			if hcpi.DateOfBirth != nil {
				dateOfBirth, err = ptypes.Timestamp(hcpi.DateOfBirth)
				if err != nil {
					return nil, err
				}
			}
			if hcpi.IsTakingPatients != nil {
				isTakingPatients = &hcpi.IsTakingPatients.Value
			}
			healthCareInfo = &HealthCareProfessionalInformation{
				DateOfBirth:                dateOfBirth,
				Email:                      hcpi.Email,
				Fellowship:                 hcpi.Fellowship,
				FirstName:                  hcpi.FirstName,
				Gender:                     Gender(hcpi.Gender),
				Initials:                   hcpi.Initials,
				InsurancesAccepted:         hcpi.InsurancesAccepted,
				LastName:                   hcpi.LastName,
				MedicalLicenseNumber:       hcpi.MedicalLicenseNumber,
				NationalProviderIdentifier: hcpi.NationalProviderIdentifier,
				Office:                 hcpi.Office,
				ProfessionalCredential: hcpi.ProfessionalCredential,
				Residency:              hcpi.Residency,
				School:                 hcpi.Residency,
				Specialty:              hcpi.Specialty,
				StandardizedTitle:      hcpi.StandardizedTitle,
				StateLicense:           hcpi.StateLicense,
				IsTakingPatients:       isTakingPatients,
			}
		}

		if len(ag.RichData.CustomFields) > 0 {
			customFields = make([]*CustomField, len(ag.RichData.CustomFields))
			for n, customField := range ag.RichData.CustomFields {
				customFields[n] = (*CustomField)(customField)
			}
		}

		if len(ag.RichData.PaymentMethods) > 0 {
			paymentMethods = make([]PaymentMethod, len(ag.RichData.PaymentMethods))
			for n, paymentMethod := range ag.RichData.PaymentMethods {
				paymentMethods[n] = PaymentMethod(paymentMethod)
			}
		}

		accountGroup.RichData = &RichData{
			TollFreeNumber:                    ag.RichData.TollFreeNumber,
			Description:                       ag.RichData.Description,
			ShortDescription:                  ag.RichData.ShortDescription,
			ServicesOffered:                   ag.RichData.ServicesOffered,
			BrandsCarried:                     ag.RichData.BrandsCarried,
			Landmark:                          ag.RichData.Landmark,
			InferredAttributes:                ag.RichData.InferredAttributes,
			PaymentMethods:                    paymentMethods,
			CustomFields:                      customFields,
			HealthCareProfessionalInformation: healthCareInfo,
		}
	}

	return accountGroup, nil
}

// GeoFromPB returns a *Geo from a geo pb
func GeoFromPB(geo *accountgroup_v1.Geo) *Geo {
	if geo == nil {
		return nil
	}
	return &Geo{
		Latitude:  geo.GetLatitude(),
		Longitude: geo.GetLongitude(),
	}
}

// AccountGroup is a container for business information, the products it has activated and more...
type AccountGroup struct {
	AccountGroupID string
	Created        time.Time
	Updated        time.Time
	Deleted        time.Time
	Version        int64
	*NAPData
	*ContactDetails
	Accounts []*Account
	*ListingDistribution
	*ListingSyncPro
	Associations []*Association
	*ExternalIdentifiers
	*SocialURLs
	*HoursOfOperation
	*Snapshots
	*LegacyProductDetails
	*Status
	*RichData
}

// Status represents status fields for the account group
type Status struct {
	Suspended bool
}

// LegacyProductDetails describes legacy fields used for the SM, MS and RM products.
type LegacyProductDetails struct {
	KeyPerson             []string
	ShareOfVoiceService   []string
	FaxNumber             string
	CommonName            []string
	CellNumber            string
	Competitor            []string
	AdminNotes            string
	SeoCategory           []string
	Email                 string
	Place                 string
	Tagline               string
	SubscribedToCampaigns bool
	FieldMask             []string
}

func (lpd *LegacyProductDetails) toProto() *accountgroup_v1.LegacyProductDetails {
	return &accountgroup_v1.LegacyProductDetails{
		KeyPerson:             lpd.KeyPerson,
		ShareOfVoiceService:   lpd.ShareOfVoiceService,
		FaxNumber:             lpd.FaxNumber,
		CommonName:            lpd.CommonName,
		CellNumber:            lpd.CellNumber,
		Competitor:            lpd.Competitor,
		AdminNotes:            lpd.AdminNotes,
		SeoCategory:           lpd.SeoCategory,
		Email:                 lpd.Email,
		Place:                 lpd.Place,
		Tagline:               lpd.Tagline,
		SubscribedToCampaigns: lpd.SubscribedToCampaigns,
	}
}

func (lpd *LegacyProductDetails) toUpdateOperation() (*accountgroup_v1.UpdateOperation, error) {

	return &accountgroup_v1.UpdateOperation{
		Operation: &accountgroup_v1.UpdateOperation_LegacyProductDetails{
			LegacyProductDetails: lpd.toProto(),
		},
		FieldMask: &accountgroup_v1.FieldMask{
			Paths: lpd.FieldMask,
		},
	}, nil
}

// Snapshots is a container for holding a list of snapshots.
type Snapshots struct {
	Snapshots []*Snapshot
}

// Snapshot describes a sales report snapshot at a specific point in time.
type Snapshot struct {
	SnapshotID string
	Created    time.Time
	Expiry     time.Time
}

// Account represents a single account for a product that has been activated on the account group.
type Account struct {
	// Whether the account is a trial account or not
	IsTrial bool

	// Tags on the account
	Tags []string

	// The marketplace app id or the legacy product id the account belongs to
	MarketplaceAppID string

	// Account ID of this account
	AccountID string

	// The date on which the account expires
	Expiry time.Time
}

// Association describes a VBC user -> account group association.
type Association struct {
	Label           string
	ProductID       string
	ProductUserID   string
	VBCUserID       string
	DefaultLocation bool
}

// ListingDistribution holds information about the latest listing distribution order
type ListingDistribution struct {
	OrderID    string
	PurchaseID string
	FromDate   time.Time
	ThruDate   time.Time
	AutoRenew  bool
}

// Geo represents a geo point
type Geo struct {
	Latitude  float64
	Longitude float64
}

// NAPData holds a business NAP data
type NAPData struct {
	CompanyName        string
	Address            string
	City               string
	State              string
	Zip                string
	Country            string
	Website            string
	WorkNumber         []string
	CallTrackingNumber []string
	Location           *Geo
	Timezone           string
	FieldMask          []string
}

func (nap *NAPData) toProto() *accountgroup_v1.AccountGroupLocation {
	var location *accountgroup_v1.Geo
	if nap.Location != nil {
		location = &accountgroup_v1.Geo{
			Latitude:  nap.Location.Latitude,
			Longitude: nap.Location.Longitude,
		}
	}
	return &accountgroup_v1.AccountGroupLocation{
		CompanyName:        nap.CompanyName,
		Address:            nap.Address,
		City:               nap.City,
		State:              nap.State,
		Zip:                nap.Zip,
		Country:            nap.Country,
		Website:            nap.Website,
		WorkNumber:         nap.WorkNumber,
		CallTrackingNumber: nap.CallTrackingNumber,
		Location:           location,
		Timezone:           nap.Timezone,
	}
}

func (nap *NAPData) toUpdateOperation() (*accountgroup_v1.UpdateOperation, error) {

	return &accountgroup_v1.UpdateOperation{
		Operation: &accountgroup_v1.UpdateOperation_AccountGroupNap{
			AccountGroupNap: nap.toProto(),
		},
		FieldMask: &accountgroup_v1.FieldMask{
			Paths: nap.FieldMask,
		},
	}, nil
}

// BillingFrequency indicates the period of listing sync pro billing.
type BillingFrequency int64

const (
	// Monthly indicates periodic monthly billing
	Monthly BillingFrequency = 0
	// Yearly indicates periodic yearly billing
	Yearly BillingFrequency = 1
	// OneTime indicates a single one time purchase that is not periodic.
	OneTime BillingFrequency = 2
)

// ServiceProviders indicates the provider that listing sync pro leverages.
type ServiceProviders int64

const (
	// Uberall provides the syndication
	Uberall ServiceProviders = 0
	// Yext provides the syndication
	Yext ServiceProviders = 1
)

// ListingSyncPro represents the current listing sync pro order.
type ListingSyncPro struct {
	PurchaseDate     time.Time
	BillingFrequency BillingFrequency
	ExpiryDate       time.Time
	Country          string
	DiscountFlag     bool
	ServiceProvider  ServiceProviders
}

// HoursOfOperation holds the business hours of operation
type HoursOfOperation struct {
	HoursOfOperation []*Span
	FieldMask        []string
}

func (hoo *HoursOfOperation) toProto() *accountgroup_v1.HoursOfOperation {
	rv := &accountgroup_v1.HoursOfOperation{
		HoursOfOperation: []*accountgroup_v1.HoursOfOperation_Span{},
	}
	for _, s := range hoo.HoursOfOperation {
		pbs := &accountgroup_v1.HoursOfOperation_Span{
			Closes:      s.Closes,
			DayOfWeek:   []string{},
			Description: s.Description,
			Opens:       s.Opens,
		}
		for _, d := range s.DayOfWeek {
			pbs.DayOfWeek = append(pbs.DayOfWeek, d)
		}
		rv.HoursOfOperation = append(rv.HoursOfOperation, pbs)
	}
	return rv
}

func (hoo *HoursOfOperation) toUpdateOperation() (*accountgroup_v1.UpdateOperation, error) {
	return &accountgroup_v1.UpdateOperation{
		Operation: &accountgroup_v1.UpdateOperation_HoursOfOperation{
			HoursOfOperation: hoo.toProto(),
		},
		FieldMask: &accountgroup_v1.FieldMask{
			Paths: hoo.FieldMask,
		},
	}, nil
}

// Span is specific time the business is order for.
type Span struct {
	DayOfWeek   []string
	Opens       string
	Closes      string
	Description string
}

// ContactDetails holds the information for the business's primary contact.
type ContactDetails struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

// ExternalIdentifiers holds the many account groups identifiers.
type ExternalIdentifiers struct {
	Origin                   string
	JobID                    []string
	CustomerIdentifier       string
	Tags                     []string
	ActionLists              []string
	SocialProfileID          string
	PartnerID                string
	MarketID                 string
	TaxIDs                   []string
	SalesPersonID            string
	AdditionalSalesPersonIDs []string
	SalesforceID             string
	FieldMask                []string
}

func (ei *ExternalIdentifiers) toProto() *accountgroup_v1.AccountGroupExternalIdentifiers {
	rv := &accountgroup_v1.AccountGroupExternalIdentifiers{
		ActionLists:        make([]string, len(ei.ActionLists)),
		CustomerIdentifier: ei.CustomerIdentifier,
		JobId:              make([]string, len(ei.JobID)),
		MarketId:           ei.MarketID,
		Origin:             ei.Origin,
		PartnerId:          ei.PartnerID,
		SalesPersonId:      ei.SalesPersonID,
		SocialProfileId:    ei.SocialProfileID,
		Tags:               make([]string, len(ei.Tags)),
		TaxIds:             make([]string, len(ei.TaxIDs)),
		AdditionalSalesPersonIds: make([]string, len(ei.AdditionalSalesPersonIDs)),
		SalesforceId:             ei.SalesforceID,
	}
	for i, v := range ei.ActionLists {
		rv.ActionLists[i] = v
	}
	for i, v := range ei.JobID {
		rv.JobId[i] = v
	}
	for i, v := range ei.Tags {
		rv.Tags[i] = v
	}
	for i, v := range ei.TaxIDs {
		rv.TaxIds[i] = v
	}
	for i, v := range ei.AdditionalSalesPersonIDs {
		rv.AdditionalSalesPersonIds[i] = v
	}
	return rv
}

func (ei *ExternalIdentifiers) toUpdateOperation() (*accountgroup_v1.UpdateOperation, error) {
	return &accountgroup_v1.UpdateOperation{
		Operation: &accountgroup_v1.UpdateOperation_AccountGroupExternalIdentifiers{
			AccountGroupExternalIdentifiers: ei.toProto(),
		},
		FieldMask: &accountgroup_v1.FieldMask{
			Paths: ei.FieldMask,
		},
	}, nil
}

// SocialURLs contains the business social urls.
type SocialURLs struct {
	GoogleplusURL string
	LinkedinURL   string
	FoursquareURL string
	TwitterURL    string
	FacebookURL   string
	RssURL        string
	YoutubeURL    string
	InstagramURL  string
	PinterestURL  string
	FieldMask     []string
}

func (su *SocialURLs) toProto() *accountgroup_v1.SocialURLs {
	return &accountgroup_v1.SocialURLs{
		FacebookUrl:   su.FacebookURL,
		FoursquareUrl: su.FoursquareURL,
		GoogleplusUrl: su.GoogleplusURL,
		InstagramUrl:  su.InstagramURL,
		LinkedinUrl:   su.LinkedinURL,
		PinterestUrl:  su.PinterestURL,
		RssUrl:        su.RssURL,
		TwitterUrl:    su.TwitterURL,
		YoutubeUrl:    su.YoutubeURL,
	}
}

func (su *SocialURLs) toUpdateOperation() (*accountgroup_v1.UpdateOperation, error) {
	return &accountgroup_v1.UpdateOperation{
		Operation: &accountgroup_v1.UpdateOperation_SocialUrls{
			SocialUrls: su.toProto(),
		},
		FieldMask: &accountgroup_v1.FieldMask{
			Paths: su.FieldMask,
		},
	}, nil
}

// RichData describes the set of extended nap data for an account group.
type RichData struct {
	TollFreeNumber                    string
	Description                       string
	ShortDescription                  string
	ServicesOffered                   []string
	BrandsCarried                     []string
	Landmark                          string
	PaymentMethods                    []PaymentMethod
	CustomFields                      []*CustomField
	HealthCareProfessionalInformation *HealthCareProfessionalInformation
	InferredAttributes                []string
	FieldMask                         []string
}

func (rd *RichData) toProto() *accountgroup_v1.RichData {
	rv := &accountgroup_v1.RichData{
		BrandsCarried:      []string{},
		CustomFields:       []*accountgroup_v1.CustomField{},
		Description:        rd.Description,
		InferredAttributes: []string{},
		Landmark:           rd.Landmark,
		PaymentMethods:     []accountgroup_v1.RichData_PaymentMethods{},
		ServicesOffered:    []string{},
		ShortDescription:   rd.ShortDescription,
		TollFreeNumber:     rd.TollFreeNumber,
	}

	if rd.HealthCareProfessionalInformation != nil {
		rv.HealthCareProfessionalInformation = rd.HealthCareProfessionalInformation.toProto()
	}
	for _, b := range rd.BrandsCarried {
		rv.BrandsCarried = append(rv.BrandsCarried, b)
	}
	for _, i := range rd.InferredAttributes {
		rv.InferredAttributes = append(rv.InferredAttributes, i)
	}
	for _, s := range rd.ServicesOffered {
		rv.ServicesOffered = append(rv.ServicesOffered, s)
	}
	for _, p := range rd.PaymentMethods {
		if val, ok := paymentToProtoMap[p]; ok {
			rv.PaymentMethods = append(rv.PaymentMethods, val)
		}
	}
	for _, c := range rd.CustomFields {
		cf := accountgroup_v1.CustomField{
			Name:  c.Name,
			Value: c.Value,
		}
		rv.CustomFields = append(rv.CustomFields, &cf)
	}

	return rv
}

func (rd *RichData) toUpdateOperation() (*accountgroup_v1.UpdateOperation, error) {
	return &accountgroup_v1.UpdateOperation{
		Operation: &accountgroup_v1.UpdateOperation_RichData{
			RichData: rd.toProto(),
		},
		FieldMask: &accountgroup_v1.FieldMask{
			Paths: rd.FieldMask,
		},
	}, nil
}

// ActionListAppend describe the name of Action List which will be appended to the array of Action List on account group
type ActionListAppend struct {
	ActionListName string
}

func (ala *ActionListAppend) toProto() *accountgroup_v1.ActionListAppend {
	return &accountgroup_v1.ActionListAppend{
		ActionListName: ala.ActionListName,
	}
}

func (ala *ActionListAppend) toUpdateOperation() (*accountgroup_v1.UpdateOperation, error) {
	return &accountgroup_v1.UpdateOperation{
		Operation: &accountgroup_v1.UpdateOperation_ActionListAppend{
			ActionListAppend: ala.toProto(),
		},
	}, nil
}

// PaymentMethod describes the available methods of payment for an account group.
type PaymentMethod int32

// Options for PaymentMethod
const (
	AmericanExpress PaymentMethod = 0
	AndroidPay      PaymentMethod = 1
	ApplePay        PaymentMethod = 2
	Cash            PaymentMethod = 3
	Check           PaymentMethod = 4
	Debit           PaymentMethod = 5
	DinersClub      PaymentMethod = 6
	Discover        PaymentMethod = 7
	Mastercard      PaymentMethod = 8
	Paypal          PaymentMethod = 9
	SamsungPay      PaymentMethod = 10
	StoreCard       PaymentMethod = 11
	TravelersCheck  PaymentMethod = 12
	VISA            PaymentMethod = 13
)

// CustomField describe a custom field stored on the account group.
type CustomField struct {
	Name  string
	Value string
}

// Gender is used for health care information data.
type Gender int32

//Options for Gender
const (
	NotSpecified Gender = 0
	Female       Gender = 1
	Male         Gender = 2
	Other        Gender = 3
)

// HealthCareProfessionalInformation describes the health care data of account group, if applicable.
type HealthCareProfessionalInformation struct {
	DateOfBirth                time.Time
	Email                      string
	Fellowship                 []string
	FirstName                  string
	Gender                     Gender
	Initials                   string
	InsurancesAccepted         []string
	LastName                   string
	MedicalLicenseNumber       string
	NationalProviderIdentifier string
	Office                     string
	ProfessionalCredential     []string
	Residency                  []string
	School                     []string
	Specialty                  []string
	StandardizedTitle          string
	StateLicense               string
	IsTakingPatients           *bool
}

func (hcpi *HealthCareProfessionalInformation) toProto() *accountgroup_v1.HealthCareProfessionalInformation {
	rv := accountgroup_v1.HealthCareProfessionalInformation{
		DateOfBirth:                nil,
		Email:                      hcpi.Email,
		Fellowship:                 []string{},
		FirstName:                  hcpi.FirstName,
		Initials:                   hcpi.Initials,
		InsurancesAccepted:         []string{},
		LastName:                   hcpi.LastName,
		MedicalLicenseNumber:       hcpi.MedicalLicenseNumber,
		NationalProviderIdentifier: hcpi.NationalProviderIdentifier,
		Office:                 hcpi.Office,
		ProfessionalCredential: []string{},
		Residency:              []string{},
		School:                 []string{},
		Specialty:              []string{},
		StandardizedTitle:      hcpi.StandardizedTitle,
		StateLicense:           hcpi.StateLicense,
	}
	if val, ok := genderToProtoMap[hcpi.Gender]; ok {
		rv.Gender = val
	}
	if hcpi.IsTakingPatients == nil {
		rv.IsTakingPatients = nil
	} else {
		rv.IsTakingPatients = &wrappers.BoolValue{Value: *hcpi.IsTakingPatients}
	}
	for _, f := range hcpi.Fellowship {
		rv.Fellowship = append(rv.Fellowship, f)
	}
	for _, i := range hcpi.InsurancesAccepted {
		rv.InsurancesAccepted = append(rv.InsurancesAccepted, i)
	}
	for _, p := range hcpi.ProfessionalCredential {
		rv.ProfessionalCredential = append(rv.ProfessionalCredential, p)
	}
	for _, r := range hcpi.Residency {
		rv.Residency = append(rv.Residency, r)
	}
	for _, s := range hcpi.School {
		rv.School = append(rv.School, s)
	}
	for _, s := range hcpi.Specialty {
		rv.Specialty = append(rv.Specialty, s)
	}
	return &rv
}
