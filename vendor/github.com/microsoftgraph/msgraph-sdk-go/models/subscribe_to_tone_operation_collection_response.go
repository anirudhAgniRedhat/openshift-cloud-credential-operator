package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubscribeToToneOperationCollectionResponse 
type SubscribeToToneOperationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSubscribeToToneOperationCollectionResponse instantiates a new SubscribeToToneOperationCollectionResponse and sets the default values.
func NewSubscribeToToneOperationCollectionResponse()(*SubscribeToToneOperationCollectionResponse) {
    m := &SubscribeToToneOperationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSubscribeToToneOperationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubscribeToToneOperationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubscribeToToneOperationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubscribeToToneOperationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubscribeToToneOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubscribeToToneOperationable, len(val))
            for i, v := range val {
                res[i] = v.(SubscribeToToneOperationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SubscribeToToneOperationCollectionResponse) GetValue()([]SubscribeToToneOperationable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SubscribeToToneOperationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SubscribeToToneOperationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *SubscribeToToneOperationCollectionResponse) SetValue(value []SubscribeToToneOperationable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SubscribeToToneOperationCollectionResponseable 
type SubscribeToToneOperationCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]SubscribeToToneOperationable)
    SetValue(value []SubscribeToToneOperationable)()
}
