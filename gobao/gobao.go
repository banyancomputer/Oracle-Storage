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
    "encode/hex"
)
ZZ
func obao_data(filepath string) *C.ObaoData {
    return C.obao_data(C.CString(filepath))
}

func ProcessFile(filepath string) (string, string) {
    // Generate Obao Data
    obao_data := obao_data(filepath)


}
