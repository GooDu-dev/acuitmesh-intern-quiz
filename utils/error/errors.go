package error

import "errors"

type CustomError interface {
	Code() string
	Error() string
	HttpStatusCode() int
	ThMessage() string
	EnMessage() string
	ErrorResponse() (int, HttpErrorResponse)
}

type HttpError struct {
	error string
}
type HttpErrorResponse struct {
	Code      string `json:"code"`
	Error     string `json:"error"`
	ThMessage string `json:"thMessage"`
	EnMessage string `json:"enMessage"`
}

func New(error string) *HttpError {
	return &HttpError{
		error: error,
	}
}
func InitError(e error) error {
	if CustomError, ok := e.(CustomError); ok {
		return CustomError
	} else {
		return errors.New(e.Error())
	}
}
func (e *HttpError) Code() string {
	if code, ok := ErrorCode[e]; ok {
		return code
	} else {
		return InternalServerError.Error()
	}
}
func (e *HttpError) Error() string {
	return e.error
}
func (e *HttpError) HttpStatusCode() int {
	if code, ok := HttpStatusCodes[e]; ok {
		return code
	} else {
		return InternalServerError.HttpStatusCode()
	}
}
func (e *HttpError) ThMessage() string {
	if Th, ok := ErrorThMessage[e]; ok {
		return Th
	} else {
		return InternalServerError.ThMessage()
	}
}
func (e *HttpError) EnMessage() string {
	if En, ok := ErrorEnMessage[e]; ok {
		return En
	} else {
		return InternalServerError.EnMessage()
	}

}
func (e *HttpError) ErrorResponse() (int, HttpErrorResponse) {
	if httpStatusCode, ok := HttpStatusCodes[e]; ok {
		return httpStatusCode, HttpErrorResponse{
			Code:      e.Code(),
			Error:     e.Error(),
			ThMessage: e.ThMessage(),
			EnMessage: e.EnMessage(),
		}
	} else {
		return InternalServerError.HttpStatusCode(), HttpErrorResponse{
			Code:      InternalServerError.Code(),
			Error:     InternalServerError.Error(),
			ThMessage: InternalServerError.ThMessage(),
			EnMessage: InternalServerError.EnMessage(),
		}
	}
}

func GetErrorResponse(e error) (int, HttpErrorResponse) {
	if http_status_codes, err := HttpStatusCodes[e]; err {
		if code, err := ErrorCode[e]; err {
			if th_message, err := ErrorThMessage[e]; err {
				if en_message, err := ErrorEnMessage[e]; err {
					return http_status_codes, HttpErrorResponse{
						Code:      code,
						Error:     e.Error(),
						ThMessage: th_message,
						EnMessage: en_message,
					}
				}
			}
		}
	}
	return InternalServerError.ErrorResponse()
}

// CustomError contain error name for error
var (
	// Normal
	InternalServerError                 CustomError = New("internal_server_error")
	DatabaseConnectionError             CustomError = New("database_connection_error")
	PageNotFoundError                   CustomError = New("page_not_found_error")
	BadRequestError                     CustomError = New("bad_request_error")
	InvalidRequestError                 CustomError = New("invalid_request_error")
	MissingRequestError                 CustomError = New("missing_request_error")
	InvalidContentTypeError             CustomError = New("invalid_content_type_error")
	MappingRequestBodyError             CustomError = New("mapping_request_body_error")
	DatabaseNameNotMatchedError         CustomError = New("database_name_not_matched_error")
	UnableToReadConfigError             CustomError = New("unable_to_read_config_error")
	MissingPathVariablesError           CustomError = New("missing_path_variables_error")
	DataTypeIsNotStructError            CustomError = New("data_type_is_not_struct_error")
	FieldContainsNilOrDefaultValueError CustomError = New("field_contains_nil_or_default_value_error")
	InvalidHeaderNotAcceptableError     CustomError = New("invalid_header_request_error")
	ContentNotFoundError                CustomError = New("content_in_database_not_found_error")
)
var ErrorCode = map[error]string{
	// Normal - 1
	InternalServerError:                 "10000",
	DatabaseConnectionError:             "10001",
	PageNotFoundError:                   "10002",
	BadRequestError:                     "10003",
	InvalidRequestError:                 "10004",
	MissingRequestError:                 "10005",
	InvalidContentTypeError:             "10006",
	MappingRequestBodyError:             "10007",
	DatabaseNameNotMatchedError:         "10008",
	UnableToReadConfigError:             "10009",
	MissingPathVariablesError:           "10010",
	DataTypeIsNotStructError:            "10011",
	FieldContainsNilOrDefaultValueError: "10012",
	InvalidHeaderNotAcceptableError:     "10013",
	ContentNotFoundError:                "10014",
}
var ErrorThMessage = map[error]string{
	// Normarl - 1
	InternalServerError:                 "เกิดข้อผิดพลาดที่เซิฟเวอร์, โปรดลองใหม่อีกครั้ง",
	DatabaseConnectionError:             "เชื่อมต่อฐานข้อมูลไม่สำเร็จ, โปรดลองใหม่อีกครั้ง",
	PageNotFoundError:                   "ไม่พบเจอหน้า, โปรดลองใหม่อีกครั้ง",
	BadRequestError:                     "คำขอไม่ถูกต้อง, โปรดลองใหม่อีกครั้ง",
	InvalidRequestError:                 "คำขอไม่ถูกต้อง, โปรดลองใหม่อีกครั้ง",
	MissingRequestError:                 "ไม่พบคำขอ, โปรดลองใหม่อีกครั้ง",
	InvalidContentTypeError:             "ประเภทของการ encode ข้อมูลผิดพลาด, โปรดลองใหม่อีกครั้ง",
	MappingRequestBodyError:             "รูปแบบคำขอไม่ถูกต้อง, โปรดลองใหม่อีกครั้ง",
	DatabaseNameNotMatchedError:         "ไม่พบชื่อฐานข้อมูล, โปรดลองใหม่อีกครั้ง",
	UnableToReadConfigError:             "ไม่สามารถอ่าน config ได้, โปรดลองใหม่อีกครั้ง",
	MissingPathVariablesError:           "ไม่พบ path variables, โปรดลองใหม่อีกครั้ง",
	DataTypeIsNotStructError:            "ประเภทของข้อมูลไม่ใช่ struct, โปรดลองใหม่อีกครั้ง",
	FieldContainsNilOrDefaultValueError: "ข้อมูลในฟิลล์ไม่มีค่าหรือเป็นค่าเริ่มต้น, โปรดลองใหม่อีกครั้ง",
	InvalidHeaderNotAcceptableError:     "ข้อมูลใน header ไม่ตรงตามเงื่อนไขของเว็บไซต์, โปรดลองใหม่อีกครั้ง",
	ContentNotFoundError:                "ไม่พบข้อมูลในฐานข้อมูล, โปรดลองใหม่อีกครั้ง",
}
var ErrorEnMessage = map[error]string{
	// Normal - 1
	InternalServerError:                 "There was an error on the server, please try again.",
	DatabaseConnectionError:             "Database connection unavailable, please try again.",
	PageNotFoundError:                   "Unable to find page, please try again.",
	BadRequestError:                     "Invalid request, please try again.",
	InvalidRequestError:                 "Invalid request, please try again.",
	MissingRequestError:                 "Missing request, please try again.",
	InvalidContentTypeError:             "Invalid content type, please try again.",
	MappingRequestBodyError:             "Mapping request body error, please try again.",
	DatabaseNameNotMatchedError:         "Database name not matched, please try again",
	UnableToReadConfigError:             "Unable to read config file, please try again",
	MissingPathVariablesError:           "Missing path variables, please try again.",
	DataTypeIsNotStructError:            "Data type is not a struct, please try again.",
	FieldContainsNilOrDefaultValueError: "Fields contain nil or default value, please try again.",
	InvalidHeaderNotAcceptableError:     "Invalud header request is not acceptable, please try again later",
	ContentNotFoundError:                "Missing content in database, please try again later",
}
var HttpStatusCodes = map[error]int{
	// Normal - 1
	InternalServerError:                 500,
	DatabaseConnectionError:             500,
	PageNotFoundError:                   404,
	BadRequestError:                     400,
	InvalidRequestError:                 400,
	MissingRequestError:                 400,
	InvalidContentTypeError:             415,
	MappingRequestBodyError:             400,
	DatabaseNameNotMatchedError:         500,
	UnableToReadConfigError:             500,
	MissingPathVariablesError:           400,
	DataTypeIsNotStructError:            400,
	FieldContainsNilOrDefaultValueError: 400,
	InvalidHeaderNotAcceptableError:     406,
	ContentNotFoundError:                404,
}
