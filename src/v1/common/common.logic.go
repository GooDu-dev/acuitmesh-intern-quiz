package common

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"runtime"
	"time"

	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/jackc/pgx/v5/pgconn"
)

func IsDefaultValueOrNil(data any) (output bool) {
	value := reflect.ValueOf(data)
	switch kind := reflect.TypeOf(data).Kind(); kind {
	case reflect.Array:
		// if array contains (nil || default_value) return true
		output = false
		for _, d := range []any{data} {
			value = reflect.ValueOf(d)
			output = value.IsNil() || value.IsZero()
			if output {
				return true
			}
			output = value.IsNil() || value.IsZero()
		}
		return output
	case reflect.Pointer:
		fmt.Println(value.IsNil(), value.IsZero())
		output = value.IsNil() || value.IsZero()
	case reflect.String:
		output = value.String() == ""
	case reflect.Int:
		output = value.IsZero()
	}
	return output
}

func DeepIsDefaultValueOrNil(data interface{}) (err error) {
	kind := reflect.TypeOf(data)
	if kind.Kind() != reflect.Struct {
		return customError.DataTypeIsNotStructError
	}

	value := reflect.ValueOf(data)
	for i := 0; i < value.NumField(); i++ {
		if canNull(kind.Field(i).Type) {
			continue
		}
		if reflect.DeepEqual(value.Field(i).Interface(), reflect.Zero(kind.Field(i).Type).Interface()) {
			return customError.FieldContainsNilOrDefaultValueError
		}
	}
	return nil
}

func canNull(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Chan:
		return true
	}
	return false
}

func GetFunctionWithPackageName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	packageName, functionName := splitFunctionName(funcName)
	return fmt.Sprintf("%s:%s", packageName, functionName)
}

func splitFunctionName(funcName string) (packageName string, functionName string) {
	for i := len(funcName) - 1; i >= 0; i-- {
		if funcName[i] == '.' {
			packageName = funcName[:i]
			functionName = funcName[i+1:]
			break
		}
	}
	return packageName, functionName
}

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func IsPostgresqlDataDup(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" // postgresql unique error for dupp data
	}
	return false
}

func GenerateToken(size int) (string, error) {
	// Generate 16 random bytes (for good entropy)
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", customError.InternalServerError
	}

	// Optionally, you can append the current timestamp or a UUID to ensure uniqueness
	timestamp := []byte(fmt.Sprintf("%d", time.Now().UnixNano())) // Using nanoseconds as timestamp
	uniqueBytes := append(randomBytes, timestamp...)              // Combine random bytes and timestamp

	// Alternatively, you can use a UUID for uniqueness
	// uuidBytes := uuid.New().Bytes()
	// uniqueBytes := append(randomBytes, uuidBytes...)

	// Generate a base64-encoded string
	token := base64.URLEncoding.EncodeToString(uniqueBytes)

	// Ensure the token is exactly 'size' characters long
	if len(token) > size {
		token = token[:size]
	} else if len(token) < size {
		// If the token is shorter than the required size, pad it
		padding := make([]byte, size-len(token))
		token = token + base64.URLEncoding.EncodeToString(padding)
	}

	return token, nil
}

func CompareTimeIsPassed(t1 time.Time, minute int) bool {
	return time.Duration(time.Since(t1).Hours()) > (time.Duration(minute) * time.Minute)
}
