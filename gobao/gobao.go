package gobao

// DYNAMICALLY LINKED LIBRARY IMPORT

/*
#cgo LDFLAGS: -L./lib -lobao
#include "./lib/obao.h"
*/


//STATIC LIBRARY IMPORT

/*
#cgo LDFLAGS: ./lib/libobao.a -ldl
#include "./lib/obao.h"
*/
import "C"

import (
    "encoding/hex"
    "io/ioutil"
    "unsafe"
    "fmt"
)

func obao_data(filepath string) *C.ObaoData {
    return C.obao_data(C.CString(filepath))
}

func ProcessFile(filepath string) (string, error) {
    // Generate Obao Data
    obao_data := obao_data(filepath)

    // Convert obao_data and hash to byte strings
    obao_bytes := C.GoBytes(unsafe.Pointer(obao_data.obao), C.int(obao_data.obao_len))
    hash_bytes := C.GoBytes(unsafe.Pointer(obao_data.hash), C.int(obao_data.hash_len))

    // Save the obao_data_bytes to a file
    err := ioutil.WriteFile(filepath + ".obao", obao_bytes, 0644)

    return hex.EncodeToString(hash_bytes), err
}

// func main() {
//     // Test with a file in the current directory
//     hash, err := ProcessFile("./test/ethereum.pdf")
//     if err != nil {
//         panic(err)
//     }
//     fmt.Println("Go Hash:", hash)
// }