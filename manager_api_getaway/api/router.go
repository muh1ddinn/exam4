package api

import (
	_ "managergetaway/api/docs" //for swagger
	"managergetaway/api/handler"
	"managergetaway/config"
	"managergetaway/pkg/grpc_client"
	"managergetaway/pkg/logger"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})
	// Assuming "r" is your Gin router instance

	// For groups
	r.POST("/groups", handler.CreateGroup)
	r.GET("/groups/:id", handler.GetGroupByID)

	// For journal
	r.POST("/journal", handler.CreateJournal)
	r.GET("/journal/:id", handler.GetJournalByID)

	// For task
	r.POST("/task", handler.CreateTask)
	r.GET("/task/:id", handler.GetTaskByID)

	// For lesson
	r.POST("/lesson", handler.CreateLesson)
	r.GET("/lesson/:id", handler.GetLessonByID)

	// For type
	r.POST("/type", handler.CreateType)
	r.GET("/type/:id", handler.GetTypeByID)

	// For schedule
	r.POST("/schedule", handler.CreateSchedules)
	r.GET("/schedule/:id", handler.GetSchedulesByID)

	// For forteacher
	r.GET("/forteacher/:id", handler.GetTeacherInfo)

	// For forsupportteacher
	r.GET("/forsupportteacher/:id", handler.GetAllSupportTeachers)

	r.POST("/students", handler.CreateStudent)
	r.GET("/students", handler.GetAllStudents)
	r.GET("/students/:id", handler.GetStudentByID)
	r.PATCH("/students", handler.UpdateStudent)
	r.DELETE("/students/:id", handler.DeleteStudent)

	r.POST("/teachers", handler.CreateTeacher)
	r.GET("/teachers", handler.GetAllTeachers)
	r.GET("/teachers/:id", handler.GetTeacherByID)
	r.PATCH("/teachers", handler.UpdateTeacher)
	r.DELETE("/teachers/:id", handler.DeleteTeacher)

	r.POST("/supportteachers", handler.CreateSupportTeacher)
	r.GET("/supportteachers", handler.GetAllSupportTeachers)
	r.GET("/supportteachers/:id", handler.GetSupportTeacherByID)
	r.PATCH("/supportteachers", handler.UpdateSupportTeacher)
	r.DELETE("/supportteachers/:id", handler.DeleteSupportTeacher)
	////////////////

	// Register routes for students
	r.POST("/superadminstudents", handler.CreateadminStudent)
	r.GET("/superadminstudents", handler.GetAlladminStudents)
	r.GET("/superadminstudents/:id", handler.GetadminStudentByID)
	r.PATCH("/superadminstudents", handler.UpdateAdminStudent)
	r.DELETE("/superadminstudents/:id", handler.DeleteAdminStudent)

	// Register routes for teachers
	r.POST("/superadminteachers", handler.CreatedminTeacher)
	r.GET("/superadminteachers", handler.GetAllAdminTeachers)
	r.GET("/superadminteachers/:id", handler.GetAdminTeacherByID)
	r.PATCH("/superadminteachers", handler.UpdateAdminTeacher)
	r.DELETE("/superadminteachers/:id", handler.DeleteAdminTeacher)

	// Register routes for support teachers
	r.POST("/superadminsupportteachers", handler.CreateSupportAdminTeacher)
	r.GET("/superadminsupportteachers", handler.GetAllSupportAdminTeachers)
	r.GET("/superadminsupportteachers/:id", handler.GetSupportAdminTeacherByID)
	r.PATCH("/superadminsupportteachers", handler.UpdateSupportAdminTeacher)
	r.DELETE("/superadminsupportteachers/:id", handler.DeleteSupportAdminTeacher)

	// Register routes for administrators
	r.POST("/administrators", handler.CreateAdministrator)
	r.GET("/administrators", handler.GetAllAdministrators)
	r.GET("/administrators/:id", handler.GetAdministratorByID)
	r.DELETE("/administrators/:id", handler.DeleteAdministrator)

	// Register routes for managers
	r.POST("/managers", handler.CreateManager)
	r.GET("/managers", handler.GetAllManagers)
	r.GET("/managers/:id", handler.GetManagerByID)
	r.PATCH("/managers", handler.UpdateManager)
	r.DELETE("/managers/:id", handler.DeleteManager)
	r.GET("/any/export/csv", handler.ExportToCSV)

	// Shipper endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
