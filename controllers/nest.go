package controllers

type Nest struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ParentId   string `json:"parentId"`
	IsLeaf     bool   `json:"isLeaf"`
	Children   []Nest `json:"children"`
	DataSource []Nest `json:"-"`
}

func NewNest(dataSource []Nest) Nest {
	return Nest{DataSource: dataSource}
}

func (d Nest) SetChild(parentIds []string, targets []Nest) {

	for i, target := range targets {
		if ContainsString(parentIds, target.Id) && target.IsLeaf == false {
			targets[i].Children = d.GetByParentId(target.Id)
		}
		d.SetChild(parentIds, targets[i].Children)
	}

}

func (d Nest) GetByParentId(parentId string) []Nest {
	list := make([]Nest, 0)
	for _, dto := range d.DataSource {
		if parentId == dto.ParentId {
			list = append(list, dto)
		}
	}
	return list
}
