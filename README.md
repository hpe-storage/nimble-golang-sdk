# HPE Nimble Storage SDK for Go
This is the Go Software Development Kit (SDK) for [HPE Nimble Storage](http://hpe.com/storage/nimblestorage) arrays. The HPE Nimble Storage array has a Representational State Transfer (REST) web service application programming interface (API). The SDK implements a simple client Go module for communicating with the HPE Nimble Storage REST API. The Go module is being used to communicate with the API over HTTPS.

The SDK provides Go modules to interact with the HPE Nimble Storage REST API. The code abstracts the lower-level API calls into Go objects that you can easily incorporate into any automation or DevOps workflows. Use it to create, modify and delete most resources like volumes, volume collections, initiator groups and more, as well as perform other tasks like snapshotting, cloning, restoring data, etc.

# Synopsis

Examples are available in [examples](examples). This is a brief Go program that creates a 5GiB volume on the specified array.

```
package main

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/service"
)

func main() {

	// Create new group service
	groupService, err := service.NewNsGroupService(
 		"192.168.1.1",              // Managment hostname or IP of array
 		"admin",                    // Username
		"admin",                    // Password
		"v1",                       // REST API version
		true)                       // Perform operations synchronous

	if err != nil {
		fmt.Printf("Unable to connect to group, %+v\n", err.Error())
		return
	}

	// Logout when finished
	defer groupService.LogoutService()

	// Initialize volume attributes
	var sizeField int64 = 5120

	// Set mandatory volume attributes
	newVolume := &nimbleos.Volume{
		Name:              param.NewString("myvol1"),
		Size:              &sizeField,
	}

	// Assign volume service instance
	volSvc := groupService.GetVolumeService()

	// Create volume
	volume, err := volSvc.CreateVolume(newVolume)

	if err != nil {
		fmt.Printf("Failed to create volume, %+v\n", err.Error())
		return
	}
	// Volume created
	fmt.Printf("Volume \"%s\" created (%s)\n", *volume.Name, *volume.ID)
}
```

Service types may be found in [pkg/service](pkg/service) and it's recommended to have a basic understanding and familiarity of the [HPE Nimble Storage REST API](https://infosight.hpe.com/InfoSight/media/cms/active/public/pubs_REST_API_Reference_NOS_53x.whz).

# Requirements

The SDK require Go version 1.13 or later. The SDK is coupled to the following NimbleOS version span.

|          | Minimum version | Maximum version |
| -------- | --------------- | --------------- |
| NimbleOS | 5.0.10          | 5.3.1           |

As future versions of NimbleOS and new endpoint services become available, the SDK will be updated.

# Contributing

Contributing guidelines are available in [CONTRIBUTING](CONTRIBUTING.md)

# Community

Join the HPE DEV slack by signing up at [slack.hpedev.io](https://slack.hpedev.io). HPE employees may sign in direct at [hpedev.slack.com](https://hpedev.slack.com) with a valid HPE address. We hang out in `#NimbleStorage`.

# License

The HPE Nimble Storage SDK for Go is released under the Apache 2.0 license, see [LICENSE](LICENSE) for details.
