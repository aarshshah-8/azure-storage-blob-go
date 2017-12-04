package azblob

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
        "net/url"
    ""
        "net/url"
    "net/http"
        "net/url"
    "context"
        "net/url"
    "fmt"
        "net/url"
    "encoding/xml"
        "net/url"
    "io/ioutil"
        "net/url"
    "bytes"
)

// ServiceClient is the client for the Service methods of the Azblob service.
type ServiceClient struct {
    ManagementClient
}
// NewServiceClient creates an instance of the ServiceClient client.
func NewServiceClient(url url.URL, p pipeline.Pipeline) ServiceClient {
    return ServiceClient{NewManagementClient(url, p)}
}

// GetProperties gets the properties of a storage account's Blob service, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client ServiceClient) GetProperties(ctx context.Context, timeout *int32, requestID *string) (*StorageServiceProperties, error) {
    if err := validate([]validation{
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.getPropertiesPreparer(timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getPropertiesResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*StorageServiceProperties), err
}

// getPropertiesPreparer prepares the GetProperties request.
func (client ServiceClient) getPropertiesPreparer(timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    if timeout != nil {
        params.Set("timeout", fmt.Sprintf("%v", *timeout))
    }
        params.Set("restype", "service")
    params.Set("comp", "properties")
    req.URL.RawQuery = params.Encode()
    req.Header.Set("x-ms-version", ServiceVersion)
    if requestID != nil {
        req.Header.Set("x-ms-client-request-id", *requestID)
    }
	return req, nil
}

// getPropertiesResponder handles the response to the GetProperties request.
func (client ServiceClient) getPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
    result:= &StorageServiceProperties{rawResponse: resp.Response()}
    if err != nil {
        return result, err
    }
    defer resp.Response().Body.Close()
    b, err:= ioutil.ReadAll(resp.Response().Body)
    if err != nil {
        return result, NewResponseError(err, resp.Response(), "failed to read response body")
    }
    if len(b) > 0 {
        err = xml.Unmarshal(b, result)
        if err != nil {
            return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
        }
    }
    return result, nil
}

// GetStats retrieves statistics related to replication for the Blob service. It is only available on the secondary
// location endpoint when read-access geo-redundant replication is enabled for the storage account.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client ServiceClient) GetStats(ctx context.Context, timeout *int32, requestID *string) (*StorageServiceStats, error) {
    if err := validate([]validation{
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.getStatsPreparer(timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getStatsResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*StorageServiceStats), err
}

// getStatsPreparer prepares the GetStats request.
func (client ServiceClient) getStatsPreparer(timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    if timeout != nil {
        params.Set("timeout", fmt.Sprintf("%v", *timeout))
    }
        params.Set("restype", "service")
    params.Set("comp", "stats")
    req.URL.RawQuery = params.Encode()
    req.Header.Set("x-ms-version", ServiceVersion)
    if requestID != nil {
        req.Header.Set("x-ms-client-request-id", *requestID)
    }
	return req, nil
}

// getStatsResponder handles the response to the GetStats request.
func (client ServiceClient) getStatsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
    result:= &StorageServiceStats{rawResponse: resp.Response()}
    if err != nil {
        return result, err
    }
    defer resp.Response().Body.Close()
    b, err:= ioutil.ReadAll(resp.Response().Body)
    if err != nil {
        return result, NewResponseError(err, resp.Response(), "failed to read response body")
    }
    if len(b) > 0 {
        err = xml.Unmarshal(b, result)
        if err != nil {
            return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
        }
    }
    return result, nil
}

// ListContainers the List Containers operation returns a list of the containers under the specified account
//
// prefix is filters the results to return only containers whose name begins with the specified prefix. marker is a
// string value that identifies the portion of the list of containers to be returned with the next listing operation.
// The operation returns the NextMarker value within the response body if the listing operation did not return all
// containers remaining to be listed with the current page. The NextMarker value can be used as the value for the
// marker parameter in a subsequent call to request the next page of list items. The marker value is opaque to the
// client. maxresults is specifies the maximum number of containers to return. If the request does not specify
// maxresults, or specifies a value greater than 5000, the server will return up to 5000 items. Note that if the
// listing operation crosses a partition boundary, then the service will return a continuation token for retrieving the
// remainder of the results. For this reason, it is possible that the service will return fewer results than specified
// by maxresults, or than the default of 5000. include is include this parameter to specify that the container's
// metadata be returned as part of the response body. timeout is the timeout parameter is expressed in seconds. For
// more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client ServiceClient) ListContainers(ctx context.Context, prefix *string, marker *string, maxresults *int32, include ListContainersIncludeType, timeout *int32, requestID *string) (*ListContainersResponse, error) {
    if err := validate([]validation{
    { targetValue: maxresults,
     constraints: []constraint{	{target: "maxresults", name: null, rule: false ,
    chain: []constraint{	{target: "maxresults", name: inclusiveMinimum, rule: 1, chain: nil },
    }}}},
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.listContainersPreparer(prefix, marker, maxresults, include, timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listContainersResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*ListContainersResponse), err
}

// listContainersPreparer prepares the ListContainers request.
func (client ServiceClient) listContainersPreparer(prefix *string, marker *string, maxresults *int32, include ListContainersIncludeType, timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    if prefix != nil {
        params.Set("prefix", *prefix)
    }
    if marker != nil {
        params.Set("marker", *marker)
    }
    if maxresults != nil {
        params.Set("maxresults", fmt.Sprintf("%v", *maxresults))
    }
    if include != ListContainersIncludeNone {
        params.Set("include", fmt.Sprintf("%v", include))
    }
    if timeout != nil {
        params.Set("timeout", fmt.Sprintf("%v", *timeout))
    }
        params.Set("comp", "list")
    req.URL.RawQuery = params.Encode()
    req.Header.Set("x-ms-version", ServiceVersion)
    if requestID != nil {
        req.Header.Set("x-ms-client-request-id", *requestID)
    }
	return req, nil
}

// listContainersResponder handles the response to the ListContainers request.
func (client ServiceClient) listContainersResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
    result:= &ListContainersResponse{rawResponse: resp.Response()}
    if err != nil {
        return result, err
    }
    defer resp.Response().Body.Close()
    b, err:= ioutil.ReadAll(resp.Response().Body)
    if err != nil {
        return result, NewResponseError(err, resp.Response(), "failed to read response body")
    }
    if len(b) > 0 {
        err = xml.Unmarshal(b, result)
        if err != nil {
            return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
        }
    }
    return result, nil
}

// SetProperties sets properties for a storage account's Blob service endpoint, including properties for Storage
// Analytics and CORS (Cross-Origin Resource Sharing) rules
//
// storageServiceProperties is the StorageService properties. timeout is the timeout parameter is expressed in seconds.
// For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> requestID is provides a client-generated, opaque value with a 1 KB
// character limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client ServiceClient) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, timeout *int32, requestID *string) (*ServiceSetPropertiesResponse, error) {
    if err := validate([]validation{
    { targetValue: storageServiceProperties,
     constraints: []constraint{	{target: "storageServiceProperties.Logging", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.Logging.Version", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.Logging.Delete", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.Logging.Read", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.Logging.Write", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.Logging.RetentionPolicy", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.Logging.RetentionPolicy.Enabled", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.Logging.RetentionPolicy.Days", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.Logging.RetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil },
    }},
    }},
    }},
    	{target: "storageServiceProperties.HourMetrics", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.HourMetrics.Version", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.HourMetrics.Enabled", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.HourMetrics.IncludeAPIs", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.HourMetrics.RetentionPolicy", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Enabled", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Days", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.HourMetrics.RetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil },
    }},
    }},
    }},
    	{target: "storageServiceProperties.MinuteMetrics", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.MinuteMetrics.Version", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.MinuteMetrics.Enabled", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.MinuteMetrics.IncludeAPIs", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Enabled", name: null, rule: true, chain: nil },
    	{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Days", name: null, rule: true ,
    chain: []constraint{	{target: "storageServiceProperties.MinuteMetrics.RetentionPolicy.Days", name: inclusiveMinimum, rule: 1, chain: nil },
    }},
    }},
    }},
    	{target: "storageServiceProperties.DefaultServiceVersion", name: null, rule: true, chain: nil }}},
    { targetValue: timeout,
     constraints: []constraint{	{target: "timeout", name: null, rule: false ,
    chain: []constraint{	{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil },
    }}}}}); err != nil {
        return nil, err
    }
	req, err := client.setPropertiesPreparer(storageServiceProperties, timeout, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.setPropertiesResponder}, req)
    if err != nil {
        return nil, err
    }
	return resp.(*ServiceSetPropertiesResponse), err
}

// setPropertiesPreparer prepares the SetProperties request.
func (client ServiceClient) setPropertiesPreparer(storageServiceProperties StorageServiceProperties, timeout *int32, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
    params := req.URL.Query()
    if timeout != nil {
        params.Set("timeout", fmt.Sprintf("%v", *timeout))
    }
        params.Set("restype", "service")
    params.Set("comp", "properties")
    req.URL.RawQuery = params.Encode()
    req.Header.Set("x-ms-version", ServiceVersion)
    if requestID != nil {
        req.Header.Set("x-ms-client-request-id", *requestID)
    }
    b, err := xml.Marshal(storageServiceProperties)
    if err != nil {
        return req, pipeline.NewError(err, "failed to marshal request body")
    }
    req.Header.Set("Content-Type", "application/xml")
    err = req.SetBody(bytes.NewReader(b))
    if err != nil {
        return req, pipeline.NewError(err, "failed to set request body")
    }
	return req, nil
}

// setPropertiesResponder handles the response to the SetProperties request.
func (client ServiceClient) setPropertiesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK,http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
    return &ServiceSetPropertiesResponse{rawResponse: resp.Response()}, err
}

