package main

import (
	"context"
)

var ctx = context.Background()

func main() {
	// sig := make(chan os.Signal, 1)
	// defer close(sig)
	// signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// // get a post
	// // TODO: add error checking
	// postId, _, _ := reddit.DefaultClient().Subreddit.getPosts(context.Background(), "AskReddit", &reddit.ListOptions{Limit: 1})

	// comments, errs, stop := reddit.DefaultClient().Stream(postId, reddit.StreamInterval(time.Second*3), reddit.StreamDiscardInitial)
	// defer stop()

	// timer := time.NewTimer(time.Minute)
	// defer timer.Stop()

	// for {
	// 	select {
	// 	case comment, ok := <-comments:
	// 		if !ok {
	// 			return
	// 		}
	// 		fmt.Printf("Received comment: %s\n", comment.Body)
	// 	case err, ok := <-errs:
	// 		if !ok {
	// 			return
	// 		}
	// 		fmt.Fprintf(os.Stderr, "Error! %v\n", err)
	// 	case rcvSig, ok := <-sig:
	// 		if !ok {
	// 			return
	// 		}
	// 		fmt.Printf("Stopping due to %s signal.\n", rcvSig)
	// 		return
	// 	case <-timer.C:
	// 		return
	// 	}
	// }
}
