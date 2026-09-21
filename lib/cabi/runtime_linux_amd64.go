//go:build !cgo && go1.26 && !go1.27

package cabi

import "unsafe"

//go:linkname runtimeCall runtime.cgocall
func runtimeCall(function, frame unsafe.Pointer) int32

//go:linkname runtimeIsCgo runtime.iscgo
var runtimeIsCgo = true

//go:linkname runtimeCrosscall runtime.set_crosscall2
var runtimeCrosscall = func() {}

//go:linkname runtimeInit _cgo_init
var runtimeInit = &bootInit

//go:linkname runtimeThreadStart _cgo_thread_start
var runtimeThreadStart = &bootThreadStart

//go:linkname runtimeNotify _cgo_notify_runtime_init_done
var runtimeNotify = &bootNotify

//go:linkname runtimeKey _cgo_pthread_key_created
var runtimeKey = &unusedKey

//go:linkname runtimeSetenv runtime._cgo_setenv
var runtimeSetenv = &bootSetenv

//go:linkname runtimeUnsetenv runtime._cgo_unsetenv
var runtimeUnsetenv = &bootUnsetenv

var unusedKey uintptr
var setG uintptr
var bootInit, bootThreadStart, bootNotify, bootSetenv, bootUnsetenv byte
var invokeABI, openABI, symbolABI, errorABI, mallocABI, freeABI byte

//go:cgo_import_dynamic goml_cabi_pthread_self pthread_self "libc.so.6"
//go:cgo_import_dynamic goml_cabi_pthread_getattr_np pthread_getattr_np "libc.so.6"
//go:cgo_import_dynamic goml_cabi_pthread_attr_init pthread_attr_init "libc.so.6"
//go:cgo_import_dynamic goml_cabi_pthread_attr_destroy pthread_attr_destroy "libc.so.6"
//go:cgo_import_dynamic goml_cabi_pthread_attr_getstack pthread_attr_getstack "libc.so.6"
//go:cgo_import_dynamic goml_cabi_pthread_attr_getstacksize pthread_attr_getstacksize "libc.so.6"
//go:cgo_import_dynamic goml_cabi_pthread_attr_setdetachstate pthread_attr_setdetachstate "libc.so.6"
//go:cgo_import_dynamic goml_cabi_pthread_create pthread_create#GLIBC_2.34 "libc.so.6"
//go:cgo_import_dynamic goml_cabi_pthread_sigmask pthread_sigmask "libc.so.6"
//go:cgo_import_dynamic goml_cabi_sigfillset sigfillset "libc.so.6"
//go:cgo_import_dynamic goml_cabi_malloc malloc "libc.so.6"
//go:cgo_import_dynamic goml_cabi_free free "libc.so.6"
//go:cgo_import_dynamic goml_cabi_abort abort "libc.so.6"
//go:cgo_import_dynamic goml_cabi_setenv setenv "libc.so.6"
//go:cgo_import_dynamic goml_cabi_unsetenv unsetenv "libc.so.6"
//go:cgo_import_dynamic goml_cabi_dlopen dlopen#GLIBC_2.34 "libc.so.6"
//go:cgo_import_dynamic goml_cabi_dlsym dlsym#GLIBC_2.34 "libc.so.6"
//go:cgo_import_dynamic goml_cabi_dlerror dlerror "libc.so.6"

func Pointer(value uint64) unsafe.Pointer
