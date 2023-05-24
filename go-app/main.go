// main package
package main

/*
#include <stdlib.h>
typedef struct {
	char* err;
} error;
*/
import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "net/http/pprof" //nolint:gosec
	"unsafe"

	"go-app/internal/logic"
	"go-app/internal/request"
	"go-app/internal/response"
)

// Execute - основная функция скрипта.
func Execute(req *request.Request) (response.Response, error) {
	// Валидация объекта структуры Request.
	if err := request.ValidateStruct(req); err != nil {
		return response.Response{}, err
	}
	// создание словаря данных.
	data := logic.NewData(req)
	// вызов методов логики.
	data.LocalCount()
	data.ResultCount()
	// создание объекта ответа.
	resp := response.NewResponse(data)

	return resp, nil
}

//export invoke
func invoke(inputJSON *C.char, buf *byte, bufSize int64) (*byte, C.error) {
	// создание буфера
	outBytes := unsafe.Slice(buf, bufSize)[:0]
	buffer := bytes.NewBuffer(outBytes)

	var req request.Request
	// парсинг входящего json
	if err := json.Unmarshal([]byte(C.GoString(inputJSON)), &req); err != nil {
		panic(err)
	}
	// вызов логики сервиса
	resp, err := Execute(&req)
	if err != nil {
		return buf, C.error{C.CString(err.Error())}
	}
	// парсинг ответа в []byte
	resBytes, _ := json.Marshal(resp)
	//наполение буфера
	buffer.Write(resBytes)
	buffer.WriteByte(0)
	// проверка на случай, если ответ будет больше заданного буфера
	if int64(buffer.Cap()) > bufSize {
		return buf, C.error{C.CString(fmt.Sprintf("err: bufsize must be %d", buffer.Len()))}
	}

	return buf, C.error{}
}

//export delError
func delError(err C.error) {
	if err.err == nil {
		return
	}
	C.free(unsafe.Pointer(err.err))
}

//// Execute - основная функция скрипта.
//func Execute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	// Парсинг входящего JSON
//	req := &request.Request{}
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	// Валидация объекта структуры Request.
//	if err := request.ValidateStruct(req); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	// создание словаря данных.
//	data := logic.NewData(req)
//	// вызов методов логики.
//	data.LocalCount()
//	data.ResultCount()
//	// создание объекта ответа.
//	resp := response.NewResponse(data)
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(resp) //nolint:errcheck,gosec
//}

func main() {
	//router := httprouter.New()
	//router.POST("/execute", Execute)
	//
	//loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	//
	//log.Println("Starting server...")
	//
	//log.Fatal(http.ListenAndServe(":8000", loggedRouter)) //nolint:gosec
}
