package instances

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ie372a3f627d97fce2cbc77ec0a65fdb0553f4c46e781d23fcb4f2d8f56abf98b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/delta"
    if0edcd02013795b841249de49316a2272c27b93c98355e1614935e50a696d53c "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendar/calendarview/item/instances/count"
)

// InstancesRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type InstancesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InstancesRequestBuilderGetQueryParameters get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns the occurrences and exceptions of the event in the specified time range.
type InstancesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// InstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InstancesRequestBuilderGetQueryParameters
}
// NewInstancesRequestBuilderInternal instantiates a new InstancesRequestBuilder and sets the default values.
func NewInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesRequestBuilder) {
    m := &InstancesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/instances{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInstancesRequestBuilder instantiates a new InstancesRequestBuilder and sets the default values.
func NewInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *InstancesRequestBuilder) Count()(*if0edcd02013795b841249de49316a2272c27b93c98355e1614935e50a696d53c.CountRequestBuilder) {
    return if0edcd02013795b841249de49316a2272c27b93c98355e1614935e50a696d53c.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns the occurrences and exceptions of the event in the specified time range.
func (m *InstancesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *InstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *InstancesRequestBuilder) Delta()(*ie372a3f627d97fce2cbc77ec0a65fdb0553f4c46e781d23fcb4f2d8f56abf98b.DeltaRequestBuilder) {
    return ie372a3f627d97fce2cbc77ec0a65fdb0553f4c46e781d23fcb4f2d8f56abf98b.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the instances (occurrences) of an event for a specified time range.  If the event is a `seriesMaster` type, this returns the occurrences and exceptions of the event in the specified time range.
func (m *InstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *InstancesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EventCollectionResponseable), nil
}
