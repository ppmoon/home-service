package bolt

import bolt "go.etcd.io/bbolt"

func GetBoltDB(dbPath string) (*bolt.DB,error) {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	//
	defer db.Close()
	return db,nil
}
