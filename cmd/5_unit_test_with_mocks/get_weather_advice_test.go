package main

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"

	"github.com/olezhek28/code-base/cmd/5_unit_test_with_mocks/mocks"
)

func TestGetWeatherAdvice(t *testing.T) {
	type weatherCenterClientMockFunc func(t *testing.T) WeatherCenterClient

	var (
		city = gofakeit.City()
	)

	tests := []struct {
		name                    string
		city                    string
		temperature             float32
		expected                string
		weatherCenterClientMock weatherCenterClientMockFunc
	}{
		{
			name:     "Температура +25 градусов",
			city:     city,
			expected: "Отличная погода для прогулок",
			weatherCenterClientMock: func(t *testing.T) WeatherCenterClient {
				mockClient := mocks.NewWeatherCenterClient(t)
				mockClient.On("GetTemperature", city).Return(float32(25), nil)

				return mockClient
			},
		},
		{
			name:     "Температура -15 градусов",
			city:     city,
			expected: "Прохладно, но можно выйти на улицу",
			weatherCenterClientMock: func(t *testing.T) WeatherCenterClient {
				mockClient := mocks.NewWeatherCenterClient(t)
				mockClient.On("GetTemperature", city).Return(float32(-15), nil)

				return mockClient
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := getWeatherAdvice(test.weatherCenterClientMock(t), test.city)
			require.NoError(t, err)
			require.Equal(t, test.expected, res)
		})
	}
}
