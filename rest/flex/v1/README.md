# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.16.1
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://flex-api.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateChannel**](docs/DefaultApi.md#createchannel) | **Post** /v1/Channels | 
*DefaultApi* | [**CreateFlexFlow**](docs/DefaultApi.md#createflexflow) | **Post** /v1/FlexFlows | 
*DefaultApi* | [**CreateWebChannel**](docs/DefaultApi.md#createwebchannel) | **Post** /v1/WebChannels | 
*DefaultApi* | [**DeleteChannel**](docs/DefaultApi.md#deletechannel) | **Delete** /v1/Channels/{Sid} | 
*DefaultApi* | [**DeleteFlexFlow**](docs/DefaultApi.md#deleteflexflow) | **Delete** /v1/FlexFlows/{Sid} | 
*DefaultApi* | [**DeleteWebChannel**](docs/DefaultApi.md#deletewebchannel) | **Delete** /v1/WebChannels/{Sid} | 
*DefaultApi* | [**FetchChannel**](docs/DefaultApi.md#fetchchannel) | **Get** /v1/Channels/{Sid} | 
*DefaultApi* | [**FetchConfiguration**](docs/DefaultApi.md#fetchconfiguration) | **Get** /v1/Configuration | 
*DefaultApi* | [**FetchFlexFlow**](docs/DefaultApi.md#fetchflexflow) | **Get** /v1/FlexFlows/{Sid} | 
*DefaultApi* | [**FetchWebChannel**](docs/DefaultApi.md#fetchwebchannel) | **Get** /v1/WebChannels/{Sid} | 
*DefaultApi* | [**ListChannel**](docs/DefaultApi.md#listchannel) | **Get** /v1/Channels | 
*DefaultApi* | [**ListFlexFlow**](docs/DefaultApi.md#listflexflow) | **Get** /v1/FlexFlows | 
*DefaultApi* | [**ListWebChannel**](docs/DefaultApi.md#listwebchannel) | **Get** /v1/WebChannels | 
*DefaultApi* | [**UpdateConfiguration**](docs/DefaultApi.md#updateconfiguration) | **Post** /v1/Configuration | 
*DefaultApi* | [**UpdateFlexFlow**](docs/DefaultApi.md#updateflexflow) | **Post** /v1/FlexFlows/{Sid} | 
*DefaultApi* | [**UpdateWebChannel**](docs/DefaultApi.md#updatewebchannel) | **Post** /v1/WebChannels/{Sid} | 


## Documentation For Models

 - [FlexV1Channel](docs/FlexV1Channel.md)
 - [FlexV1Configuration](docs/FlexV1Configuration.md)
 - [FlexV1FlexFlow](docs/FlexV1FlexFlow.md)
 - [FlexV1WebChannel](docs/FlexV1WebChannel.md)
 - [ListChannelResponse](docs/ListChannelResponse.md)
 - [ListChannelResponseMeta](docs/ListChannelResponseMeta.md)
 - [ListFlexFlowResponse](docs/ListFlexFlowResponse.md)
 - [ListWebChannelResponse](docs/ListWebChannelResponse.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Author

support@twilio.com

