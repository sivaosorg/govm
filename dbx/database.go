package dbx

import "github.com/sivaosorg/govm/utils"

func NewDbx() *Dbx {
	d := &Dbx{}
	return d
}

func (d *Dbx) SetConnected(value bool) *Dbx {
	d.IsConnected = value
	return d
}

func (d *Dbx) SetError(err error) *Dbx {
	d.Error = err
	return d
}

func (d *Dbx) SetMessage(value string) *Dbx {
	d.Message = value
	return d
}

func (d *Dbx) SetDatabase(value string) *Dbx {
	d.Database = value
	return d
}

func (d *Dbx) SetDebugMode(value bool) *Dbx {
	d.DebugMode = value
	return d
}

func (d *Dbx) SetNewInstance(value bool) *Dbx {
	d.IsNewInstance = value
	return d
}

func (d *Dbx) SetPid(value int) *Dbx {
	d.Pid = value
	return d
}

func (d *Dbx) Json() string {
	return utils.ToJson(d)
}
