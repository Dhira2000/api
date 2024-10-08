package internal

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	operatorsv1 "github.com/Dhira2000/api/pkg/operators/v1"
	"github.com/Dhira2000/api/pkg/validation/errors"
	"sigs.k8s.io/yaml"
)

func TestValidateOperatorGroup(t *testing.T) {
	cases := []struct {
		validatorFuncTest
		operatorGroupPath string
	}{
		{
			validatorFuncTest{
				description: "successfully validated",
			},
			filepath.Join("testdata", "correct.og.yaml"),
		},
		{
			validatorFuncTest{
				description: "invalid annotation name for operator group",
				wantErr:     true,
				errors: []errors.Error{
					errors.ErrFailedValidation("provided annotation olm.providedapis uses wrong case and should be olm.providedAPIs instead", "nginx-hbvsw"),
				},
			},
			filepath.Join("testdata", "badAnnotationNames.og.yaml"),
		},
	}
	for _, c := range cases {
		b, err := ioutil.ReadFile(c.operatorGroupPath)
		if err != nil {
			t.Fatalf("Error reading OperatorGroup path %s: %v", c.operatorGroupPath, err)
		}
		og := operatorsv1.OperatorGroup{}
		if err = yaml.Unmarshal(b, &og); err != nil {
			t.Fatalf("Error unmarshalling OperatorGroup at path %s: %v", c.operatorGroupPath, err)
		}
		result := validateOperatorGroupV1(&og)
		c.check(t, result)
	}
}
