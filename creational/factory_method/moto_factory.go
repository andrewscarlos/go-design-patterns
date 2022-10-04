package factorymethod

type MotoFactory struct {
}

func (mf *MotoFactory) GetVeiculo(name string) Veiculos {
	return &Moto{
		Name: name,
	}
}
