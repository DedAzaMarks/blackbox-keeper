package configuration

type Config struct {
	Apps map[string]struct {
		Command     string `yaml:"command"`
		HealthCheck struct {
			Http struct {
				Path                string `yaml:"path"`
				WaitAfterStartMilli int64  `yaml:"wait-after-start"`
				RepeatAfterMilli    int64  `yaml:"repeat-after"`
				TimeoutMilli        int64  `yaml:"timeout"`
			} `yaml:"http"`
			StopAction struct {
				Signal struct {
					SignalType   string `yaml:"signal-type"`
					TimeoutMilli int64  `yaml:"timeout"`
				} `yaml:"signal"`
			} `yaml:"stop-action"`
		} `yaml:"health-check"`
	} `yaml:"apps"`
}
