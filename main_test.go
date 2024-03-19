package main;

import "fmt";
import "testing";

func TestCombineVideoDifferentCodecs(t *testing.T) {
	fmt.println("TestCombineVideoDifferentCodecs");

	var video1 string = "test_files/v1.mkv";
	var video2 string = "test_files/v2.mp4";
	var output string = "test_files/output.mp4";
	var baseVideo uint8 = 1;

	var err = combineVideo(video1, video2, output, baseVideo);
	if err != nil {
		t.Errorf("Error combining videos: %v", err);
	}
}