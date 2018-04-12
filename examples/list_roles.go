//
// Copyright (c) 2017 Joey <majunjiev@gmail.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"time"

	ovirtsdk4 "gopkg.in/imjoey/go-ovirt.v4"
)

func listRoles() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("qwer1234").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %v\n", err)
		return
	}
	defer conn.Close()

	// To use `Must` methods, you should recover it if panics
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panics occurs, try the non-Must methods to find the reason")
		}
	}()

	roleService := conn.SystemService().RolesService()
	resp, err := roleService.List().Send()
	if err != nil {
		fmt.Printf("Failed to get role list, reason: %v\n", err)
		return
	}

	if roleSlice, ok := resp.Roles(); ok {
		for _, role := range roleSlice.Slice() {
			fmt.Printf("Role: (")
			if roleName, ok := role.Name(); ok {
				fmt.Printf(" name: %v", roleName)
			}
			if roleDesc, ok := role.Description(); ok {
				fmt.Printf(" desc: %v", roleDesc)
			}
			fmt.Println(")")
		}
	}

}
