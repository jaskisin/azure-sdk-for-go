package computervisionapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.1/computervision"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gofrs/uuid"
	"io"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	AnalyzeImage(ctx context.Context, imageURL computervision.ImageURL, visualFeatures []computervision.VisualFeatureTypes, details []computervision.Details, language string, descriptionExclude []computervision.DescriptionExclude) (result computervision.ImageAnalysis, err error)
	AnalyzeImageByDomain(ctx context.Context, model string, imageURL computervision.ImageURL, language string) (result computervision.DomainModelResults, err error)
	AnalyzeImageByDomainInStream(ctx context.Context, model string, imageParameter io.ReadCloser, language string) (result computervision.DomainModelResults, err error)
	AnalyzeImageInStream(ctx context.Context, imageParameter io.ReadCloser, visualFeatures []computervision.VisualFeatureTypes, details []computervision.Details, language string, descriptionExclude []computervision.DescriptionExclude) (result computervision.ImageAnalysis, err error)
	DescribeImage(ctx context.Context, imageURL computervision.ImageURL, maxCandidates *int32, language string, descriptionExclude []computervision.DescriptionExclude) (result computervision.ImageDescription, err error)
	DescribeImageInStream(ctx context.Context, imageParameter io.ReadCloser, maxCandidates *int32, language string, descriptionExclude []computervision.DescriptionExclude) (result computervision.ImageDescription, err error)
	DetectObjects(ctx context.Context, imageURL computervision.ImageURL) (result computervision.DetectResult, err error)
	DetectObjectsInStream(ctx context.Context, imageParameter io.ReadCloser) (result computervision.DetectResult, err error)
	GenerateThumbnail(ctx context.Context, width int32, height int32, imageURL computervision.ImageURL, smartCropping *bool) (result computervision.ReadCloser, err error)
	GenerateThumbnailInStream(ctx context.Context, width int32, height int32, imageParameter io.ReadCloser, smartCropping *bool) (result computervision.ReadCloser, err error)
	GetAreaOfInterest(ctx context.Context, imageURL computervision.ImageURL) (result computervision.AreaOfInterestResult, err error)
	GetAreaOfInterestInStream(ctx context.Context, imageParameter io.ReadCloser) (result computervision.AreaOfInterestResult, err error)
	GetReadResult(ctx context.Context, operationID uuid.UUID) (result computervision.ReadOperationResult, err error)
	ListModels(ctx context.Context) (result computervision.ListModelsResult, err error)
	Read(ctx context.Context, imageURL computervision.ImageURL, language computervision.OcrDetectionLanguage) (result autorest.Response, err error)
	ReadInStream(ctx context.Context, imageParameter io.ReadCloser, language computervision.OcrDetectionLanguage) (result autorest.Response, err error)
	RecognizePrintedText(ctx context.Context, detectOrientation bool, imageURL computervision.ImageURL, language computervision.OcrLanguages) (result computervision.OcrResult, err error)
	RecognizePrintedTextInStream(ctx context.Context, detectOrientation bool, imageParameter io.ReadCloser, language computervision.OcrLanguages) (result computervision.OcrResult, err error)
	TagImage(ctx context.Context, imageURL computervision.ImageURL, language string) (result computervision.TagResult, err error)
	TagImageInStream(ctx context.Context, imageParameter io.ReadCloser, language string) (result computervision.TagResult, err error)
}

var _ BaseClientAPI = (*computervision.BaseClient)(nil)