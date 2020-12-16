package models
import (
	"encoding/json"
	"log"
	"net/http"
)

type emailParamsStruct struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type errorResponseStruct struct {
	Code    int
	Message string
}

type successResponseStruct struct {
	Code     int
	Message  string
	Response interface{}
}

// RenderHome Rendering the Home Page
func RenderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/home.html")
}

// SendEmailHandler Will be used to send emails
func SendEmailHandler(response http.ResponseWriter, request *http.Request) {
	var emailRequest emailParamsStruct
	var errorResponse = errorResponseStruct{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&emailRequest)
	defer request.Body.Close()

	if decoderErr != nil {
		returnErrorResponse(response, request, errorResponse)
	} else {
		errorResponse.Code = http.StatusBadRequest
		if emailRequest.Email == "" {
			errorResponse.Message = "Email can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if emailRequest.Name == "" {
			errorResponse.Message = "Password can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else if emailRequest.Message == "" {
			errorResponse.Message = "Message can't be empty"
			returnErrorResponse(response, request, errorResponse)
		} else {

			log.Println(emailRequest)

			var to = []string{
				emailRequest.Email,
			}

			sent, err := SendEmail(emailRequest.Message, to)

			log.Println(sent)

			log.Println(err)

			if err != nil && !sent {
				errorResponse.Message = "It's not you it's me."
				returnErrorResponse(response, request, errorResponse)
			} else {
				var successResponse = successResponseStruct{
					Code:    http.StatusOK,
					Message: "You email is sent, check your outbox.",
				}

				successJSONResponse, jsonError := json.Marshal(successResponse)

				if jsonError != nil {
					returnErrorResponse(response, request, errorResponse)
				}
				response.Header().Set("Content-Type", "application/json")
				response.Write(successJSONResponse)
			}
		}
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage errorResponseStruct) {
	httpResponse := &errorResponseStruct{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
