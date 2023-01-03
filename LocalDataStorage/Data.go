package localdatastorage

import (
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/ShivamTripathi21/Go/internal/model"
	"github.com/google/uuid"
)

var ClusterSet map[uuid.UUID]*model.ClusterInfo

func ClusterSize() int {
	num := len(ClusterSet)
	return num
}

func addCluster(newCluster *model.ClusterInfo) (bool, string) {
	_, isExist := ClusterSet[newCluster.ClusterID]
	if isExist {
		return false, "Cluster ID already exist"
	} else {
		ClusterSet[newCluster.ClusterID] = newCluster
		return true, "Cluster is added"
	}
}

func removeCluster(guid uuid.UUID) (bool, string) {
	_, isExist := ClusterSet[guid]

	if !isExist {
		return false, "Cluster Not found"
	} else {
		delete(ClusterSet, guid)
		return true, "Cluster is deleted"
	}
}

func getMyCluster(guid uuid.UUID) (*model.ClusterInfo, bool) {
	cluster, isExist := ClusterSet[guid]
	if isExist {
		return cluster, true
	} else {
		return nil, false
	}
}

func createDefaultCluster(isForTesting bool) (*model.ClusterInfo, error) {
	GID, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	} else {

		//cli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			panic(err)
			return nil, err
		}
		//configs, err := cli.Info(context.Background())
		if err != nil {
			panic(err)
			return nil, err
		}
		//get cluster type
		cluster := model.ClusterInfo{
			ClusterID:                     GID,
			ClusterName:                   "Cluster-1",
			ClusterType:                   &model.ClusterType{},
			ClusterROMSize:                0,
			ClusterRAMSize:                0,
			ClusterGPUSize:                0,
			ClusterSwapMemorySize:         0,
			IsGPUAvilable:                 false,
			ClusterAvilableROMSize:        0,
			ClusterAvilableRAMSize:        0,
			ClusterAvilableGPUSize:        0,
			ClusterAvilableSwapMemorySize: 0,
			TotalNumberOfPies:             0,
			ClusterCreatedByID:            guid.GUID{},
			ClusterCreateTime:             time.Now(),
			ClusterResourceProvider:       "",
			ClusterResourceURL:            "",
			IsLocalCluster:                false,
			IsClusterUP:                   false,
			IsClusterInUse:                false,
			IsBackupClusterEnabled:        false,
			BackupClusterID:               guid.GUID{},
		}
		return &cluster, nil
	}
}
