package handler

import (
	"context"
	"managergetaway/genproto/manager_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSupportTeacher godoc
// @Security ApiKeyAuth
// @Router      /supportteachers [POST]
// @Summary     Create a support teacher
// @Description Create a new support teacher
// @Tags        supportteachers
// @Accept      json
// @Produce     json
// @Param       supportteacher body manager_service.CreateManagerSupportTeacher true "Support Teacher"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateSupportTeacher(c *gin.Context) {
	var req manager_service.CreateManagerSupportTeacher

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.grpcClient.SupportTeacherService().Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create support teacher"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetSupportTeacherByID godoc
// @Security ApiKeyAuth
// @Router      /supportteachers/{id} [GET]
// @Summary     Get a support teacher by its ID
// @Description Retrieves a support teacher by its ID
// @Tags        supportteachers
// @Accept      json
// @Produce     json
// @Param       id path string true "Support Teacher ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetSupportTeacherByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SupportTeacherService().GetByID(context.Background(), &manager_service.ManagerSupportTeacherPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve support teacher"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllSupportTeachers godoc
// @Security ApiKeyAuth
// @Router      /supportteachers [GET]
// @Summary     Get all support teachers
// @Description Retrieves information about all support teachers
// @Tags        supportteachers
// @Accept      json
// @Produce     json
// @Param       search query string true "Search"
// @Param       page query uint64 false "Page"
// @Param       limit query uint64 false "Limit"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetAllSupportTeachers(c *gin.Context) {
	var req manager_service.GetListManagerSupportTeacherRequest

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

	resp, err := h.grpcClient.SupportTeacherService().GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve support teachers"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateSupportTeacher godoc
// @Security ApiKeyAuth
// @Router      /supportteachers [PATCH]
// @Summary     Update a support teacher
// @Description Update an existing support teacher
// @Tags        supportteachers
// @Accept      json
// @Produce     json
// @Param       supportteacher body manager_service.UpdateManagerSupportTeacher true "Support Teacher"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) UpdateSupportTeacher(c *gin.Context) {
	var req manager_service.UpdateManagerSupportTeacher

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.grpcClient.SupportTeacherService().Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update support teacher"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteSupportTeacher godoc
// @Security ApiKeyAuth
// @Router      /supportteachers/{id} [DELETE]
// @Summary     Delete a support teacher by its ID
// @Description Deletes a support teacher by its ID
// @Tags        supportteachers
// @Accept      json
// @Produce     json
// @Param       id path string true "Support Teacher ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) DeleteSupportTeacher(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	_, err := h.grpcClient.SupportTeacherService().Delete(context.Background(), &manager_service.ManagerSupportTeacherPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete support teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Support teacher deleted successfully"})
}
