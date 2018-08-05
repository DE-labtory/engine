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

package blockchain

import (
	"io/ioutil"
	"os"
	"time"

	"encoding/json"
)

func CreateGenesisBlock(genesisconfFilePath string) (DefaultBlock, error) {

	//declare
	GenesisBlock := &DefaultBlock{}
	validator := DefaultValidator{}

	//set basic
	err := setBlockWithConfig(genesisconfFilePath, GenesisBlock)

	if err != nil {
		return DefaultBlock{}, ErrSetConfig
	}

	//build
	Seal, err := validator.BuildSeal(GenesisBlock.Timestamp, GenesisBlock.PrevSeal, GenesisBlock.TxSeal, GenesisBlock.Creator)

	if err != nil {
		return DefaultBlock{}, ErrBuildingSeal
	}

	//set seal
	GenesisBlock.SetSeal(Seal)

	return *GenesisBlock, nil
}

func setBlockWithConfig(filePath string, block *DefaultBlock) error {

	// load
	jsonFile, err := os.Open(filePath)
	defer jsonFile.Close()

	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return err
	}

	GenesisConfig := &GenesisConfig{}

	err = json.Unmarshal(byteValue, GenesisConfig)
	if err != nil {
		return err
	}

	// set
	const longForm = "Jan 1, 2006 at 0:00am (MST)"

	timeStamp, err := time.Parse(longForm, GenesisConfig.TimeStamp)

	if err != nil {
		return err
	}

	block.SetPrevSeal(make([]byte, 0))
	block.SetHeight(uint64(GenesisConfig.Height))
	block.SetTxSeal(make([][]byte, 0))
	block.SetTimestamp(timeStamp)
	block.SetCreator([]byte(GenesisConfig.Creator))
	block.SetState(Created)

	return nil
}

type GenesisConfig struct {
	Organization string
	NedworkId    string
	Height       int
	TimeStamp    string
	Creator      string
}

func CreateProposedBlock(prevSeal []byte, height uint64, txList []*DefaultTransaction, Creator []byte) (DefaultBlock, error) {

	//declare
	ProposedBlock := &DefaultBlock{}
	validator := DefaultValidator{}
	TimeStamp := time.Now().Round(0)

	//build
	for _, tx := range txList {
		ProposedBlock.PutTx(tx)
	}

	txSeal, err := validator.BuildTxSeal(ConvertTxType(txList))

	if err != nil {
		return DefaultBlock{}, ErrBuildingTxSeal
	}

	Seal, err := validator.BuildSeal(TimeStamp, prevSeal, txSeal, Creator)

	if err != nil {
		return DefaultBlock{}, ErrBuildingSeal
	}

	//set
	ProposedBlock.SetSeal(Seal)
	ProposedBlock.SetPrevSeal(prevSeal)
	ProposedBlock.SetHeight(height)
	ProposedBlock.SetTxSeal(txSeal)
	ProposedBlock.SetTimestamp(TimeStamp)
	ProposedBlock.SetCreator(Creator)
	ProposedBlock.SetState(Created)

	return *ProposedBlock, nil
}
