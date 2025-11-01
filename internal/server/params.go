package server

import "strings"

func parseParams(paramsStr string) map[string]string {
	params := map[string]string{}

	if paramsStr == "" {
		return params
	}

	for p := range strings.SplitSeq(paramsStr, ",") {
		if p == "" {
			continue
		}

		kv := strings.SplitN(p, "_", 2)
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		} else {
			params[kv[0]] = ""
		}
	}

	return params
}
