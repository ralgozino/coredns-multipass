package multipass

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type VMDetails struct {
	IPv4    []string `json:"ipv4"`
	Name    string   `json:"name"`
	Release string   `json:"release"`
	State   string   `json:"state"`
}

type List struct {
	List []VMDetails `json:"list"`
}

func vmList() (map[string][]string, error) {
	vmList := map[string][]string{}
	cmd := exec.Command("multipass", "list", "--format=json")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error while invoking multipass: %s", err)
	}
	list := new(List)
	err = json.Unmarshal([]byte(out.String()), &list)
	if err != nil {
		return nil, fmt.Errorf("error while serializing response from multipass: %s", err)
	}

	for i := range list.List {
		vmList[list.List[i].Name] = list.List[i].IPv4
	}
	return vmList, nil
}
