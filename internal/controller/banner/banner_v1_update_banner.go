package banner

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/banner/v1"
)

func (c *ControllerV1) UpdateBanner(ctx context.Context, req *v1.UpdateBannerReq) (res *v1.UpdateBannerRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}