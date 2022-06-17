package composite

import "fmt"

type component interface {
	search(string) bool
}

type file struct {
	name string
}

func (f *file) search(keyword string) bool {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
	if f.name == keyword {
		return true
	}
	return false
}

func (f *file) getName() string {
	return f.name
}

type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) bool {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		if composite.search(keyword) {
			return true
		}
	}
	return false
}

func (f *folder) add(c component) *folder {
	f.components = append(f.components, c)
	return f
}
