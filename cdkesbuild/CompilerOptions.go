package cdkesbuild


type CompilerOptions struct {
	ImportsNotUsedAsValues *string `field:"optional" json:"importsNotUsedAsValues" yaml:"importsNotUsedAsValues"`
	JsxFactory *string `field:"optional" json:"jsxFactory" yaml:"jsxFactory"`
	JsxFragmentFactory *string `field:"optional" json:"jsxFragmentFactory" yaml:"jsxFragmentFactory"`
	PreserveValueImports *bool `field:"optional" json:"preserveValueImports" yaml:"preserveValueImports"`
	UseDefineForClassFields *bool `field:"optional" json:"useDefineForClassFields" yaml:"useDefineForClassFields"`
}

