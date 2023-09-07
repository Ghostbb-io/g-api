package v1

import (
	"github.com/Ghostbb-io/g-api/app/system/service"
	"github.com/Ghostbb-io/g-api/pkg/ginx"
)

type MenuApi struct {
	service.Menu
}

func (m *MenuApi) Register(ver ginx.VersionFunc) {

}
