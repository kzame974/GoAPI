package models

import (
	"encoding/hex"
	"fmt"
	"time"
)

type SensorTemperatureModel struct {
	Temperature float64
	Humidity    float64
	Timestamp   string
}

func DecodeHexData(hexData string) (SensorTemperatureModel, error) {
	var sensorData SensorTemperatureModel

	// décoder les données hexadécimales en bytes
	data, err := hex.DecodeString(hexData)
	if err != nil {
		return sensorData, err
	}

	// vérifier la longueur des données décodées
	if len(data) != 16 {
		return sensorData, fmt.Errorf("donnée hexadécimale trop grande, attendue 16 bytes mais rendu %d", len(data))
	}

	// extraire les donnée de température et d'humidité à partir des bytes décodés,
	//on part aussi du principe que le capteur devra TOUJOURS envoyer les données dans cette ordre
	//nb: la location n'est pas une valeur que des capteurs peuvent faire, pas pertinent
	sensorData.Temperature = float64(data[0]) + float64(data[1])/100
	sensorData.Humidity = float64(data[2]) + float64(data[3])/100
	sensorData.Timestamp = time.Now().Format(time.RFC3339)

	return sensorData, nil
}
