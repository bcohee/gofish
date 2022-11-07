//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var sensorBody = `{
		"@odata.context": "/redfish/v1/$metadata#Sensor.Sensor",
		"@odata.type": "#Sensor.v1_0_0.Sensor",
		"@odata.id": "/redfish/v1/Sensor",
		"Id": "Sensor-1",
		"Name": "SensorOne",
		"Description": "Sensor One",
		"Fans": [{
			"Id": "Fan1",
			"FanName": "Fan One",
			"Assembly": {
				"@odata.id": "/redfish/v1/Assemblies/1"
			},
			"HotPluggable": true,
			"IndicatorLED": "Lit",
			"LowerThresholdCritical": 10,
			"LowerThresholdFatal": 0,
			"LowerThresholdNonCritical": 11,
			"Manufacturer": "Acme Fans",
			"MaxReadingRange": 100,
			"MemberId": "Fan1",
			"MinReadingRange": 10,
			"Model": "Fan2000",
			"Name": "Charlie",
			"PartNumber": "F123",
			"PhysicalContext": "Exhaust",
			"Reading": 1000,
			"ReadingUnits": "RPM",
			"Redundancy": [],
			"Redundancy@odata.count": 0,
			"RelatedItem": [],
			"RelatedItem@odata.count": 0,
			"SensorNumber": 1,
			"SerialNumber": "12345",
			"SparePartNumber": "F120",
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			},
			"UpperThresholdCritical": 9999,
			"UpperThresholdFatal": 10000,
			"UpperThresholdNonCritical": 9998
		}],
		"Fan@odata.count": 1,
		"Redundancy": [],
		"Redundancy@odata.count": 0,
		"Status": {
			"State": "Enabled",
			"Health": "OK"
		},
		"Temperatures": [{
			"@odata.id": "/redfish/v1/Temp",
			"Id": "Temp1",
			"AdjustedMaxAllowableOperatingValue": 60,
			"AdjustedMinAllowableOperatingValue": 1,
			"DeltaPhysicalContext": "Exhaust",
			"DeltaReadingCelsius": 35,
			"LowerThresholdCritical": 1,
			"LowerThresholdFatal": 0,
			"LowerThresholdNonCritical": 2,
			"MaxAllowableOperatingValue": 45,
			"MaxReadingRangeTemp": 45,
			"MemberId": "Thermal1",
			"MinAllowableOperatingValue": -5,
			"MinReadingRangeTemp": -12,
			"Name": "Thermal Temp One",
			"PhysicalContext": "Exhaust",
			"ReadingCelsius": 32,
			"RelatedItem": [],
			"RelatedItem@odata.count": 0,
			"SensorNumber": 1,
			"Status": {
				"State": "Enabled",
				"Health": "OK"
			},
			"UpperThresholdCritical": 9999,
			"UpperThresholdFatal": 10000,
			"UpperThresholdNonCritical": 9998
		}],
		"Temperatures@odata.count": 1
	}`

// TestSensor tests the parsing of Sensor objects.
func TestSensor(t *testing.T) {
	var result Sensor
	err := json.NewDecoder(strings.NewReader(sensorBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Sensor-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "SensorOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.Fans[0].Name != "Fan One" {
		t.Errorf("Invalid fan name: %s", result.Fans[0].Name)
	}
}
