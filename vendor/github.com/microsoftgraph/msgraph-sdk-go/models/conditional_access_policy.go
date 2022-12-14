package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessPolicy provides operations to manage the collection of agreementAcceptance entities.
type ConditionalAccessPolicy struct {
    Entity
    // The conditions property
    conditions ConditionalAccessConditionSetable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // Specifies a display name for the conditionalAccessPolicy object.
    displayName *string
    // Specifies the grant controls that must be fulfilled to pass the policy.
    grantControls ConditionalAccessGrantControlsable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the session controls that are enforced after sign-in.
    sessionControls ConditionalAccessSessionControlsable
    // The state property
    state *ConditionalAccessPolicyState
}
// NewConditionalAccessPolicy instantiates a new conditionalAccessPolicy and sets the default values.
func NewConditionalAccessPolicy()(*ConditionalAccessPolicy) {
    m := &ConditionalAccessPolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.conditionalAccessPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateConditionalAccessPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessPolicy(), nil
}
// GetConditions gets the conditions property value. The conditions property
func (m *ConditionalAccessPolicy) GetConditions()(ConditionalAccessConditionSetable) {
    return m.conditions
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
func (m *ConditionalAccessPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description property
func (m *ConditionalAccessPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Specifies a display name for the conditionalAccessPolicy object.
func (m *ConditionalAccessPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conditions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConditionalAccessConditionSetFromDiscriminatorValue , m.SetConditions)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["grantControls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConditionalAccessGrantControlsFromDiscriminatorValue , m.SetGrantControls)
    res["modifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetModifiedDateTime)
    res["sessionControls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConditionalAccessSessionControlsFromDiscriminatorValue , m.SetSessionControls)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseConditionalAccessPolicyState , m.SetState)
    return res
}
// GetGrantControls gets the grantControls property value. Specifies the grant controls that must be fulfilled to pass the policy.
func (m *ConditionalAccessPolicy) GetGrantControls()(ConditionalAccessGrantControlsable) {
    return m.grantControls
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
func (m *ConditionalAccessPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetSessionControls gets the sessionControls property value. Specifies the session controls that are enforced after sign-in.
func (m *ConditionalAccessPolicy) GetSessionControls()(ConditionalAccessSessionControlsable) {
    return m.sessionControls
}
// GetState gets the state property value. The state property
func (m *ConditionalAccessPolicy) GetState()(*ConditionalAccessPolicyState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ConditionalAccessPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("grantControls", m.GetGrantControls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sessionControls", m.GetSessionControls())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConditions sets the conditions property value. The conditions property
func (m *ConditionalAccessPolicy) SetConditions(value ConditionalAccessConditionSetable)() {
    m.conditions = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
func (m *ConditionalAccessPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description property
func (m *ConditionalAccessPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Specifies a display name for the conditionalAccessPolicy object.
func (m *ConditionalAccessPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGrantControls sets the grantControls property value. Specifies the grant controls that must be fulfilled to pass the policy.
func (m *ConditionalAccessPolicy) SetGrantControls(value ConditionalAccessGrantControlsable)() {
    m.grantControls = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Readonly.
func (m *ConditionalAccessPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetSessionControls sets the sessionControls property value. Specifies the session controls that are enforced after sign-in.
func (m *ConditionalAccessPolicy) SetSessionControls(value ConditionalAccessSessionControlsable)() {
    m.sessionControls = value
}
// SetState sets the state property value. The state property
func (m *ConditionalAccessPolicy) SetState(value *ConditionalAccessPolicyState)() {
    m.state = value
}
