package handler

import (
	"context"
	"fmt"
	"managergetaway/genproto/manager_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateStudent godoc
// @Security ApiKeyAuth
// @Router      /students [POST]
// @Summary     Create a student
// @Description Create a new student
// @Tags        students
// @Accept      json
// @Produce     json
// @Param       student body manager_service.CreateManagerStudent true "student"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateStudent(c *gin.Context) {
	var req manager_service.CreateManagerStudent

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.StudentService().CreateStudent(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating handler")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllStudents godoc
// @Security ApiKeyAuth
// @Router      /students [GET]
// @Summary     Get all students
// @Description Retrieves information about all students
// @Tags        students
// @Accept      json
// @Produce     json
// @Param       search query string true "Search"
// @Param       page query uint64 false "Page"
// @Param       limit query uint64 false "Limit"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetAllStudents(c *gin.Context) {
	var req manager_service.GetListManagerStudentRequest

	req.Search = c.Query("search")
	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	req.Page = page

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	req.Limit = limit

	resp, err := h.grpcClient.StudentService().GetAllStudents(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	c.JSON(http.StatusOK, resp)

}

// GetStudentByID godoc
// @Security ApiKeyAuth
// @Router      /students/{id} [GET]
// @Summary     Get a student by its ID
// @Description Retrieves a student by its ID
// @Tags        students
// @Accept      json
// @Produce     json
// @Param       id path string true "Student ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.StudentService().GetStudentByID(context.Background(), &manager_service.ManagerStudentPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)

}

// UpdateStudent godoc
// @Security ApiKeyAuth
// @Router      /students [PATCH]
// @Summary     Update a student
// @Description Update an existing student
// @Tags        students
// @Accept      json
// @Produce     json
// @Param       student body manager_service.UpdateManagerStudent true "student"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) UpdateStudent(c *gin.Context) {
	var req manager_service.UpdateManagerStudent

	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	resp, err := h.grpcClient.StudentService().UpdateStudent(context.Background(), &req)
	if err != nil {
		return
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)

}

// DeleteStudent godoc
// @Security ApiKeyAuth
// @Router      /students/{id} [DELETE]
// @Summary     Delete a student by its ID
// @Description Deletes a student by its ID
// @Tags        students
// @Accept      json
// @Produce     json
// @Param       id path string true "Student ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.StudentService().DeleteStudent(context.Background(), &manager_service.ManagerStudentPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while delete")
		return
	}

	c.JSON(http.StatusOK, resp)
}
