package grade

import (
	"context"
	"fmt"
	"time"

	"github.com/asynccnu/grade_service_v2/handler"
	"github.com/asynccnu/grade_service_v2/pkg/errno"
	pb "github.com/asynccnu/grade_service_v2/rpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetDetail ... 获取平时和期末成绩
func GetDetail(c *gin.Context) {
	var r GradeDetailRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		fmt.Println(err)
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	client := pb.GetClient()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	detail, err := client.GetUndergraduateGradeDetail(ctx, &pb.GradeDetailRequest{
		Sid:      c.MustGet("Sid").(string),
		Password: c.MustGet("Password").(string),
		Course:   r.Course,
		ClassId:  r.ClassID,
		Term:     r.Term,
		Year:     r.Year,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			if st.Code() == codes.Unauthenticated {
				handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
				return
			}
		}
		handler.SendError(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, &GradeDetailResponse{
		Course:          r.Course,
		UsualGrade:      detail.UsualGrade,
		FinalGrade:      detail.FinalGrade,
		UsualPercentage: detail.UsualPercentage,
		FinalPercentage: detail.FinalPercentage,
	})
}
