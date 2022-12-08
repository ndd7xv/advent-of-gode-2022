package day07

// A SystemResource is either a file or a directory.
// In the case that it is a file, Resources should be nil. If a directory, then it should contain
// the SystemResources that directory holds.
type SystemResource struct {
	Name      string
	Size      *int
	Resources []*SystemResource
}

// NewResource initializes a new resource. Directories start with a nil fileSize since their size
// cannot be calculated until all the subdirectory, and files within it have had their size
// calculated.
func NewResource(name string, fileSize *int, isDir bool) *SystemResource {
	if isDir {
		return &SystemResource{name, fileSize, []*SystemResource{}}
	}
	return &SystemResource{name, fileSize, nil}
}

// AddFile adds a file to a directory's Resources field.
func (root *SystemResource) AddFile(child *SystemResource) {
	if root == nil {
		panic("Adding file to nil SystemResource root!")
	} else if root.Resources == nil {
		panic("Adding a file to another file - should be directory!")
	} else if child == nil {
		panic("Adding nil child!")
	}

	root.Resources = append(root.Resources, child)
}

// GetSize Returns the size of the SystemResource.
// In the case of a directory that does not have a calculated size, this function recursively
// finds the sizes of subdirectories and adds it with the files to get and store the file size of
// the SystemResource so the work to calculate doesn't need to be done again.
func (root *SystemResource) GetSize() int {
	if root.Size != nil {
		return *root.Size
	}

	total := 0

	for _, resource := range root.Resources {
		if resource.Resources == nil {
			total += *resource.Size
		} else {
			total += resource.GetSize()
		}
	}

	root.Size = Of(total)
	return *root.Size
}

// Of is a function that allows us to have a pointer to an int, for fileSize.
// See https://stackoverflow.com/questions/30716354/how-do-i-do-a-literal-int64-in-go for more.
func Of[E any](e E) *E {
	return &e
}
