package cover

import "testing"

func Test_1(t *testing.T) {
	videoPath := "./1.mp4"
	imgPath := "./1.png"
	VideoSaveToImg(videoPath, imgPath, 100)
}
