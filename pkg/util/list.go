package util

func Deduplicate[T comparable](list []T) []T {
	seen := make(map[T]struct{})
	var result []T

	for _, v := range list {
		if _, found := seen[v]; !found {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
