package codegen

// Walks the type definitions up to the root of each type, following the `Ref`
// field.
//
// Returns the child type if it is not a reference.
func RootRef(childType TypeDefinition, allTypes map[string]TypeDefinition) TypeDefinition {
	currentType := childType

	// Unrolled recursive walk of the types.
	for {
		if !currentType.Schema.IsRef() {
			return childType
		}

		currentType = allTypes[currentType.Schema.RefType]

		// In the unlikely case of a circular reference, we break the loop.
		if currentType.TypeName == childType.TypeName {
			return childType
		}
	}
}
