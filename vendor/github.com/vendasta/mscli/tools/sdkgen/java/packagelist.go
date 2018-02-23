package java

type PackageList struct {
	packageMap map[string]bool
}

func newPackageList() PackageList {
	return PackageList{packageMap: map[string]bool{}}
}

func (p *PackageList) Add(pkg string) {
	p.packageMap[pkg] = true
}

func (p *PackageList) Merge(pkgList PackageList) {
	for k := range pkgList.packageMap {
		p.Add(k)
	}
}

func (p *PackageList) List() []string {
	rv := []string{}
	for k := range p.packageMap {
		rv = append(rv, k)
	}
	return rv
}
