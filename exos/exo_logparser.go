package exos

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type LogEntry struct {
	IP     string
	Method string
	Path   string
	Status int
}

type PathCount struct {
	Path  string
	Count int
}

// ParseLine transforme une ligne en LogEntry. Retourne une erreur si la ligne est malformée.
func ParseLine(line string) (LogEntry, error) {
	var logentry LogEntry

	splitline := strings.Fields(line)
	if len(splitline) != 4 {
		return logentry, errors.New("incorrect len")
	}
	logentry.IP = splitline[0]
	logentry.Method = splitline[1]
	logentry.Path = splitline[2]
	status, err := strconv.Atoi(splitline[3])

	if err != nil {
		return logentry, err
	}
	logentry.Status = status

	return logentry, nil
}

// TopErrors retourne les n endpoints (Path) ayant généré le plus d'erreurs (status >= 400),
// triés par nombre d'erreurs décroissant, départage par Path alphabétique.
func TopErrors(entries []LogEntry, n int) []PathCount {
	paths := make(map[string]int)
	for _, entrie := range entries {
		if entrie.Status >= 400 {
			paths[entrie.Path]++
		}
	}

	var pathcounts []PathCount

	for k, v := range paths {

		pathcounts = append(pathcounts, PathCount{k, v})
	}

	sort.Slice(pathcounts, func(i, j int) bool {

		if pathcounts[i].Count != pathcounts[j].Count {

			//Ordre decroissant
			return pathcounts[i].Count > pathcounts[j].Count

		}
		return pathcounts[i].Path < pathcounts[j].Path
	})

	return pathcounts[:min(n, len(pathcounts))]

}
