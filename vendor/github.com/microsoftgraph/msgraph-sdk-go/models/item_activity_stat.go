package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActivityStat 
type ItemActivityStat struct {
    Entity
    // Statistics about the access actions in this interval. Read-only.
    access ItemActionStatable
    // Exposes the itemActivities represented in this itemActivityStat resource.
    activities []ItemActivityable
    // Statistics about the create actions in this interval. Read-only.
    create ItemActionStatable
    // Statistics about the delete actions in this interval. Read-only.
    delete ItemActionStatable
    // Statistics about the edit actions in this interval. Read-only.
    edit ItemActionStatable
    // When the interval ends. Read-only.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates that the statistics in this interval are based on incomplete data. Read-only.
    incompleteData IncompleteDataable
    // Indicates whether the item is 'trending.' Read-only.
    isTrending *bool
    // Statistics about the move actions in this interval. Read-only.
    move ItemActionStatable
    // When the interval starts. Read-only.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewItemActivityStat instantiates a new itemActivityStat and sets the default values.
func NewItemActivityStat()(*ItemActivityStat) {
    m := &ItemActivityStat{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.itemActivityStat";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateItemActivityStatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActivityStatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActivityStat(), nil
}
// GetAccess gets the access property value. Statistics about the access actions in this interval. Read-only.
func (m *ItemActivityStat) GetAccess()(ItemActionStatable) {
    return m.access
}
// GetActivities gets the activities property value. Exposes the itemActivities represented in this itemActivityStat resource.
func (m *ItemActivityStat) GetActivities()([]ItemActivityable) {
    return m.activities
}
// GetCreate gets the create property value. Statistics about the create actions in this interval. Read-only.
func (m *ItemActivityStat) GetCreate()(ItemActionStatable) {
    return m.create
}
// GetDelete gets the delete property value. Statistics about the delete actions in this interval. Read-only.
func (m *ItemActivityStat) GetDelete()(ItemActionStatable) {
    return m.delete
}
// GetEdit gets the edit property value. Statistics about the edit actions in this interval. Read-only.
func (m *ItemActivityStat) GetEdit()(ItemActionStatable) {
    return m.edit
}
// GetEndDateTime gets the endDateTime property value. When the interval ends. Read-only.
func (m *ItemActivityStat) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActivityStat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["access"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemActionStatFromDiscriminatorValue , m.SetAccess)
    res["activities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemActivityFromDiscriminatorValue , m.SetActivities)
    res["create"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemActionStatFromDiscriminatorValue , m.SetCreate)
    res["delete"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemActionStatFromDiscriminatorValue , m.SetDelete)
    res["edit"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemActionStatFromDiscriminatorValue , m.SetEdit)
    res["endDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEndDateTime)
    res["incompleteData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIncompleteDataFromDiscriminatorValue , m.SetIncompleteData)
    res["isTrending"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsTrending)
    res["move"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemActionStatFromDiscriminatorValue , m.SetMove)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    return res
}
// GetIncompleteData gets the incompleteData property value. Indicates that the statistics in this interval are based on incomplete data. Read-only.
func (m *ItemActivityStat) GetIncompleteData()(IncompleteDataable) {
    return m.incompleteData
}
// GetIsTrending gets the isTrending property value. Indicates whether the item is 'trending.' Read-only.
func (m *ItemActivityStat) GetIsTrending()(*bool) {
    return m.isTrending
}
// GetMove gets the move property value. Statistics about the move actions in this interval. Read-only.
func (m *ItemActivityStat) GetMove()(ItemActionStatable) {
    return m.move
}
// GetStartDateTime gets the startDateTime property value. When the interval starts. Read-only.
func (m *ItemActivityStat) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *ItemActivityStat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("access", m.GetAccess())
        if err != nil {
            return err
        }
    }
    if m.GetActivities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetActivities())
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("create", m.GetCreate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("edit", m.GetEdit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("incompleteData", m.GetIncompleteData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTrending", m.GetIsTrending())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("move", m.GetMove())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccess sets the access property value. Statistics about the access actions in this interval. Read-only.
func (m *ItemActivityStat) SetAccess(value ItemActionStatable)() {
    m.access = value
}
// SetActivities sets the activities property value. Exposes the itemActivities represented in this itemActivityStat resource.
func (m *ItemActivityStat) SetActivities(value []ItemActivityable)() {
    m.activities = value
}
// SetCreate sets the create property value. Statistics about the create actions in this interval. Read-only.
func (m *ItemActivityStat) SetCreate(value ItemActionStatable)() {
    m.create = value
}
// SetDelete sets the delete property value. Statistics about the delete actions in this interval. Read-only.
func (m *ItemActivityStat) SetDelete(value ItemActionStatable)() {
    m.delete = value
}
// SetEdit sets the edit property value. Statistics about the edit actions in this interval. Read-only.
func (m *ItemActivityStat) SetEdit(value ItemActionStatable)() {
    m.edit = value
}
// SetEndDateTime sets the endDateTime property value. When the interval ends. Read-only.
func (m *ItemActivityStat) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetIncompleteData sets the incompleteData property value. Indicates that the statistics in this interval are based on incomplete data. Read-only.
func (m *ItemActivityStat) SetIncompleteData(value IncompleteDataable)() {
    m.incompleteData = value
}
// SetIsTrending sets the isTrending property value. Indicates whether the item is 'trending.' Read-only.
func (m *ItemActivityStat) SetIsTrending(value *bool)() {
    m.isTrending = value
}
// SetMove sets the move property value. Statistics about the move actions in this interval. Read-only.
func (m *ItemActivityStat) SetMove(value ItemActionStatable)() {
    m.move = value
}
// SetStartDateTime sets the startDateTime property value. When the interval starts. Read-only.
func (m *ItemActivityStat) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
