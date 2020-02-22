package challenge14

type SampleInput struct {
	input    string
	expected int
}

var example1 = SampleInput{
	`10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`,
	31,
}

var example2 = SampleInput{
	`9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL`,
	165,
}

var example3 = SampleInput{
	`157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`,
	13312,
}

var example4 = SampleInput{
	`2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF`,
	180697,
}

var example5 = SampleInput{
	`171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`,
	2210736,
}

var puzzleInput = SampleInput{
	`4 SRWZ, 3 ZGSFW, 1 HVJVQ, 1 RWDSX, 12 BDHX, 1 GDPKF, 23 WFPSM, 1 MPKC => 6 VCWNW
3 BXVJK, 3 WTPN => 4 GRQC
5 KWFD => 9 NMZND
1 DNZQ, 5 CDSP => 3 PFDBV
4 VSPSC, 34 MPKC, 9 DFNVL => 9 PZWSP
5 NTXHM => 9 DBKN
4 JNSP, 4 TCKR, 7 PZWSP => 7 DLHG
12 CNBS, 3 FNPC => 2 SRWZ
3 RWDSX, 4 NHSTB, 2 JNSP => 8 TCKR
24 PGHF, 1 NMZND => 3 RWDSX
1 DLHG => 9 QSVN
6 HVJVQ => 2 QSNCW
4 CHDTJ => 9 FDVNC
1 HBXF, 1 RWDSX => 7 BWSPN
2 ZGSFW, 1 KWFD => 8 JNSP
2 BWSPN, 7 GDPKF, 1 BXVJK => 6 FVQM
2 MHBH => 6 FNPC
2 WTPN, 15 GRQC => 3 ZGSFW
9 LXMLX => 6 CLZT
5 DFNVL, 1 KHCQ => 4 MHLBR
21 CNTFK, 3 XHST => 9 CHDTJ
1 CNTFK => 7 MHBH
1 GMQDW, 34 GDPKF, 2 ZDGPL, 1 HVJVQ, 13 QSVN, 1 QSNCW, 1 BXVJK => 2 SGLGN
1 BMVRK, 1 XHST => 8 XHLNT
23 CXKN => 1 BDKN
121 ORE => 9 XHST
4 NTXHM, 4 FNPC, 15 VCMVN => 8 MPKC
2 ZDGPL, 7 JNSP, 3 FJVMD => 4 GMQDW
1 LXMLX, 2 BWSPN => 2 DNZQ
6 WTPN => 9 KCMH
20 CDSP => 2 VSPSC
2 QSNCW, 1 BDHX, 3 HBXF, 8 PFDBV, 17 ZDGPL, 1 MHLBR, 9 ZGSFW => 8 FDWSG
2 VSFTG, 2 DLHG => 9 BDHX
174 ORE => 5 BMVRK
2 BMVRK => 2 KWFD
3 WTPN, 9 TVJPG => 9 CDSP
191 ORE => 2 CNTFK
9 FDVNC, 1 MHBH => 8 NTXHM
3 NHSTB, 2 BXVJK, 1 JNSP => 1 WFPSM
7 FJVMD => 9 CXKN
3 GDPKF, 10 QSNCW => 7 ZDGPL
7 LPXM, 11 VSPSC => 1 LXMLX
6 RWDSX, 2 NMZND, 1 MPKC => 1 KHCQ
6 RWDSX => 4 QMJK
15 MHBH, 28 DBKN, 12 CNBS => 4 PGHF
20 NMZND, 1 PGHF, 1 BXVJK => 2 LPXM
1 CDSP, 17 BXVJK => 5 NHSTB
12 HVJVQ => 3 VSFTG
2 PGHF, 3 VCMVN, 2 NHSTB => 1 DFNVL
5 FNPC => 9 HBXF
3 DPRL => 4 FJVMD
1 KWFD, 1 TVJPG => 8 VCMVN
1 FDWSG, 1 VCWNW, 4 BDKN, 14 FDVNC, 1 CLZT, 62 SGLGN, 5 QMJK, 26 ZDGPL, 60 KCMH, 32 FVQM, 15 SRWZ => 1 FUEL
3 XHLNT => 8 TVJPG
5 HBXF => 2 HVJVQ
3 CHDTJ, 15 KWFD => 9 WTPN
7 CNTFK => 7 CNBS
1 CNBS => 2 JPDF
5 JNSP => 8 DPRL
11 NTXHM => 8 GDPKF
10 JPDF => 9 BXVJK`,
	751038,
}
