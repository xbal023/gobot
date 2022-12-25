package lib

import (
	"os"
	"log"
	"os/exec"
	"gobot/utils/helper"
)

func ImgToStick(buffer []byte) []byte {
	image := helper.GenerateID() + ".png"
	err := os.WriteFile(image, buffer, 0600)
	if err != nil {
		log.Fatal("Failed to save image: %v", err)
	}
	
}