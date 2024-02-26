// Package sliceutils offers utility functions for slice manipulation and analysis in Go.
package sliceutils

// PowerSet returns a power-set of the input in 'src'.
func PowerSet[T any, TS ~[]T](src TS) []TS {
	powerSet := make([]TS, 1<<uint(len(src)))
	powerSet[0] = TS{}
	p := 1
	for _, element := range src {
		for i := range p {
			existingSubset := powerSet[i]
			newSubset := make(TS, len(existingSubset)+1)
			copy(newSubset, existingSubset)
			newSubset[len(existingSubset)] = element
			powerSet[p] = newSubset
			p++
		}
	}

	return powerSet
}

// Product generates the Cartesian product of variable number of slices.
func Product[T any, TS ~[]T](slices ...TS) []TS {
	if len(slices) == 0 {
		return []TS{}
	}

	product := []TS{{}}
	for _, slice := range slices {
		var tempProduct []TS
		for _, existingCombo := range product {
			for _, item := range slice {
				newCombo := make(TS, len(existingCombo)+1)
				copy(newCombo, existingCombo)
				newCombo[len(existingCombo)] = item
				tempProduct = append(tempProduct, newCombo)
			}
		}
		product = tempProduct
	}

	return product
}
