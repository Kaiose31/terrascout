package utils

import (
	"backend/wildlife"
	"fmt"
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

// func FilterObservations2(observations []wildlife.Observation, swle, nele, swlo, nwlo float64) []wildlife.Observation {
// 	var filtered []wildlife.Observation

// 	for _, obs := range observations {
// 		if obs.Latitude > swle && obs.Latitude < nele && obs.Longitude > swlo && obs.Longitude < nwlo {
// 			filtered = append(filtered, obs)
// 		}
// 	}

// 	return filtered
// }

func RunML(imageData []byte) (string, error) {

	result := "species, iconic taxon name"
	return result, nil
}
