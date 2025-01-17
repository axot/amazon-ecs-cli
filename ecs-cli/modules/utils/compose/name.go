// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package utils

import (
	"fmt"
	"strings"
)

const (
	// the ~/.cache/ecscompose directory in which to store cached task definitions
	// changing this will cause user's caches to break.
	ProjectTDCache = "tdcache"

	// prefix for grouping tasks
	TaskGroupPrefix = "task"
)

// GetServiceName returns a service name
func GetServiceName(prefix, projectName string) string {
	return fmt.Sprintf("%s%s", prefix, projectName)
}

// GetTaskGroup returns <task-group-prefix><task-definition-family> used for grouping tasks
func GetTaskGroup(prefix, projectName string) string {
	return fmt.Sprintf("%s:%s%s", TaskGroupPrefix, prefix, projectName)
}

// GetFormattedContainerName returns the container name with task id prepended to it
func GetFormattedContainerName(ecsTaskId, ecsContainerName string) string {
	return fmt.Sprintf("%s/%s", ecsTaskId, ecsContainerName)
}

// GetIdFromArn parses an ARN and returns only the family name and revision
func GetIdFromArn(arn string) string {
	parts := strings.SplitN(arn, "/", 2)
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

// GetIdFromArn parses an ARN and returns only the uuid part
func GetTaskIdFromArn(arn string) string {
	parts := strings.SplitN(arn, "/", 3)
	if len(parts) != 3 {
		return ""
	}
	return parts[2]
}

// GetAwsAccountIdFromArn parses an ARN and returns only the accountId portion
func GetAwsAccountIdFromArn(arn string) string {
	parts := strings.SplitN(arn, ":", 7)
	if len(parts) != 7 {
		return ""
	}
	return parts[4]
}
