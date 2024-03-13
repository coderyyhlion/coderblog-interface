package article

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"coderblog-interface/api/article/v1"
)

func (c *ControllerV1) UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (res *v1.UpdateArticleRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
