package endpoint

type IdValidator struct{}

func (*IdValidator) GenerateSnowflakeId() (any, error) {
	return nil, nil
}
