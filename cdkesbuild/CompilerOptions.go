package cdkesbuild


type CompilerOptions struct {
	AlwaysStrict *bool `field:"optional" json:"alwaysStrict" yaml:"alwaysStrict"`
	BaseUrl *bool `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	ExperimentalDecorators *bool `field:"optional" json:"experimentalDecorators" yaml:"experimentalDecorators"`
	ImportsNotUsedAsValues *string `field:"optional" json:"importsNotUsedAsValues" yaml:"importsNotUsedAsValues"`
	Jsx *string `field:"optional" json:"jsx" yaml:"jsx"`
	JsxFactory *string `field:"optional" json:"jsxFactory" yaml:"jsxFactory"`
	JsxFragmentFactory *string `field:"optional" json:"jsxFragmentFactory" yaml:"jsxFragmentFactory"`
	JsxImportSource *string `field:"optional" json:"jsxImportSource" yaml:"jsxImportSource"`
	Paths *map[string]*[]*string `field:"optional" json:"paths" yaml:"paths"`
	PreserveValueImports *bool `field:"optional" json:"preserveValueImports" yaml:"preserveValueImports"`
	Strict *bool `field:"optional" json:"strict" yaml:"strict"`
	Target *string `field:"optional" json:"target" yaml:"target"`
	UseDefineForClassFields *bool `field:"optional" json:"useDefineForClassFields" yaml:"useDefineForClassFields"`
	VerbatimModuleSyntax *bool `field:"optional" json:"verbatimModuleSyntax" yaml:"verbatimModuleSyntax"`
}

