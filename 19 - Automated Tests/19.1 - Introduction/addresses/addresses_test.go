package addresses_test

// #IMPORTANT
// tests have the exception to have a different packge inside the same folder, but for now let's continue with the same "addresses" package
import (
	"introduction-tests/addresses"
	"testing"
)

// UNIT TEST: x_test (require _test in the file name)
// requires "Test" + "Function Name" (not required to be the same function name, but it's a good practice)

// to run: go inside the folder and: go test
// PASS
// ok      introduction-tests/addresses    0.006s

type testScenario struct {
	insertedAddress string
	expectedReturn  string
}

func TestFake(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Error")
	}
}

func TestTypeOfAddress(t *testing.T) {
	// making tests run in parallel, you need to put the same into other test functions
	t.Parallel()

	scenariosOfTest := []testScenario{
		{"Street Catanduva", "Street"},
		{"Avenue Paulista", "Avenue"},
		{"Road XYZ", "Road"},
		{"HIGHWAY ABC", "Highway"},
		{"", "Invalid Type"},
		{"This is not valid", "Invalid Type"},
	}

	for _, scenario := range scenariosOfTest {
		// if we had this test file using the same package "address" then we would not need to import addresses or use it,
		// just call function actualReturn := TypeOfAddress(scenario.insertedAddress)

		// another option to not use addresses.Function is to import . "package/addresses"
		actualReturn := addresses.TypeOfAddress(scenario.insertedAddress)

		if actualReturn != scenario.expectedReturn {
			// t.Error("Actual Type is different from the expected one!")
			t.Errorf("Actual Type is different from the expected one! Expected %s but received %s",
				scenario.expectedReturn,
				actualReturn, // mandatory
			)
		}
	}
}

// REPLACED BY STRUCT ABOVE
// func TestTypeOfAddress(t *testing.T) {
// 	addressForTest := "Avenue Paulista"
// 	expectedType := "Avenue"
// 	actualType := TypeOfAddress(addressForTest)

// 	if expectedType != actualType {
// 		// t.Error("Actual Type is different from the expected one!")
// 		t.Errorf("Actual Type is different from the expected one! Expected %s but received %s",
// 			expectedType,
// 			actualType
// 		)
// 	}
// }
