// +build windows

package main

import (
	"golang.org/x/sys/windows/registry"
	"fmt"
	"errors"
)

func main() {
	//HKEY_CURRENT_USER\Software\ODBC\ODBC.INI\softdent\Server: WEAVE-3F2ABA82D\PWNGSQL
	//HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Microsoft SQL Server\InstalledInstances: PWNGSQL
	//HKEY_CURRENT_USER\Software\DTMsoft\Connections\ConStr: One path to the full connection string
	//HKEY_CURRENT_USER\Software\DTMsoft\SchemaInspector\Connections\ConStr0: There are a series here, 0 - N, might be useful.
	//key, err := registry.OpenKey(registry.CURRENT_USER, `Software\DTMsoft\Connections`, registry.READ)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//defer key.Close()
	//
	//fmt.Println("Key: ", key)
	//subs, err := key.ReadSubKeyNames(0)
	//if err != nil {
	//	fmt.Println("sub key read failed: " + err.Error())
	//	return
	//}
	//for _, str := range subs {
	//	fmt.Println(str)
	//}
	//fmt.Println(key.GetStringValue("ConStrs"))
	//
	//return windows.GetRegistryValueData(registry.USERS, `S-1-5-21-1960408961-436374069-854245398-1003\Software\DTMsoft\Connections`, `ConStr`, registry.QUERY_VALUE)
	//if err != nil {
	//	mlog.Error("failed to locate connection string")
	//}
	//mlog.Infoln("Connect String: " + connectString)
	//return connectString

	str, err := searchForStringValue(registry.USERS, "ConStr")
	fmt.Println("error: ", err)
	fmt.Println("value: ", str)
}
func searchForStringValue(key registry.Key, value string) (string, error) {
	s, _, e := key.GetStringValue(value)
	if e == nil {
		return s, e
	}

	subs, er := key.ReadSubKeyNames(0)
	if er != nil {
		return "", errors.New("sub key read failed: " + er.Error())
	}
	for _, str := range subs {
		nKey, err := registry.OpenKey(key, str, registry.READ)
		if err == nil {
			str, err := searchForStringValue(nKey, value)
			if err == nil {
				return str, err
			}
		}
	}
	return "", errors.New("unable to locate " + value)
}
