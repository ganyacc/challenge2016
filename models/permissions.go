package models

type Permission struct {
	Type   string
	Region string
}

func NewPermission(permType, region string) Permission {
	return Permission{
		Type:   permType,
		Region: region,
	}
}
