package main

type Array []string

func (arr Array) hasProperty(str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func ShortestPath(graph map[string][]string, start string, end string, path Array) []string {
	if _, exist := graph[start]; !exist {
		return path
	}

	path = append(path, start)
	if start == end {
		return path
	}

	shortest := make([]string, 0)
	for _, node := range graph[start] {
		if !path.hasProperty(node) {
			newPath := ShortestPath(graph, node, end, path)
			if len(newPath) > 0 {
				if len(shortest) == 0 || (len(newPath) < len(shortest)) {
					shortest = newPath
				}
			}
		}
	}
	return shortest
}
