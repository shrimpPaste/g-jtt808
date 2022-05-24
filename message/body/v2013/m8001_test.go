package v2013

import "testing"

func TestM0100_CityID(t *testing.T) {
	var m M0100
	m.SetCityID(1234)
	if m.CityID() != 1234 {
		t.Fatalf("设置市县域ID错误，应为%d，实际为%d", 1234, m.CityID())
	}
}

func TestM0100_TermID(t *testing.T) {
	var m M0100
	m.SetTermID("A1B2C3")
	if m.TermID() != "A1B2C3" {
		t.Fatalf("设置终端ID错误，应为%s，实际为%s", "A1B2C3", m.TermID())
	}
}

func TestM0100_TermModel(t *testing.T) {
	var m M0100
	m.SetTermModel("A001")
	if m.TermModel() != "A001" {
		t.Fatalf("设置终端型号错误，应为%s，实际为%s", "A001", m.TermModel())
	}
}

func TestM0100_ManufacturerID(t *testing.T) {
	var m M0100
	m.SetManufacturerID("12345")
	if m.ManufacturerID() != "12345" {
		t.Fatalf("设置制造商ID错误，应为%s，实际为%s", "12345", m.ManufacturerID())
	}
}

func TestM0100_ProvincialID(t *testing.T) {
	var m M0100
	m.SetProvincialID(44)
	if m.ProvincialID() != 44 {
		t.Fatalf("设置省域ID错误，应为%d，实际为%d", 44, m.ProvincialID())
	}
}

func TestM0100_LicensePlateColor(t *testing.T) {
	var m M0100
	m.SetLicensePlateColor(200)
	if m.LicensePlateColor() != 200 {
		t.Fatalf("设置车牌颜色错误，应为%d，实际为%d", 200, m.LicensePlateColor())
	}
}

func TestM0100_LicensePlate(t *testing.T) {
	var m M0100
	m.SetLicensePlate("粤T88888")
	if m.LicensePlate() != "粤T88888" {
		t.Fatalf("设置车牌标识错误，应为%s，实际为%s", "粤T88888", m.LicensePlate())
	}
}
