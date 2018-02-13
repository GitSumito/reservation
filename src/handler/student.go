//modify!!

package handler

import (
    "net/http"
    "strconv"

    "api-sample/src/domain/entity"

    "github.com/gin-gonic/gin"
)

/*
Stundets - Response Stundet List
*/
func Stundets(c *gin.Context) {
    /*
        students, err := models.GetStudents()
        if err != nil {
            c.JSON(http.StatusInternalServerError, &APIResponse{Status: AppStatusError, Response: "DB Error"})
            return
        }
        c.JSON(http.StatusOK, &APIResponse{Status: AppStatusOK, Response: students})
    */
    c.JSON(http.StatusInternalServerError, &APIResponse{Status: AppStatusError, Response: "DB Error"})
}

/*
Stundet - Response Stundet Detail
*/
func Stundet(c *gin.Context) {

    sid, err := strconv.Atoi(c.Param("sid"))
    if err != nil {
        c.JSON(http.StatusBadRequest, &APIResponse{Status: AppStatusError, Response: ""})
        return
    }

    student := entity.NewStudent()
    err = student.GetStudentBySID(sid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, &APIResponse{Status: AppStatusError, Response: "DB Error"})
        return
    }
    c.JSON(http.StatusOK, &APIResponse{Status: AppStatusOK, Response: student})
}