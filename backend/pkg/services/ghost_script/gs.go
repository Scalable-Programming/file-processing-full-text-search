package gs

import (
	"os/exec"
	"strings"
)

func GenerateThumbnail(path string) (string, error) {
	splitPath := strings.SplitAfter(path, ".")
	outputFile := strings.Join(splitPath[:len(splitPath)-1], "") + "thumb.jpg"

	gsPathExecutable, err := exec.LookPath("gs")

	if err != nil {
		return "", err
	}

	args := []string{"-o", outputFile, "-sDEVICE=jpeg", "-dDEVICEHEIGHT=720", "-sPageList=1", "-dPDFFitPage", path}

	cmd := exec.Command(gsPathExecutable, args...)

	err = cmd.Run()

	if err != nil {
		return "", err
	}

	return outputFile, nil

}
