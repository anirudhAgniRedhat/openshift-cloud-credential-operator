package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceGeoLocation device location
type DeviceGeoLocation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceGeoLocation instantiates a new deviceGeoLocation and sets the default values.
func NewDeviceGeoLocation()(*DeviceGeoLocation) {
    m := &DeviceGeoLocation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceGeoLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceGeoLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceGeoLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceGeoLocation) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAltitude gets the altitude property value. Altitude, given in meters above sea level
func (m *DeviceGeoLocation) GetAltitude()(*float64) {
    val, err := m.GetBackingStore().Get("altitude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DeviceGeoLocation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceGeoLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["altitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltitude(val)
        }
        return nil
    }
    res["heading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeading(val)
        }
        return nil
    }
    res["horizontalAccuracy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHorizontalAccuracy(val)
        }
        return nil
    }
    res["lastCollectedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCollectedDateTime(val)
        }
        return nil
    }
    res["latitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatitude(val)
        }
        return nil
    }
    res["longitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongitude(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["speed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeed(val)
        }
        return nil
    }
    res["verticalAccuracy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerticalAccuracy(val)
        }
        return nil
    }
    return res
}
// GetHeading gets the heading property value. Heading in degrees from true north
func (m *DeviceGeoLocation) GetHeading()(*float64) {
    val, err := m.GetBackingStore().Get("heading")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetHorizontalAccuracy gets the horizontalAccuracy property value. Accuracy of longitude and latitude in meters
func (m *DeviceGeoLocation) GetHorizontalAccuracy()(*float64) {
    val, err := m.GetBackingStore().Get("horizontalAccuracy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetLastCollectedDateTime gets the lastCollectedDateTime property value. Time at which location was recorded, relative to UTC
func (m *DeviceGeoLocation) GetLastCollectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCollectedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLatitude gets the latitude property value. Latitude coordinate of the device's location
func (m *DeviceGeoLocation) GetLatitude()(*float64) {
    val, err := m.GetBackingStore().Get("latitude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetLongitude gets the longitude property value. Longitude coordinate of the device's location
func (m *DeviceGeoLocation) GetLongitude()(*float64) {
    val, err := m.GetBackingStore().Get("longitude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceGeoLocation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSpeed gets the speed property value. Speed the device is traveling in meters per second
func (m *DeviceGeoLocation) GetSpeed()(*float64) {
    val, err := m.GetBackingStore().Get("speed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetVerticalAccuracy gets the verticalAccuracy property value. Accuracy of altitude in meters
func (m *DeviceGeoLocation) GetVerticalAccuracy()(*float64) {
    val, err := m.GetBackingStore().Get("verticalAccuracy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceGeoLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("altitude", m.GetAltitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("heading", m.GetHeading())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("horizontalAccuracy", m.GetHorizontalAccuracy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastCollectedDateTime", m.GetLastCollectedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("longitude", m.GetLongitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("speed", m.GetSpeed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("verticalAccuracy", m.GetVerticalAccuracy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceGeoLocation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAltitude sets the altitude property value. Altitude, given in meters above sea level
func (m *DeviceGeoLocation) SetAltitude(value *float64)() {
    err := m.GetBackingStore().Set("altitude", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceGeoLocation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetHeading sets the heading property value. Heading in degrees from true north
func (m *DeviceGeoLocation) SetHeading(value *float64)() {
    err := m.GetBackingStore().Set("heading", value)
    if err != nil {
        panic(err)
    }
}
// SetHorizontalAccuracy sets the horizontalAccuracy property value. Accuracy of longitude and latitude in meters
func (m *DeviceGeoLocation) SetHorizontalAccuracy(value *float64)() {
    err := m.GetBackingStore().Set("horizontalAccuracy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastCollectedDateTime sets the lastCollectedDateTime property value. Time at which location was recorded, relative to UTC
func (m *DeviceGeoLocation) SetLastCollectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCollectedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLatitude sets the latitude property value. Latitude coordinate of the device's location
func (m *DeviceGeoLocation) SetLatitude(value *float64)() {
    err := m.GetBackingStore().Set("latitude", value)
    if err != nil {
        panic(err)
    }
}
// SetLongitude sets the longitude property value. Longitude coordinate of the device's location
func (m *DeviceGeoLocation) SetLongitude(value *float64)() {
    err := m.GetBackingStore().Set("longitude", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceGeoLocation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSpeed sets the speed property value. Speed the device is traveling in meters per second
func (m *DeviceGeoLocation) SetSpeed(value *float64)() {
    err := m.GetBackingStore().Set("speed", value)
    if err != nil {
        panic(err)
    }
}
// SetVerticalAccuracy sets the verticalAccuracy property value. Accuracy of altitude in meters
func (m *DeviceGeoLocation) SetVerticalAccuracy(value *float64)() {
    err := m.GetBackingStore().Set("verticalAccuracy", value)
    if err != nil {
        panic(err)
    }
}
// DeviceGeoLocationable 
type DeviceGeoLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAltitude()(*float64)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetHeading()(*float64)
    GetHorizontalAccuracy()(*float64)
    GetLastCollectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLatitude()(*float64)
    GetLongitude()(*float64)
    GetOdataType()(*string)
    GetSpeed()(*float64)
    GetVerticalAccuracy()(*float64)
    SetAltitude(value *float64)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetHeading(value *float64)()
    SetHorizontalAccuracy(value *float64)()
    SetLastCollectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLatitude(value *float64)()
    SetLongitude(value *float64)()
    SetOdataType(value *string)()
    SetSpeed(value *float64)()
    SetVerticalAccuracy(value *float64)()
}
