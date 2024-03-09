package role

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/role/v1"
)

func (c *ControllerV1) GetOneRole(ctx context.Context, req *v1.GetOneRoleReq) (res *v1.GetOneRoleRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}