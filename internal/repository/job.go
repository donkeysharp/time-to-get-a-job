package repository

import "github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/models"

type JobPostRepository struct{}

func NewJobPostRepository() *JobPostRepository {
	return &JobPostRepository{}
}

func (me *JobPostRepository) Get(id int) (*models.JobPost, error) {
	return nil, nil
}

func (me *JobPostRepository) GetAll() ([]*models.JobPost, error) {
	return nil, nil
}

func (me *JobPostRepository) Create(item *models.JobPost) error {
	return nil
}

func (me *JobPostRepository) Update(item *models.JobPost) error {
	return nil
}

func (me *JobPostRepository) Delete(item *models.JobPost) error {
	return nil
}

func (me *JobPostRepository) Apply(userId int, jobPost *models.JobPost) error {
	return nil
}

func (me *JobPostRepository) GetApplication(applicationId int) (*models.JobPostApplication, error) {
	return nil, nil
}

func (me *JobPostRepository) GetApplications(userId int) ([]*models.JobPostApplication, error) {
	return nil, nil
}
