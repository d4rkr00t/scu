package target

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
	"strings"
)

type TargetsMap = map[string]Target

type Target struct {
	Cmd     string
	Deps    []string
	Outputs []string
}

func (t Target) String() string {
	return fmt.Sprintf("%s:%s:%v", t.Cmd, strings.Join(t.Deps, ","), strings.Join(t.Outputs, ","))
}

func MergeTargets(targets ...*TargetsMap) TargetsMap {
	var mergedTargets = TargetsMap{}

	for _, tm := range targets {
		for _, target := range GetSortedTargetsNames(tm) {
			mergedTargets[target] = (*tm)[target]
		}
	}

	return mergedTargets
}

func GetSortedTargetsNames(tm *TargetsMap) []string {
	var targetsNames = []string{}
	for groupName := range *tm {
		targetsNames = append(targetsNames, groupName)
	}
	sort.Strings(targetsNames)
	return targetsNames
}

func GetAllTargetsOutputs(tm *TargetsMap) []string {
	var outputs = []string{}

	for _, rule := range *tm {
		outputs = append(outputs, rule.Outputs...)
	}

	return outputs
}

func GetTargetsHash(tm *TargetsMap) string {
	var h = sha1.New()
	for _, targetName := range GetSortedTargetsNames(tm) {
		io.WriteString(h, (*tm)[targetName].String())
	}
	return hex.EncodeToString(h.Sum(nil))
}
