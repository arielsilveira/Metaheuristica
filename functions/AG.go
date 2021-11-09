package functions

import (
	"math"
	"math/rand"
	"time"
)

func GeneticAlgorithm() {
	var population [][]int

}

func roleta(nind int, fo_pop []float64) (choisen int) {
	// Rever a função cria_vetor_float
	var faptidao []float64
	var fracao []float64
	var escala []float64
	sum := 0.0
	fo_min := math.MaxFloat64
	fo_max := -math.MaxFloat64

	for i := 0; i < nind; i++ {
		if fo_pop[i] < fo_min {
			fo_min = fo_pop[i]
		}
		if fo_pop[i] > fo_max {
			fo_max = fo_pop[i]
		}
	}

	tg_alfa := 100 / (fo_max - fo_min)

	for i := 0; i < nind; i++ {
		faptidao[i] = tg_alfa * (fo_max - fo_pop[i])
		sum += faptidao[i]
	}

	for i := 0; i < nind; i++ {
		fracao[i] = faptidao[i] / sum
	}

	escala[0] = fracao[0]

	for i := 1; i < nind; i++ {
		escala[i] = escala[i-1] + fracao[i]
	}

	rand.Seed(time.Now().UnixNano())
	aux := rand.Float64()

	i := 0
	for escala[i] < aux {
		choisen++
	}

	return choisen
}

func roleta_scaling(nind int, fo_pop []float64) int {
	register int j;
  float aux, alfa, soma, media, somatorio, desvio, desvio_aux;
  float *escala, *fracao, *faptidao;
  int escolhido;

  fracao = cria_vetor_float(nind);
  escala = cria_vetor_float(nind);
  faptidao = cria_vetor_float(nind);
  soma = 0;
  somatorio = 0;
  alfa = 2;
  desvio = 0;

  for (int j = 0; j < nind; j++)
   somatorio += fo_pop[j];

  media = somatorio / nind;

  desvio = calcula_desvio_padrao(fo_pop, nind) * media;

  for (int j = 0; j < nind; j++){
    faptidao[j] = media + alfa * desvio - fo_pop[j];
  }

  for (int j = 0; j < nind; j++) {
    if (faptidao[j] < 0) faptidao[j] = 0;
    soma += faptidao[j];
  }

/*
  for (int j = 0; j < nind; j++){
    printf("faptidao[%3d] = %8.2f  fo_pop[%3d] = %8.2f \n",j,faptidao[j],j,fo_pop[j]);
  }
*/
  for (int j = 0; j < nind; j++)
    fracao[j] = faptidao[j] / soma;

  escala[0] = fracao[0];

  for (int j = 1; j < nind; j++)
    escala[j] = escala[j-1] + fracao[j];

  aux = randomico(0,1);
  j = 0;
  while (escala[j] < aux) j++;
  escolhido = j;
  free(fracao);
  free(escala);
  free(faptidao);
	return escolhido
}


/* Esta rotina devolve o individuo escolhido pelo mecanismo da roleta */
func mutacao(s []int, n int) []int {
   var i, j, cid int

   i = rand.Intn(n)
   j = rand.Intn(n)

   for i == j {
     j = rand.Intn(n)
   }

   cid = s[i]
   s[i] = s[j]
   s[j] = cid

   return s
}

/* Operador Crossover OX */
func crossover_OX(pai1 []int,
                  pai2 []int,
                  filho1 []int,
                  filho2 []int,
                  n int)
{

  int ponto_de_corte_1, ponto_de_corte_2, i, j;
  bool existe;
  int *lista_pai1, *lista_pai2;

  ponto_de_corte_1 = randomico(2,(n-1)/2);
  ponto_de_corte_2 = randomico((n+1)/2,n-3);

  /* Copia os genes entre os 2 pontos de corte em cada filho */
  for (int i = ponto_de_corte_1; i <= ponto_de_corte_2; i++) {
    filho1[i] = pai1[i];
    filho2[i] = pai2[i];
  }

  int tam_lista = n - (ponto_de_corte_2 - ponto_de_corte_1 + 1);

  /* Cria uma lista com os genes do outro pai a serem inseridos */
  lista_pai1 = cria_vetor(tam_lista);
  lista_pai2 = cria_vetor(tam_lista);

  i = ponto_de_corte_2 + 1;
  j = 0;
  do {
    /* procura a cidade pai1[i] no filho2*/
    existe = false;
    for(int k = ponto_de_corte_1; k <= ponto_de_corte_2; k++) {
      if (filho2[k] == pai1[i]) existe = true;
    }
    if (! existe) {
      lista_pai1[j] = pai1[i];
      j++;
    }
    if (i == (n-1)) i = 0;
    else i++;

  } while(j < tam_lista);

  i = ponto_de_corte_2 + 1;
  j = 0;
  do {
    /* procura a cidade pai2[i] no filho1*/
    existe = false;
    for(int k = ponto_de_corte_1; k <= ponto_de_corte_2; k++) {
      if (filho1[k] == pai2[i]) existe = true;
    }
    if (! existe) {
      lista_pai2[j] = pai2[i];
      j++;
    }

    if (i == (n-1)) i = 0;
    else i++;

  } while(j < tam_lista);

  /* Completa os genes que faltam */
  i = 0;
  j = ponto_de_corte_2 + 1;
  do {
    filho1[j] = lista_pai2[i];
    filho2[j] = lista_pai1[i];
    i++;
    if (j == (n-1)) j = 0;
    else j++;
  } while (i < tam_lista);

  libera_vetor(lista_pai1);
  libera_vetor(lista_pai2);
}

/* Operador Crossover ERX */
func crossover_ERX( pai1 []int,
                    pai2 []int,
                    filho1 []int,
                    filho2 []int,
                    n int)
{
  int **arestas1, **arestas2, pos_cid_p1, pos_cid_p2, j, m;
  m = 5;
  arestas1 = cria_matriz(n,m);
  for(int i = 0; i < n; i++) inicializa_vetor(arestas1[i], m);
  arestas2 = cria_matriz(n,m);
  for(int i = 0; i < n; i++) inicializa_vetor(arestas2[i], m);

  /* Preenche a matriz de arestas */
  for(int i = 0; i < n; i++) {
    pos_cid_p1 = busca_pos_valor(pai1, i, n);
    pos_cid_p2 = busca_pos_valor(pai2, i, n);
    j = 1;
    if (pos_cid_p1 == 0) {
      arestas1[i][j] = pai1[pos_cid_p1+1];
      arestas2[i][j] = pai1[pos_cid_p1+1];
      j++;
      arestas1[i][j] = pai1[n-1];
      arestas2[i][j] = pai1[n-1];
      j++;
    }
    else if (pos_cid_p1 == (n-1)) {
      arestas1[i][j] = pai1[pos_cid_p1-1];
      arestas2[i][j] = pai1[pos_cid_p1-1];
      j++;
      arestas1[i][j] = pai1[0];
      arestas2[i][j] = pai1[0];
      j++;
    }
    else {
      arestas1[i][j] = pai1[pos_cid_p1-1];
      arestas2[i][j] = pai1[pos_cid_p1-1];
      j++;
      arestas1[i][j] = pai1[pos_cid_p1+1];
      arestas2[i][j] = pai1[pos_cid_p1+1];
      j++;
    }
    if (pos_cid_p2 == 0 ) {
      arestas1[i][j] = pai2[pos_cid_p2+1];
      arestas2[i][j] = pai2[pos_cid_p2+1];
      j++;
      arestas1[i][j] = pai2[n-1];
      arestas2[i][j] = pai2[n-1];
      j++;
    }
    else if (pos_cid_p2 == (n-1)) {
      arestas1[i][j] = pai2[pos_cid_p2-1];
      arestas2[i][j] = pai2[pos_cid_p2-1];
      j++;
      arestas1[i][j] = pai2[0];
      arestas2[i][j] = pai2[0];
      j++;
    }
    else {
      arestas1[i][j] = pai2[pos_cid_p2-1];
      arestas2[i][j] = pai2[pos_cid_p2-1];
      j++;
      arestas1[i][j] = pai2[pos_cid_p2+1];
      arestas2[i][j] = pai2[pos_cid_p2+1];
      j++;
    }

    /* Retira cidades repetidas e armazena o nº de arestas */
    if (arestas1[i][3] == arestas1[i][1] || arestas1[i][3] == arestas1[i][2]) {
      arestas1[i][3] = -1;
      arestas2[i][3] = -1;
      j--;
    }

    if (arestas1[i][4] == arestas1[i][1] || arestas1[i][4] == arestas1[i][2]) {
      arestas1[i][4] = -1;
      arestas2[i][4] = -1;
      j--;
    }
    arestas1[i][0] = j - 1;
    arestas2[i][0] = j - 1;
  }
  int prox_cid, cid_corrente, cid_rota, cid_aleatoria, num_arestas;

  /* Construção do 1º filho */
  filho1[0] = pai1[0];
  atualiza_arestas(arestas1,n,m,0);
  for (int i = 1; i < n; i++) {
    num_arestas = INT_MAX;
    cid_rota = filho1[i-1];
    prox_cid = -1;
    for (int j = 1; j < m; j++) {
      cid_corrente = arestas1[cid_rota][j];
      if (cid_corrente != -1) {
        if (arestas1[cid_corrente][0] < num_arestas) {
          num_arestas = arestas1[cid_corrente][0];
          prox_cid = cid_corrente;
        }
        else if (arestas1[cid_corrente][0] == num_arestas)
          if (cid_corrente < prox_cid)
            prox_cid = cid_corrente;
      }
    }
    if (prox_cid == -1) {
      do {
        cid_aleatoria = (int) rand()%n;
      } while (foi_inserida(filho1, cid_aleatoria, i));
      prox_cid = cid_aleatoria;
    }
    filho1[i] = prox_cid;
    atualiza_arestas(arestas1,n,m,prox_cid);
  }

  /* Construção do 2º filho */
  filho2[0] = pai2[0];
  atualiza_arestas(arestas2,n,m,0);
  for (int i = 1; i < n; i++) {
    num_arestas = INT_MAX;
    cid_rota = filho2[i-1];
    prox_cid = -1;
    for (int j = 1; j < m; j++) {
      cid_corrente = arestas2[cid_rota][j];
      if (cid_corrente != -1) {
        if (arestas2[cid_corrente][0] < num_arestas) {
          num_arestas = arestas2[cid_corrente][0];
          prox_cid = cid_corrente;
        }
        else if (arestas2[cid_corrente][0] == num_arestas)
          if (cid_corrente < prox_cid)
            prox_cid = cid_corrente;
      }
    }
    if (prox_cid == -1) {
      do {
        cid_aleatoria = (int) rand()%n;
      } while (foi_inserida(filho2, cid_aleatoria, i));
      prox_cid = cid_aleatoria;
    }
    filho2[i] = prox_cid;
    atualiza_arestas(arestas2,n,m,prox_cid);
  }
  libera_matriz(arestas1,n);
  libera_matriz(arestas2,n);
}

func binaryTournamentSelection(nind int, fo_pop []float64, jpai1 []int, jpai2 []int){

  /* Cálculo da aptidão */
  float *faptidao = cria_vetor_float(nind);
  float fo_min = INT_MAX;
  float fo_max = INT_MIN;

  for (int j = 0; j < nind; j++){
    if (fo_pop[j] < fo_min) fo_min = fo_pop[j];
    if (fo_pop[j] > fo_max) fo_max = fo_pop[j];
  }

  float tg_alfa = 100 / (fo_max - fo_min);

  for (int j = 0; j < nind; j++){
    faptidao[j] = tg_alfa * (fo_max - fo_pop[j]);
  }
  
  int aux1, aux2;
  
  do{
    /* Selecao do Primeiro pai */
    do{
        aux1 = rand() % (nind/2);
        aux2 = rand() % (nind/2);
    }while(aux1 == aux2);
    (*jpai1) = (faptidao[aux1] >= faptidao[aux2] ? aux1 : aux2);
    
    /* Seleção do Segundo pai */
    do{
        aux1 = rand() % (nind/2);
        aux2 = rand() % (nind/2);
    }while(aux1 == aux2);
    (*jpai2) = (faptidao[aux1] >= faptidao[aux2] ? aux1 : aux2);
    
  }while((*jpai1) != (*jpai2));
}