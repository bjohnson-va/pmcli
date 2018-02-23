package subject

import (
	"github.com/vendasta/gosdks/iam/subjectcontext"
)

func ToSalesPerson(c *subjectcontext.Context, s *subject) *SalesPerson {
	salesPerson := &SalesPerson{
		subject:   s,
		PartnerID: c.Namespace,
	}

	return salesPerson
}

// SalesPerson represents a Sales and Success Center user.
type SalesPerson struct {
	*subject

	PartnerID string

	FirstName string `attribute:"first_name"`
	LastName  string `attribute:"last_name"`

	JobTitle   string `attribute:"job_title"`
	Department string `attribute:"department"`
	Region     string `attribute:"region"`
	Supervisor string `attribute:"supervisor"`
	MarketID   string `attribute:"market_id"`

	PhoneNumbers            []string `attribute:"phone_numbers"`
	Address                 string   `attribute:"address"`
	City                    string   `attribute:"city"`
	State                   string   `attribute:"state"`
	Country                 string   `attribute:"country"`
	ZipCode                 string   `attribute:"zip_code"`
	TitleChoice             string   `attribute:"title_choice"`
	PictureName             string   `attribute:"picture_name"`
	PictureServingURL       string   `attribute:"picture_serving_url"`
	PictureServingURLSecure string   `attribute:"picture_serving_url_secure"`
}

func (s *SalesPerson) Context() *subjectcontext.Context {
	return subjectcontext.New("sales_person", s.PartnerID)
}

// NewSalesPerson creates a new SalesPerson struct
func NewSalesPerson(s *subject, PartnerID, FirstName, LastName, JobTitle, Department, Region, Supervisor,
					MarketID string, PhoneNumbers []string, Address, City, State, Country, ZipCode, TitleChoice,
					PictureName, PictureServingURL, PictureServingURLSecure string) *SalesPerson {
	return &SalesPerson{
		subject:   s,
		PartnerID: PartnerID,
		FirstName: FirstName,
		LastName: LastName,
		JobTitle: JobTitle,
		Department: Department,
		Region: Region,
		Supervisor: Supervisor,
		MarketID: MarketID,
		PhoneNumbers: PhoneNumbers,
		Address: Address,
		City: City,
		State: State,
		Country: Country,
		ZipCode: ZipCode,
		TitleChoice: TitleChoice,
		PictureName: PictureName,
		PictureServingURL: PictureServingURL,
		PictureServingURLSecure: PictureServingURLSecure,
	}
}

