package xml

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"testing"
)

type CountriesResponse struct {
	Counties []Country `json:"counties"`
}

type Country struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type GetCountriesAvailable struct {
	CountryCode []CountryCode `xml:"Body>GetCountriesAvailableResponse>GetCountriesAvailableResult>CountryCode"`
}

type CountryCode struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

func (g GetCountriesAvailable) ToJSON() CountriesResponse {
	counties := make([]Country, len(g.CountryCode))
	for index := range g.CountryCode {
		counties[index] = Country{
			Code:        g.CountryCode[index].Code,
			Description: g.CountryCode[index].Description,
		}
	}
	return CountriesResponse{
		Counties: counties,
	}
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

func Test_ToJson_Should_Be_Json(t *testing.T) {
	expected := `{"counties":[{"code":"Canada","description":"Canada"},{"code":"GreatBritain","description":"Great Britain and Wales"},{"code":"IrelandNorthern","description":"Northern Ireland"},{"code":"IrelandRepublicOf","description":"Republic of Ireland"},{"code":"Scotland","description":"Scotland"},{"code":"UnitedStates","description":"United States"}]}`
	getCountriesAvailable := GetCountriesAvailable{
		CountryCode: []CountryCode{
			CountryCode{"Canada", "Canada"},
			CountryCode{"GreatBritain", "Great Britain and Wales"},
			CountryCode{"IrelandNorthern", "Northern Ireland"},
			CountryCode{"IrelandRepublicOf", "Republic of Ireland"},
			CountryCode{"Scotland", "Scotland"},
			CountryCode{"UnitedStates", "United States"},
		},
	}
	countiesResults := getCountriesAvailable.ToJSON()
	actual, _ := json.Marshal(countiesResults)
	if expected != string(actual) {
		t.Errorf("Expect %s but it got %s", expected, actual)
	}
}
//ทำเอาเสื้อเฉยๆ
