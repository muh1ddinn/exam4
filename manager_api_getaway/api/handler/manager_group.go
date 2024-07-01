package handler

import (
	"context"
	"fmt"
	"managergetaway/api/models"
	"managergetaway/genproto/manager_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Creategroup godoc
// @Security ApiKeyAuth
// @Router      /groups [POST]
// @Summary     Create a group
// @Description Create a new group
// @Tags        groups
// @Accept      json
// @Produce     json
// @Param       group body manager_service.CreateGroupRequest true "group"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateGroup(c *gin.Context) {
	var req manager_service.CreateGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.GroupServiceC().CreateGroup(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating handler")
		return
	}
	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "Created successfully")

}

// GetgroupByID godoc
// @Security ApiKeyAuth
// @Router      /groups/{id} [GET]
// @Summary     Get a group by its ID
// @Description Retrieves a group by its ID
// @Tags        groups
// @Accept      json
// @Produce     json
// @Param       id path string true "group ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetGroupByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.GroupServiceC().GetGroupByID(context.Background(), &manager_service.GroupID{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}

	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "GETBYID successfully")

}

// Createjournal godoc
// @Security ApiKeyAuth
// @Router      /journal [POST]
// @Summary     Create a journal
// @Description Create a new journal
// @Tags        journal
// @Accept      json
// @Produce     json
// @Param       group body manager_service.CreateJournalRequest true "journal"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateJournal(c *gin.Context) {
	var req manager_service.CreateJournalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.GroupServiceC().CreateJournal(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating handler")
		return
	}
	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "Created successfully")

}

// GetgroupByID godoc
// @Security ApiKeyAuth
// @Router      /journal/{id} [GET]
// @Summary     Get a journal by its ID
// @Description Retrieves a journal by its ID
// @Tags        journal
// @Accept      json
// @Produce     json
// @Param       id path string true "journal ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetJournalByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.GroupServiceC().GetGroupByID(context.Background(), &manager_service.GroupID{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}

	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "GETBYID successfully")

}

// Createjournal godoc
// @Security ApiKeyAuth
// @Router      /task [POST]
// @Summary     Create a task
// @Description Create a new task
// @Tags        task
// @Accept      json
// @Produce     json
// @Param       group body manager_service.CreateTaskRequest true "task"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateTask(c *gin.Context) {
	var req manager_service.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.GroupServiceC().CreateTask(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating handler")
		return
	}
	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "Created successfully")

}

// GetgroupByID godoc
// @Security ApiKeyAuth
// @Router      /task/{id} [GET]
// @Summary     Get a task by its ID
// @Description Retrieves a task by its ID
// @Tags        task
// @Accept      json
// @Produce     json
// @Param       id path string true "task ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.GroupServiceC().GetTaskByID(context.Background(), &manager_service.TaskID{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}

	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "GETBYID successfully")

}

// Createjournal godoc
// @Security ApiKeyAuth
// @Router      /lesson [POST]
// @Summary     Create a lesson
// @Description Create a new lesson
// @Tags        lesson
// @Accept      json
// @Produce     json
// @Param       group body manager_service.CreateLessonRequest true "talessonsk"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateLesson(c *gin.Context) {
	var req manager_service.CreateLessonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.GroupServiceC().CreateLesson(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating handler")
		return
	}
	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "Created successfully")

}

// GetgroupByID godoc
// @Security ApiKeyAuth
// @Router      /lesson/{id} [GET]
// @Summary     Get a lesson by its ID
// @Description Retrieves a lesson by its ID
// @Tags        lesson
// @Accept      json
// @Produce     json
// @Param       id path string true "lesson ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetLessonByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.GroupServiceC().GetLessonByID(context.Background(), &manager_service.LessonID{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}

	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "GETBYID successfully")

}

// Createjournal godoc
// @Security ApiKeyAuth
// @Router      /type [POST]
// @Summary     Create a type
// @Description Create a new type
// @Tags        type
// @Accept      json
// @Produce     json
// @Param       group body manager_service.CreateTypeRequest true "type "
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateType(c *gin.Context) {
	var req manager_service.CreateTypeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.GroupServiceC().CreateType(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating handler")
		return
	}
	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "Created successfully")

}

// GetgroupByID godoc
// @Security ApiKeyAuth
// @Router      /type/{id} [GET]
// @Summary     Get a type by its ID
// @Description Retrieves a type by its ID
// @Tags        type
// @Accept      json
// @Produce     json
// @Param       id path string true "type ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetTypeByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.GroupServiceC().GetTypeByID(context.Background(), &manager_service.TypeID{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}

	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "GETBYID successfully")

}

// Createjournal godoc
// @Security ApiKeyAuth
// @Router      /schedule [POST]
// @Summary     Create a schedule
// @Description Create a new schedule
// @Tags        schedule
// @Accept      json
// @Produce     json
// @Param       group body manager_service.CreateSchedulesRequest true "schedule "
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) CreateSchedules(c *gin.Context) {
	var req manager_service.CreateSchedulesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while binding body")
		return
	}

	resp, err := h.grpcClient.GroupServiceC().CreateSchedules(context.Background(), &req)
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating handler")
		return
	}
	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "Created successfully")

}

// GetgroupByID godoc
// @Security ApiKeyAuth
// @Router      /schedule/{id} [GET]
// @Summary     Get a schedule by its ID
// @Description Retrieves a schedule by its ID
// @Tags        schedule
// @Accept      json
// @Produce     json
// @Param       id path string true "schedule ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetSchedulesByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.GroupServiceC().GetSchedulesByID(context.Background(), &manager_service.SchedulesID{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}

	fmt.Println(resp)
	handleGrpcErrWithDescription(c, h.log, err, "GETBYID successfully")

}

// GetgroupByID godoc
// @Security ApiKeyAuth
// @Router      /forteacher/{id} [GET]
// @Summary     Get a forteacher by its ID
// @Description Retrieves a forteacher by its ID
// @Tags        forteacher
// @Accept      json
// @Produce     json
// @Param       id path string true "forteacher ID"
// @Success     200 {object} manager_service.TeacherGetTeacherInfoResponse
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetTeacherInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	resp, err := h.grpcClient.GroupServiceC().GetTeacherInfo(context.Background(), &manager_service.TeacherGetTeacherInfoRequest{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while creating")
		return
	}
	// Assuming resp is the correct response object that matches models.ResponseSuccess
	c.JSON(http.StatusOK, resp)
}

// GetSupporTeacherInfo godoc
// @Security ApiKeyAuth
// @Router      /forsupportteacher/{id} [GET]
// @Summary     Get a forsupportteacher by its ID
// @Description Retrieves a forsupportteacher by its ID
// @Tags        forsupportteacher
// @Accept      json
// @Produce     json
// @Param       id path string true "forsupportteacher ID"
// @Success     200 {object} models.ResponseSuccess
// @Failure     400 {object} models.ResponseError
// @Failure     404 {object} models.ResponseError
// @Failure     500 {object} models.ResponseError
func (h *handler) GetSupporTeacherInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{Error: "ID is required"})
		return
	}

	resp, err := h.grpcClient.GroupServiceC().GetSupporTeacherInfo(context.Background(), &manager_service.TeacherGetTeacherInfoRequest{Id: id})
	if err != nil {
		handleGrpcErrWithDescription(c, h.log, err, "error while retrieving support teacher info")
		return
	}

	// Assuming resp is the correct response object that matches models.ResponseSuccess
	c.JSON(http.StatusOK, resp)
}
