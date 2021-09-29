package addresses

// tests have the exception to have a different packge inside the same folder, but for now let's continue with the same "addresses" package
import "testing"

// UNIT TEST: x_test (require _test in the file name)
// requires "Test" + "Function Name" (not required to be the same function name, but it's a good practice)

// to run: go inside the folder and: go test
// PASS
// ok      introduction-tests/addresses    0.006s

type testScenario struct {
	insertedAddress string
	expectedReturn  string
}

func TestTypeOfAddress(t *testing.T) {
	scenariosOfTest := []testScenario{
		{"Street Catanduva", "Street"},
		{"Avenue Paulista", "Avenue"},
		{"Road XYZ", "Road"},
		{"HIGHWAY ABC", "Highway"},
		{"", "Invalid Type"},
		{"This is not valid", "Invalid Type"},
	}

	for _, scenario := range scenariosOfTest {
		actualReturn := TypeOfAddress(scenario.insertedAddress)
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
