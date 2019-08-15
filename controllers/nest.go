package controllers

type Nest struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ParentId   string `json:"parentId"`
	IsLeaf     bool   `json:"isLeaf"`
	Level      int    `json:"level"`
	Children   []Nest `json:"children"`
	DataSource []Nest `json:"-"`
}

func NewNest(dataSource []Nest) *Nest {
	return &Nest{DataSource: dataSource}
}

func (d *Nest) WithFilter(level int) *Nest {
	if level == 0 {
		return d
	}
	for k := len(d.DataSource) - 1; k >= 0; k-- {
		if d.DataSource[k].Level == level {
			d.DataSource = append(d.DataSource[:k], d.DataSource[k+1:]...)
		}
	}
	return d
}
func (d *Nest) WithObject(nest Nest) *Nest {
	d.DataSource = append(d.DataSource, nest)
	return d
}

func (d *Nest) CallSetChild(parentIds []string, targets []Nest) {

	for i, target := range targets {
		if ContainsString(parentIds, target.Id) && target.IsLeaf == false {
			targets[i].Children = d.GetByParentId(target.Id)
		}
		d.CallSetChild(parentIds, targets[i].Children)
	}

}

func (d *Nest) GetByParentId(parentId string) []Nest {
	list := make([]Nest, 0)
	for _, dto := range d.DataSource {
		if parentId == dto.ParentId {
			list = append(list, dto)
		}
	}
	return list
}

func (d Nest) GetById(id string) Nest {
	for _, dto := range d.DataSource {
		if id == dto.Id {
			return dto
		}
	}
	return Nest{}
}
