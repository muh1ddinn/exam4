package handler

import (
	"context"
	"managergetaway/genproto/manager_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTeacher godoc
// @Security ApiKeyAuth
// @Router      /teachers [POST]
// @Summary     Create a teacher
// @Description Create a new teacher
// @Tags        teachers
// @Accept      json
// @Produce     json
// @Param       teacher body manager_service.CreateManagerTeacher true "Teacher"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateTeacher(c *gin.Context) {
	var req manager_service.CreateManagerTeacher

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.grpcClient.TeacherService().Createteacher(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create teacher"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetTeacherByID godoc
// @Security ApiKeyAuth
// @Router      /teachers/{id} [GET]
// @Summary     Get a teacher by its ID
// @Description Retrieves a teacher by its ID
// @Tags        teachers
// @Accept      json
// @Produce     json
// @Param       id path string true "Teacher ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.TeacherService().GetteacherByI(context.Background(), &manager_service.ManagerTeacherPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teacher"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllTeachers godoc
// @Security ApiKeyAuth
// @Router      /teachers [GET]
// @Summary     Get all teachers
// @Description Retrieves information about all teachers
// @Tags        teachers
// @Accept      json
// @Produce     json
// @Param       search query string true "Search"
// @Param       page query uint64 false "Page"
// @Param       limit query uint64 false "Limit"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetAllTeachers(c *gin.Context) {
	var req manager_service.GetListManagerTeacherRequest

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

	resp, err := h.grpcClient.TeacherService().GetAllteacher(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teachers"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateTeacher godoc
// @Security ApiKeyAuth
// @Router      /teachers [PATCH]
// @Summary     Update a teacher
// @Description Update an existing teacher
// @Tags        teachers
// @Accept      json
// @Produce     json
// @Param       teacher body manager_service.UpdateManagerTeacher true "Teacher"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) UpdateTeacher(c *gin.Context) {
	var req manager_service.UpdateManagerTeacher

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.grpcClient.TeacherService().Updateteacher(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update teacher"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteTeacher godoc
// @Security ApiKeyAuth
// @Router      /teachers/{id} [DELETE]
// @Summary     Delete a teacher by its ID
// @Description Deletes a teacher by its ID
// @Tags        teachers
// @Accept      json
// @Produce     json
// @Param       id path string true "Teacher ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	_, err := h.grpcClient.TeacherService().Deleteteacher(context.Background(), &manager_service.ManagerTeacherPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}
