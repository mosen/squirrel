package munki

// ManifestStore is the interface for accessing manifests in a database or filesystem
type ManifestStore interface {
	AllManifests() (*ManifestCollection, error)
	Manifest(name string) (*Manifest, error)
	NewManifest(name string) (*Manifest, error)
	SaveManifest(path string, manifest *Manifest) error
	DeleteManifest(name string) error
}

// Manifest represents the structure of a munki manifest
// This is what would be serialized in a datastore
type Manifest struct {
	Filename          string      `plist:"-" json:"-"`
	Catalogs          []string    `plist:"catalogs,omitempty" json:"catalogs,omitempty"`
	DisplayName       string      `plist:"display_name,omitempty" json:"display_name,omitempty"`
	IncludedManifests []string    `plist:"included_manifests,omitempty" json:"included_manifests,omitempty"`
	Notes             string      `plist:"notes,omitempty" json:"notes,omitempty"`
	User              string      `plist:"user,omitempty" json:"user,omitempty"`
	ConditionalItems  []condition `plist:"conditional_items,omitempty" json:"conditional_items,omitempty"`
	manifestItems
}

type manifestItems struct {
	OptionalInstalls  []string `plist:"optional_installs,omitempty" json:"optional_installs,omitempty"`
	ManagedInstalls   []string `plist:"managed_installs,omitempty" json:"managed_installs,omitempty"`
	ManagedUninstalls []string `plist:"managed_uninstalls,omitempty" json:"managed_uninstalls,omitempty"`
	ManagedUpdates    []string `plist:"managed_updates,omitempty" json:"managed_updates,omitempty"`
}

type condition struct {
	Condition string `plist:"condition" json:"condition"`
	manifestItems
}

// ManifestCollection represents a list of manifests
type ManifestCollection []*Manifest

// UpdateFromPayload updates a manifest from a ManifestPayload
func (m *Manifest) UpdateFromPayload(payload *ManifestPayload) {
	if payload.Catalogs != nil {
		m.Catalogs = *payload.Catalogs
	}

	if payload.DisplayName != nil {
		m.DisplayName = *payload.DisplayName
	}

	if payload.IncludedManifests != nil {
		m.IncludedManifests = *payload.IncludedManifests
	}

	if payload.OptionalInstalls != nil {
		m.OptionalInstalls = *payload.OptionalInstalls
	}

	if payload.ManagedInstalls != nil {
		m.ManagedInstalls = *payload.ManagedInstalls
	}

	if payload.ManagedUninstalls != nil {
		m.ManagedUninstalls = *payload.ManagedUninstalls
	}

	if payload.ManagedUpdates != nil {
		m.ManagedUpdates = *payload.ManagedUpdates
	}

	if payload.Notes != nil {
		m.Notes = *payload.Notes
	}

	if payload.User != nil {
		m.User = *payload.User
	}

	if payload.ConditionalItems != nil {
		m.ConditionalItems = *payload.ConditionalItems
	}
}
