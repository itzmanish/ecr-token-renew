package k8s

import (
	"context"
	"regexp"
	"strings"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespaces(envVar string) ([]string, error) {
	list := strings.Split(envVar, ",")
	single := make([]string, 0, len(list))
	rgx := make([]*regexp.Regexp, 0, len(list))

	for _, val := range list {
		if hasWildCard(val) {
			r, err := getRegex(val)
			if err != nil {
				return nil, err
			}
			rgx = append(rgx, r)
		} else {
			single = append(single, val)
		}
	}

	matchedNamespaces, err := findNamespaces(rgx)
	if err != nil {
		return nil, err
	}

	return unique(append(single, matchedNamespaces...)), nil
}

func unique(values []string) []string {
	result := make([]string, 0, len(values))
	check := map[string]bool{}
	for _, val := range values {
		_, ok := check[val]
		if !ok {
			check[val] = true
			result = append(result, val)
		}
	}
	return result
}

func hasWildCard(val string) bool {
	for _, r := range []rune{'*', '?'} {
		if strings.ContainsRune(val, r) {
			return true
		}
	}
	return false
}

func getRegex(val string) (*regexp.Regexp, error) {
	rgx := strings.Replace(val, "*", ".*", -1)
	rgx = strings.Replace(rgx, "?", ".", -1)
	rgx = "^" + rgx + "$"
	return regexp.Compile(rgx)
}

func findNamespaces(rgx []*regexp.Regexp) ([]string, error) {
	if len(rgx) == 0 {
		return nil, nil
	}

	namespaces, err := getAllNamespaces()
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(namespaces))
	for _, ns := range namespaces {
		if isAnyMatch(ns, rgx) {
			result = append(result, ns)
		}
	}

	return result, nil
}

func isAnyMatch(ns string, regexes []*regexp.Regexp) bool {
	for _, r := range regexes {
		if r.MatchString(ns) {
			return true
		}
	}

	return false
}

func getAllNamespaces() ([]string, error) {
	var result []string

	client, err := GetClient()
	if nil != err {
		return nil, err
	}

	opts := metaV1.ListOptions{}
	first := true

	for first || opts.Continue != "" {
		first = false
		res, err := client.CoreV1().Namespaces().List(context.TODO(), opts)
		if nil != err {
			return nil, err
		}

		opts.Continue = res.Continue
		newNames := make([]string, len(res.Items))
		for i, item := range res.Items {
			newNames[i] = item.Name
		}

		result = append(result, newNames...)
	}

	return result, nil
}
