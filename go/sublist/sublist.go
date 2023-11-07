package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	var swapped bool
	if equal(l1, l2) {
		return RelationEqual
	}
	if len(l2) == 0 {
		return RelationSuperlist
	}
	if len(l1) == 0 {
		return RelationSublist
	}

	if len(l1) < len(l2) {
		l1, l2 = l2, l1
		swapped = true
	}
	for i := 0; i <= len(l1)-len(l2); i++ {
		if equal(l1[i:i+len(l2)], l2) {
			if swapped {
				return RelationSublist
			} else {
				return RelationSuperlist
			}
		}
	}

	return RelationUnequal
}

func equal(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if (l1)[i] != (l2)[i] {
			return false
		}
	}
	return true
}
