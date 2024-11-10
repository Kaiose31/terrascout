package utils

import (
	"backend/wildlife"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

func ConvertStringsToFloats(strs []string) ([]float64, error) {
	var floats []float64

	for _, str := range strs {
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to float64: %v", str, err)
		}
		floats = append(floats, floatVal)
	}

	return floats, nil
}

func FilterObservations(observations *[]wildlife.Observation, swle, swlo, nele, nelo float64) {
	obs := *observations
	n := 0

	for i := 0; i < len(obs); i++ {
		if obs[i].Latitude >= swle && obs[i].Latitude <= nele && obs[i].Longitude >= swlo && obs[i].Longitude <= nelo {
			obs[n] = obs[i]
			n++
		}
	}

	*observations = obs[:n]
}

func FilterObsByName(observations *[]wildlife.Observation, scientificName string) {
	obs := *observations
	n := 0

	for i := 0; i < len(obs); i++ {
		if obs[i].ScientificName == scientificName {
			obs[n] = obs[i]
			n++
		}
	}

	*observations = obs[:n]
}

func RunML(imageData []byte) (string, error) {

	var requestBody bytes.Buffer

	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("image", "image.jpg")
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}

	_, err = io.Copy(part, bytes.NewReader(imageData))
	if err != nil {
		return "", fmt.Errorf("failed to write image data: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}
	resp, err := http.Post("http://127.0.0.1:5001/inference", writer.FormDataContentType(), &requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to send post request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(body), nil
}

func FilterTaxon(taxa *[]wildlife.Taxon, taxon_id int) {
	obs := *taxa
	n := 0

	for i := 0; i < len(obs); i++ {
		if obs[i].TaxonID == taxon_id {
			obs[n] = obs[i]
			n++
		}
	}

	*taxa = obs[:n]
}
