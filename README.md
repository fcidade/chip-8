# Chip-8

## Memória

**Tamanho:** 4KB

**Alcance:** 0x000 até 0xFFF

**Bytes não utilizados:** 0x000 até 0x1FF

**Início dos programas:** 0x200

O Chip-8 é capaz de acessar um alcance de 4KB (0x000 a 0xFFF) de memória.
Os primeiros 512 bytes, de 0x000 até 0x1FF, são onde o interpretador era armazenado originalmente e não deve ser sobrescrito por programas.

A maioria dos programas começam na posição **0x200**.

## Registradores

Registradores funcionam como "variáveis globais" do seu programa.

**Registradores:** V0 até VF, 16 registradores de 8-bit cada

**Flag VF:** O registrador VF é normalmente utilizado como uma flag para algumas das instruções.

**Ponteiro de memória (I):** Registrador de 16-bit utilizado como um ponteiro para um endereço da memória, como a memória do chip-8 vai de 0x000 até 0xFFF, o registrador I só utiliza 3 bytes.

Dois registradores de 8-bit também são utilizados para o uso de delay e som:

- **Delay Timer (DT):** Sempre que o valor de DT for diferente de zero, o timer subtrai "1" do próprio valor em uma velocidade de atualização de 60Hz. Quando DT chega a zero, ele desativa.
- **Sound Timer (ST):** O ST funciona da mesma forma que o DT, porém, enquanto o valor de ST for diferente de zero, será ativado um *buzzer (ou um som qualquer).* O som é definido pelo desenvolvedor do interpretador.

O Chip-8 tem também alguns registradores que não são acessíveis pelos programas, mas são importantes para o funcionamento do interpretador:

- **Program Counter (PC):** Armazena o endereço da instrução atual. Costuma ser 16-bit.
- **Stack Pointer (SP):** Armazena o valor do topo da pilha de execução. O SP armazena 16 valores de 16 bits que empilham os endereços gerados pelas subrotinas do interpretador. É capaz de armazenar até 16 níveis de subrotinas.

## Teclado / Input

Originalmente, os computadores que utilizavam o interpretador Chip-8 tinham um teclado com o seguinte layout:

```sql
| 1 | 2 | 3 | c |
| 4 | 5 | 6 | d |
| 7 | 8 | 9 | e |
| a | 0 | b | f |
```

E estas teclas podem ser mapeadas para encaixar no seu teclado, por exemplo:

```sql
| 1 | 2 | 3 | 4 |
| q | w | e | r |
| a | s | d | f |
| z | x | c | v |
```

## Gráficos

A implementação original do Chip-8 possui um display de 64x32 pixels monocromático.

A única forma de desenhar na tela do Chip-8 é através de sprites. Um sprite é composto por um grupo de bytes e podem ter um tamanho de 5x1 até 5x15.

O Chip-8 por padrão carrega no endereço 0x000 os sprites que representam os caracteres do 0 ao F.

Para saber mais:

[Cowgod's Chip-8 Technical Reference](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.4)

## Instruções

O Chip-8 possui um total de 36 instruções.

Todas as suas instruções possuem 2 bytes, onde o primeiro byte é o mais significante.

### Legenda:

- nnn ou addr - Um valor entre 0x000 e 0xFFF indicando um endereço da memória.
- n ou nibble - Um valor entre 0x0 e 0xF.
- x - Um valor entre 0x0 e 0xF normalmente utilizado para indicar um registrador entre V0 e VF.
- y - Um valor entre 0x0 e 0xF também normalmente utilizado para indicar um registrador V0 e VF.
- kk ou byte - Um valor entre 0x00 e 0xFF.

### 0nnn - SYS addr

Esta instrução era utilizada apenas nos computadores que originalmente implementavam o Chip-8, então você pode ignorar completamente.

### 00E0 - CLS

Limpa a tela.

### 00EE - RET

Volta de uma subrotina.

Atribui o valor do topo do Stack Pointer ao Program Counter e remove o mesmo da pilha.

### 1nnn - JP addr

Pula o programa para um endereço específico na memória.

Atribui o valor de *nnn* para o Program Counter.

### 2nnn - CALL addr

Chama subrotina no endereço *nnn*.

Adiciona o endereço atual PC na Stack e atribui o valor de *nnn* ao PC.

### 3xkk - SE Vx, byte

Pula a próxima instrução se V*x* for igual a *kk.*

### 4xkk - SNE Vx, byte

Pula a próxima instrução se Vx for **diferente** de *kk.*

### *5xy0 - SE vx, vy*

Pula a próxima instrução se V*x* for igual a Vy*.*

### 6xkk - LD Vx, byte

Atribui o valor de *kk* no registrador V*x.*

### 7xkk - ADD Vx, byte

Adiciona o valor de *kk* ao valor já existente do registrador *x*.

### 8xy0 - LD Vx, Vy

Atribui o valor de V*y* ao registrador V*x*.

### 8xy1 - OR Vx, Vy

Realiza uma operação *bitwise or* no entre os registradores V*x* e Vy e atribui o resultado no registrador V*x.*

### 8xy2 - AND Vx, Vy

Realiza uma operação *bitwise and* no entre os registradores V*x* e Vy e atribui o resultado no registrador V*x.*

### 8xy3 - XOR Vx, Vy

Realiza uma operação *bitwise xor* no entre os registradores V*x* e Vy e atribui o resultado no registrador V*x.*

### 8xy4 - ADD Vx, Vy

Adiciona o valor de Vy ao valor já existente do registrador V*x*.

Se o resultado da soma for maior que 255, atribui 1 ao registrador Vf e 0 caso não seja. O mesmo será utilizado como um *carry flag*.

### 8xy5 - SUB Vx, Vy

Subtrai o valor de Vy ao valor já existente do registrador V*x*.

Se Vx for maior que Vy, atribui 1 ao registrador Vf, utilizando como um *no borrow flag*.

### 8xy6 - SHR Vx {, Vy}

Move os *bits* do registrador Vx para a direita.

Se o bit menos significante de Vx for igual a 1, atribui 1 ao registrador Vf, se não, 0.

### 8xy7 - SUBN Vx, Vy

Atribui ao registrador Vx o valor de Vy menos Vx.

Se Vy for maior que Vx, atribui 1 ao registrador Vf, utilizando como um *no borrow flag*.

### 8xyE - SHL Vx {, Vy}

Move os *bits* do registrador Vx para a esquerda.

Se o bit mais significante de Vx for igual a 1, atribui 1 ao registrador Vf, se não, 0.

### 9xy0 - SNE Vx, Vy

Pula a próxima instrução se Vx for **diferente** de Vy.

### Annn - LD I, addr

Atribui o valor de *nnn* ao registrador I.

### Bnnn - JP V0, addr

Pula o programa para a localização de *nnn* mais o valor do registrador V0 (nnn + V0).

O valor da soma entre o *nnn* e o V0 é atribuido ao PC.

### Cxkk - RND Vx, byte

Gera um valor randômico entre 0x00 e 0xFF e realiza uma operação binária *bitwise and* com o valor de *kk*.

Em seguida atribui este valor ao registrador V*x.*

### Dxyn - DRW Vx, Vy, nibble

O interpretador pega todos os bytes contidos na memória entre o registrador I até I + nibble.

O *nibble* também pode ser visto como a altura do sprite, já que a largura sempre é 5.

Os bytes carregados são desenhados na tela.

Caso algum byte seja apagado no momento em que um pixel é desenhado, o registrador VF tem seu valor modificado para 1.

### Ex9E - SKP Vx

Pula a próxima instrução se a tecla com o valor Vx estiver pressionada.

### ExA1 - SKNP Vx

Pula a próxima instrução se a tecla com o valor Vx **não** estiver pressionada.

### Fx07 - LD Vx, DT

Atribui o valor de DT ao registrador Vx.

### Fx0A - LD Vx, K

Repete a instrução indefinidamente até o usuário pressionar algum botão, logo em seguida, o valor do botão é armazenado em Vx.

### Fx15 - LD DT, Vx

Atribui o valor do registrador Vx ao registrador DT.

### Fx18 - LD ST, Vx

Atribui o valor do registrador Vx ao registrador ST.

### Fx1E - ADD I, Vx

Soma e atribui o valor de Vx ao valor já existente do registrador I.

### Fx29 - LD F, Vx

Atribui ao registrador I o endereço onde o sprite para o dígito de Vx se encontra.

### Fx33 - LD B, Vx

Salva o digito da casa das centenas de Vx na memória no endereço de I

Salva o digito da casa das dezenas de Vx na memória no endereço de I + 1

Salva o último digito de Vx na memória no endereço de I + 2

### Fx55 - LD [I], Vx

Persiste os valores dos registradores de V0 até Vx na memória, começando pelo endereço do registrador I.

### Fx65 - LD Vx, [I]

Lê valores na memória de os persiste nos registradores de V0 até Vx, começando pelo endereço do registrador I.

# Referências

Todo este artigo foi baseado na documentação provida em: 

[Cowgod's Chip-8 Technical Reference](http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#2.5)

[CHIP-8 - Wikipedia](https://en.wikipedia.org/wiki/CHIP-8)