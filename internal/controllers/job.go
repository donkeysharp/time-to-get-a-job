package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type JobPostController struct {
	ControllerShared
}

func NewJobPostController(name string) *JobPostController {
	return &JobPostController{
		ControllerShared: ControllerShared{Name: name},
	}
}

func (me *JobPostController) List(c echo.Context) error {
	return c.String(http.StatusOK, "ListJobPosts")
}

func (me *JobPostController) Get(c echo.Context) error {
	return c.String(http.StatusOK, "GetJobPost")
}

func (me *JobPostController) Create(c echo.Context) error {
	return c.String(http.StatusOK, "CreateJobPost")
}

func (me *JobPostController) Update(c echo.Context) error {
	return c.String(http.StatusOK, "UpdateJobPost")
}

func (me *JobPostController) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteJobPost")
}

func (me *JobPostController) Apply(c echo.Context) error {
	return c.String(http.StatusOK, "Apply")
}

func (me *JobPostController) ListApplications(c echo.Context) error {
	return c.String(http.StatusOK, "ListApplications")
}

func (me *JobPostController) GetApplication(c echo.Context) error {
	return c.String(http.StatusOK, "GetApplication")
}

func (me *JobPostController) RegisterRoutes(e *echo.Echo) {
	e.GET("/jobs/", me.List)
	e.GET("/jobs/:id", me.Get)
	e.POST("/jobs", me.Create)
	e.PUT("/jobs/:id", me.Update)
	e.DELETE("/jobs/:id", me.Delete)
	e.POST("/jobs/:id/apply", me.Apply)
	e.GET("/applications/", me.ListApplications)
	e.GET("/applications/:id", me.GetApplication)

}
