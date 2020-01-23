package main

var inputBucketPath string
var destBucketPath string

func getFileURL(fileName string) string {
	return ""
}

func getFileFromBucket(filename string) string {
	return getFileURL("toBeTranscodedVideos/" + filename)
}

type toTranscode struct {
	endTime   *int
	startTime *int
	url       string
	srcPath   string
	destPath  string
}

func createFileToTranscode(startTime *int, endTime *int) *toTranscode {
	var fileURL = getFileURL("fileName.fileType")
	return &toTranscode{
		endTime:   endTime,
		startTime: startTime,
		url:       fileURL,
		srcPath:   inputBucketPath,
		destPath:  destBucketPath,
	}
}

func transcodeFile(file *toTranscode) error {
	// Do the transcoding here
	// Move the file to the destination bucket
	return nil
}

func workerProcess(file *toTranscode) {
	err := transcodeFile(file)
	if err != nil {
		// print error
	}
}

func main() {

}
