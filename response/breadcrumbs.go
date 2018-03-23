package response

import "fmt"

type BreadCrumb interface {
	ToString() string // Deprecated
	AndField(fieldName string) BreadCrumb
	Dive() BreadCrumb
	Indexed(index int) BreadCrumb
	AsSuperCrumb() string
}

func Initial() BreadCrumb {
	return breadCrumbImpl{
		stringRepr: "",
		isStart: true,
		individualizer: 1,
	}
}

type breadCrumbImpl struct {
	isStart bool
	stringRepr string
	individualizer int
}

func (b breadCrumbImpl) ToString() string {
	return b.stringRepr
}

func (b breadCrumbImpl) AndField(fieldName string) BreadCrumb {
	if b.isStart {
		return breadCrumbImpl{
			isStart: false,
			stringRepr: fieldName,
			individualizer: b.individualizer,
		}
	}
	return breadCrumbImpl{
		isStart: false,
		stringRepr: b.stringRepr + "." + fieldName,
		individualizer: b.individualizer,
	}
}

func (b breadCrumbImpl) Dive() BreadCrumb {
	return breadCrumbImpl{
		isStart: false,
		stringRepr: b.stringRepr,
		individualizer: b.individualizer * 10,
	}
}

func (b breadCrumbImpl) Indexed(index int) BreadCrumb {
	return breadCrumbImpl{
		isStart: false,
		stringRepr: b.stringRepr,
		individualizer: b.individualizer + index + 1,
	}
}

func (b breadCrumbImpl) AsSuperCrumb() string {
	return fmt.Sprintf("%s%d", b.stringRepr, b.individualizer)
}

