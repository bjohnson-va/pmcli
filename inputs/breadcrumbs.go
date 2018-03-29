package inputs

import "fmt"

type BreadCrumb interface {
	ToString() string // Deprecated
	AndField(fieldName string) BreadCrumb
	Indexed(index int) BreadCrumb
	AsSuperCrumb() string
	GetIndividualizer() int
}

func InitialBreadCrumb() BreadCrumbImpl {
	return BreadCrumbImpl{
		stringRepr:     "",
		isStart:        true,
		individualizer: 0,
	}
}

type BreadCrumbImpl struct {
	isStart        bool
	stringRepr     string
	individualizer int
}

func (b BreadCrumbImpl) ToString() string {
	return b.stringRepr
}

func (b BreadCrumbImpl) AndField(fieldName string) BreadCrumb {
	if b.isStart {
		return BreadCrumbImpl{
			isStart:        false,
			stringRepr:     fieldName,
			individualizer: b.individualizer + 1,
		}
	}
	return BreadCrumbImpl{
		isStart:        false,
		stringRepr:     b.stringRepr + "." + fieldName,
		individualizer: b.individualizer*10 + 1,
	}
}

func (b BreadCrumbImpl) Indexed(index int) BreadCrumb {
	if b.isStart {
		return b
	}
	return BreadCrumbImpl{
		isStart:        false,
		stringRepr:     b.stringRepr,
		individualizer: ((b.individualizer / 10) * 10) + index + 1,
	}
}

func (b BreadCrumbImpl) AsSuperCrumb() string {
	return fmt.Sprintf("%s%d", b.stringRepr, b.individualizer)
}

func (b BreadCrumbImpl) GetIndividualizer() int {
	return b.individualizer
}
