package incident_test

import (
	"testing"
	"time"

	"github.com/involvestecnologia/statuspage/component"

	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/involvestecnologia/statuspage/models"

	"github.com/stretchr/testify/assert"
)

func TestNewIncidentService(t *testing.T) {
	logger := mock.NewLogRepositoryMock()
	incidentDAO := mock.NewMockIncidentDAO()
	componentService := component.NewService(mock.NewMockComponentDAO(), logger)
	s := incident.NewService(incidentDAO, componentService, logger)
	assert.Implements(t, (*incident.Service)(nil), s)
}
func TestIncidentService_ListIncidents(t *testing.T) {
	s := incident.NewService(
		mock.NewMockIncidentDAO(),
		component.NewService(mock.NewMockComponentDAO(), mock.NewLogRepositoryMock()),
		mock.NewLogRepositoryMock(),
	)

	params := models.ListIncidentQueryParameters{
		EndDate:    "",
		StartDate:  "",
		Unresolved: true,
	}

	i, err := s.ListIncidents(params)
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		if assert.IsType(t, []models.Incident{}, i) {
			assert.Equal(t, mock.ZeroTimeHex, i[0].ComponentRef)
			assert.Equal(t, mock.OneSecTimeHex, i[len(i)-1].ComponentRef)
		}
	}

	incs, err := s.ListIncidents(params)
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		if assert.IsType(t, []models.Incident{}, i) {
			for _, inc := range incs {
				assert.False(t, inc.Resolved)
			}
		}
	}

	params.StartDate = time.Now().Add(-24 * time.Hour).Format(time.RFC3339)
	params.Unresolved = false
	// Valid StartDate
	monthOnly, err := s.ListIncidents(params)
	if assert.Nil(t, err) && assert.NotNil(t, monthOnly) {
		assert.Equal(t, 1, int(monthOnly[0].Date.Month()))
	}

	params.EndDate = time.Now().Add(-12 * time.Hour).Format(time.RFC3339)
	// Valid EndDate
	yearOnly, err := s.ListIncidents(params)
	if assert.Nil(t, err) && assert.NotNil(t, yearOnly) {
		assert.Equal(t, 1, yearOnly[0].Date.Year())
	}

	params.StartDate = time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	// Invalid StartDate
	_, err = s.ListIncidents(params)
	assert.NotNil(t, err)

	params.EndDate = time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	// Invalid EndDate
	_, err = s.ListIncidents(params)
	assert.NotNil(t, err)

	params.StartDate = time.Now().Add(-24 * time.Hour).Format(time.RFC3339)
	params.EndDate = time.Now().Add(-36 * time.Hour).Format(time.RFC3339)
	// Invalid EndDate
	_, err = s.ListIncidents(params)
	assert.NotNil(t, err)

	params.StartDate = time.Now().Add(-24 * time.Hour).Format(time.RFC1123)
	// Invalid StartDate format
	_, err = s.ListIncidents(params)
	assert.NotNil(t, err)

	params.StartDate = time.Now().Add(-24 * time.Hour).Format(time.RFC3339)
	params.EndDate = time.Now().Add(-12 * time.Hour).Format(time.RFC1123)
	// Invalid EndDate format
	_, err = s.ListIncidents(params)
	assert.NotNil(t, err)
}
func TestIncidentService_CreateIncident(t *testing.T) {
	s := incident.NewService(
		mock.NewMockIncidentDAO(),
		component.NewService(mock.NewMockComponentDAO(), mock.NewLogRepositoryMock()),
		mock.NewLogRepositoryMock(),
	)

	i := models.Incident{
		ComponentRef: mock.ZeroTimeHex,
		Status:       models.IncidentStatusOutage,
		Description:  "",
		Date:         time.Now(),
	}

	err := s.CreateIncidents(i)
	assert.Nil(t, err)

	inc, err := s.FindIncidents(map[string]interface{}{"component_ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		assert.Equal(t, i, inc[0])
	}

	i.Status = models.IncidentStatusUnstable
	err = s.CreateIncidents(i)
	assert.NotNil(t, err)

	i.ComponentRef = "Empty Component"
	err = s.CreateIncidents(i)
	assert.Nil(t, err)

	i.Status = models.IncidentStatusOK
	err = s.CreateIncidents(i)
	assert.Nil(t, err)

	i.Status = models.IncidentStatusUnstable
	err = s.CreateIncidents(i)
	assert.Nil(t, err)

	i.Status = models.IncidentStatusOutage
	err = s.CreateIncidents(i)
	assert.Nil(t, err)

	i.Status = 5
	err = s.CreateIncidents(i)
	assert.NotNil(t, err)

	i.Status = models.IncidentStatusOK

	i.ComponentRef = "Invalid Component Ref"
	err = s.CreateIncidents(i)
	assert.NotNil(t, err)
}
func TestIncidentService_FindIncidents(t *testing.T) {
	s := incident.NewService(
		mock.NewMockIncidentDAO(),
		component.NewService(mock.NewMockComponentDAO(), mock.NewLogRepositoryMock()),
		mock.NewLogRepositoryMock(),
	)

	i, err := s.FindIncidents(map[string]interface{}{"component_ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, i) {
		if assert.IsType(t, []models.Incident{}, i) {
			assert.Equal(t, i[0].Status, models.IncidentStatusOK)
			assert.Equal(t, i[len(i)-1].Status, models.IncidentStatusOK)
		}
	}

	_, err = s.FindIncidents(map[string]interface{}{"component_ref": "Invalid Component Ref"})
	assert.NotNil(t, err)

	_, err = s.FindIncidents(map[string]interface{}{"component_ref": bson.NewObjectId().Hex()})
	assert.NotNil(t, err)

}
func TestIncidentService_ValidateMonth(t *testing.T) {
	s := incident.NewService(
		mock.NewMockIncidentDAO(),
		component.NewService(mock.NewMockComponentDAO(), mock.NewLogRepositoryMock()),
		mock.NewLogRepositoryMock(),
	)

	start := time.Time{}
	end := time.Now()
	err := s.ValidateDate(start, end)
	assert.Nil(t, err)

	end, err = time.Parse(time.RFC3339, end.Add(24*time.Hour).Format(time.RFC3339))
	assert.Nil(t, err)
	err = s.ValidateDate(start, end)
	assert.Nil(t, err)

	end, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	assert.Nil(t, err)
	start, err = time.Parse(time.RFC3339, end.Add(2*time.Hour).Format(time.RFC3339))
	assert.Nil(t, err)
	err = s.ValidateDate(start, end)
	assert.NotNil(t, err)

	end, err = time.Parse(time.RFC3339, time.Now().Add(-6*time.Hour).Format(time.RFC3339))
	assert.Nil(t, err)
	start, err = time.Parse(time.RFC3339, end.Add(2*time.Hour).Format(time.RFC3339))
	assert.Nil(t, err)
	err = s.ValidateDate(start, end)
	assert.NotNil(t, err)

}
