package cmlclient

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
)

// {
// 	"boot_disk_size": 0,
// 	"compute_id": "9c2519bf-dda6-4d31-942e-8068a6349b5e",
// 	"configuration": "bridge0",
// 	"cpu_limit": 100,
// 	"cpus": 0,
// 	"data_volume": 0,
// 	"hide_links": false,
// 	"id": "9efb1503-7e2a-4d2a-959e-865209f1acc0",
// 	"image_definition": null,
// 	"lab_id": "52d5c824-e10c-450a-b9c5-b700bd3bc17a",
// 	"label": "ext-conn-0",
// 	"node_definition": "external_connector",
// 	"ram": 0,
// 	"tags": [],
// 	"vnc_key": "",
// 	"x": 317,
// 	"y": 341,
// 	"config_filename": "noname",
// 	"config_mediatype": "ISO",
// 	"config_image_path": "/var/local/virl2/images/52d5c824-e10c-450a-b9c5-b700bd3bc17a/9efb1503-7e2a-4d2a-959e-865209f1acc0/config.img",
// 	"cpu_model": null,
// 	"data_image_path": "/var/local/virl2/images/52d5c824-e10c-450a-b9c5-b700bd3bc17a/9efb1503-7e2a-4d2a-959e-865209f1acc0/data.img",
// 	"disk_image": null,
// 	"disk_image_2": null,
// 	"disk_image_3": null,
// 	"disk_image_path": null,
// 	"disk_image_path_2": null,
// 	"disk_image_path_3": null,
// 	"disk_driver": null,
// 	"driver_id": "external_connector",
// 	"efi_boot": false,
// 	"image_dir": "/var/local/virl2/images/52d5c824-e10c-450a-b9c5-b700bd3bc17a/9efb1503-7e2a-4d2a-959e-865209f1acc0",
// 	"libvirt_image_dir": "/var/lib/libvirt/images/virl-base-images",
// 	"nic_driver": null,
// 	"number_of_serial_devices": 0,
// 	"serial_devices": [],
// 	"video_memory": 0,
// 	"video_model": null,
// 	"state": "BOOTED",
// 	"boot_progress": "Booted"
//   }

type Node struct {
	ID              string       `json:"id"`
	Label           string       `json:"label"`
	X               int          `json:"x"`
	Y               int          `json:"y"`
	NodeDefinition  string       `json:"node_definition"`
	ImageDefinition string       `json:"image_definition"`
	Configuration   string       `json:"configuration"`
	CPUs            int          `json:"cpus"`
	RAM             int          `json:"ram"`
	State           string       `json:"state"`
	DataVolume      int          `json:"data_volume"`
	Interfaces      interfaceMap `json:"interfaces"`

	// extra
	lab *Lab
}

func (nmap nodeMap) MarshalJSON() ([]byte, error) {
	nodeList := []*Node{}
	for _, node := range nmap {
		nodeList = append(nodeList, node)
	}
	// we want this as a stable sort by node UUID
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].ID < nodeList[j].ID
	})

	return json.Marshal(nodeList)
}

func (c *Client) getNodesForLab(ctx context.Context, lab *Lab) error {
	api := fmt.Sprintf("labs/%s/nodes", lab.ID)

	nodeIDlist := &IDlist{}
	err := c.jsonGet(ctx, api, nodeIDlist)
	if err != nil {
		return err
	}

	nodeMap := make(nodeMap)
	for _, nodeID := range *nodeIDlist {
		api = fmt.Sprintf("labs/%s/nodes/%s", lab.ID, nodeID)
		node := &Node{lab: lab}
		err := c.jsonGet(ctx, api, node)
		if err != nil {
			return err
		}
		nodeMap[nodeID] = node
	}
	lab.Nodes = nodeMap
	return nil
}
