package grade

import (
	"context"
	"log"
	"net/http"
	"time"

	. "github.com/asynccnu/grade_service_v2/handler"
	"github.com/asynccnu/grade_service_v2/pkg/errno"
	pb "github.com/asynccnu/grade_service_v2/rpc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Get(c *gin.Context) {
	var r Reqeust
	if err := c.ShouldBindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(viper.GetString("data_service_url"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDataProviderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	table, err := client.GetUndergraduateGrade(ctx, &pb.GradeRequest{
		Sid:      c.MustGet("Sid").(string),
		Password: c.MustGet("Password").(string),
		Xqm:      r.XQM,
		Xnm:      r.XNM,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			if st.Code() == codes.Unauthenticated {
				c.JSON(http.StatusOK, Response{
					Code:    errno.ErrPasswordIncorrect.Code,
					Message: st.Message(),
					Data:    nil,
				})
				return
			}
		}
		SendError(c, err, nil)
		return
	}

	gradeList := make([]GradeItem, 0)

	for _, item := range table.Lists {
		gradeList = append(gradeList, GradeItem{
			Course:   item.Kcmc,
			Grade:    item.Cj,
			Credit:   item.Xf,
			Category: item.Kclbmc,
			Type:     item.Kcgsmc,
			Kcxzmc:   item.Kcxzmc,
			Xnm:      item.Xnm,
		})
	}

	SendResponse(c, nil, &gradeList)
}
