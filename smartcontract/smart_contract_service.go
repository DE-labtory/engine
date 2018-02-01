package smartcontract

import (
	"errors"
	"it-chain/service/blockchain"
	"strings"
	"fmt"
	"os"
	"time"
)

const (
	GITHUB_TOKEN string = "20fc6f01c0cab8f97ac3e8433a946c8b869702cd"

)

type SmartContract struct {
	ReposName string
	OriginReposPath string
	SmartContractPath string
}

type SmartContractService struct {
	GithubID string
	SmartContractDirPath string
	SmartContractMap map[string]SmartContract
}

func Init() {
}

func (scs *SmartContractService) Deploy(ReposPath string) (string, error) {
	origin_repos_name := strings.Split(ReposPath, "/")[1]
	new_repos_name := strings.Replace(ReposPath, "/", "_", -1)

	_, ok := scs.keyByValue(ReposPath)
	if ok {
		// 버전 업데이트 기능 추가 필요
		return "", errors.New("Already exist smart contract ID")
	}

	repos, err := GetRepos(ReposPath)
	if err != nil {
		return "", errors.New("An error occured while getting repos!")
	}
	if repos.Message == "Bad credentials" {
		return "", errors.New("Not Exist Repos!")
	}

	err = os.MkdirAll(scs.SmartContractDirPath + "/" + new_repos_name, 0755)
	if err != nil {
		return "", errors.New("An error occured while make repository's directory!")
	}

	err = CloneRepos(ReposPath, scs.SmartContractDirPath + "/" + new_repos_name)
	if err != nil {
		return "", errors.New("An error occured while cloning repos!")
	}

	_, err = CreateRepos(new_repos_name, GITHUB_TOKEN)
	if err != nil {
		return "", errors.New(err.Error())//"An error occured while creating repos!")
	}

	err = ChangeRemote(scs.GithubID + "/" + new_repos_name, scs.SmartContractDirPath + "/" + new_repos_name + "/" + origin_repos_name)
	if err != nil {
		return "", errors.New("An error occured while cloning repos!")
	}

	// 버전 관리를 위한 파일 추가
	now := time.Now().Format("2006-01-02 15:04:05");
	file, err := os.OpenFile(scs.SmartContractDirPath + "/" + new_repos_name + "/" + origin_repos_name + "/version", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return "", errors.New("An error occured while creating or opening file!")
	}

	_, err = file.WriteString("Deployed at " + now + "\n")
	if err != nil {
		return "", errors.New("An error occured while writing file!")
	}
	err = file.Close()
	if err != nil {
		return "", errors.New("An error occured while closing file!")
	}

	err = CommitAndPush(scs.SmartContractDirPath + "/" + new_repos_name + "/" + origin_repos_name, "It-Chain Smart Contract \"" + new_repos_name + "\" Deploy")
	if err != nil {
		return "", errors.New(err.Error())
		//return "", errors.New("An error occured while committing and pushing!")
	}

	githubResponseCommits, err := GetReposCommits(scs.GithubID + "/" + new_repos_name)
	if err != nil {
		return "", errors.New("An error occured while getting commit log!")
	}


	reposDirPath := scs.SmartContractDirPath + "/" + new_repos_name + "/" + githubResponseCommits[0].Sha
	err = os.Rename(scs.SmartContractDirPath + "/" + new_repos_name + "/" + origin_repos_name, reposDirPath)
	if err != nil {
		return "", errors.New("An error occured while renaming directory!")
	}

	scs.SmartContractMap[githubResponseCommits[0].Sha] = SmartContract{new_repos_name, ReposPath, ""}

	return githubResponseCommits[0].Sha, nil
}
/***************************************************
 *	1. smartcontract 검사
 *	2. smartcontract -> sc.tar : 애초에 풀 받을 때 압축해 둘 수 있음
 *	3. go 버전에 맞는 docker image를 Create
 *	4. sc.tar를 docker container로 복사
 *	5. docker container Start
 *	6. docker에서 smartcontract 실행
 ****************************************************/
func (scs *SmartContractService) Query(transaction blockchain.Transaction) (error) {
	tmpDir := "/tmp"
	sc, ok := scs.SmartContractMap[transaction.TxData.ContractID];
	if !ok {
		return errors.New("Not exist contract ID")
	}

	_, err := os.Stat(sc.SmartContractPath)
	if os.IsNotExist(err) {
		fmt.Println("File or Directory Not Exist")
		return errors.New("File or Directory Not Exist")
	}

	err = MakeTar(sc.SmartContractPath, tmpDir)
	if err != nil {
		return errors.New("An error occured while archiving file!")
	}

	PullAndCopyAndRunDocker("docker.io/library/node", tmpDir+"/"+transaction.TxData.ContractID+".tar")

	return nil
}


func (scs *SmartContractService) Invoke() {

}

func (scs *SmartContractService) keyByValue(OriginReposPath string) (key string, ok bool) {
	contractName := strings.Replace(OriginReposPath, "/", "^", -1)
	for k, v := range scs.SmartContractMap {
		if contractName == v.OriginReposPath {
			key = k
			ok = true
			return key, ok
		}
	}
	return "", false
}