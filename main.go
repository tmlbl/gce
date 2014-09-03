package main

import (
	jwt "code.google.com/p/goauth2/oauth/jwt"
	compute "code.google.com/p/google-api-go-client/compute/v1"
	"fmt"
)

func main() {
	fmt.Println("Started the application")
	service, err := NewComputeService()
	if err != nil {
		panic(err)
	}
	disk_list := GetDisks()
	fmt.Printf("Found %d disks\n", len(disk_list.Items))
	for i, disk := range disk_list.Items {
		fmt.Printf("%d: %s\n", i+1, disk.Name)
	}
	disk := disk_list.Items[0]
	fmt.Printf("Using %s as boot disk\n", disk.Name)
	attDisk := AttachBootDisk(disk)
	netInterface := GetNetworkInterface()
	instance := compute.Instance{
		Name:              "myresource",
		Kind:              "coreos-alpha",
		Zone:              "us-central1-a",
		MachineType:       "https://www.googleapis.com/compute/v1/projects/gokubernetes/zones/us-central1-a/machineTypes/f1-micro",
		Disks:             []*compute.AttachedDisk{attDisk},
		NetworkInterfaces: []*compute.NetworkInterface{netInterface},
	}
	create_call := service.Instances.Insert(
		"gokubernetes",
		"us-central1-a",
		&instance,
	)
	res, err := create_call.Do()
	if err != nil {
		fmt.Println("Failed to create instance")
		panic(err)
	}
	fmt.Println(res)
	call := service.Instances.List("gokubernetes", "us-central1-a")
	list, err := call.Do()
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}

func NewComputeService() (*compute.Service, error) {
	iss := credentials.ISS
	scope := "https://www.googleapis.com/auth/compute"
	key := []byte(credentials.PrivateKey)
	token := jwt.NewToken(iss, scope, key)
	if transport, err := jwt.NewTransport(token); err == nil {
		service, err := compute.New(transport.Client())
		return service, err
	} else {
		return &compute.Service{}, err
	}
}

func GetDisks() *compute.DiskList {
	client, _ := NewComputeService()
	call := client.Disks.List("gokubernetes", "us-central1-a")
	res, err := call.Do()
	if err != nil {
		panic(err)
	}
	return res
}

func AttachBootDisk(disk *compute.Disk) *compute.AttachedDisk {
	return &compute.AttachedDisk{
		Source: disk.SelfLink,
		Type:   disk.Type,
		Index:  int64(disk.Id),
		Boot:   true,
	}
}

func GetNetworkInterface() *compute.NetworkInterface {
	client, _ := NewComputeService()
	call := client.Networks.List("gokubernetes")
	res, err := call.Do()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found %d networks\n", len(res.Items))
	for i, network := range res.Items {
		fmt.Printf("%d: %s\n", i+1, network.SelfLink)
	}
	network := res.Items[0]
	return &compute.NetworkInterface{
		Network: network.SelfLink,
	}
}

func GetMachineType() *compute.MachineTypeList {
	client, _ := NewComputeService()
	call := client.MachineTypes.List("gokubernetes", "us-central1-a")
	res, err := call.Do()
	if err != nil {
		panic(err)
	}
	return res
}
