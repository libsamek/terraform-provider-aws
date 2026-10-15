// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodb_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfdynamodb "github.com/hashicorp/terraform-provider-aws/internal/service/dynamodb"
)

func TestARNForNewRegion(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		TestName      string
		ARN           string
		NewRegion     string
		ExpectedARN   string
		ErrorExpected bool
	}{
		{
			TestName:      acctest.CtBasic,
			ARN:           "arn:aws:dynamodb:us-west-2:786648903940:table/tf-acc-test-7864711876941043153", //lintignore:AWSAT003,AWSAT005
			NewRegion:     "us-east-2",                                                                     //lintignore:AWSAT003
			ExpectedARN:   "arn:aws:dynamodb:us-east-2:786648903940:table/tf-acc-test-7864711876941043153", //lintignore:AWSAT003,AWSAT005
			ErrorExpected: false,
		},
		{
			TestName:      "basic2",
			ARN:           "arn:aws:dynamodb:us-west-2:786648903940:table/tf-acc-test-7864711876941043153", //lintignore:AWSAT003,AWSAT005
			NewRegion:     "us-east-1",                                                                     //lintignore:AWSAT003
			ExpectedARN:   "arn:aws:dynamodb:us-east-1:786648903940:table/tf-acc-test-7864711876941043153", //lintignore:AWSAT003,AWSAT005
			ErrorExpected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			t.Parallel()

			got, err := tfdynamodb.ARNForNewRegion(testCase.ARN, testCase.NewRegion)

			if err != nil && !testCase.ErrorExpected {
				t.Errorf("did not expect an error but got one: %s", err)
			}

			if err == nil && testCase.ErrorExpected {
				t.Error("expected an error but got none")
			}

			if got != testCase.ExpectedARN {
				t.Errorf("got %s, expected %s", got, testCase.ExpectedARN)
			}
		})
	}
}

func TestIsRecoveryPointARN(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ARN      string
		Expected bool
	}{
		"recovery point": {
			ARN:      "arn:aws:backup:us-west-2:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45", //lintignore:AWSAT003,AWSAT005
			Expected: true,
		},
		"recovery point other partition": {
			ARN:      "arn:aws-us-gov:backup:us-gov-west-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45", //lintignore:AWSAT003,AWSAT005
			Expected: true,
		},
		"DynamoDB backup": {
			ARN:      "arn:aws:dynamodb:us-west-2:123456789012:table/test/backup/01234567890123-abcdefgh", //lintignore:AWSAT003,AWSAT005
			Expected: false,
		},
		"backup vault": {
			ARN:      "arn:aws:backup:us-west-2:123456789012:backup-vault:test", //lintignore:AWSAT003,AWSAT005
			Expected: false,
		},
		"not an ARN": {
			ARN:      "recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45",
			Expected: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got, want := tfdynamodb.IsRecoveryPointARN(testCase.ARN), testCase.Expected; got != want {
				t.Errorf("IsRecoveryPointARN(%q) = %t, want %t", testCase.ARN, got, want)
			}
		})
	}
}
