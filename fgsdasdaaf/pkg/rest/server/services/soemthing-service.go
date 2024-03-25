package services

import (
	"github.com/pavan-intelops/testinglol12/fgsdasdaaf/pkg/rest/server/daos"
	"github.com/pavan-intelops/testinglol12/fgsdasdaaf/pkg/rest/server/models"
)

type SoemthingService struct {
	soemthingDao *daos.SoemthingDao
}

func NewSoemthingService() (*SoemthingService, error) {
	soemthingDao, err := daos.NewSoemthingDao()
	if err != nil {
		return nil, err
	}
	return &SoemthingService{
		soemthingDao: soemthingDao,
	}, nil
}

func (soemthingService *SoemthingService) UpdateSoemthing(id int64, soemthing *models.Soemthing) (*models.Soemthing, error) {
	return soemthingService.soemthingDao.UpdateSoemthing(id, soemthing)
}

func (soemthingService *SoemthingService) DeleteSoemthing(id int64) error {
	return soemthingService.soemthingDao.DeleteSoemthing(id)
}
