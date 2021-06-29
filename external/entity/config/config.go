package config

type YAMLConfig struct {
	Databse struct {
		TableName string `yaml:"tableName"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
	} `yaml:"databse"`
}