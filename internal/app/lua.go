package app

import lua "github.com/yuin/gopher-lua"

func RunLuaScript(script string) (string, error) {
	L := lua.NewState()
	defer L.Close()

	var output string
	L.SetGlobal("print", L.NewFunction(func(L *lua.LState) int {
		output += L.ToString(1) + "\n"
		return 0
	}))

	if err := L.DoString(script); err != nil {
		return "", err
	}

	return output, nil
}
