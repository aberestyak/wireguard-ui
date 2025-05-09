package jsondb

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/aberestyak/wireguard-ui/model"
	"github.com/aberestyak/wireguard-ui/util"
)

func (o *JsonDB) GetWakeOnLanHosts() ([]model.WakeOnLanHost, error) {
	var hosts []model.WakeOnLanHost

	// read all client json file in "hosts" directory
	records, err := o.conn.ReadAll(model.WakeOnLanHostCollectionName)
	if err != nil {
		return hosts, err
	}

	// build the ClientData list
	for _, f := range records {
		host := model.WakeOnLanHost{}

		// get client info
		if err := json.Unmarshal(f, &host); err != nil {
			return hosts, fmt.Errorf("cannot decode client json structure: %v", err)
		}

		// create the list of hosts and their qrcode data
		hosts = append(hosts, host)
	}

	return hosts, nil
}

func (o *JsonDB) GetWakeOnLanHost(macAddress string) (*model.WakeOnLanHost, error) {
	host := &model.WakeOnLanHost{
		MacAddress: macAddress,
	}
	resourceName, err := host.ResolveResourceName()
	if err != nil {
		return nil, err
	}

	err = o.conn.Read(model.WakeOnLanHostCollectionName, resourceName, host)
	if err != nil {
		host = nil
	}
	return host, err
}

func (o *JsonDB) DeleteWakeOnHostLanHost(macAddress string) error {
	host := &model.WakeOnLanHost{
		MacAddress: macAddress,
	}
	resourceName, err := host.ResolveResourceName()
	if err != nil {
		return err
	}

	return o.conn.Delete(model.WakeOnLanHostCollectionName, resourceName)
}

func (o *JsonDB) SaveWakeOnLanHost(host model.WakeOnLanHost) error {
	resourceName, err := host.ResolveResourceName()
	if err != nil {
		return err
	}

	wakeOnLanHostPath := path.Join(path.Join(o.dbPath, model.WakeOnLanHostCollectionName), resourceName+".json")
	output := o.conn.Write(model.WakeOnLanHostCollectionName, resourceName, host)
	err = util.ManagePerms(wakeOnLanHostPath)
	if err != nil {
		return err
	}

	return output
}

func (o *JsonDB) DeleteWakeOnHost(host model.WakeOnLanHost) error {
	resourceName, err := host.ResolveResourceName()
	if err != nil {
		return err
	}

	return o.conn.Delete(model.WakeOnLanHostCollectionName, resourceName)
}
