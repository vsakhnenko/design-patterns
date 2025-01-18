package main

import (
	"fmt"
	"streamer"
)

func main() {
	//number of workers
	const numJobs = 4
	const numWorkers = 4

	//create channels
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	wp := streamer.New(videoQueue, numWorkers)

	//start worker pool
	wp.Run()

	//mp4
	video1 := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	//should fail
	video2 := wp.NewVideo(2, "./input/bad.txt", "./output", "mp4", notifyChan, nil)

	//hls
	ops := &streamer.VideoOptions{
		RenameOutput:    true,
		SegmentDuration: 10,
		MaxRate1080p:    "1200k",
		MaxRate720p:     "600k",
		MaxRate480p:     "400k",
	}
	video3 := wp.NewVideo(3, "./input/puppy2.mp4", "./output", "hls", notifyChan, ops)

	//mp4
	video4 := wp.NewVideo(4, "./input/puppy2.mp4", "./output", "mp4", notifyChan, nil)

	//send the videos to worker pool
	videoQueue <- streamer.VideoProcessingJob{Video: video1}
	videoQueue <- streamer.VideoProcessingJob{Video: video2}
	videoQueue <- streamer.VideoProcessingJob{Video: video3}
	videoQueue <- streamer.VideoProcessingJob{Video: video4}

	//print out results
	for i := 1; i <= numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i:", i, "msg:", msg)
	}
	fmt.Println("Done!")
}
