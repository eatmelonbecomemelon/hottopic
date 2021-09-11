package mongoplus

import (
	cm "common"
	cf "config"
	"encoding/json"
	"fmt"
	"mgo"
	"mgo/bson"
	"strings"
	"time"
)

type MgoDB struct {
	Session *mgo.Session
}

var MgoEntry MgoDB

func MongoInit() MgoDB {
	var mongoPara []string
	for _, value := range cf.SysConfig.MongoDB {
		one := fmt.Sprintf("%s:%s",
			value.IP,
			value.Port,
		)
		mongoPara = append(mongoPara, one)
	}
	mongoUser := cf.SysConfig.MongoDB["1"].User
	mongoPasswd := cf.SysConfig.MongoDB["1"].PassWord
	mongoPoolLimit := cf.SysConfig.MongoDB["1"].MongoPoolLimit
	if mongoPoolLimit <= 0 {
		mongoPoolLimit = 128
	}
	MgoEntry.MgoConn(mongoPara, mongoUser, mongoPasswd, mongoPoolLimit)
	return MgoEntry
}

func (m *MgoDB) MgoConn(addrs []string, user, passwd string, mongoPoolLimit int) {

	dialInfo := &mgo.DialInfo{
		Addrs:    addrs,
		Username: user,
		Password: passwd,
		Timeout:  5 * time.Second,
	}
	for {
		session, err := mgo.DialWithInfo(dialInfo)
		if err != nil {
			cm.Warningf("mgo failed:", err, dialInfo)
			time.Sleep(1 * time.Second)
			continue
		}
		session.SetPoolLimit(mongoPoolLimit)
		//session.SetMode(mgo.Monotonic, true)

		m.Session = session
		cm.Info("mgo conn ok")
		break

	}
	return
}

func CreateIndex(coll string, indexes []string) (ret int) {
	var newSession = MgoEntry.Session.Copy()
	defer newSession.Close()

	var c = newSession.DB(cf.SysConfig.MongoDB["1"].DBName).C(coll)

	if len(indexes) == 0 {
		ret = -1
		return
	}
	var err error
	err = c.EnsureIndexKey(indexes...)
	if err != nil {
		cm.Error("create index err", coll, err)
		ret = -1
		return
	}

	return
}

func InsertRecords(reportData interface{}, reportCol string) (ret int) {
	MgoEntry.InsertReport(reportCol, reportData)
	return
}

func (m *MgoDB) InsertReport(reportCol string, reportData interface{}) int {
	newSession := m.Session.Copy()
	defer newSession.Close()
	var reportData2 []interface{}
	temp, _ := json.Marshal(reportData)
	err := json.Unmarshal(temp, &reportData2)
	if err != nil {
		cm.Error(err)
	}
	fmt.Println("start InsertReport", len(reportData2))
	c := newSession.DB(cf.SysConfig.MongoDB["1"].DBName).C(reportCol)
	err = c.Insert(reportData2...)
	if nil != err {
		cm.Error(err)
		return cm.ErrInnerFault
	}
	fmt.Println("end InsertReport", len(reportData2))
	return cm.Success
}

func (m *MgoDB) QueryOne(dataCol string, filter bson.M) (result map[string]interface{}, err error) {
	newSession := m.Session.Copy()
	defer newSession.Close()
	c := newSession.DB(cf.SysConfig.MongoDB["1"].DBName).C(dataCol)
	err = c.Find(filter).One(&result)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		} else {
			return
		}
	}
	return
}

func (m *MgoDB) Query(dataCol string, filter bson.M, sort string, limit int) (result []map[string]interface{}, err error) {
	newSession := m.Session.Copy()
	defer newSession.Close()
	c := newSession.DB(cf.SysConfig.MongoDB["1"].DBName).C(dataCol)
	err = c.Find(filter).Sort(sort).Limit(limit).All(&result)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		} else {
			return
		}
	}
	return
}

func (m *MgoDB) Upsert(dataCol string, filter bson.M, data interface{}) (err error) {
	newSession := m.Session.Copy()
	defer newSession.Close()
	c := newSession.DB(cf.SysConfig.MongoDB["1"].DBName).C(dataCol)
	_, err = c.Upsert(filter, data)
	if err != nil {
		return
	}
	return
}
