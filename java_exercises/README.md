## Programação Concorrente 2021.1
### Aquecimento shared-memory - 1

1. *Máquinas de estados* - Discutimos a máquina de estados de processos de um sistema operacional UNIX. Agora, indique quais os estados que threads Java podem assumir. Discuta que eventos (incluindo chamadas de métodos) podem provocar mudanças de estado.

  
2. *Fork-sleep-join* - Crie um programa que recece um número inteiro *n* como argumento e cria *n* threads. Cada uma dessas threads deve dormir por um tempo aleatório de no máximo 5 segundos. A main-thread deve esperar todas as threads filhas terminarem de executar para em seguida escrever na saída padrão o valor de *n*.
  

3. *Two-phase sleep* - Crie um programa que recebe um número inteiro *n* como argumento e cria *n* threads. Cada uma dessas threads deve dormir por um tempo aleatório de no máximo 5 segundos. Depois que  acordar, cada thread deve sortear um outro número aleatório *s* (entre 0 e 10). *Somente depois de todas* as *n* threads terminarem suas escolhas (ou seja, ao fim da primeira fase), começamos a segunda fase. Nesta segunda fase, a n-ésima thread criada deve dormir pelo tempo *s* escolhido pela thread n - 1 (faça a contagem de maneira modula, ou seja, a primeira thread dorme conforme o número sorteado pela última).

Obs.1 - Para as questões 2 e 3, seus programas devem ser escrito em Java, C ou rust (que tal escrever uma versão em cada linguagem?).  
Obs.2 - Para a questão 2, faça a thread-mãe esperar as filhas de duas maneiras: 1) usando o equivalente à função join em c, como mostrado em aula; 2) usando semáforos.  
Obs.3 - Para a questão 3 use semáforos para coordenar o trabalho entre as threads.