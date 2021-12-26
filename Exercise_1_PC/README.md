## Programa��o Concorrente 2021.1
### Aquecimento shared-memory - 1

1. *M�quinas de estados* - Discutimos a m�quina de estados de processos de um sistema operacional UNIX. Agora, indique quais os estados que threads Java podem assumir. Discuta que eventos (incluindo chamadas de m�todos) podem provocar mudan�as de estado.

  
2. *Fork-sleep-join* - Crie um programa que recece um n�mero inteiro *n* como argumento e cria *n* threads. Cada uma dessas threads deve dormir por um tempo aleat�rio de no m�ximo 5 segundos. A main-thread deve esperar todas as threads filhas terminarem de executar para em seguida escrever na sa�da padr�o o valor de *n*.
  

3. *Two-phase sleep* - Crie um programa que recebe um n�mero inteiro *n* como argumento e cria *n* threads. Cada uma dessas threads deve dormir por um tempo aleat�rio de no m�ximo 5 segundos. Depois que  acordar, cada thread deve sortear um outro n�mero aleat�rio *s* (entre 0 e 10). *Somente depois de todas* as *n* threads terminarem suas escolhas (ou seja, ao fim da primeira fase), come�amos a segunda fase. Nesta segunda fase, a n-�sima thread criada deve dormir pelo tempo *s* escolhido pela thread n - 1 (fa�a a contagem de maneira modula, ou seja, a primeira thread dorme conforme o n�mero sorteado pela �ltima).

Obs.1 - Para as quest�es 2 e 3, seus programas devem ser escrito em Java, C ou rust (que tal escrever uma vers�o em cada linguagem?).  
Obs.2 - Para a quest�o 2, fa�a a thread-m�e esperar as filhas de duas maneiras: 1) usando o equivalente � fun��o join em c, como mostrado em aula; 2) usando sem�foros.  
Obs.3 - Para a quest�o 3 use sem�foros para coordenar o trabalho entre as threads.