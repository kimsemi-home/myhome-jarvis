package externalbootstrap

import "strings"

func childPublicSafetyPattern() string {
	oldOwner := strings.Join([]string{"kim", "jooyoon"}, "")
	oldTeam := strings.Join([]string{"kim-joo", "-yoon"}, "")
	localPath := "/" + "Users" + "/"
	tokenPrefix := strings.Join([]string{
		"github_" + "pat_",
		"gh" + "p_",
		"gh" + "o_",
		"gh" + "u_",
		"gh" + "s_",
	}, "|")
	secretShapes := strings.Join([]string{
		"BEGIN (RSA|OPENSSH|DSA|EC|PRIVATE) KEY",
		"AKIA[0-9A-Z]{16}",
		"xox[baprs]-",
		"sk-[A-Za-z0-9]{20,}",
	}, "|")
	return strings.Join([]string{oldOwner, oldTeam, localPath, tokenPrefix, secretShapes}, "|")
}
