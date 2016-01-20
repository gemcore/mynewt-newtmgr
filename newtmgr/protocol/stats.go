/*
 Copyright 2015 Runtime Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package protocol

import (
	"encoding/json"
	"fmt"

	"git-wip-us.apache.org/repos/asf/incubator-mynewt-newt/util"
)

const (
	STATS_NMGR_OP_READ = 0
)

const (
	NMGR_GROUP_ID_STATS = 2
)

type StatsReadReq struct {
	Name string `json:"n"`
}

type StatsReadRsp struct {
	Name   string                 `json:"n"`
	Fields map[string]interface{} `json:"f"`
}

func NewStatsReadReq() (*StatsReadReq, error) {
	s := &StatsReadReq{}
	s.Name = ""

	return s, nil
}

func (sr *StatsReadReq) EncodeWriteRequest() (*NmgrReq, error) {
	nmr, err := NewNmgrReq()
	if err != nil {
		return nil, err
	}

	nmr.Op = NMGR_OP_READ
	nmr.Flags = 0
	nmr.Group = NMGR_GROUP_ID_STATS
	nmr.Id = STATS_NMGR_OP_READ

	srr := &StatsReadReq{
		Name: sr.Name,
	}

	data, _ := json.Marshal(srr)
	nmr.Data = data
	nmr.Len = uint16(len(data))

	return nmr, nil
}

func DecodeStatsReadResponse(data []byte) (*StatsReadRsp, error) {
	var sr StatsReadRsp
	err := json.Unmarshal(data, &sr)
	if err != nil {
		return nil, util.NewNewtError(fmt.Sprintf("Invalid incoming json: %s",
			err.Error()))
	}

	return &sr, nil
}
