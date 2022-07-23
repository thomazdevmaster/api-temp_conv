package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func Converte(c *gin.Context) {
	// Recebe a escala e o valor pela url
	valor := c.Param("val")
	scale := c.Param("scale")
	// Verifica e converte o valor em float
	val, err := strconv.ParseFloat(valor, 32)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Valor deve ser numérico",
		})
		return
	}

	// Cria as variáveis definindo o tipo
	var cels tempconv.Celsius;
	var fahr tempconv.Fahrenheit;
	// var kelv tempconv.Kelvin;

	// Verifica a escala atual e faz as conversões
	switch scale {
	case "celsius":
		cels = tempconv.Celsius(val);
		fahr = tempconv.CToF(tempconv.Celsius(val));
		// kelv = tempconv.CToK(tempconv.Celsius(val));
	case "fahrenheit":
		cels = tempconv.FToC(tempconv.Fahrenheit(val));
		fahr = tempconv.Fahrenheit(val);
		// kelv = tempconv.FToK(tempconv.Fahrenheit(val));
	// case "kelvin":
	// 	cels = tempconv.KToC(tempconv.Kelvin(val));
	// 	fahr = tempconv.KToF(tempconv.Kelvin(val));
	// 	kelv = tempconv.Kelvin(val));
	}
	
	// Retorna o objeto
	c.JSON(200, gin.H{
		"celsius": cels,
		"fahrenheit": fahr,
		// "kelvin": kelv,
	})
}