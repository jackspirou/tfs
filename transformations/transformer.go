package transformations

// Transformer describes a terraform state file transformation.
type Transformer interface {
	Transform(string, error)
}
