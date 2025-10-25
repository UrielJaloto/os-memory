Trabalho 2: Gerenciamento de Memória com Paginação
Introdução
O objetivo desta atividade é exercitar conceitos relacionados a gerência de memória explorados na Unidade 2 da disciplina. Neste trabalho, você deverá implementar um programa que simula o gerenciamento de memória usando paginação.

Descrição do Trabalho
O programa deve implementar o mecanismo de paginação para alocação não contígua de memória para processos. Isto inclui a implementação dos algoritmos e das estruturas de dados necessárias. A memória física pode ser representada por um vetor de bytes, cujo tamanho deve ser configurável. A memória lógica também pode ser representada por um vetor de bytes, porém o tamanho deverá ser definido no momento da criação do processo. O tamanho máximo de um processo, isto é, da sua memória lógica, deve ser configurável no programa. O tamanho de uma página e, consequentemente, de um quadro também devem ser configuráveis no programa. Assuma, para os tamanhos das memórias física e lógica e da página (quadro), valores que são potências de dois.

Na criação de um processo, sua memória lógica deve ser inicializada (preferencialmente, com valores aleatórios), as páginas da memória lógica devem ser carregadas para quadros da memória física e uma tabela de páginas deve ser criada e inicializada com o mapeamento entre páginas e quadros. Nesta atividade, é suficiente que cada entrada na tabela de páginas possua apenas o número do quadro. Isto é, não é necessário incluir bits auxiliares. Também não é necessário, neste trabalho, implementar algoritmos de substituição de páginas. Para alocar memória para um processo, o gerenciador de memória deverá manter uma estrutura de dados contendo os quadros livres. Pode ser utilizada uma lista encadeada ou um mapa de bits.

Interface do Simulador com o Usuário
A interface do simulador pode ser por linha de comando (terminal) e deve oferecer as seguintes opções para o usuário:

Visualizar memória: esta opção deve exibir o porcentual de memória livre e cada quadro da memória física, com seu respectivo valor. Inicialmente, todos os quadros devem estar livres e vazios.
Criar processo: para criação do processo, o usuário deve informar um número inteiro que identifica o processo e o tamanho do processo em bytes. Se o tamanho informado for maior que o tamanho máximo configurado, uma mensagem deve ser exibida e um novo valor deve ser solicitado. Se não houver memória suficiente para alocar o processo, uma mensagem deve ser exibida e o usuário deve poder solicitar outra opção.
Visualizar tabela de páginas: esta opção deve exibir o tamanho do processo e a tabela de páginas para o processo identificado pelo número inteiro informado pelo usuário.
Requisitos
O programa pode ser escrito em qualquer linguagem que vocês consigam apresentar código e execução em sala de aula;
As seguintes informações devem ser configuráveis no programa:
Tamanho da memória física;
Tamanho da página (quadro);
Tamanho máximo de um processo.
Entrega do Trabalho
O trabalho pode ser realizado em grupos de até três participantes. A entrega deste trabalho consistirá de um arquivo compactado (formato ZIP), contendo: código-fonte do programa, relatório descrevendo brevemente as principais partes do código, as instruções para compilação e execução e os casos de teste executados, com as respectivas saídas observadas na execução.

O relatório deve conter um link para um vídeo de 5 a 10 minutos onde vocês apresentam o código, explicam a implementação e demonstram a execução.

Importante: para que o trabalho seja avaliado, é preciso que o professor consiga compilar e executar o código enviado. Portanto, assegurem que toda a informação necessária está no relatório e todo código necessário foi entregue no Moodle.
