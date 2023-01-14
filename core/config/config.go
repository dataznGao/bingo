package config

type Configuration struct {
	InputPath   string         `yaml:"inputPath"`
	OutputPath  string         `yaml:"outputPath"`
	FaultPoints []*FaultConfig `yaml:"faultPoints"`
}
type LocationPatternP struct {
	PackageP  *FaultPatternP `yaml:"packageP"`
	StructP   *FaultPatternP `yaml:"structP"`
	MethodP   *FaultPatternP `yaml:"methodP"`
	VariableP *FaultPatternP `yaml:"variableP"`
}
type FaultConfig struct {
	FaultType        string              `yaml:"faultType"`
	LocationPatterns []*LocationPatternP `yaml:"locationPattern"`
	TargetValue      interface{}         `yaml:"targetValue"`
}
type FaultPatternP struct {
	Name           string `yaml:"name"`
	ActivationRate string `yaml:"activationRate"`
}

var Config *Configuration
