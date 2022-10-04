package factorymethod

type AbstractVeiculosFactory interface {
	GetVeiculo(name string) Veiculos
}
