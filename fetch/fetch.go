package fetch

import (
	myError "fetch-data/errors"

	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

/*
 *
 *
 */

/* FetchData ------------------------------------------------------------------------------------------------- */

type FetchError struct {
	Status     string
	StatusCode int
	Response   *http.Response
}

type FetchBody struct {
	body []byte
}

/* To convert the response body to String */
func (fb *FetchBody) String() string {
	return string(fb.body)
}

/* To convert the response body to Json. `dataHolder` should be a pointer. */
func (fb *FetchBody) Json(dataHolder any) error {
	err := json.Unmarshal(fb.body, dataHolder)
	myError.ErrorHandlerWithInfo(err, "FetchData.Json >> json.Unmarshal [json conversion]")

	return err
}

/* Custom Struct for the response */
type FetchData struct {
	*http.Response
	Body FetchBody
}

/* FetchData ------------------------------------------------------------------------------------------------- */

/*
 *
 *
 */

/* local functions ------------------------------------------------------------------------------------------- */

/* To convert the response body to Byte Array */
func convertToByte(body io.ReadCloser) []byte {
	responseData, err := ioutil.ReadAll(body)
	myError.ErrorHandlerWithInfo(err, "FetchData.Byte >> ioutil.ReadAll [to byte conversion ]")

	return responseData
}

/* Create an instance of an FetchData */
func createFetchData(res *http.Response) *FetchData {
	return &FetchData{Response: res, Body: FetchBody{convertToByte(res.Body)}}
}

/* check the status code for the error */
func checkStatus(response *FetchData) {
	if response.StatusCode >= 400 {
		panic(
			&FetchError{
				Status:     response.Status,
				StatusCode: response.StatusCode,
				Response:   response.Response,
			},
		)

	}
}

/* local functions ------------------------------------------------------------------------------------------- */

/*
 *
 *
 */

/* HTTP Methods ********************************************************************************************** */

/* HTTP GET Method */
func Get(url string) *FetchData {
	response, err := http.Get(url)
	myError.ErrorHandlerWithInfo(err, "fetch http.Get(url)")

	data := createFetchData(response)
	checkStatus(data)

	return data
}

/* HTTP Methods ********************************************************************************************** */
