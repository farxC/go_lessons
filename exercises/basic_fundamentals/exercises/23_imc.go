package exercises

func Imc(weight float32, height float32) float32 {
	return weight / (height * height)
}

func GetImcClassification(imc float32) string {
	threshold := []float32{
		18.5,
		24.99,
		29.99,
		34.99,
		39.99,
		40,
	}

	classifications := []string{
		"Baixo peso",
		"Peso saúdavel",
		"Sobrepeso",
		"Obesidade grau I",
		"Obesidade grau II",
		"Obesidade grau III",
	}

	for i := range threshold {
		if imc < threshold[i] {
			return classifications[i]
		}
	}

	return classifications[len(classifications)-1]
}
