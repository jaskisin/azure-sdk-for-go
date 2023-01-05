//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azquery_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/monitor/azquery"
	"github.com/stretchr/testify/require"
)

var query string = "let dt = datatable (DateTime: datetime, Bool:bool, Guid: guid, Int: int, Long:long, Double: double, String: string, Timespan: timespan, Decimal: decimal, Dynamic: dynamic)\n" + "[datetime(2015-12-31 23:59:59.9), false, guid(74be27de-1e4e-49d9-b579-fe0b331d3642), 12345, 1, 12345.6789, 'string value', 10s, decimal(0.10101), dynamic({\"a\":123, \"b\":\"hello\", \"c\":[1,2,3], \"d\":{}})];" + "range x from 1 to 100 step 1 | extend y=1 | join kind=fullouter dt on $left.y == $right.Long"

type queryTest struct {
	Bool   bool
	Long   int64
	String string
}

func TestLogsClient(t *testing.T) {
	client, err := azquery.NewLogsClient(credential, nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	c := cloud.Configuration{
		ActiveDirectoryAuthorityHost: "https://...",
		Services: map[cloud.ServiceName]cloud.ServiceConfiguration{
			cloud.ResourceManager: {
				Audience: "",
				Endpoint: "",
			},
		},
	}
	opts := azcore.ClientOptions{Cloud: c}
	cloudClient, err := azquery.NewLogsClient(credential, &azquery.LogsClientOptions{ClientOptions: opts})
	require.Error(t, err)
	require.Equal(t, err.Error(), "provided Cloud field is missing Azure Monitor Logs configuration")
	require.Nil(t, cloudClient)
}

func TestQueryWorkspace_BasicQuerySuccess(t *testing.T) {
	client := startLogsTest(t)
	body := azquery.Body{
		Query:    to.Ptr(query),
		Timespan: to.Ptr("2015-12-31/2016-01-01"),
	}
	testSerde(t, &body)

	res, err := client.QueryWorkspace(context.Background(), workspaceID, body, nil)
	require.NoError(t, err)
	require.Nil(t, res.Error)
	require.Nil(t, res.Render)
	require.Nil(t, res.Statistics)
	require.Len(t, res.Tables, 1)
	require.Len(t, res.Tables[0].Rows, 100)

	var queryResults []queryTest
	for _, table := range res.Tables {
		queryResults = make([]queryTest, len(table.Rows))
		indexLong := table.ColumnIndexLookup["Long"]
		indexString := table.ColumnIndexLookup["String"]
		indexBool := table.ColumnIndexLookup["Bool"]

		for index, row := range table.Rows {
			queryResults[index] = queryTest{
				Long:   int64(row[indexLong].(float64)),
				String: row[indexString].(string),
				Bool:   row[indexBool].(bool),
			}
		}
	}

	require.Len(t, queryResults, 100)
	require.False(t, queryResults[99].Bool)
	require.Equal(t, queryResults[99].String, "string value")
	require.Equal(t, queryResults[99].Long, int64(1))

	testSerde(t, &res)
}

func TestQueryWorkspace_BasicQueryFailure(t *testing.T) {
	client := startLogsTest(t)

	res, err := client.QueryWorkspace(context.Background(), workspaceID, azquery.Body{Query: to.Ptr("not a valid query")}, nil)
	require.Error(t, err)
	require.Nil(t, res.Error)
	require.Nil(t, res.Tables)

	var httpErr *azcore.ResponseError
	require.ErrorAs(t, err, &httpErr)
	require.Equal(t, httpErr.ErrorCode, "BadArgumentError")
	require.Equal(t, httpErr.StatusCode, 400)

	testSerde(t, &res)
}

func TestQueryWorkspace_PartialError(t *testing.T) {
	client := startLogsTest(t)
	query := "let Weight = 92233720368547758; range x from 1 to 3 step 1 | summarize percentilesw(x, Weight * 100, 50)"

	res, err := client.QueryWorkspace(context.Background(), workspaceID, azquery.Body{Query: &query}, nil)
	require.NoError(t, err)
	require.NotNil(t, res.Error)
	require.Equal(t, res.Error.Code, "PartialError")
	require.Contains(t, res.Error.Error(), "PartialError")

	testSerde(t, &res)
}

// tests for special options: timeout, statistics, visualization
func TestQueryWorkspace_AdvancedQuerySuccess(t *testing.T) {
	client := startLogsTest(t)
	prefer := "wait=180,include-statistics=true,include-render=true"

	res, err := client.QueryWorkspace(context.Background(), workspaceID, azquery.Body{Query: &query}, &azquery.LogsClientQueryWorkspaceOptions{Prefer: &prefer})
	require.NoError(t, err)
	require.Nil(t, res.Error)
	require.NotNil(t, res.Tables)
	require.NotNil(t, res.Render)
	require.NotNil(t, res.Statistics)
	testSerde(t, &res)
}

func TestQueryWorkspace_MultipleWorkspaces(t *testing.T) {
	client := startLogsTest(t)
	workspaces := []*string{&workspaceID2}
	body := azquery.Body{
		Query:      &query,
		Workspaces: workspaces,
	}
	testSerde(t, &body)

	res, err := client.QueryWorkspace(context.Background(), workspaceID, body, nil)
	require.NoError(t, err)
	require.Nil(t, res.Error)
	require.Len(t, res.Tables[0].Rows, 100)
}

func TestQueryBatch_QuerySuccess(t *testing.T) {
	client := startLogsTest(t)
	query1, query2 := query, query+" | take 2"

	batchRequest := azquery.BatchRequest{[]*azquery.BatchQueryRequest{
		{Body: &azquery.Body{Query: to.Ptr(query1)}, ID: to.Ptr("1"), Workspace: to.Ptr(workspaceID)},
		{Body: &azquery.Body{Query: to.Ptr(query2)}, ID: to.Ptr("2"), Workspace: to.Ptr(workspaceID)},
	}}
	testSerde(t, &batchRequest)

	res, err := client.QueryBatch(context.Background(), batchRequest, nil)
	require.NoError(t, err)
	require.Len(t, res.Responses, 2)
	for _, resp := range res.Responses {
		require.Nil(t, resp.Body.Error)
		require.NotNil(t, resp.Body.Tables)
		if *resp.ID == "1" && len(resp.Body.Tables[0].Rows) != 100 {
			t.Fatal("expected 100 rows from batch request 1")
		}
		if *resp.ID == "2" && len(resp.Body.Tables[0].Rows) != 2 {
			t.Fatal("expected 2 rows from batch request 2")
		}
	}
	testSerde(t, &res)
}

func TestQueryBatch_BasicQueryFailure(t *testing.T) {
	client := startLogsTest(t)
	query1, query2 := query, query+" | take 2"

	batchRequest := azquery.BatchRequest{[]*azquery.BatchQueryRequest{
		{Body: &azquery.Body{Query: to.Ptr(query1)}, ID: to.Ptr("1"), Workspace: to.Ptr(workspaceID)},
		{Body: &azquery.Body{Query: to.Ptr(query2)}, ID: to.Ptr("2"), Workspace: to.Ptr(workspaceID)},
	}}
	testSerde(t, &batchRequest)

	res, err := client.QueryBatch(context.Background(), azquery.BatchRequest{}, nil)
	require.Error(t, err)
	require.Nil(t, res.Responses)

	var httpErr *azcore.ResponseError
	require.ErrorAs(t, err, &httpErr)
	require.Equal(t, httpErr.ErrorCode, "BadArgumentError")
	require.Equal(t, httpErr.StatusCode, 400)
}

func TestQueryBatch_AdvancedQuerySuccess(t *testing.T) {
	client := startLogsTest(t)
	batchPrefer := "wait=180,include-statistics=true,include-render=true"
	headers := map[string]*string{"prefer": &batchPrefer}

	batchRequestAdvanced := azquery.BatchRequest{[]*azquery.BatchQueryRequest{
		{Body: &azquery.Body{Query: to.Ptr(query)}, ID: to.Ptr("1"), Workspace: to.Ptr(workspaceID2), Headers: headers},
		{Body: &azquery.Body{Query: to.Ptr(query)}, ID: to.Ptr("2"), Workspace: to.Ptr(workspaceID2), Headers: headers},
	}}
	testSerde(t, &batchRequestAdvanced)

	res, err := client.QueryBatch(context.Background(), batchRequestAdvanced, nil)
	require.NoError(t, err)
	require.Len(t, res.Responses, 2)
	for _, resp := range res.Responses {
		require.Nil(t, resp.Body.Error)
		require.NotNil(t, resp.Body.Tables)
		require.NotNil(t, resp.Body.Render)
		require.NotNil(t, resp.Body.Statistics)
		require.Len(t, resp.Body.Tables[0].Rows, 100)
	}
	testSerde(t, &res)
}

func TestQueryBatch_PartialError(t *testing.T) {
	client := startLogsTest(t)

	batchRequest := azquery.BatchRequest{[]*azquery.BatchQueryRequest{
		{Body: &azquery.Body{Query: to.Ptr("not a valid query")}, ID: to.Ptr("1"), Workspace: to.Ptr(workspaceID)},
		{Body: &azquery.Body{Query: to.Ptr(query)}, ID: to.Ptr("2"), Workspace: to.Ptr(workspaceID)},
	}}

	res, err := client.QueryBatch(context.Background(), batchRequest, nil)
	require.NoError(t, err)
	require.Len(t, res.Responses, 2)
	for _, resp := range res.Responses {
		if *resp.ID == "1" {
			require.NotNil(t, resp.Body.Error)
			require.Equal(t, resp.Body.Error.Code, "BadArgumentError")
			require.Contains(t, resp.Body.Error.Error(), "BadArgumentError")
		}
		if *resp.ID == "2" {
			require.Nil(t, resp.Body.Error)
			require.Len(t, resp.Body.Tables[0].Rows, 100)
		}
	}
}

func TestLogConstants(t *testing.T) {
	batchMethod := []azquery.BatchQueryRequestMethod{azquery.BatchQueryRequestMethodPOST}
	batchMethodRes := azquery.PossibleBatchQueryRequestMethodValues()
	require.Equal(t, batchMethod, batchMethodRes)

	batchPath := []azquery.BatchQueryRequestPath{azquery.BatchQueryRequestPathQuery}
	batchPathRes := azquery.PossibleBatchQueryRequestPathValues()
	require.Equal(t, batchPath, batchPathRes)

	logsColumnType := []azquery.LogsColumnType{azquery.LogsColumnTypeBool, azquery.LogsColumnTypeDatetime, azquery.LogsColumnTypeDecimal, azquery.LogsColumnTypeDynamic, azquery.LogsColumnTypeGUID, azquery.LogsColumnTypeInt, azquery.LogsColumnTypeLong, azquery.LogsColumnTypeReal, azquery.LogsColumnTypeString, azquery.LogsColumnTypeTimespan}
	logsColumnTypeRes := azquery.PossibleLogsColumnTypeValues()
	require.Equal(t, logsColumnType, logsColumnTypeRes)
}