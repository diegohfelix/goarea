package goarea // Por padrão o nome do package deve ser o mesmo da pasta

/*
O diretório /home/diegohfelix/go, que é onde está este arquivo, faz parte do workspace do GOPATH
Sendo assim poderemos utilizar este pacote em outros pacotes/projetos
O arquivo compilado (.a) referente a este arquivo (.go) é gerado na pasta pkg do diretório /home/diegohfelix/go
*/

import "math"

// Pi é uma proporção numérica definida pela relação entre o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

// Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é uma função pública, pois inicia com _
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
