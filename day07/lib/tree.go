package day07

// A SystemResource is either a file or a directory.
// In the case that it is a file, Resources should be nil. If a directory, then it should contain
// the SystemResources that directory holds.
type SystemResource struct {
	Name      string
	Size      *int
	Resources []*SystemResource
}

func NewResource(name string, fileSize *int, isDir bool) *SystemResource {
	if isDir {
		return &SystemResource{name, fileSize, []*SystemResource{}}
	}
	return &SystemResource{name, fileSize, nil}
}

func (root *SystemResource) AddFile(child *SystemResource) {
	if root == nil {
		panic("Adding file to nil SystemResource root!")
	} else if child == nil {
		panic("Adding nil child!")
	}

	root.Resources = append(root.Resources, child)
}

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

func Of[E any](e E) *E {
	return &e
}
