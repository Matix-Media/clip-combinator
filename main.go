package main

import (
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	fmt.Println("Not implemented yet")
}

func combineVideo(video1 string, video2 string, output string, baseVideo uint8) error {
	//var videoStream1 *ffmpeg.Stream = ffmpeg.Input(video1)
	//var videoStream2 *ffmpeg.Stream = ffmpeg.Input(video2)
	var videoMain string
	var videoSec string
	var videoMainInfo string
	var videoSecInfo string

	if baseVideo == 1 {
		videoMain = video1
		videoSec = video2
	} else {
		videoMain = video2
		videoSec = video1
	}

	var err error
	videoMainInfo, err = ffmpeg.Probe(videoMain)
	if err != nil {
		fmt.Println("Error reading main video: " + err.Error())
		return err
	}
	videoSecInfo, err = ffmpeg.Probe(videoSec)
	if err != nil {
		fmt.Println("Error reading secondary video: " + err.Error())
		return err
	}

	fmt.Println("Main video: " + videoMainInfo)
	fmt.Println("Secondary video: " + videoSecInfo)

	return nil
}
