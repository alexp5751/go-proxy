package handler

import "testing"

func TestTimeMachineWeather(t *testing.T) {
	var w Weather
	dsr, err := w.getTimeMachineWeather("60", "60", 1507325139)
	if err != nil {
		t.Error(err)
	}
	if dsr.Latitude != 60 || dsr.Longitude != 60 {
		t.Errorf("Latitude was %f, should have been 60.", dsr.Latitude)
		t.Errorf("Longitude was %f, should have been 60.", dsr.Longitude)
	}
}
func TestAsyncGetAllWeather(t *testing.T) {
	var w Weather
	dsrs, err := w.getAllWeather("60", "60")
	if err != nil {
		t.Error(err)
	}
	if len(dsrs) != 7 {
		t.Errorf("Should have length of 7. Length was %d.", len(dsrs))
	}
	if dsrs[0].Latitude != 60 || dsrs[5].Longitude != 60 {
		t.Errorf("Latitude was %f, should have been 60.", dsrs[0].Latitude)
		t.Errorf("Longitude was %f, should have been 60.", dsrs[5].Longitude)
	}
	if dsrs[0].TimeZone != "Asia/Yekaterinburg" {
		t.Errorf("TimeZone should have been 'Asia/Yekaterinburg', but was %s.", dsrs[0].TimeZone)
	}
}
