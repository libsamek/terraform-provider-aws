// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodb

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

func arnForNewRegion(rn string, newRegion string) (string, error) {
	parsedARN, err := arn.Parse(rn)
	if err != nil {
		return "", err
	}

	parsedARN.Region = newRegion

	return parsedARN.String(), nil
}

func regionFromARN(rn string) (string, error) {
	parsedARN, err := arn.Parse(rn)
	if err != nil {
		return "", err
	}

	return parsedARN.Region, nil
}

func tableNameFromARN(rn string) (string, error) {
	parsedARN, err := arn.Parse(rn)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(parsedARN.Resource, "table/"), nil
}

// isRecoveryPointARN reports whether rn is an AWS Backup recovery point ARN.
// Tables backed up by AWS Backup with advanced DynamoDB backup enabled get these
// instead of DynamoDB backup ARNs, and only AWS Backup can restore them.
func isRecoveryPointARN(rn string) bool {
	parsedARN, err := arn.Parse(rn)
	if err != nil {
		return false
	}

	return parsedARN.Service == "backup" && strings.HasPrefix(parsedARN.Resource, "recovery-point:")
}
