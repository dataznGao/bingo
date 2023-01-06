package constant

import "errors"

func NewNoFaultTypeError(message string) error {
	return errors.New("[NoFaultTypeError] this fault type is not exist, please check your fault point, fault type: " + message)
}

func NewLocationPatternFormatError(message string) error {
	return errors.New("[LocationPatternFormatError] this location pattern is wrong, please check your location pattern, location patter: " + message)
}

func NewProbabilityError(message string) error {
	return errors.New("[ProbabilityError] this probability is wrong, please check your probability, make sure it is a fractional number and <= 1, location patter: " + message)
}
