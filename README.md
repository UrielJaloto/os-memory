# Simulador de Gerenciamento de Memória com Paginação

Este projeto é uma implementação em Go de um simulador de gerenciamento de memória que utiliza o método de paginação. O programa simula a alocação não contígua de memória para processos, gerenciando quadros livres e tabelas de páginas.

## Funcionalidades

  * **Visualizar Memória:** Exibe o status de todos os quadros da memória física (livre ou ocupado), o percentual de memória livre e uma amostra dos dados contidos nos quadros ocupados.
  * **Criar Processo:** Aloca um novo processo na memória. O usuário informa um ID e o tamanho em bytes. O simulador calcula as páginas necessárias, encontra quadros livres e atualiza a tabela de páginas do processo.
  * **Visualizar Tabela de Páginas:** Exibe o mapeamento de páginas para quadros de um processo específico, dado seu ID.
  * **Validações:** O sistema valida falta de memória, tamanhos de processo que excedem o máximo configurado, IDs de processo duplicados e entradas inválidas.

## Configuração

O simulador é configurado através do arquivo `config.json`, que deve estar na raiz do projeto.

```json
{
  "tamanho_memoria_fisica": 128,
  "tamanho_pagina": 16,
  "tamanho_maximo_processo": 64
}
```

  * `tamanho_memoria_fisica`: O tamanho total da memória física em bytes.
  * `tamanho_pagina`: O tamanho de cada página (e, consequentemente, de cada quadro) em bytes.
  * `tamanho_maximo_processo`: O tamanho máximo permitido para a memória lógica de um único processo.

**Importante:** Todos os valores de configuração devem ser potências de dois. O programa não será iniciado se os valores não atenderem a este requisito.

## Como Executar

Para compilar e executar o projeto, utilize o seguinte comando Go a partir do diretório raiz:

```sh
go run main.go
```

Certifique-se de que o `config.json` existe e está formatado corretamente antes de executar.