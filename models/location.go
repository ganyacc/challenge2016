package models

type Location struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}

type LocationDB struct {
	Locations map[string]*Location
}

func NewLocationDB() *LocationDB {
	return &LocationDB{
		Locations: make(map[string]*Location),
	}
}

func (db *LocationDB) AddLocation(cityCode, provinceCode, countryCode, cityName, provinceName, countryName string) {
	// Create composite keys for different levels
	countryKey := countryCode
	provinceKey := provinceCode + "-" + countryCode
	cityKey := cityCode + "-" + provinceCode + "-" + countryCode

	// Add all levels to the locations map
	db.Locations[countryKey] = &Location{
		CountryCode: countryCode,
		CountryName: countryName,
	}

	db.Locations[provinceKey] = &Location{
		ProvinceCode: provinceCode,
		CountryCode:  countryCode,
		ProvinceName: provinceName,
		CountryName:  countryName,
	}

	db.Locations[cityKey] = &Location{
		CityCode:     cityCode,
		ProvinceCode: provinceCode,
		CountryCode:  countryCode,
		CityName:     cityName,
		ProvinceName: provinceName,
		CountryName:  countryName,
	}
}

func (db *LocationDB) IsValidRegion(code string) bool {
	_, exists := db.Locations[code]
	return exists
}

// Global instance
var GlobalLocationDB = NewLocationDB()
