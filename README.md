channelMq
=========

Overview
--------
A light-weight, performant byte-slice MessageQueue implemented with go channels.

channelMq supports N-many subscribable message channels based on string subscription keys.
Clients subscribing with an empty set of keys can either receive none of the
messages on the system or all of them.
```go
    mq := channelMq.NewMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)
    // or
	mq.EmptyKeyGets(channelMq.NO_MESSAGES)
```

Example Code
------------
```go
import (
    "fmt"
	"github.com/JoeReid/channelMq"
)

func main() {
    // Lets create a messageQueue
    mq := channelMq.NewMQ()
	mq.EmptyKeyGets(channelMq.ALL_MESSAGES)

    // Now lets subscribe with an empty set of keys
	keys := []string{}
	ch := mq.Subscribe(keys)

    // Now lets send some messages
    mq.Send([]byte("Hello, World!"), []string{"someRandomKey", "someOtherKey"})

    // To receive a message, simply receive from the channel
    receivedMessage := <-ch

    fmt.Println(string(receivedMessage)) // Prints "Hello, World!"
}
```
