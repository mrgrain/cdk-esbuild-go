package cdkesbuild


type ProviderTransformOptions struct {
	Banner *string `field:"optional" json:"banner" yaml:"banner"`
	// Documentation: https://esbuild.github.io/api/#charset.
	Charset *string `field:"optional" json:"charset" yaml:"charset"`
	// Documentation: https://esbuild.github.io/api/#color.
	Color *bool `field:"optional" json:"color" yaml:"color"`
	// Documentation: https://esbuild.github.io/api/#define.
	Define *map[string]*string `field:"optional" json:"define" yaml:"define"`
	// Documentation: https://esbuild.github.io/api/#drop.
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
	Footer *string `field:"optional" json:"footer" yaml:"footer"`
	// Documentation: https://esbuild.github.io/api/#format.
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Documentation: https://esbuild.github.io/api/#global-name.
	GlobalName *string `field:"optional" json:"globalName" yaml:"globalName"`
	// Documentation: https://esbuild.github.io/api/#ignore-annotations.
	IgnoreAnnotations *bool `field:"optional" json:"ignoreAnnotations" yaml:"ignoreAnnotations"`
	// Documentation: https://esbuild.github.io/api/#jsx.
	Jsx *string `field:"optional" json:"jsx" yaml:"jsx"`
	// Documentation: https://esbuild.github.io/api/#jsx-development.
	JsxDev *bool `field:"optional" json:"jsxDev" yaml:"jsxDev"`
	// Documentation: https://esbuild.github.io/api/#jsx-factory.
	JsxFactory *string `field:"optional" json:"jsxFactory" yaml:"jsxFactory"`
	// Documentation: https://esbuild.github.io/api/#jsx-fragment.
	JsxFragment *string `field:"optional" json:"jsxFragment" yaml:"jsxFragment"`
	// Documentation: https://esbuild.github.io/api/#jsx-import-source.
	JsxImportSource *string `field:"optional" json:"jsxImportSource" yaml:"jsxImportSource"`
	// Documentation: https://esbuild.github.io/api/#jsx-side-effects.
	JsxSideEffects *bool `field:"optional" json:"jsxSideEffects" yaml:"jsxSideEffects"`
	// Documentation: https://esbuild.github.io/api/#keep-names.
	KeepNames *bool `field:"optional" json:"keepNames" yaml:"keepNames"`
	// Documentation: https://esbuild.github.io/api/#legal-comments.
	LegalComments *string `field:"optional" json:"legalComments" yaml:"legalComments"`
	Loader *string `field:"optional" json:"loader" yaml:"loader"`
	// Documentation: https://esbuild.github.io/api/#log-level.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Documentation: https://esbuild.github.io/api/#log-limit.
	LogLimit *float64 `field:"optional" json:"logLimit" yaml:"logLimit"`
	// Documentation: https://esbuild.github.io/api/#log-override.
	LogOverride *map[string]*string `field:"optional" json:"logOverride" yaml:"logOverride"`
	// Documentation: https://esbuild.github.io/api/#mangle-props.
	MangleCache *map[string]interface{} `field:"optional" json:"mangleCache" yaml:"mangleCache"`
	// Documentation: https://esbuild.github.io/api/#mangle-props.
	MangleProps interface{} `field:"optional" json:"mangleProps" yaml:"mangleProps"`
	// Documentation: https://esbuild.github.io/api/#mangle-props.
	MangleQuoted *bool `field:"optional" json:"mangleQuoted" yaml:"mangleQuoted"`
	// Documentation: https://esbuild.github.io/api/#minify.
	Minify *bool `field:"optional" json:"minify" yaml:"minify"`
	// Documentation: https://esbuild.github.io/api/#minify.
	MinifyIdentifiers *bool `field:"optional" json:"minifyIdentifiers" yaml:"minifyIdentifiers"`
	// Documentation: https://esbuild.github.io/api/#minify.
	MinifySyntax *bool `field:"optional" json:"minifySyntax" yaml:"minifySyntax"`
	// Documentation: https://esbuild.github.io/api/#minify.
	MinifyWhitespace *bool `field:"optional" json:"minifyWhitespace" yaml:"minifyWhitespace"`
	// Documentation: https://esbuild.github.io/api/#platform.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Documentation: https://esbuild.github.io/api/#pure.
	Pure *[]*string `field:"optional" json:"pure" yaml:"pure"`
	// Documentation: https://esbuild.github.io/api/#mangle-props.
	ReserveProps interface{} `field:"optional" json:"reserveProps" yaml:"reserveProps"`
	Sourcefile *string `field:"optional" json:"sourcefile" yaml:"sourcefile"`
	// Documentation: https://esbuild.github.io/api/#sourcemap.
	Sourcemap interface{} `field:"optional" json:"sourcemap" yaml:"sourcemap"`
	// Documentation: https://esbuild.github.io/api/#source-root.
	SourceRoot *string `field:"optional" json:"sourceRoot" yaml:"sourceRoot"`
	// Documentation: https://esbuild.github.io/api/#sources-content.
	SourcesContent *bool `field:"optional" json:"sourcesContent" yaml:"sourcesContent"`
	// Documentation: https://esbuild.github.io/api/#supported.
	Supported *map[string]*bool `field:"optional" json:"supported" yaml:"supported"`
	// Documentation: https://esbuild.github.io/api/#target.
	Target interface{} `field:"optional" json:"target" yaml:"target"`
	// Documentation: https://esbuild.github.io/api/#tree-shaking.
	TreeShaking *bool `field:"optional" json:"treeShaking" yaml:"treeShaking"`
	TsconfigRaw interface{} `field:"optional" json:"tsconfigRaw" yaml:"tsconfigRaw"`
}

