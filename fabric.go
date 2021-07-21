package logger

func CreatePrintLogTarget(levels []string, categories []string) *Target {
	var target Target

	target = &PrintLogTarget{
		BaseLogTarget{
			Levels: levels, Categories: categories,
		},
	}

	return &target
}

func CreateFileLogTarget(FilePath string, levels []string, categories []string) *Target {
	var target Target

	target = &FileLogTarget{
		BaseLogTarget{
			Levels: levels, Categories: categories,
		},
		FilePath,
	}

	return &target
}
