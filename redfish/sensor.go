//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	
	"github.com/apex/log"

	"github.com/bcohee/gofish/common"
)

type redfishSensorThresholdReadingType struct {
	Reading         float64
}

type redfishSensorThresholdType struct {
	LowerCaution     redfishSensorThresholdReadingType
	LowerCritical    redfishSensorThresholdReadingType
	UpperCaution     redfishSensorThresholdReadingType
	UpperCritical    redfishSensorThresholdReadingType
}

type redfishSensorType struct {
	common.Entity
	Reading         float64
	ReadingRangeMax float64
	ReadingRangeMin float64
	ReadingType     string
	ReadingUnits    string
	Status          common.Status
	Thresholds      redfishSensorThresholdType
}

type redfishSensorMembers struct {
	CPU                        redfishSensorType
	Memory                     redfishSensorType
	Storage_Internal           redfishSensorType
	Storage_RW                 redfishSensorType
	chassis_eﬀiciency          redfishSensorType
	chassis_input_current      redfishSensorType
	chassis_input_power        redfishSensorType
	chassis_input_voltage      redfishSensorType
	chassis_output_current     redfishSensorType
	chassis_output_power       redfishSensorType
	chassis_output_voltage     redfishSensorType
	chassis_temperature        redfishSensorType
	p0_ambient                 redfishSensorType
	p0_exhaust                 redfishSensorType
	p0_fan1                    redfishSensorType
	p0_hotspot                 redfishSensorType
	p0_iin                     redfishSensorType
	p0_iout                    redfishSensorType
	p0_pin                     redfishSensorType
	p0_pout                    redfishSensorType
	p0_vin                     redfishSensorType
	p0_vout                    redfishSensorType

	p1_ambient                 redfishSensorType
	p1_exhaust                 redfishSensorType
	p1_fan1                    redfishSensorType
	p1_hotspot                 redfishSensorType
	p1_iin                     redfishSensorType
	p1_iout                    redfishSensorType
	p1_pin                     redfishSensorType
	p1_pout                    redfishSensorType
	p1_vin                     redfishSensorType
	p1_vout                    redfishSensorType

	p2_ambient                 redfishSensorType
	p2_exhaust                 redfishSensorType
	p2_fan1                    redfishSensorType
	p2_hotspot                 redfishSensorType
	p2_iin                     redfishSensorType
	p2_iout                    redfishSensorType
	p2_pin                     redfishSensorType
	p2_pout                    redfishSensorType
	p2_vin                     redfishSensorType
	p2_vout                    redfishSensorType

	p3_ambient                 redfishSensorType
	p3_exhaust                 redfishSensorType
	p3_fan1                    redfishSensorType
	p3_hotspot                 redfishSensorType
	p3_iin                     redfishSensorType
	p3_iout                    redfishSensorType
	p3_pin                     redfishSensorType
	p3_pout                    redfishSensorType
	p3_vin                     redfishSensorType
	p3_vout                    redfishSensorType

	p4_ambient                 redfishSensorType
	p4_exhaust                 redfishSensorType
	p4_fan1                    redfishSensorType
	p4_hotspot                 redfishSensorType
	p4_iin                     redfishSensorType
	p4_iout                    redfishSensorType
	p4_pin                     redfishSensorType
	p4_pout                    redfishSensorType
	p4_vin                     redfishSensorType
	p4_vout                    redfishSensorType

	p5_ambient                 redfishSensorType
	p5_exhaust                 redfishSensorType
	p5_fan1                    redfishSensorType
	p5_hotspot                 redfishSensorType
	p5_iin                     redfishSensorType
	p5_iout                    redfishSensorType
	p5_pin                     redfishSensorType
	p5_pout                    redfishSensorType
	p5_vin                     redfishSensorType
	p5_vout                    redfishSensorType
}

// Sensor is used to represent a *custom LITEON PMC + PSU sensor metrics resource for a Redfish
// implementation.
type Sensor struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	Members redfishSensorMembers

	Oem json.RawMessage
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals an object from the raw JSON.
func (sensor *Sensor) UnmarshalJSON(b []byte) error {
	type temp Sensor
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sensor = Sensor(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	sensor.rawData = b

	return nil
}

// // Update commits updates to this object's properties to the running system.
// func (sensor *Sensor) Update() error {

// 	// Get a representation of the object's original state so we can find what
// 	// to update.
// 	original := new(Sensor)
// 	original.UnmarshalJSON(sensor.rawData)

// 	readWriteFields := []string{
// 		"SensorFans",
// 		"SensorTemperatures",
// 	}

// 	originalElement := reflect.ValueOf(original).Elem()
// 	currentElement := reflect.ValueOf(sensor).Elem()

// 	return sensor.Entity.Update(originalElement, currentElement, readWriteFields)
// }

// GetSensor will get a Sensor instance from the service.
func GetSensor(c common.Client, uri string) (*Sensor, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var sensor Sensor
	err = json.NewDecoder(resp.Body).Decode(&sensor)
	if err != nil {
		return nil, err
	}

	sensor.SetClient(c)
	return &sensor, nil
}

// ListReferencedSensors gets the collection of Sensor from a provided reference.
func ListReferencedSensors(c common.Client, link string) ([]*Sensor, error) { //nolint:dupl
	var result []*Sensor
	log.Debugf("gofish/sensor/ListReferencedSensors: link = %s", link)
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, sensorLink := range links.ItemLinks {
		sensor, err := GetSensor(c, sensorLink)
		if err != nil {
			collectionError.Failures[sensorLink] = err
		} else {
			result = append(result, sensor)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
