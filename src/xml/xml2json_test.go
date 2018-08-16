package xml

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

type Envelope struct {
	xmlName xml.Name `xml:http://schemas.xmlsoap.org/soap/envelope/ Envelop`
	Body    Body
}
type Body struct {
	xmlName                       xml.Name `xml:http://schemas.xmlsoap.org/soap/envelope/ Body`
	GetCountriesAvailableResponse GetCountriesAvailableResponse
}
type GetCountriesAvailableResponse struct {
	xmlName                     xml.Name `xml:http://www.holidaywebservice.com/HolidayService_v2/ GetCountriesAvailableResponse`
	GetCountriesAvailableResult GetCountriesAvailableResult
}
type GetCountriesAvailableResult struct {
	xmlName     xml.Name `xml:http://www.holidaywebservice.com/HolidayService_v2/ GetCountriesAvailableResult`
	CountryCode []CountryCode
}
type CountryCode struct {
	Code        string `xml:"Code"`
	Description string `xml:"Description"`
}

func Test_ConvertXML_Input_XML_Should_Be_Struct(t *testing.T) {
	var actual Envelope
	xmlFile, _ := ioutil.ReadFile("./reponse.xml")
	expected := Envelope{
		Body: Body{
			GetCountriesAvailableResponse: GetCountriesAvailableResponse{
				GetCountriesAvailableResult: GetCountriesAvailableResult{
					CountryCode: []CountryCode{
						CountryCode{"Canada", "Canada"},
						CountryCode{"GreatBritain", "Great Britain and Wales"},
						CountryCode{"IrelandNorthern", "Northern Ireland"},
						CountryCode{"IrelandRepublicOf", "Republic of Ireland"},
						CountryCode{"Scotland", "Scotland"},
						CountryCode{"UnitedStates", "United States"},
					},
				},
			},
		},
	}

	xml.Unmarshal(xmlFile, &actual)

	for index, actualCountryCode := range actual.Body.GetCountriesAvailableResponse.GetCountriesAvailableResult.CountryCode {
		expectedCountryCode := expected.Body.GetCountriesAvailableResponse.GetCountriesAvailableResult.CountryCode[index]
		if actualCountryCode != expectedCountryCode {
			t.Errorf("Expect at index: %d %s but it got %s", index, expectedCountryCode, actualCountryCode)
		}
	}

}
