package app

/*
#cgo CFLAGS: -I/opt/homebrew/opt/lua@5.4/include/lua
#cgo LDFLAGS: -L/opt/homebrew/opt/lua@5.4/lib -llua
#include <stdlib.h>
#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>

// Forward declaration of the Go function
extern void go_print(char*);

// Wrapper to call the Go function from Lua
extern int lua_go_print(lua_State *L);

// Helper function to execute Lua code from a string
static int run_lua_string(lua_State *L, const char *code) {
    int status = luaL_loadstring(L, code);
    if (status != LUA_OK) {
        return status;
    }
    status = lua_pcallk(L, 0, LUA_MULTRET, 0, 0, NULL);
    return status;
}

// Function to create a limited Lua environment
static void create_sandbox(lua_State *L) {
    luaL_openlibs(L);
    lua_pushnil(L);  // Remove the "os" library
    lua_setglobal(L, "os");
    lua_pushnil(L);  // Remove the "io" library
    lua_setglobal(L, "io");
    lua_pushnil(L);  // Remove the "package" library
    lua_setglobal(L, "package");
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export go_print
func go_print(msg *C.char) {
	fmt.Println(C.GoString(msg))
}

//export lua_go_print
func lua_go_print(L *C.lua_State) C.int {
	str := C.luaL_checklstring(L, 1, nil)
	go_print((*C.char)(str))
	return 0
}

func RegisterLuaFunctions(L *C.lua_State) {
	// Register the Go function to Lua
	funcName := C.CString("go_print")
	defer C.free(unsafe.Pointer(funcName))
	C.lua_pushcclosure(L, (C.lua_CFunction)(unsafe.Pointer(C.lua_go_print)), 0)
	C.lua_setglobal(L, funcName)
}

func RunLuaScript(script string) (string, error) {
	// Initialize a new Lua state
	L := C.luaL_newstate()
	if L == nil {
		return "", fmt.Errorf("failed to create Lua state")
	}
	defer C.lua_close(L)

	// Create a sandboxed Lua environment
	C.create_sandbox(L)

	// Register the Go function to Lua
	RegisterLuaFunctions(L)

	// Prepare the Lua script
	cScript := C.CString(script)
	defer C.free(unsafe.Pointer(cScript))

	// Capture the output
	var output string
	goPrint = func(msg *C.char) {
		output += C.GoString(msg) + "\n"
	}

	// Run the Lua script
	if status := C.luaL_loadstring(L, cScript); status != C.LUA_OK {
		err := C.lua_tolstring(L, -1, nil)
		return "", fmt.Errorf("error loading Lua script: %s", C.GoString(err))
	}
	if status := C.lua_pcallk(L, 0, C.LUA_MULTRET, 0, 0, nil); status != C.LUA_OK {
		err := C.lua_tolstring(L, -1, nil)
		return "", fmt.Errorf("error running Lua script: %s", C.GoString(err))
	}

	return output, nil
}

var goPrint func(*C.char)
