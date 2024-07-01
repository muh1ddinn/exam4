package handler

import (
	"context"
	"fmt"
	"managergetaway/genproto/superadmin_service"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateAdminStudent godoc
// @Security ApiKeyAuth
// @Router      /superadminstudents [POST]
// @Summary     Create a superadminstudents
// @Description Create a new superadminstudents
// @Tags        superadminstudents
// @Accept      json
// @Produce     json
// @Param       student body superadmin_service.CreateManagerStudent true "student"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateadminStudent(c *gin.Context) {
	var req superadmin_service.CreateManagerStudent

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().CreateStudent(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating handler")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllAdminStudents godoc
// @Security ApiKeyAuth
// @Router      /superadminstudents [GET]
// @Summary     Get all superadminstudents
// @Description Retrieves information about all students
// @Tags        superadminstudents
// @Accept      json
// @Produce     json
// @Param       search query string true "Search"
// @Param       page query uint64 false "Page"
// @Param       limit query uint64 false "Limit"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetAlladminStudents(c *gin.Context) {
	var req superadmin_service.GetListManagerStudentRequest

	req.Search = c.Query("search")
	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	req.Page = int32(page)

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	req.Limit = int32(limit)

	resp, err := h.grpcClient.SuperadminService().GetAllStudents(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	c.JSON(http.StatusOK, resp)

}

// GetAdminStudentByID godoc
// @Security ApiKeyAuth
// @Router      /superadminstudents/{id} [GET]
// @Summary     Get a superadminstudents by its ID
// @Description Retrieves a superadminstudents by its ID
// @Tags        studensuperadminstudentsts
// @Accept      json
// @Produce     json
// @Param       id path string true "Student ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetadminStudentByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.SuperadminService().GetStudentByID(context.Background(), &superadmin_service.ManagerStudentPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)

}

// UpdateStudent godoc
// @Security ApiKeyAuth
// @Router      /superadminstudents [PATCH]
// @Summary     Update a student
// @Description Update an existing superadminstudents
// @Tags        superadminstudents
// @Accept      json
// @Produce     json
// @Param       superadminstudents body superadmin_service.UpdateManagerStudent true "student"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) UpdateAdminStudent(c *gin.Context) {
	var req superadmin_service.UpdateManagerStudent

	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}

	resp, err := h.grpcClient.SuperadminService().UpdateStudent(context.Background(), &req)
	if err != nil {
		return
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)

}

// DeleteStudent godoc
// @Security ApiKeyAuth
// @Router      /superadminstudents/{id} [DELETE]
// @Summary     Delete a superadminstudents by its ID
// @Description Deletes a superadminstudents by its ID
// @Tags        superadminstudents
// @Accept      json
// @Produce     json
// @Param       id path string true "Student ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) DeleteAdminStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.SuperadminService().DeleteStudent(context.Background(), &superadmin_service.ManagerStudentPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while delete")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreatedminTeacher godoc
// @Security ApiKeyAuth
// @Router /superadminteachers [POST]
// @Summary Create a teacher
// @Description Create a new teacher
// @Tags teacher
// @Accept json
// @Produce json
// @Param teacher body superadmin_service.CreateManagerTeacher true "Teacher"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) CreatedminTeacher(c *gin.Context) {
	var req superadmin_service.CreateManagerTeacher

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().CreateTeacher(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllAdminTeachers godoc
// @Security ApiKeyAuth
// @Router /superadminteachers [GET]
// @Summary Get all teachers
// @Description Retrieves information about all teachers
// @Tags teacher
// @Accept json
// @Produce json
// @Param search query string true "Search"
// @Param page query uint64 false "Page"
// @Param limit query uint64 false "Limit"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetAllAdminTeachers(c *gin.Context) {
	var req superadmin_service.GetListManagerTeacherRequest

	req.Search = c.Query("search")
	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	req.Page = int32(page)

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	req.Limit = int32(limit)

	resp, err := h.grpcClient.SuperadminService().GetAllTeachers(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting all teachers")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAdminTeacherByID godoc
// @Security ApiKeyAuth
// @Router /superadminteachers/{id} [GET]
// @Summary Get a teacher by its ID
// @Description Retrieves a teacher by its ID
// @Tags teacher
// @Accept json
// @Produce json
// @Param id path string true "Teacher ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetAdminTeacherByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SuperadminService().GetTeacherByID(context.Background(), &superadmin_service.ManagerTeacherPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting teacher by ID")
		return
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// UpdateAdminTeacher godoc
// @Security ApiKeyAuth
// @Router /superadminteachers [PATCH]
// @Summary Update a teacher
// @Description Update an existing teacher
// @Tags teacher
// @Accept json
// @Produce json
// @Param teacher body superadmin_service.UpdateManagerTeacher true "Teacher"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) UpdateAdminTeacher(c *gin.Context) {
	var req superadmin_service.UpdateManagerTeacher

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().UpdateTeacher(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating teacher")
		return
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// DeleteAdminTeacher godoc
// @Security ApiKeyAuth
// @Router /superadminteachers/{id} [DELETE]
// @Summary Delete a teacher by its ID
// @Description Deletes a teacher by its ID
// @Tags teacher
// @Accept json
// @Produce json
// @Param id path string true "Teacher ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) DeleteAdminTeacher(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SuperadminService().DeleteTeacher(context.Background(), &superadmin_service.ManagerTeacherPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting teacher")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateSupportAdminTeacher godoc
// @Security ApiKeyAuth
// @Router /superadminsupportteachers [POST]
// @Summary Create a support teacher
// @Description Create a new support teacher
// @Tags supportteacher
// @Accept json
// @Produce json
// @Param supportteacher body superadmin_service.CreateManagerSupportTeacher true "Support Teacher"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) CreateSupportAdminTeacher(c *gin.Context) {
	var req superadmin_service.CreateManagerSupportTeacher

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().CreateSupportTeacher(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating support teacher")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllSupportAdminTeachers godoc
// @Security ApiKeyAuth
// @Router /superadminsupportteachers [GET]
// @Summary Get all support teachers
// @Description Retrieves information about all support teachers
// @Tags supportteacher
// @Accept json
// @Produce json
// @Param search query string true "Search"
// @Param page query uint64 false "Page"
// @Param limit query uint64 false "Limit"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetAllSupportAdminTeachers(c *gin.Context) {
	var req superadmin_service.GetListManagerSupportTeacherRequest

	req.Search = c.Query("search")
	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	req.Page = int32(page)

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	req.Limit = int32(limit)

	resp, err := h.grpcClient.SuperadminService().GetAllSupportTeachers(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting all support teachers")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetSupportAdminTeacherByID godoc
// @Security ApiKeyAuth
// @Router /superadminsupportteachers/{id} [GET]
// @Summary Get a support teacher by its ID
// @Description Retrieves a support teacher by its ID
// @Tags supportteacher
// @Accept json
// @Produce json
// @Param id path string true "Support Teacher ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetSupportAdminTeacherByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SuperadminService().GetSupportTeacherByID(context.Background(), &superadmin_service.ManagerSupportTeacherPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting support teacher by ID")
		return
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// UpdateSupportAdminTeacher godoc
// @Security ApiKeyAuth
// @Router /superadminsupportteachers [PATCH]
// @Summary Update a support teacher
// @Description Update an existing support teacher
// @Tags supportteacher
// @Accept json
// @Produce json
// @Param supportteacher body superadmin_service.UpdateManagerSupportTeacher true "Support Teacher"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) UpdateSupportAdminTeacher(c *gin.Context) {
	var req superadmin_service.UpdateManagerSupportTeacher

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().UpdateSupportTeacher(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating support teacher")
		return
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// DeleteSupportAdminTeacher godoc
// @Security ApiKeyAuth
// @Router /superadminsupportteachers/{id} [DELETE]
// @Summary Delete a support teacher by its ID
// @Description Deletes a support teacher by its ID
// @Tags supportteacher
// @Accept json
// @Produce json
// @Param id path string true "Support Teacher ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) DeleteSupportAdminTeacher(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SuperadminService().DeleteSupportTeacher(context.Background(), &superadmin_service.ManagerSupportTeacherPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting support teacher")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateAdministrator godoc
// @Security ApiKeyAuth
// @Router /administrators [POST]
// @Summary Create an administrator
// @Description Create a new administrator
// @Tags administrator
// @Accept json
// @Produce json
// @Param administrator body superadmin_service.CreateManagerAdministrator true "Administrator"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) CreateAdministrator(c *gin.Context) {
	var req superadmin_service.CreateManagerAdministrator

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().CreateAdministrator(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating administrator")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllAdministrators godoc
// @Security ApiKeyAuth
// @Router /administrators [GET]
// @Summary Get all administrators
// @Description Retrieves information about all administrators
// @Tags administrator
// @Accept json
// @Produce json
// @Param search query string true "Search"
// @Param page query uint64 false "Page"
// @Param limit query uint64 false "Limit"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetAllAdministrators(c *gin.Context) {
	var req superadmin_service.GetListManagerAdministratorRequest

	req.Search = c.Query("search")
	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	req.Page = int32(page)

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	req.Limit = int32(limit)

	resp, err := h.grpcClient.SuperadminService().GetAllAdministrators(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting all administrators")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAdministratorByID godoc
// @Security ApiKeyAuth
// @Router /administrators/{id} [GET]
// @Summary Get an administrator by its ID
// @Description Retrieves an administrator by its ID
// @Tags administrator
// @Accept json
// @Produce json
// @Param id path string true "Administrator ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetAdministratorByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SuperadminService().GetAdministratorByID(context.Background(), &superadmin_service.ManagerAdministratorPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting administrator by ID")
		return
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// DeleteAdministrator godoc
// @Security ApiKeyAuth
// @Router /administrators/{id} [DELETE]
// @Summary Delete an administrator by its ID
// @Description Deletes an administrator by its ID
// @Tags administrator
// @Accept json
// @Produce json
// @Param id path string true "Administrator ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) DeleteAdministrator(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SuperadminService().DeleteAdministrator(context.Background(), &superadmin_service.ManagerAdministratorPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting administrator")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateManager godoc
// @Security ApiKeyAuth
// @Router /managers [POST]
// @Summary Create a manager
// @Description Create a new manager
// @Tags manager
// @Accept json
// @Produce json
// @Param manager body superadmin_service.CreateManagerRequest true "Manager"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) CreateManager(c *gin.Context) {
	var req superadmin_service.CreateManagerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().CreateManager(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating manager")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllManagers godoc
// @Security ApiKeyAuth
// @Router /managers [GET]
// @Summary Get all managers
// @Description Retrieves information about all managers
// @Tags manager
// @Accept json
// @Produce json
// @Param search query string true "Search"
// @Param page query uint64 false "Page"
// @Param limit query uint64 false "Limit"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetAllManagers(c *gin.Context) {
	var req superadmin_service.GetAllManagersRequest

	req.Search = c.Query("search")
	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}
	req.Page = int32(page)

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}
	req.Limit = int32(limit)

	resp, err := h.grpcClient.SuperadminService().GetAllManagers(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting all managers")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetManagerByID godoc
// @Security ApiKeyAuth
// @Router /managers/{id} [GET]
// @Summary Get a manager by its ID
// @Description Retrieves a manager by its ID
// @Tags manager
// @Accept json
// @Produce json
// @Param id path string true "Manager ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) GetManagerByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SuperadminService().GetManagerByID(context.Background(), &superadmin_service.ManagerPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while getting manager by ID")
		return
	}

	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// UpdateManager godoc
// @Security ApiKeyAuth
// @Router /managers [PATCH]
// @Summary Update a manager
// @Description Update an existing manager
// @Tags manager
// @Accept json
// @Produce json
// @Param manager body superadmin_service.UpdateManagerRequest true "Manager"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) UpdateManager(c *gin.Context) {
	var req superadmin_service.UpdateManagerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.SuperadminService().UpdateManager(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while updating manager")
		return
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}

// DeleteManager godoc
// @Security ApiKeyAuth
// @Router /managers/{id} [DELETE]
// @Summary Delete a manager by its ID
// @Description Deletes a manager by its ID
// @Tags manager
// @Accept json
// @Produce json
// @Param id path string true "Manager ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handler) DeleteManager(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	resp, err := h.grpcClient.SuperadminService().DeleteManager(context.Background(), &superadmin_service.ManagerPrimaryKey{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while deleting manager")
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *handler) ExportToCSV(c *gin.Context) {
	var req superadmin_service.TableName

	// Parse query parameters
	req.TableName = c.Query("table")
	req.Outfilename = c.Query("outfilename")

	// Ensure outfilename ends with ".csv"
	if !strings.HasSuffix(req.Outfilename, ".csv") {
		req.Outfilename += ".csv"
	}

	// Call service method to export to CSV
	_, err := h.grpcClient.SuperadminService().ExportToCSV(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(req.Outfilename)
	// Read the file after it's exported (assuming it exists)
	data, err := os.ReadFile(req.Outfilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	fmt.Println(err, "err")
	fmt.Println(data, "dataaaaaaaaa")

	// Set response headers for CSV download
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(req.Outfilename)))

	// Return CSV data as response
	c.Data(http.StatusOK, "text/csv", data)
}
