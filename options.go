package bulpherjs

// ApplicationOptions is a structure which holds application's options.
type ApplicationOptions struct {
	// Bulma is a Bulma CSS framework specific options.
	Bulma *BulmaOptions
	// Name is an application name as shown in title.
	Name string
}

// BulmaOptions is a structure that controls everything related to Bulma
// CSS framework.
type BulmaOptions struct {
	// Version is a version of bulma to use. Warning: incorrectly set
	// version might break your application.
	Version string
}
