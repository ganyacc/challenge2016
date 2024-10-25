package models

type Distributor struct {
	Name        string
	Parent      *Distributor
	Permissions []Permission
}

func NewDistributor(name string, parent *Distributor) *Distributor {
	return &Distributor{
		Name:        name,
		Parent:      parent,
		Permissions: make([]Permission, 0),
	}
}

func (d *Distributor) AddPermission(permType, region string) {
	d.Permissions = append(d.Permissions, NewPermission(permType, region))
}

func (d *Distributor) HasPermission(region string) bool {
	// Check parent permissions first
	if d.Parent != nil && !d.Parent.HasPermission(region) {
		return false
	}

	targetRegion := NewRegion(region)
	allowed := false

	// Process includes first
	for _, perm := range d.Permissions {
		if perm.Type == "INCLUDE" {
			permRegion := NewRegion(perm.Region)
			if permRegion.Contains(targetRegion) {
				allowed = true
				break
			}
		}
	}

	// If not included by this distributor but has parent, check parent's includes
	if !allowed && d.Parent != nil {
		allowed = d.Parent.HasPermission(region)
	}

	// Process excludes
	if allowed {
		for _, perm := range d.Permissions {
			if perm.Type == "EXCLUDE" {
				permRegion := NewRegion(perm.Region)
				if permRegion.Contains(targetRegion) {
					return false
				}
			}
		}
	}

	return allowed
}
