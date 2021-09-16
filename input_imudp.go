package main

import (
	"encoding/json"
	"fmt"
)

type inputIMUDP struct {
	Name     string `json:"name"`
	Recvmmsg int64  `json:"called.recvmmsg"`
	Recvmsg  int64  `json:"called.recvmsg"`
	Received int64  `json:"msgs.received"`
}

func newInputIMUDPFromJSON(b []byte) (*inputIMUDP, error) {
	var pstat inputIMUDP
	err := json.Unmarshal(b, &pstat)
	if err != nil {
		return nil, fmt.Errorf("error decoding input stat `%v`: %v", string(b), err)
	}
	return &pstat, nil
}

func (i *inputIMUDP) toPoints() []*point {
	points := make([]*point, 3)

	points[0] = &point{
		Name:        "input_called_recvmmsg",
		Type:        counter,
		Value:       i.Recvmmsg,
		Description: "Number of recvmmsg called",
		LabelName:   "worker",
		LabelValue:  i.Name,
	}
	points[1] = &point{
		Name:        "input_called_recvmsg",
		Type:        counter,
		Value:       i.Recvmsg,
		Description: "Number of recvmmsg called",
		LabelName:   "worker",
		LabelValue:  i.Name,
	}

	points[2] = &point{
		Name:        "input_received",
		Type:        counter,
		Value:       i.Received,
		Description: "messages received",
		LabelName:   "worker",
		LabelValue:  i.Name,
	}

	return points
}
