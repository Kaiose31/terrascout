package wildlife

type Observation struct {
	ID                       int         `json:"id"`
	ObservedOnString         string      `json:"observed_on_string"`
	ObservedOn               string      `json:"observed_on"`
	TimeObservedAt           string      `json:"time_observed_at"`
	TimeZone                 string      `json:"time_zone"`
	CreatedAt                string      `json:"created_at"`
	UpdatedAt                string      `json:"updated_at"`
	QualityGrade             string      `json:"quality_grade"`
	URL                      string      `json:"url"`
	ImageURL                 string      `json:"image_url"`
	Description              string      `json:"description"`
	CaptiveCultivated        bool        `json:"captive_cultivated"`
	PlaceGuess               string      `json:"place_guess"`
	Latitude                 float64     `json:"latitude"`
	Longitude                float64     `json:"longitude"`
	PositionalAccuracy       interface{} `json:"positional_accuracy"`
	PrivatePlaceGuess        string      `json:"private_place_guess"`
	PrivateLatitude          string      `json:"private_latitude"`
	PrivateLongitude         string      `json:"private_longitude"`
	PublicPositionalAccuracy interface{} `json:"public_positional_accuracy"`
	Geoprivacy               string      `json:"geoprivacy"`
	TaxonGeoprivacy          string      `json:"taxon_geoprivacy"`
	CoordinatesObscured      bool        `json:"coordinates_obscured"`
	PositioningMethod        string      `json:"positioning_method"`
	PositioningDevice        string      `json:"positioning_device"`
	SpeciesGuess             string      `json:"species_guess"`
	ScientificName           string      `json:"scientific_name"`
	CommonName               string      `json:"common_name"`
	IconicTaxonName          string      `json:"iconic_taxon_name"`
	TaxonID                  int         `json:"taxon_id"`
}
