package middleware

func IsExtNameValide(extName string) bool {

	allowExtMap := map[string]bool{
		".log": true,
		".txt": true,
		".wsd": true,
		".xlm": true,
		".csv": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return true
	}

	return false
}
