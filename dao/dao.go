package dao

func Init(){

}

func Exit(){
	GameInst.exit()
	UserInst.exit()
}

type exportModels struct {
	models []interface{}
}

func NewExportModels() *exportModels {
	em := &exportModels{
		make([]interface{}),
	}
	append(em.models, NewUser())
	append(em.models, NewGame())
	return em
}

func (em *exportModels)Exit(){
	for m in em.models {

	}
}