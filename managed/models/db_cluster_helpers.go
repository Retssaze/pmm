package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
	"strings"
)

type DBClusterFilters struct {
	Name                string
	KubernetesClusterID string
	ClusterType         DBClusterType
}

func FindDBClusters(q *reform.Querier, filters DBClusterFilters) ([]*DBCluster, error) {
	var args []interface{}
	var andConds []string
	idx := 1
	if filters.Name != "" {
		andConds = append(andConds, fmt.Sprintf("name = %s", q.Placeholder(idx)))
		args = append(args, filters.Name)
		idx += 1
	}
	if filters.KubernetesClusterID != "" {
		andConds = append(andConds, fmt.Sprintf("kubernetes_cluster_id = %s", q.Placeholder(idx)))
		args = append(args, filters.KubernetesClusterID)
		idx += 1
	}
	if filters.ClusterType != "" {
		andConds = append(andConds, fmt.Sprintf("cluster_type = %s", q.Placeholder(idx)))
		args = append(args, filters.ClusterType)
		idx += 1
	}

	var tail strings.Builder

	if len(andConds) != 0 {
		tail.WriteString("WHERE ")
		tail.WriteString(strings.Join(andConds, " AND "))
		tail.WriteRune(' ')
	}
	tail.WriteString("ORDER BY created_at DESC")

	structs, err := q.SelectAllFrom(DBClusterTable, tail.String(), args...)
	if err != nil {
		return nil, err
	}
	dbClusters := make([]*DBCluster, len(structs))
	for i, s := range structs {
		dbClusters[i] = s.(*DBCluster)
	}
	return dbClusters, nil
}

// DBClusterParams params for add/update db cluster.
type DBClusterParams struct {
	KubernetesClusterID string
	Name                string
	Exposed             bool
	InstalledImage      string
	PSMDBClusterParams  *PSMDBClusterParams
	PXCClusterParams    *PXCClusterParams
}

// CreateOrUpdateDBCluster creates DB Cluster with given type.
func CreateOrUpdateDBCluster(q *reform.Querier, dbClusterType DBClusterType, params *DBClusterParams) (*DBCluster, error) {
	_, err := FindKubernetesClusterByID(q, params.KubernetesClusterID)
	if err != nil {
		return nil, err
	}

	dbClusters, err := FindDBClusters(q, DBClusterFilters{
		Name:                params.Name,
		KubernetesClusterID: params.KubernetesClusterID,
		ClusterType:         dbClusterType,
	})
	if err != nil {
		return nil, err
	}

	row := &DBCluster{
		ClusterType:         dbClusterType,
		KubernetesClusterID: params.KubernetesClusterID,
		Name:                params.Name,
		Exposed:             params.Exposed,
		InstalledImage:      params.InstalledImage,
		PSMDBClusterParams:  params.PSMDBClusterParams,
		PXCClusterParams:    params.PXCClusterParams,
	}

	if len(dbClusters) == 0 {
		id := "/dbcluster_id/" + uuid.New().String()
		if err := checkUniqueDBClusterID(q, id); err != nil {
			return nil, err
		}

		row.ID = id

		if err := q.Insert(row); err != nil {
			return nil, errors.WithStack(err)
		}
	} else {
		row.ID = dbClusters[0].ID
		if err := q.Save(row); err != nil {
			return nil, err
		}
	}

	return row, nil
}

func checkUniqueDBClusterID(q *reform.Querier, id string) error {
	if id == "" {
		panic("empty DB Cluster ID")
	}

	agent := &DBCluster{ID: id}
	switch err := q.Reload(agent); err {
	case nil:
		return status.Errorf(codes.AlreadyExists, "DB Cluster with ID %q already exists.", id)
	case reform.ErrNoRows:
		return nil
	default:
		return errors.WithStack(err)
	}
}