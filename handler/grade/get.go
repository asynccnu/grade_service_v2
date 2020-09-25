package grade

import (
	"context"
	"time"

	. "github.com/asynccnu/grade_service_v2/handler"
	"github.com/asynccnu/grade_service_v2/pkg/errno"
	pb "github.com/asynccnu/grade_service_v2/rpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Get(c *gin.Context) {
	var r Reqeust
	if err := c.ShouldBindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	client := pb.GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	grades, err := client.GetUndergraduateGrade(ctx, &pb.GradeRequest{
		Sid:      c.MustGet("Sid").(string),
		Password: c.MustGet("Password").(string),
		Xqm:      r.XQM,
		Xnm:      r.XNM,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			if st.Code() == codes.Unauthenticated {
				SendResponse(c, errno.ErrPasswordIncorrect, nil)
				return
			}
		}
		SendError(c, err, nil)
		return
	}

	gradeList := make([]GradeItem, 0)

	for _, item := range grades.Lists {
		gradeList = append(gradeList, GradeItem{
			Course:   item.Kcmc,
			Grade:    item.Cj,
			Credit:   item.Xf,
			Category: item.Kclbmc,
			Type:     item.Kcgsmc,
			Kcxzmc:   item.Kcxzmc,
			Xnm:      item.Xnm,
			JxbID:    item.JxbID,
		})
	}

	SendResponse(c, nil, &gradeList)
}
