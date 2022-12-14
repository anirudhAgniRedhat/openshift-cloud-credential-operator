package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingBusiness 
type BookingBusiness struct {
    Entity
    // The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
    address PhysicalAddressable
    // All the appointments of this business. Read-only. Nullable.
    appointments []BookingAppointmentable
    // The hours of operation for the business.
    businessHours []BookingWorkHoursable
    // The type of business.
    businessType *string
    // The set of appointments of this business in a specified date range. Read-only. Nullable.
    calendarView []BookingAppointmentable
    // All the customers of this business. Read-only. Nullable.
    customers []BookingCustomerBaseable
    // All the custom questions of this business. Read-only. Nullable.
    customQuestions []BookingCustomQuestionable
    // The code for the currency that the business operates in on Microsoft Bookings.
    defaultCurrencyIso *string
    // The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
    displayName *string
    // The email address for the business.
    email *string
    // The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
    isPublished *bool
    // The languageTag property
    languageTag *string
    // The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
    phone *string
    // The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
    publicUrl *string
    // Specifies how bookings can be created for this business.
    schedulingPolicy BookingSchedulingPolicyable
    // All the services offered by this business. Read-only. Nullable.
    services []BookingServiceable
    // All the staff members that provide services in this business. Read-only. Nullable.
    staffMembers []BookingStaffMemberBaseable
    // The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
    webSiteUrl *string
}
// NewBookingBusiness instantiates a new BookingBusiness and sets the default values.
func NewBookingBusiness()(*BookingBusiness) {
    m := &BookingBusiness{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.bookingBusiness";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBookingBusinessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingBusinessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingBusiness(), nil
}
// GetAddress gets the address property value. The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingBusiness) GetAddress()(PhysicalAddressable) {
    return m.address
}
// GetAppointments gets the appointments property value. All the appointments of this business. Read-only. Nullable.
func (m *BookingBusiness) GetAppointments()([]BookingAppointmentable) {
    return m.appointments
}
// GetBusinessHours gets the businessHours property value. The hours of operation for the business.
func (m *BookingBusiness) GetBusinessHours()([]BookingWorkHoursable) {
    return m.businessHours
}
// GetBusinessType gets the businessType property value. The type of business.
func (m *BookingBusiness) GetBusinessType()(*string) {
    return m.businessType
}
// GetCalendarView gets the calendarView property value. The set of appointments of this business in a specified date range. Read-only. Nullable.
func (m *BookingBusiness) GetCalendarView()([]BookingAppointmentable) {
    return m.calendarView
}
// GetCustomers gets the customers property value. All the customers of this business. Read-only. Nullable.
func (m *BookingBusiness) GetCustomers()([]BookingCustomerBaseable) {
    return m.customers
}
// GetCustomQuestions gets the customQuestions property value. All the custom questions of this business. Read-only. Nullable.
func (m *BookingBusiness) GetCustomQuestions()([]BookingCustomQuestionable) {
    return m.customQuestions
}
// GetDefaultCurrencyIso gets the defaultCurrencyIso property value. The code for the currency that the business operates in on Microsoft Bookings.
func (m *BookingBusiness) GetDefaultCurrencyIso()(*string) {
    return m.defaultCurrencyIso
}
// GetDisplayName gets the displayName property value. The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
func (m *BookingBusiness) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email address for the business.
func (m *BookingBusiness) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingBusiness) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePhysicalAddressFromDiscriminatorValue , m.SetAddress)
    res["appointments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingAppointmentFromDiscriminatorValue , m.SetAppointments)
    res["businessHours"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingWorkHoursFromDiscriminatorValue , m.SetBusinessHours)
    res["businessType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBusinessType)
    res["calendarView"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingAppointmentFromDiscriminatorValue , m.SetCalendarView)
    res["customers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingCustomerBaseFromDiscriminatorValue , m.SetCustomers)
    res["customQuestions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingCustomQuestionFromDiscriminatorValue , m.SetCustomQuestions)
    res["defaultCurrencyIso"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultCurrencyIso)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["email"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmail)
    res["isPublished"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsPublished)
    res["languageTag"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLanguageTag)
    res["phone"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPhone)
    res["publicUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublicUrl)
    res["schedulingPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateBookingSchedulingPolicyFromDiscriminatorValue , m.SetSchedulingPolicy)
    res["services"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingServiceFromDiscriminatorValue , m.SetServices)
    res["staffMembers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateBookingStaffMemberBaseFromDiscriminatorValue , m.SetStaffMembers)
    res["webSiteUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebSiteUrl)
    return res
}
// GetIsPublished gets the isPublished property value. The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
func (m *BookingBusiness) GetIsPublished()(*bool) {
    return m.isPublished
}
// GetLanguageTag gets the languageTag property value. The languageTag property
func (m *BookingBusiness) GetLanguageTag()(*string) {
    return m.languageTag
}
// GetPhone gets the phone property value. The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) GetPhone()(*string) {
    return m.phone
}
// GetPublicUrl gets the publicUrl property value. The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
func (m *BookingBusiness) GetPublicUrl()(*string) {
    return m.publicUrl
}
// GetSchedulingPolicy gets the schedulingPolicy property value. Specifies how bookings can be created for this business.
func (m *BookingBusiness) GetSchedulingPolicy()(BookingSchedulingPolicyable) {
    return m.schedulingPolicy
}
// GetServices gets the services property value. All the services offered by this business. Read-only. Nullable.
func (m *BookingBusiness) GetServices()([]BookingServiceable) {
    return m.services
}
// GetStaffMembers gets the staffMembers property value. All the staff members that provide services in this business. Read-only. Nullable.
func (m *BookingBusiness) GetStaffMembers()([]BookingStaffMemberBaseable) {
    return m.staffMembers
}
// GetWebSiteUrl gets the webSiteUrl property value. The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
func (m *BookingBusiness) GetWebSiteUrl()(*string) {
    return m.webSiteUrl
}
// Serialize serializes information the current object
func (m *BookingBusiness) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    if m.GetAppointments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppointments())
        err = writer.WriteCollectionOfObjectValues("appointments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessHours() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetBusinessHours())
        err = writer.WriteCollectionOfObjectValues("businessHours", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("businessType", m.GetBusinessType())
        if err != nil {
            return err
        }
    }
    if m.GetCalendarView() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCalendarView())
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomers())
        err = writer.WriteCollectionOfObjectValues("customers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomQuestions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomQuestions())
        err = writer.WriteCollectionOfObjectValues("customQuestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultCurrencyIso", m.GetDefaultCurrencyIso())
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
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedulingPolicy", m.GetSchedulingPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetServices())
        err = writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStaffMembers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetStaffMembers())
        err = writer.WriteCollectionOfObjectValues("staffMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webSiteUrl", m.GetWebSiteUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The street address of the business. The address property, together with phone and webSiteUrl, appear in the footer of a business scheduling page. The attribute type of physicalAddress is not supported in v1.0. Internally we map the addresses to the type others.
func (m *BookingBusiness) SetAddress(value PhysicalAddressable)() {
    m.address = value
}
// SetAppointments sets the appointments property value. All the appointments of this business. Read-only. Nullable.
func (m *BookingBusiness) SetAppointments(value []BookingAppointmentable)() {
    m.appointments = value
}
// SetBusinessHours sets the businessHours property value. The hours of operation for the business.
func (m *BookingBusiness) SetBusinessHours(value []BookingWorkHoursable)() {
    m.businessHours = value
}
// SetBusinessType sets the businessType property value. The type of business.
func (m *BookingBusiness) SetBusinessType(value *string)() {
    m.businessType = value
}
// SetCalendarView sets the calendarView property value. The set of appointments of this business in a specified date range. Read-only. Nullable.
func (m *BookingBusiness) SetCalendarView(value []BookingAppointmentable)() {
    m.calendarView = value
}
// SetCustomers sets the customers property value. All the customers of this business. Read-only. Nullable.
func (m *BookingBusiness) SetCustomers(value []BookingCustomerBaseable)() {
    m.customers = value
}
// SetCustomQuestions sets the customQuestions property value. All the custom questions of this business. Read-only. Nullable.
func (m *BookingBusiness) SetCustomQuestions(value []BookingCustomQuestionable)() {
    m.customQuestions = value
}
// SetDefaultCurrencyIso sets the defaultCurrencyIso property value. The code for the currency that the business operates in on Microsoft Bookings.
func (m *BookingBusiness) SetDefaultCurrencyIso(value *string)() {
    m.defaultCurrencyIso = value
}
// SetDisplayName sets the displayName property value. The name of the business, which interfaces with customers. This name appears at the top of the business scheduling page.
func (m *BookingBusiness) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email address for the business.
func (m *BookingBusiness) SetEmail(value *string)() {
    m.email = value
}
// SetIsPublished sets the isPublished property value. The scheduling page has been made available to external customers. Use the publish and unpublish actions to set this property. Read-only.
func (m *BookingBusiness) SetIsPublished(value *bool)() {
    m.isPublished = value
}
// SetLanguageTag sets the languageTag property value. The languageTag property
func (m *BookingBusiness) SetLanguageTag(value *string)() {
    m.languageTag = value
}
// SetPhone sets the phone property value. The telephone number for the business. The phone property, together with address and webSiteUrl, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetPhone(value *string)() {
    m.phone = value
}
// SetPublicUrl sets the publicUrl property value. The URL for the scheduling page, which is set after you publish or unpublish the page. Read-only.
func (m *BookingBusiness) SetPublicUrl(value *string)() {
    m.publicUrl = value
}
// SetSchedulingPolicy sets the schedulingPolicy property value. Specifies how bookings can be created for this business.
func (m *BookingBusiness) SetSchedulingPolicy(value BookingSchedulingPolicyable)() {
    m.schedulingPolicy = value
}
// SetServices sets the services property value. All the services offered by this business. Read-only. Nullable.
func (m *BookingBusiness) SetServices(value []BookingServiceable)() {
    m.services = value
}
// SetStaffMembers sets the staffMembers property value. All the staff members that provide services in this business. Read-only. Nullable.
func (m *BookingBusiness) SetStaffMembers(value []BookingStaffMemberBaseable)() {
    m.staffMembers = value
}
// SetWebSiteUrl sets the webSiteUrl property value. The URL of the business web site. The webSiteUrl property, together with address, phone, appear in the footer of a business scheduling page.
func (m *BookingBusiness) SetWebSiteUrl(value *string)() {
    m.webSiteUrl = value
}
