package main

import (
	"bufio"
	"fmt"
	"log"
	mrand "math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf("\nResult: %s\n", NewHoptass())
}

func diceRoll() int {
	mrand.Seed(time.Now().UnixNano())
	return mrand.Intn(5)
}

func NewHoptass() string {
	preorg := bufio.NewReader(os.Stdin)
	fmt.Print("Type to change str to passphrase: ")
	org, err := preorg.ReadString('\n')
	if err != nil {
		panic(err)
	}
	org = strings.Trim(org, "\n")
	return hoptChanger(org)
}

func hoptChanger(org string) string {
	str := make(map[string]string)
	var ful string

	for _, nex := range org {
		s_nex := string(nex)
		s_nex = strings.ToLower(s_nex)

		str["a_1"] = `a`
		str["a_2"] = `A`
		str["a_3"] = `@`
		str["a_4"] = `o.`

		str["b_1"] = `b`
		str["b_2"] = `B`
		str["b_3"] = `'o`
		str["b_4"] = `|.`

		str["c_1"] = `c`
		str["c_2"] = `C`
		str["c_3"] = `<`
		str["c_4"] = `|=`

		str["d_1"] = `d`
		str["d_2"] = `D`
		str["d_3"] = `o'`
		str["d_4"] = `,|`

		str["e_1"] = `e`
		str["e_2"] = `E`
		str["e_3"] = `o_`
		str["e_4"] = `.=`

		str["f_1"] = `f`
		str["f_2"] = `F`
		str["f_3"] = `|=`
		str["f_4"] = `\`

		str["g_1"] = `g`
		str["g_2"] = `G`
		str["g_3"] = `_o'`
		str["g_4"] = `&'`

		str["h_1"] = `h`
		str["h_2"] = `H`
		str["h_3"] = `|]`
		str["h_4"] = `'n`

		str["i_1"] = `i`
		str["i_2"] = `I`
		str["i_3"] = `|'`
		str["i_4"] = `:`

		str["j_1"] = `j`
		str["j_2"] = `j`
		str["j_3"] = `_/'`
		str["j_4"] = `;`

		str["k_1"] = `k`
		str["k_2"] = `K`
		str["k_3"] = `|<`
		str["k_4"] = `[`

		str["l_1"] = `l`
		str["l_2"] = `L`
		str["l_3"] = `|`
		str["l_4"] = `!`

		str["m_1"] = `m`
		str["m_2"] = `M`
		str["m_3"] = `nx2`
		str["m_4"] = `]]]`

		str["n_1"] = `n`
		str["n_2"] = `N`
		str["n_3"] = `|'|`
		str["n_4"] = `[]`

		str["o_1"] = `o`
		str["o_2"] = `O`
		str["o_3"] = `.`
		str["o_4"] = `,`

		str["p_1"] = `p`
		str["p_2"] = `P`
		str["p_3"] = `,o`
		str["p_4"] = `,@`

		str["q_1"] = `q`
		str["q_2"] = `Q`
		str["q_3"] = `o,`
		str["q_4"] = `9`

		str["r_1"] = `r`
		str["r_2"] = `R`
		str["r_3"] = `|^`
		str["r_4"] = `[`

		str["s_1"] = `s`
		str["s_2"] = `S`
		str["s_3"] = `/'`
		str["s_4"] = `,/`

		str["t_1"] = `t`
		str["t_2"] = `T`
		str["t_3"] = `+`
		str["t_4"] = `*`

		str["u_1"] = `u`
		str["u_2"] = `U`
		str["u_3"] = `|_|`
		str["u_4"] = `]]`

		str["v_1"] = `v`
		str["v_2"] = `V`
		str["v_3"] = `\/`
		str["v_4"] = `<`

		str["w_1"] = `w`
		str["w_2"] = `W`
		str["w_3"] = `\\/`
		str["w_4"] = `vx2`

		str["x_1"] = `x`
		str["x_2"] = `X`
		str["x_3"] = `><`
		str["x_4"] = `}{`

		str["y_1"] = `y`
		str["y_2"] = `Y`
		str["y_3"] = `7'`
		str["y_4"] = `,v`

		str["z_1"] = `z`
		str["z_2"] = `Z`
		str["z_3"] = `2'`
		str["z_4"] = `'/,`

		str["1_1"] = `1`
		str["1_2"] = `|`
		str["1_3"] = `l`
		str["1_4"] = `!`

		str["2_1"] = `2`
		str["2_2"] = `%`
		str["2_3"] = `'/_`
		str["2_4"] = `z`

		str["3_1"] = `3`
		str["3_2"] = `~`
		str["3_3"] = `>>`
		str["3_4"] = `}`

		str["4_1"] = `4`
		str["4_2"] = `/+`
		str["4_3"] = `^|`
		str["4_4"] = `<|`

		str["5_1"] = `5`
		str["5_2"] = `}'`
		str["5_3"] = `3'`
		str["5_4"] = `^>`

		str["6_1"] = `6'`
		str["6_2"] = `.'`
		str["6_3"] = `@'`
		str["6_4"] = `o'`

		str["7_1"] = `7`
		str["7_2"] = `'/`
		str["7_3"] = `y`
		str["7_4"] = `>`

		str["8_1"] = `8`
		str["8_2"] = `oo`
		str["8_3"] = `&`
		str["8_4"] = `:`

		str["9_1"] = `9`
		str["9_2"] = `?`
		str["9_3"] = `_q`
		str["9_4"] = `,q`

		str["0_1"] = `0`
		str["0_2"] = `o`
		str["0_3"] = `O`
		str["0_4"] = `.`

		var dice int
		if diceRoll() == diceRoll() {
			dice = 1
		} else if diceRoll() == 0 {
			dice = 4
		} else if diceRoll()%2 == 0 {
			dice = 2
		} else {
			dice = 3
		}
		str_nex := fmt.Sprintf("%v", s_nex) + "_" + fmt.Sprintf("%v", dice)
		//fmt.Println(str_nex)
		sne := str[str_nex]
		if sne == "" {
			log.Fatal("bad format\nYou must use, abcdefghijklmnopqrstuvwxyz0123456789")
		}
		ful = ful + sne
	}
	return ful
}
