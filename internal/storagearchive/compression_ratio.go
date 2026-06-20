package storagearchive

func compressionRatioPercent(inputBytes int64, outputBytes int64) int {
	if inputBytes <= 0 || outputBytes <= 0 {
		return 0
	}
	return int((outputBytes*100 + inputBytes - 1) / inputBytes)
}
