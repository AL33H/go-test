## Teste de seleção para vaga de Java

### Faça um fork desse projeto e siga as instruções a seguir utilizando esse projeto.

### Pré-requisitos

Implementar apenas a API (Backend)
Versão Java +8 (caso seja Java 8, verificar compatibilidade da sua IDE)
Versão Spring Boot >= 2.4
Banco de dados fica a seu critério (Sql, NoSql)
Seu projeto deve obrigatoriamente ter as anotações: @Repository, @Entity e @Controller
Documentação mínima da API (Swagger ou documento PDF)

### Objetivo
Implementar para empresa de transporte de cargas SigaBem o endpoint para o cálculo do preço do frete:

Você deve calcular o valor total do frete e a data prevista da entrega.

Considerar regras para calcular o valor do frete:
 * CEPs com DDDs iguais tem 50% de desconto no valor do frete e entrega prevista de 1 dia
 * CEPs de estados iguais tem 75% de desconto no valor do frete e entrega prevista de 3 dias
 * CEPs de estados diferentes não deve ser aplicado o desconto no valor do frete e entrega prevista de 10 dias
 * O valor do frete é cobrado pelo peso da encomenda, o valor para cada KG é R$1,00

Seu input de entrada deve ser “peso”, “cepOrigem”, “cepDestino” e “nomeDestinatario“

Você utilizará a API gratuita de consulta de CEP abaixo: 
Documentação da API: https://viacep.com.br/
Exemplo do GET: https://viacep.com.br/ws/<CEP_A_CONSULTAR>/json/

Endpoint pode ser público
Response/Output deve possuir: “vlTotalFrete” e “dataPrevistaEntrega”, “cepOrigem” e “cepDestino”
Deve ser persistido no banco os valores da cotação os valores consultados: “peso”, “cepOrigem”, “cepDestino”, “nomeDestinatario”, “vlTotalFrete”, “dataPrevistaEntrega” e “dataConsulta”



### Critérios de avaliação:
 * Implementação das regras de negócios para o cálculo do frete
 * Boas práticas de programação, arquitetura  e padrões de projetos

### Entrega: 
 * Disponibilizar um link do repositório no GitHub e encaminhar para developer@cd2.com.br


# A Proposta inicial era fazer em Java, contudo repliquei o teste para Go para práticar.


# Tecnologias 
- Language: Go Lang
- Banco de dados: PostgreSQL-dev/prod


# Iniciar Aplicação

### 1. Selecione o diretório abra o prompt de sua preferência, e insira;

```xsdregexp
git clone https://github.com/AL33H/go-test.git
```
### 2. Após Clonar o diretório, deverá conectar um banco de dados localmente;

- Com docker instalado, execute o docker
``` docker-compose -up``` para instânciar um novo banco de dados localmente.

### 3. Para executar o programa, deverá executar o comando ```go run main.go```

### 4. Para testar via postman, basta importar o arquivo.
```json 
{

	"info": {
		"_postman_id": "f4860117-4f6b-4121-9ba0-c0d984d84c05",
		"name": "API-GO-TEST",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "New",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n            \"nome\":\"Aleff R\",\r\n\t\t\t\"cepOrigem\":\"71699-310\",\r\n\t\t\t\"cepDestino\":\"71699-310\",\r\n\t\t\t\"peso\": 32.53\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8080/consultar/new",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"consultar",
						"new"
					],
					"query": [
						{
							"key": null,
							"value": "1",
							"disabled": true
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "FindAll",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/consultar/",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"consultar",
						""
					]
				}
			},
			"response": []
		},
		{
			"name": "FindById",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/consultar/1",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"consultar",
						"1"
					]
				}
			},
			"response": []
		},
		{
			"name": "deleteById",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "localhost:8080/consultar/delete/1",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"consultar",
						"delete",
						"1"
					]
				}
			},
			"response": []
		}
	]
}
```


