package factorymethod

type CarFactory struct {
}

func (cf *CarFactory) GetVeiculo(name string) Veiculos {
	return &Car{
		Name: name,
	}
}
