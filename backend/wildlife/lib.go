package wildlife

import "time"

type Observation struct {
	ID                       int       `json:"id"`
	ObservedOnString         string    `json:"observed_on_string"`
	ObservedOn               time.Time `json:"observed_on"`
	TimeObservedAt           time.Time `json:"time_observed_at"`
	TimeZone                 string    `json:"time_zone"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	QualityGrade             string    `json:"quality_grade"`
	URL                      string    `json:"url"`
	ImageURL                 string    `json:"image_url"`
	Description              string    `json:"description"`
	CaptiveCultivated        bool      `json:"captive_cultivated"`
	OAuthApplicationID       int       `json:"oauth_application_id"`
	PlaceGuess               string    `json:"place_guess"`
	Latitude                 float64   `json:"latitude"`
	Longitude                float64   `json:"longitude"`
	PositionalAccuracy       float64   `json:"positional_accuracy"`
	PrivatePlaceGuess        string    `json:"private_place_guess"`
	PrivateLatitude          float64   `json:"private_latitude"`
	PrivateLongitude         float64   `json:"private_longitude"`
	PublicPositionalAccuracy float64   `json:"public_positional_accuracy"`
	Geoprivacy               string    `json:"geoprivacy"`
	TaxonGeoprivacy          string    `json:"taxon_geoprivacy"`
	CoordinatesObscured      bool      `json:"coordinates_obscured"`
	PositioningMethod        string    `json:"positioning_method"`
	PositioningDevice        string    `json:"positioning_device"`
	SpeciesGuess             string    `json:"species_guess"`
	ScientificName           string    `json:"scientific_name"`
	CommonName               string    `json:"common_name"`
	IconicTaxonName          string    `json:"iconic_taxon_name"`
	TaxonID                  int       `json:"taxon_id"`
}
