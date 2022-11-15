// Deprecated: We’re retiring the Azure Video Analyzer preview service; you're advised to transition your applications off of Video Analyzer by 01 December 2022. This SDK is no longer maintained and won’t work after the service is retired. To learn how to transition off, please refer to: https://aka.ms/azsdk/videoanalyzer/transitionoffguidance.
package videoanalyzer

import "github.com/Azure/azure-sdk-for-go/version"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
	return "Azure-SDK-For-Go/" + Version() + " videoanalyzer/2021-05-01-preview"
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return version.Number
}