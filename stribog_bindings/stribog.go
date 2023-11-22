package stribog_bindings

/*
#include "stribog.h"

typedef unsigned long long ullong;
typedef const unsigned char *cpuchar;
typedef unsigned char *puchar;
*/
import "C"
import "unsafe"

func Hash256(message []byte) []byte {
	out := make([]byte, 256/8)
	C.hash_256(C.cpuchar(unsafe.Pointer(&message[0])), C.ullong(len(message))*8, C.puchar(unsafe.Pointer(&out[0])))
	return out
}

func Hash512(message []byte) []byte {
	out := make([]byte, 512/8)
	C.hash_512(C.cpuchar(unsafe.Pointer(&message[0])), C.ullong(len(message))*8, C.puchar(unsafe.Pointer(&out[0])))
	return out
}
