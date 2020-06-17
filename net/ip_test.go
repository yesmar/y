// y net © 2019, Ramsey Dow. All rights reserved.
// SPDX-License-Identifier: BSD-2-Clause

package net

import (
	"fmt"
	"log"
	"math/big"
	"net"
	"testing"
)

const (
	ipv4addr = "1.2.3.4"
	ipv6addr = "2402:9400:0000:0000:0000:0000:0000:0001"
)

func TestNipsv4(t *testing.T) {
	testpairs := []struct {
		input    string
		expected uint64
	}{
		{input: fmt.Sprintf("%s/0", ipv4addr), expected: 4294967296},
		{input: fmt.Sprintf("%s/1", ipv4addr), expected: 2147483648},
		{input: fmt.Sprintf("%s/2", ipv4addr), expected: 1073741824},
		{input: fmt.Sprintf("%s/3", ipv4addr), expected: 536870912},
		{input: fmt.Sprintf("%s/4", ipv4addr), expected: 268435456},
		{input: fmt.Sprintf("%s/5", ipv4addr), expected: 134217728},
		{input: fmt.Sprintf("%s/6", ipv4addr), expected: 67108864},
		{input: fmt.Sprintf("%s/7", ipv4addr), expected: 33554432},
		{input: fmt.Sprintf("%s/8", ipv4addr), expected: 16777214},
		{input: fmt.Sprintf("%s/9", ipv4addr), expected: 8388608},
		{input: fmt.Sprintf("%s/10", ipv4addr), expected: 4194304},
		{input: fmt.Sprintf("%s/11", ipv4addr), expected: 2097152},
		{input: fmt.Sprintf("%s/12", ipv4addr), expected: 1048576},
		{input: fmt.Sprintf("%s/13", ipv4addr), expected: 524288},
		{input: fmt.Sprintf("%s/14", ipv4addr), expected: 262144},
		{input: fmt.Sprintf("%s/15", ipv4addr), expected: 131072},
		{input: fmt.Sprintf("%s/16", ipv4addr), expected: 65536},
		{input: fmt.Sprintf("%s/17", ipv4addr), expected: 32768},
		{input: fmt.Sprintf("%s/18", ipv4addr), expected: 16384},
		{input: fmt.Sprintf("%s/19", ipv4addr), expected: 8192},
		{input: fmt.Sprintf("%s/20", ipv4addr), expected: 4096},
		{input: fmt.Sprintf("%s/21", ipv4addr), expected: 2048},
		{input: fmt.Sprintf("%s/22", ipv4addr), expected: 1024},
		{input: fmt.Sprintf("%s/23", ipv4addr), expected: 512},
		{input: fmt.Sprintf("%s/24", ipv4addr), expected: 256},
		{input: fmt.Sprintf("%s/25", ipv4addr), expected: 128},
		{input: fmt.Sprintf("%s/26", ipv4addr), expected: 64},
		{input: fmt.Sprintf("%s/27", ipv4addr), expected: 32},
		{input: fmt.Sprintf("%s/28", ipv4addr), expected: 16},
		{input: fmt.Sprintf("%s/29", ipv4addr), expected: 8},
		{input: fmt.Sprintf("%s/30", ipv4addr), expected: 4},
		{input: fmt.Sprintf("%s/31", ipv4addr), expected: 2},
		{input: fmt.Sprintf("%s/32", ipv4addr), expected: 1},
	}

	for _, pair := range testpairs {
		_, ipn, err := net.ParseCIDR(pair.input)
		if err != nil {
			t.Error("failed to parse", pair.input)
		}
		actual, err := Nipsv4(ipn)
		if err != nil {
			t.Errorf("%s: expected %v, got %v", pair.input, pair.expected, actual)
			log.Fatal(err)
		}
	}
}

func TestNilNipsv4(t *testing.T) {
	actual, err := Nipsv4(nil)
	if err == nil {
		t.Error("Expected nil, got", actual)
	}
}

func TestNipsv4WithIPv6(t *testing.T) {
	input := fmt.Sprintf("%s/64", ipv6addr)
	_, ipn, err := net.ParseCIDR(input)
	if err != nil {
		t.Error("failed to parse", input)
	}
	actual, err := Nipsv4(ipn)
	if err == nil {
		t.Error("Expected nil, got", actual)
	}
}

func TestNipsv6(t *testing.T) {
	// /8—/65 are all too big to fit within a big.Int. As a workaround, we use the
	// big.Float Int() conversion hack. Note that we skip /0–/7 because, at present,
	// the number of IPs within those subnets is unrealistic to work with.

	/* /8 */
	bf := big.NewFloat(1329227995784915872903807060280344576)
	capacity8 := new(big.Int)
	bf.Int(capacity8)

	/* /9 */
	bf = big.NewFloat(664613997892457936451903530140172288)
	capacity9 := new(big.Int)
	bf.Int(capacity9)

	/* /10 */
	bf = big.NewFloat(332306998946228968225951765070086144)
	capacity10 := new(big.Int)
	bf.Int(capacity10)

	/* /11 */
	bf = big.NewFloat(166153499473114484112975882535043072)
	capacity11 := new(big.Int)
	bf.Int(capacity11)

	/* /12 */
	bf = big.NewFloat(83076749736557242056487941267521536)
	capacity12 := new(big.Int)
	bf.Int(capacity12)

	/* /13 */
	bf = big.NewFloat(41538374868278621028243970633760768)
	capacity13 := new(big.Int)
	bf.Int(capacity13)

	/* /14 */
	bf = big.NewFloat(20769187434139310514121985316880384)
	capacity14 := new(big.Int)
	bf.Int(capacity14)

	/* /15 */
	bf = big.NewFloat(10384593717069655257060992658440192)
	capacity15 := new(big.Int)
	bf.Int(capacity15)

	/* /16 */
	bf = big.NewFloat(5192296858534827628530496329220096)
	capacity16 := new(big.Int)
	bf.Int(capacity16)

	/* /17 */
	bf = big.NewFloat(2596148429267413814265248164610048)
	capacity17 := new(big.Int)
	bf.Int(capacity17)

	/* /18 */
	bf = big.NewFloat(1298074214633706907132624082305024)
	capacity18 := new(big.Int)
	bf.Int(capacity18)

	/* /19 */
	bf = big.NewFloat(649037107316853453566312041152512)
	capacity19 := new(big.Int)
	bf.Int(capacity19)

	/* /20 */
	bf = big.NewFloat(324518553658426726783156020576256)
	capacity20 := new(big.Int)
	bf.Int(capacity20)

	/* /21 */
	bf = big.NewFloat(162259276829213363391578010288128)
	capacity21 := new(big.Int)
	bf.Int(capacity21)

	/* /22 */
	bf = big.NewFloat(81129638414606681695789005144064)
	capacity22 := new(big.Int)
	bf.Int(capacity22)

	/* /23 */
	bf = big.NewFloat(40564819207303340847894502572032)
	capacity23 := new(big.Int)
	bf.Int(capacity23)

	/* /24 */
	bf = big.NewFloat(20282409603651670423947251286016)
	capacity24 := new(big.Int)
	bf.Int(capacity24)

	/* /25 */
	bf = big.NewFloat(10141204801825835211973625643008)
	capacity25 := new(big.Int)
	bf.Int(capacity25)

	/* /26 */
	bf = big.NewFloat(5070602400912917605986812821504)
	capacity26 := new(big.Int)
	bf.Int(capacity26)

	/* /27 */
	bf = big.NewFloat(2535301200456458802993406410752)
	capacity27 := new(big.Int)
	bf.Int(capacity27)

	/* /28 */
	bf = big.NewFloat(1267650600228229401496703205376)
	capacity28 := new(big.Int)
	bf.Int(capacity28)

	/* /29 */
	bf = big.NewFloat(633825300114114700748351602688)
	capacity29 := new(big.Int)
	bf.Int(capacity29)

	/* /30 */
	bf = big.NewFloat(316912650057057350374175801344)
	capacity30 := new(big.Int)
	bf.Int(capacity30)

	/* /31 */
	bf = big.NewFloat(158456325028528675187087900672)
	capacity31 := new(big.Int)
	bf.Int(capacity31)

	/* /32 */
	bf = big.NewFloat(79228162514264337593543950336)
	capacity32 := new(big.Int)
	bf.Int(capacity32)

	/* /33 */
	bf = big.NewFloat(39614081257132168796771975168)
	capacity33 := new(big.Int)
	bf.Int(capacity33)

	/* /34 */
	bf = big.NewFloat(19807040628566084398385987584)
	capacity34 := new(big.Int)
	bf.Int(capacity34)

	/* /35 */
	bf = big.NewFloat(9903520314283042199192993792)
	capacity35 := new(big.Int)
	bf.Int(capacity35)

	/* /36 */
	bf = big.NewFloat(4951760157141521099596496896)
	capacity36 := new(big.Int)
	bf.Int(capacity36)

	/* /37 */
	bf = big.NewFloat(2475880078570760549798248448)
	capacity37 := new(big.Int)
	bf.Int(capacity37)

	/* /38 */
	bf = big.NewFloat(1237940039285380274899124224)
	capacity38 := new(big.Int)
	bf.Int(capacity38)

	/* /39 */
	bf = big.NewFloat(618970019642690137449562112)
	capacity39 := new(big.Int)
	bf.Int(capacity39)

	/* /40 */
	bf = big.NewFloat(309485009821345068724781056)
	capacity40 := new(big.Int)
	bf.Int(capacity40)

	/* /41 */
	bf = big.NewFloat(154742504910672534362390528)
	capacity41 := new(big.Int)
	bf.Int(capacity41)

	/* /42 */
	bf = big.NewFloat(77371252455336267181195264)
	capacity42 := new(big.Int)
	bf.Int(capacity42)

	/* /43 */
	bf = big.NewFloat(38685626227668133590597632)
	capacity43 := new(big.Int)
	bf.Int(capacity43)

	/* /44 */
	bf = big.NewFloat(19342813113834066795298816)
	capacity44 := new(big.Int)
	bf.Int(capacity44)

	/* /45 */
	bf = big.NewFloat(9671406556917033397649408)
	capacity45 := new(big.Int)
	bf.Int(capacity45)

	/* /46 */
	bf = big.NewFloat(4835703278458516698824704)
	capacity46 := new(big.Int)
	bf.Int(capacity46)

	/* /47 */
	bf = big.NewFloat(2417851639229258349412352)
	capacity47 := new(big.Int)
	bf.Int(capacity47)

	/* /48 */
	bf = big.NewFloat(1208925819614629174706176)
	capacity48 := new(big.Int)
	bf.Int(capacity48)

	/* /49 */
	bf = big.NewFloat(604462909807314587353088)
	capacity49 := new(big.Int)
	bf.Int(capacity49)

	/* /50 */
	bf = big.NewFloat(302231454903657293676544)
	capacity50 := new(big.Int)
	bf.Int(capacity50)

	/* /51 */
	bf = big.NewFloat(151115727451828646838272)
	capacity51 := new(big.Int)
	bf.Int(capacity51)

	/* /52 */
	bf = big.NewFloat(75557863725914323419136)
	capacity52 := new(big.Int)
	bf.Int(capacity52)

	/* /53 */
	bf = big.NewFloat(37778931862957161709568)
	capacity53 := new(big.Int)
	bf.Int(capacity53)

	/* /54 */
	bf = big.NewFloat(18889465931478580854784)
	capacity54 := new(big.Int)
	bf.Int(capacity54)

	/* /55 */
	bf = big.NewFloat(9444732965739290427392)
	capacity55 := new(big.Int)
	bf.Int(capacity55)

	/* /56 */
	bf = big.NewFloat(4722366482869645213696)
	capacity56 := new(big.Int)
	bf.Int(capacity56)

	/* /57 */
	bf = big.NewFloat(2361183241434822606848)
	capacity57 := new(big.Int)
	bf.Int(capacity57)

	/* /58 */
	bf = big.NewFloat(1180591620717411303424)
	capacity58 := new(big.Int)
	bf.Int(capacity58)

	/* /59 */
	bf = big.NewFloat(590295810358705651712)
	capacity59 := new(big.Int)
	bf.Int(capacity59)

	/* /60 */
	bf = big.NewFloat(295147905179352825856)
	capacity60 := new(big.Int)
	bf.Int(capacity60)

	/* /61 */
	bf = big.NewFloat(147573952589676412928)
	capacity61 := new(big.Int)
	bf.Int(capacity61)

	/* /62 */
	bf = big.NewFloat(73786976294838206464)
	capacity62 := new(big.Int)
	bf.Int(capacity62)

	/* /63 */
	bf = big.NewFloat(36893488147419103232)
	capacity63 := new(big.Int)
	bf.Int(capacity63)

	/* /64 */
	bf = big.NewFloat(18446744073709551616)
	capacity64 := new(big.Int)
	bf.Int(capacity64)

	/* /65 */
	bf = big.NewFloat(9223372036854775808)
	capacity65 := new(big.Int)
	bf.Int(capacity65)

	// /66—/128 all fit within a big.Int, so we don't need to use the big.Float (Int) conversion hack.

	testpairs := []struct {
		input    string
		expected *big.Int
	}{
		{input: fmt.Sprintf("%s/8", ipv6addr), expected: capacity8},
		{input: fmt.Sprintf("%s/9", ipv6addr), expected: capacity9},
		{input: fmt.Sprintf("%s/10", ipv6addr), expected: capacity10},
		{input: fmt.Sprintf("%s/11", ipv6addr), expected: capacity11},
		{input: fmt.Sprintf("%s/12", ipv6addr), expected: capacity12},
		{input: fmt.Sprintf("%s/13", ipv6addr), expected: capacity13},
		{input: fmt.Sprintf("%s/14", ipv6addr), expected: capacity14},
		{input: fmt.Sprintf("%s/15", ipv6addr), expected: capacity15},
		{input: fmt.Sprintf("%s/16", ipv6addr), expected: capacity16},
		{input: fmt.Sprintf("%s/17", ipv6addr), expected: capacity17},
		{input: fmt.Sprintf("%s/18", ipv6addr), expected: capacity18},
		{input: fmt.Sprintf("%s/19", ipv6addr), expected: capacity19},
		{input: fmt.Sprintf("%s/20", ipv6addr), expected: capacity20},
		{input: fmt.Sprintf("%s/21", ipv6addr), expected: capacity21},
		{input: fmt.Sprintf("%s/22", ipv6addr), expected: capacity22},
		{input: fmt.Sprintf("%s/23", ipv6addr), expected: capacity23},
		{input: fmt.Sprintf("%s/24", ipv6addr), expected: capacity24},
		{input: fmt.Sprintf("%s/25", ipv6addr), expected: capacity25},
		{input: fmt.Sprintf("%s/26", ipv6addr), expected: capacity26},
		{input: fmt.Sprintf("%s/27", ipv6addr), expected: capacity27},
		{input: fmt.Sprintf("%s/28", ipv6addr), expected: capacity28},
		{input: fmt.Sprintf("%s/29", ipv6addr), expected: capacity29},
		{input: fmt.Sprintf("%s/30", ipv6addr), expected: capacity30},
		{input: fmt.Sprintf("%s/31", ipv6addr), expected: capacity31},
		{input: fmt.Sprintf("%s/32", ipv6addr), expected: capacity32},
		{input: fmt.Sprintf("%s/33", ipv6addr), expected: capacity33},
		{input: fmt.Sprintf("%s/34", ipv6addr), expected: capacity34},
		{input: fmt.Sprintf("%s/35", ipv6addr), expected: capacity35},
		{input: fmt.Sprintf("%s/36", ipv6addr), expected: capacity36},
		{input: fmt.Sprintf("%s/37", ipv6addr), expected: capacity37},
		{input: fmt.Sprintf("%s/38", ipv6addr), expected: capacity38},
		{input: fmt.Sprintf("%s/39", ipv6addr), expected: capacity39},
		{input: fmt.Sprintf("%s/40", ipv6addr), expected: capacity40},
		{input: fmt.Sprintf("%s/41", ipv6addr), expected: capacity41},
		{input: fmt.Sprintf("%s/42", ipv6addr), expected: capacity42},
		{input: fmt.Sprintf("%s/43", ipv6addr), expected: capacity43},
		{input: fmt.Sprintf("%s/44", ipv6addr), expected: capacity44},
		{input: fmt.Sprintf("%s/45", ipv6addr), expected: capacity45},
		{input: fmt.Sprintf("%s/46", ipv6addr), expected: capacity46},
		{input: fmt.Sprintf("%s/47", ipv6addr), expected: capacity47},
		{input: fmt.Sprintf("%s/48", ipv6addr), expected: capacity48},
		{input: fmt.Sprintf("%s/49", ipv6addr), expected: capacity49},
		{input: fmt.Sprintf("%s/50", ipv6addr), expected: capacity50},
		{input: fmt.Sprintf("%s/51", ipv6addr), expected: capacity51},
		{input: fmt.Sprintf("%s/52", ipv6addr), expected: capacity52},
		{input: fmt.Sprintf("%s/53", ipv6addr), expected: capacity53},
		{input: fmt.Sprintf("%s/54", ipv6addr), expected: capacity54},
		{input: fmt.Sprintf("%s/55", ipv6addr), expected: capacity55},
		{input: fmt.Sprintf("%s/56", ipv6addr), expected: capacity56},
		{input: fmt.Sprintf("%s/57", ipv6addr), expected: capacity57},
		{input: fmt.Sprintf("%s/58", ipv6addr), expected: capacity58},
		{input: fmt.Sprintf("%s/59", ipv6addr), expected: capacity59},
		{input: fmt.Sprintf("%s/60", ipv6addr), expected: capacity60},
		{input: fmt.Sprintf("%s/61", ipv6addr), expected: capacity61},
		{input: fmt.Sprintf("%s/62", ipv6addr), expected: capacity62},
		{input: fmt.Sprintf("%s/63", ipv6addr), expected: capacity63},
		{input: fmt.Sprintf("%s/64", ipv6addr), expected: capacity64},
		{input: fmt.Sprintf("%s/65", ipv6addr), expected: capacity65},
		{input: fmt.Sprintf("%s/66", ipv6addr), expected: big.NewInt(4611686018427387904)},
		{input: fmt.Sprintf("%s/67", ipv6addr), expected: big.NewInt(2305843009213693952)},
		{input: fmt.Sprintf("%s/68", ipv6addr), expected: big.NewInt(1152921504606846976)},
		{input: fmt.Sprintf("%s/69", ipv6addr), expected: big.NewInt(576460752303423488)},
		{input: fmt.Sprintf("%s/70", ipv6addr), expected: big.NewInt(288230376151711744)},
		{input: fmt.Sprintf("%s/71", ipv6addr), expected: big.NewInt(144115188075855872)},
		{input: fmt.Sprintf("%s/72", ipv6addr), expected: big.NewInt(72057594037927936)},
		{input: fmt.Sprintf("%s/73", ipv6addr), expected: big.NewInt(36028797018963968)},
		{input: fmt.Sprintf("%s/74", ipv6addr), expected: big.NewInt(18014398509481985)},
		{input: fmt.Sprintf("%s/75", ipv6addr), expected: big.NewInt(9007199254740992)},
		{input: fmt.Sprintf("%s/76", ipv6addr), expected: big.NewInt(4503599627370496)},
		{input: fmt.Sprintf("%s/77", ipv6addr), expected: big.NewInt(2251799813685248)},
		{input: fmt.Sprintf("%s/78", ipv6addr), expected: big.NewInt(1125899906842624)},
		{input: fmt.Sprintf("%s/79", ipv6addr), expected: big.NewInt(562949953421312)},
		{input: fmt.Sprintf("%s/80", ipv6addr), expected: big.NewInt(281474976710656)},
		{input: fmt.Sprintf("%s/81", ipv6addr), expected: big.NewInt(140737488355328)},
		{input: fmt.Sprintf("%s/82", ipv6addr), expected: big.NewInt(70368744177664)},
		{input: fmt.Sprintf("%s/83", ipv6addr), expected: big.NewInt(35184372088832)},
		{input: fmt.Sprintf("%s/84", ipv6addr), expected: big.NewInt(17592186044416)},
		{input: fmt.Sprintf("%s/85", ipv6addr), expected: big.NewInt(8796093022208)},
		{input: fmt.Sprintf("%s/86", ipv6addr), expected: big.NewInt(4398046511104)},
		{input: fmt.Sprintf("%s/87", ipv6addr), expected: big.NewInt(2199023255552)},
		{input: fmt.Sprintf("%s/88", ipv6addr), expected: big.NewInt(1099511627776)},
		{input: fmt.Sprintf("%s/89", ipv6addr), expected: big.NewInt(549755813888)},
		{input: fmt.Sprintf("%s/90", ipv6addr), expected: big.NewInt(274877906944)},
		{input: fmt.Sprintf("%s/91", ipv6addr), expected: big.NewInt(137438953472)},
		{input: fmt.Sprintf("%s/92", ipv6addr), expected: big.NewInt(68719476736)},
		{input: fmt.Sprintf("%s/93", ipv6addr), expected: big.NewInt(34359738368)},
		{input: fmt.Sprintf("%s/94", ipv6addr), expected: big.NewInt(17179869184)},
		{input: fmt.Sprintf("%s/95", ipv6addr), expected: big.NewInt(8589934592)},
		{input: fmt.Sprintf("%s/96", ipv6addr), expected: big.NewInt(4294967296)},
		{input: fmt.Sprintf("%s/97", ipv6addr), expected: big.NewInt(2147483648)},
		{input: fmt.Sprintf("%s/98", ipv6addr), expected: big.NewInt(1073741824)},
		{input: fmt.Sprintf("%s/99", ipv6addr), expected: big.NewInt(536870912)},
		{input: fmt.Sprintf("%s/100", ipv6addr), expected: big.NewInt(268435456)},
		{input: fmt.Sprintf("%s/101", ipv6addr), expected: big.NewInt(134217728)},
		{input: fmt.Sprintf("%s/102", ipv6addr), expected: big.NewInt(67108864)},
		{input: fmt.Sprintf("%s/103", ipv6addr), expected: big.NewInt(33554432)},
		{input: fmt.Sprintf("%s/104", ipv6addr), expected: big.NewInt(16777216)},
		{input: fmt.Sprintf("%s/105", ipv6addr), expected: big.NewInt(8388608)},
		{input: fmt.Sprintf("%s/106", ipv6addr), expected: big.NewInt(4194304)},
		{input: fmt.Sprintf("%s/107", ipv6addr), expected: big.NewInt(2097152)},
		{input: fmt.Sprintf("%s/108", ipv6addr), expected: big.NewInt(1048576)},
		{input: fmt.Sprintf("%s/109", ipv6addr), expected: big.NewInt(524288)},
		{input: fmt.Sprintf("%s/110", ipv6addr), expected: big.NewInt(262144)},
		{input: fmt.Sprintf("%s/111", ipv6addr), expected: big.NewInt(131072)},
		{input: fmt.Sprintf("%s/112", ipv6addr), expected: big.NewInt(65536)},
		{input: fmt.Sprintf("%s/113", ipv6addr), expected: big.NewInt(32768)},
		{input: fmt.Sprintf("%s/114", ipv6addr), expected: big.NewInt(16384)},
		{input: fmt.Sprintf("%s/115", ipv6addr), expected: big.NewInt(8192)},
		{input: fmt.Sprintf("%s/116", ipv6addr), expected: big.NewInt(4096)},
		{input: fmt.Sprintf("%s/117", ipv6addr), expected: big.NewInt(2048)},
		{input: fmt.Sprintf("%s/118", ipv6addr), expected: big.NewInt(1024)},
		{input: fmt.Sprintf("%s/119", ipv6addr), expected: big.NewInt(512)},
		{input: fmt.Sprintf("%s/120", ipv6addr), expected: big.NewInt(256)},
		{input: fmt.Sprintf("%s/121", ipv6addr), expected: big.NewInt(128)},
		{input: fmt.Sprintf("%s/122", ipv6addr), expected: big.NewInt(64)},
		{input: fmt.Sprintf("%s/123", ipv6addr), expected: big.NewInt(32)},
		{input: fmt.Sprintf("%s/124", ipv6addr), expected: big.NewInt(16)},
		{input: fmt.Sprintf("%s/125", ipv6addr), expected: big.NewInt(8)},
		{input: fmt.Sprintf("%s/126", ipv6addr), expected: big.NewInt(4)},
		{input: fmt.Sprintf("%s/127", ipv6addr), expected: big.NewInt(2)},
		{input: fmt.Sprintf("%s/128", ipv6addr), expected: big.NewInt(1)},
	}

	for _, pair := range testpairs {
		_, ipn, err := net.ParseCIDR(pair.input)
		if err != nil {
			t.Error("failed to parse", pair.input)
		}
		actual, err := Nipsv6(ipn)
		if err != nil {
			t.Errorf("%s: expected %v, got %v", pair.input, pair.expected, actual)
			log.Fatal(err)
		}
	}
}

func TestNilNipsv6(t *testing.T) {
	actual, err := Nipsv6(nil)
	if err == nil {
		t.Error("Expected nil, got", actual)
	}
}

func TestNipsv6WithIPv4(t *testing.T) {
	input := fmt.Sprintf("%s/24", ipv4addr)
	_, ipn, err := net.ParseCIDR(input)
	if err != nil {
		t.Error("failed to parse", input)
	}
	actual, err := Nipsv6(ipn)
	if err == nil {
		t.Error("Expected nil, got", actual)
	}
}
