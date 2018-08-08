/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tesseract_test

import (
	"errors"
	"os"
	"testing"

	"github.com/it-chain/engine/icode"
	"github.com/it-chain/engine/icode/infra/tesseract"
	"github.com/stretchr/testify/assert"
)

func setContainer(t *testing.T) (*tesseract.ContainerService, func()) {
	GOPATH := os.Getenv("GOPATH")

	if GOPATH == "" {
		t.Fatal(errors.New("need go path"))
		return nil, func() {}
	}

	containerService := tesseract.NewContainerService()

	meta := icode.Meta{
		ICodeID:        "1",
		RepositoryName: "test icode",
		Path:           GOPATH + "/src/github.com/it-chain/engine/icode/mock/",
		GitUrl:         "github.com/mock",
	}

	err := containerService.StartContainer(meta)
	assert.NoError(t, err)

	return containerService, func() {
		err := containerService.StopContainer(meta.ICodeID)
		assert.NoError(t, err)
	}
}

func TestNewTesseractContainerService(t *testing.T) {
	_, tearDown := setContainer(t)
	defer tearDown()
}

func TestTesseractContainerService_ExecuteRequest(t *testing.T) {
	cs, tearDown := setContainer(t)
	defer tearDown()

	// success case
	result, err := cs.ExecuteRequest(icode.Request{
		ICodeID:  "1",
		Function: "initA",
		Type:     "invoke",
		Args:     []string{},
	})
	assert.NoError(t, err)
	assert.Equal(t, result.Err, "")

	// success case
	result, err = cs.ExecuteRequest(icode.Request{
		ICodeID:  "1",
		Function: "getA",
		Type:     "query",
		Args:     []string{},
	})

	assert.NoError(t, err)
	assert.Equal(t, result.Data["A"], "0")
	assert.Equal(t, result.Err, "")

	// no corresponding icode id
	result, err = cs.ExecuteRequest(icode.Request{
		ICodeID:  "2",
		Function: "initA",
		Type:     "invoke",
		Args:     []string{},
	})
	assert.Equal(t, err, tesseract.ErrContainerDoesNotExist)

	// invalid type
	result, err = cs.ExecuteRequest(icode.Request{
		ICodeID:  "1",
		Function: "initA",
		Type:     "invoke2",
		Args:     []string{},
	})

	assert.NotEqual(t, result.Err, "")

	// invalid func
	result, err = cs.ExecuteRequest(icode.Request{
		ICodeID:  "1",
		Function: "initAb",
		Type:     "invoke",
		Args:     []string{},
	})
	assert.NotEqual(t, result.Err, "")
}
