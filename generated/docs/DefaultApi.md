# \DefaultApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTest**](DefaultApi.md#GetTest) | **Get** /test | Your GET endpoint
[**GetTest2**](DefaultApi.md#GetTest2) | **Get** /test2 | Your GET endpoint
[**PostHoge**](DefaultApi.md#PostHoge) | **Post** /hoge | hogetest



## GetTest

> map[string]interface{} GetTest(ctx, )

Your GET endpoint

### Required Parameters

This endpoint does not need any parameter.

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTest2

> GetTest2(ctx, )

Your GET endpoint

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostHoge

> PostHoge(ctx, optional)

hogetest

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostHogeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostHogeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testmodel** | [**optional.Interface of Testmodel**](Testmodel.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

