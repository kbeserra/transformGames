package reelGames

type SymbolSequence struct {
	Name       string                `json:"name"`
	OffSet     int                   `json:"offSet"`
	Components []map[string]struct{} `json:"components"`
}

func (s SymbolSequence) matchesLine(line []string) bool {
	if len(line) < s.OffSet+len(s.Components) {
		return false
	}

	for i, sym := range line[s.OffSet:] {
		if len(s.Components) <= i {
			break
		}
		if _, in := s.Components[i][sym]; !in {
			return false
		}
	}

	return true
}

func (s SymbolSequence) numWays(cols [][]string) int {
	if len(cols) < s.OffSet+len(s.Components) {
		return 0
	}

	count, col_count := 1, 0
	for i, col := range cols[s.OffSet:] {
		if len(s.Components) <= i {
			break
		}
		col_count = 0
		for _, sym := range col {
			if _, in := s.Components[i][sym]; in {
				col_count++
			}
		}
		count *= col_count
		if count == 0 {
			break
		}
	}

	return count
}

func (s SymbolSequence) subSetEq(t SymbolSequence) bool {

	// t must be "shorter" than s. That is, t makes few restrictions.
	if t.OffSet < s.OffSet ||
		s.OffSet+len(s.Components) < t.OffSet+len(t.Components) {
		return false
	}

	// each of t's Components must make fewer restrictions.
	for i, a := range t.Components {
		// a is representing the (t.OffSet + i)th compnent.
		// we must compare that to the corresponding compnent of s.
		b := s.Components[t.OffSet-s.OffSet+i]
		// need b to be a subset of a
		for sym := range b {
			if _, in := a[sym]; !in {
				return false
			}
		}
	}

	return true
}

//
//
//func PaySequenceFromShortCode(shortCode string, length int, cache bool) (*PaySequence, error) {
//
//	if length <= 0 {
//		return nil, errors.New("sequence length must be positive")
//	}
//
//	codeSplit := strings.Split(shortCode, SymbolSequenceShortCodeOAKDelimiter)
//	if len(codeSplit) != 2 {
//		return nil, fmt.Errorf("there must be exactly 1 \"of a kind\" delimiter, %s", SymbolSequenceShortCodeOAKDelimiter)
//	}
//
//	oak, err := strconv.Atoi(codeSplit[0])
//	if err != nil {
//		return nil, err
//	}
//	if oak <= 0 {
//		return nil, errors.New("the number of a kind to match must be positive")
//	}
//	if oak > length {
//		return nil, errors.New("the number of a kind to match must less than or equal the length of the pay sequence")
//	}
//
//	codeSplit = strings.Split(codeSplit[1], SymbolSequenceShortCodeDirectionDelimiter)
//	if len(codeSplit) > 2 {
//		return nil, fmt.Errorf("there must no more than 1 direction delimiter, %s", SymbolSequenceShortCodeDirectionDelimiter)
//	}
//
//	direction := SymbolSequenceShortCodeDirectionLeft
//	if len(codeSplit) == 2 {
//		direction = codeSplit[1]
//	}
//	if direction != SymbolSequenceShortCodeDirectionLeft && direction != SymbolSequenceShortCodeDirectionRight {
//		return nil, fmt.Errorf("direction must be one of %s or %s", SymbolSequenceShortCodeDirectionLeft, SymbolSequenceShortCodeDirectionRight)
//	}
//
//	symb := codeSplit[0]
//
//	code := make([]string, length)
//
//	var i int = 0
//	var d int = 1
//	if direction != SymbolSequenceShortCodeDirectionLeft {
//		i = length - 1
//		d = -1
//	}
//
//	for ; i >= 0 && i < length; i += d {
//		if oak > 0 {
//			code[i] = symb
//		} else if oak == 0 {
//			if direction == SymbolSequenceShortCodeDirectionLeft {
//				code[i] = SymbolSequenceCodeNegativeBackward
//			} else {
//				code[i] = SymbolSequenceCodeNegativeForward
//			}
//		} else {
//			code[i] = SymbolSequenceCodeWild
//		}
//		oak--
//	}
//
//	return &PaySequence{Code: code, Cache: cache}, nil
//}
//
//func PaySequenceFromCode(code string, cache bool) *PaySequence {
//	return &PaySequence{Code: strings.Split(code, seqCodeDelimiter), Cache: cache}
//}
//
//func IsSeqCodeSymbol(symb string) bool {
//	return symb == SymbolSequenceCodeWild ||
//		symb == SymbolSequenceCodeNegativeForward ||
//		symb == SymbolSequenceCodeNegativeBackward
//}
//
//func (seq PaySequence) IsLeft() bool {
//	return !IsSeqCodeSymbol(seq.Code[0])
//}
//
//func (seq PaySequence) IsRight() bool {
//	return !IsSeqCodeSymbol(seq.Code[len(seq.Code)-1])
//}
//
//func (seq PaySequence) SeqLen() int {
//	count := 0
//	for _, symb := range seq.Code {
//		if !IsSeqCodeSymbol(symb) {
//			count++
//		}
//	}
//	return count
//}
//
//func (seq PaySequence) SeqSymbol() (string, error) {
//	foundNoneCode := false
//	var s string
//	for _, symb := range seq.Code {
//		if !IsSeqCodeSymbol(symb) {
//			if !foundNoneCode {
//				foundNoneCode = true
//				s = symb
//			}
//			if s != symb {
//				return SymbolSequenceUnknownSymbol, errors.New("more than 1 none codeing symbols")
//			}
//		}
//	}
//
//	if !foundNoneCode {
//		return SymbolSequenceUnknownSymbol, errors.New("no none codeing symbols")
//	}
//	return s, nil
//}
//
//func (seq PaySequence) CodeString() string {
//	return strings.Join(seq.Code, seqCodeDelimiter)
//}
//
//func (seq PaySequence) ShortCode() (string, error) {
//	symb, err := seq.SeqSymbol()
//	if err != nil {
//		return "", err
//	}
//	directions := []string{}
//	if seq.IsLeft() {
//		directions = append(directions, SymbolSequenceShortCodeDirectionLeft)
//	}
//	if seq.IsRight() {
//		directions = append(directions, SymbolSequenceShortCodeDirectionRight)
//	}
//	direction := strings.Join(directions, SymbolSequenceCodeOr)
//	return fmt.Sprintf("%d%s%s%s%s", seq.SeqLen(), SymbolSequenceShortCodeOAKDelimiter, symb, SymbolSequenceShortCodeDirectionDelimiter, direction), nil
//}
//
//func symbContainsOr(symb string) bool {
//	return strings.Contains(symb, SymbolSequenceCodeOr)
//}
//
//func (seq PaySequence) interpretSeqSymbol(i int, symb string, symbols []string, wildSymbols []string, wildSubstitutable map[string]bool) []string {
//
//	if symbContainsOr(symb) {
//		rtn := make([]string, 0)
//		for _, subSymb := range strings.Split(symb, SymbolSequenceCodeOr) {
//			tmp := seq.interpretSeqSymbol(i, subSymb, symbols, wildSymbols, wildSubstitutable)
//			rtn = append(rtn, tmp...)
//		}
//		return rtn
//	}
//
//	if !IsSeqCodeSymbol(symb) {
//		rtn := []string{symb}
//		if canSub, in := wildSubstitutable[symb]; in && canSub {
//			rtn = append(rtn, wildSymbols...)
//		}
//		return rtn
//	}
//
//	if symb == SymbolSequenceCodeWild {
//		return symbols
//	}
//
//	rtn := make([]string, 0)
//
//	var j int
//	if symb == SymbolSequenceCodeNegativeBackward {
//		j = i - 1
//	} else {
//		j = i + 1
//	}
//	other := seq.interpretSeqSymbol(j, seq.Code[j], symbols, wildSymbols, wildSubstitutable)
//
//	for _, s := range symbols {
//		isInOther := false
//		for _, t := range other {
//			if s == t {
//				isInOther = true
//				break
//			}
//			if !isInOther {
//				rtn = append(rtn, s)
//			}
//		}
//
//	}
//
//	return rtn
//}
//
//func (seq *PaySequence) includeCord(i int) bool {
//	return !IsSeqCodeSymbol(seq.Code[i])
//}
//
//func (seq *PaySequence) CacheSymbs(symbols []string, wildSymbols []string, wildSubstitutable map[string]bool) {
//	if seq.Cache && (seq.cacheSymbols == nil || seq.cacheInclude == nil) {
//		// fmt.Println("generating cache ")
//		seq.cacheSymbols = make([][]string, len(seq.Code))
//		seq.cacheInclude = make([]bool, len(seq.Code))
//		for i, symb := range seq.Code {
//			seq.cacheSymbols[i] = seq.interpretSeqSymbol(i, symb, symbols, wildSymbols, wildSubstitutable)
//			seq.cacheInclude[i] = seq.includeCord(i)
//		}
//	}
//}
//
//func (seq *PaySequence) IsLineMatch(line []string, symbols []string, wildSymbols []string, wildSubstitutable map[string]bool) bool {
//
//	seq.CacheSymbs(symbols, wildSymbols, wildSubstitutable)
//
//	if len(line) != len(seq.Code) {
//		return false
//	}
//
//	for i, symb := range seq.Code {
//
//		var searchSymbs []string
//		if seq.Cache {
//			searchSymbs = seq.cacheSymbols[i]
//		} else {
//			searchSymbs = seq.interpretSeqSymbol(i, symb, symbols, wildSymbols, wildSubstitutable)
//		}
//
//		isIn := false
//		for _, s := range searchSymbs {
//			if line[i] == s {
//				isIn = true
//				break
//			}
//		}
//
//		if !isIn {
//			return false
//		}
//	}
//
//	return true
//}
//
//func (seq *PaySequence) GetLineCells(line []int) [][2]int {
//	cells := make([][2]int, seq.SeqLen())
//	cellIndex := 0
//	for i := range line {
//		if seq.includeCord(i) {
//			// cells[cellIndex] = [2]int{i, line[i]}
//			cells[cellIndex][boardCellRowIndex] = line[i]
//			cells[cellIndex][boardCellColumnIndex] = i
//			cellIndex++
//		}
//	}
//	return cells
//}
//
//func (seq *PaySequence) GetLineMultiplier(line []string) float64 {
//	rtn := 1.
//	for i := range line {
//		if seq.includeCord(i) {
//			if r, in := seq.MultiplyingSymbols[line[i]]; in {
//				rtn *= r
//			}
//		}
//	}
//	return rtn
//}
//
//func (seq *PaySequence) CountWaysMatch(b board, symbols []string, wildSymbols []string, wildSubstitutable map[string]bool) uint {
//
//	seq.CacheSymbs(symbols, wildSymbols, wildSubstitutable)
//
//	if len(b.symbols) != len(seq.Code) {
//		return 0
//	}
//
//	ways := uint(1)
//	for i, symb := range seq.Code {
//
//		var searchSymbs []string
//		var include bool
//
//		if seq.Cache {
//			searchSymbs, include = seq.cacheSymbols[i], seq.cacheInclude[i]
//		} else {
//			searchSymbs, include = seq.interpretSeqSymbol(i, symb, symbols, wildSymbols, wildSubstitutable), seq.includeCord(i)
//		}
//
//		count := uint(0)
//		for _, s := range b.symbols[i] {
//			for _, t := range searchSymbs {
//				if s == t {
//					count++
//					break
//				}
//			}
//		}
//
//		if count == 0 {
//			return 0
//		}
//		if include {
//			ways *= count
//		}
//
//	}
//
//	return ways
//}
//
//func (seq *PaySequence) GetWaysCells(b board, symbols []string, wildSymbols []string, wildSubstitutable map[string]bool) [][2]int {
//
//	if len(b.symbols) != len(seq.Code) {
//		return nil
//	}
//	var searchSymbs []string
//	var include bool
//
//	dom := make([]int, seq.SeqLen())
//	coords := make([][]int, seq.SeqLen())
//
//	cellIndex := 0
//	for i, col := range b.symbols {
//		if seq.Cache {
//			searchSymbs, include = seq.cacheSymbols[i], seq.cacheInclude[i]
//		} else {
//			searchSymbs, include = seq.interpretSeqSymbol(i, seq.Code[i], symbols, wildSymbols, wildSubstitutable), seq.includeCord(i)
//		}
//		if include {
//			dom[cellIndex] = i
//			coords[cellIndex] = make([]int, 0)
//			for j, s := range col {
//				for _, t := range searchSymbs {
//					if t == s {
//						coords[cellIndex] = append(coords[cellIndex], j)
//						break
//					}
//				}
//			}
//			cellIndex++
//		}
//	}
//
//	n := 1
//	for _, c := range coords {
//		n *= len(c)
//	}
//	if n == 0 {
//		return nil
//	}
//
//	cells := make([][2]int, n*len(coords))
//	bounds := func(i int) int { return len(coords[i]) }
//	cellIndex = 0
//	for ran := make([]int, len(coords)); ran[0] < bounds(0); IncrementIndexTuple(ran, bounds) {
//
//		for i, r := range ran {
//			// cells[cellIndex] = [2]int{dom[i], coords[i][r]}
//			cells[cellIndex][boardCellRowIndex] = coords[i][r]
//			cells[cellIndex][boardCellColumnIndex] = dom[i]
//			cellIndex++
//		}
//	}
//
//	return cells
//}
//
//func (seq *PaySequence) GetWaysMultiplier(b board, symbols []string, wildSymbols []string, wildSubstitutable map[string]bool) float64 {
//
//	if len(b.symbols) != len(seq.Code) {
//		return 0
//	}
//	var searchSymbs []string
//	var include bool
//
//	dom := make([]int, seq.SeqLen())
//	coords := make([][]int, seq.SeqLen())
//
//	cellIndex := 0
//	for i, col := range b.symbols {
//		if seq.Cache {
//			searchSymbs, include = seq.cacheSymbols[i], seq.cacheInclude[i]
//		} else {
//			searchSymbs, include = seq.interpretSeqSymbol(i, seq.Code[i], symbols, wildSymbols, wildSubstitutable), seq.includeCord(i)
//		}
//		if include {
//			dom[cellIndex] = i
//			coords[cellIndex] = make([]int, 0)
//			for j, s := range col {
//				for _, t := range searchSymbs {
//					if t == s {
//						coords[cellIndex] = append(coords[cellIndex], j)
//						break
//					}
//				}
//			}
//			cellIndex++
//		}
//	}
//
//	n := 1
//	for _, c := range coords {
//		n *= len(c)
//	}
//	if n == 0 {
//		return 0
//	}
//
//	// multiplicative option.
//	// rtn := 1.
//	// for c, rowIndices := range coords {
//	// 	for _, r := range rowIndices {
//	// 		if m, in := seq.multiplyingSymbols[b.symbols[c][r]]; in {
//	// 			rtn *= m
//	// 		}
//	// 	}
//	// }
//
//	rtn := 0.
//	found := false
//	for c, rowIndices := range coords {
//		for _, r := range rowIndices {
//			if m, in := seq.MultiplyingSymbols[b.symbols[c][r]]; in {
//				rtn += m
//				found = true
//			}
//		}
//	}
//
//	if !found {
//		return 1
//	} else {
//		return rtn
//	}
//}
//
//func (p PaySequence) isInitialSegmentOf(q PaySequence) bool {
//	equivalentCodes := func(a, b string) bool {
//		aS := strings.Split(a, SymbolSequenceCodeOr)
//		bS := strings.Split(b, SymbolSequenceCodeOr)
//		if len(aS) != len(bS) {
//			return false
//		}
//		for _, x := range aS {
//			isIn := false
//			for _, y := range bS {
//				if x == y {
//					isIn = true
//					break
//				}
//			}
//			if !isIn {
//				return false
//			}
//		}
//		return true
//	}
//
//	if p.SeqLen() >= q.SeqLen() {
//		return false
//	}
//
//	if p.IsLeft() {
//		if !q.IsLeft() {
//			return false
//		}
//
//		for i := 0; i < p.SeqLen(); i++ {
//			if IsSeqCodeSymbol(p.Code[i]) {
//				break
//			}
//			if IsSeqCodeSymbol(q.Code[i]) {
//				return false
//			}
//			if !equivalentCodes(p.Code[i], q.Code[i]) {
//				return false
//			}
//		}
//		return true
//
//	} else if p.IsRight() {
//		if !q.IsRight() {
//			return false
//		}
//
//		for i := p.SeqLen() - 1; i >= 0; i-- {
//			if IsSeqCodeSymbol(p.Code[i]) {
//				break
//			}
//			if IsSeqCodeSymbol(q.Code[i]) {
//				return false
//			}
//			if !equivalentCodes(p.Code[i], q.Code[i]) {
//				return false
//			}
//		}
//		return true
//	}
//	return false
//}
