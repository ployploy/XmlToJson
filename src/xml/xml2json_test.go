package xml

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

type GetCountriesAvailable struct {
	CountryCode []CountryCode `xml:"Body>GetCountriesAvailableResponse>GetCountriesAvailableResult>CountryCode"`
}

type CountryCode struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

func Test_ConvertXML_Input_XML_Should_Be_Struct(t *testing.T) {
	var actual GetCountriesAvailable
	xmlFile, _ := ioutil.ReadFile("./reponse.xml")
	expected := GetCountriesAvailable{
		CountryCode: []CountryCode{
			CountryCode{"Canada", "Canada"},
			CountryCode{"GreatBritain", "Great Britain and Wales"},
			CountryCode{"IrelandNorthern", "Northern Ireland"},
			CountryCode{"IrelandRepublicOf", "Republic of Ireland"},
			CountryCode{"Scotland", "Scotland"},
			CountryCode{"UnitedStates", "United States"},
		},
	}

	xml.Unmarshal(xmlFile, &actual)

	for index, actualCountryCode := range actual.CountryCode {
		expectedCountryCode := expected.CountryCode[index]
		if actualCountryCode != expectedCountryCode {
			t.Errorf("Expect at index: %d %s but it got %s", index, expectedCountryCode, actualCountryCode)
		}
	}

}
