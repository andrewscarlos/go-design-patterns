package factorymethod

import "fmt"

// Inteção oficial

// Fabricas são simplesmente operaçoes que criam objetos.

/* definir uma interface para criar um objecto, mas deixar as bsubclasses decidirem que classe instanciar.
O factory Method permite adiar a instanciação para as subclasses. */

// oculta a logica de instanciação do codigo cliente. o metodo fabrica sera responsavel por instanciar as classes desejadas
// é obtido atraves de herança. o metodo fabrica pode ser criado ou sobrescrito por subclasses
// da flexibilidade ao codigo cliente permitindo a criação de novas factories sem a necessidade a alterar codigo ja escrito (OCP)
// pode usar parametros para determinar o tipo dos objetos a serem criados ou os parametros a serem enviados aos objetos sendo criados

type Veiculos interface {
	PegarCliente(cliente string)
	Stop()
}

type Car struct {
	Name string
}

func NewCar(name string) Car {
	return Car{
		Name: name,
	}
}

func (c *Car) PegarCliente(cliente string) {
	fmt.Println(c.Name + " esta buscando o " + cliente)
}

func (c *Car) Stop() {
	fmt.Println("o carro parou")
}

type Moto struct {
	Name string
}

func NewMoto(name string) Moto {
	return Moto{
		Name: name,
	}
}

func (c *Moto) PegarCliente(cliente string) {
	fmt.Println(c.Name + " esta buscando o " + cliente)
}

func (c *Moto) Stop() {
	fmt.Println("a moto parou")
}
