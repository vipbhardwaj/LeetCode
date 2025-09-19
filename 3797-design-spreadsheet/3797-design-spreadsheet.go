type Spreadsheet struct {
    sheet [][]int
}

func Constructor(rows int) Spreadsheet {
    excel := Spreadsheet {
        sheet: make([][]int, rows+1),
    }
    for i, row := range excel.sheet {
        row = make([]int, 26)
        excel.sheet[i] = row
    }
    return excel
}

func (this *Spreadsheet) SetCell(cell string, value int)  {
    row, col, _ := getCoord(cell)
    this.sheet[row][col] = value
}

func (this *Spreadsheet) ResetCell(cell string)  {
    row, col, _ := getCoord(cell)
    this.sheet[row][col] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
    row1, col1, i := getCoord(formula)
    row2, col2, _ := getCoord(formula[i+1:])
    if row1 == -1 && row2 == -1 {
        return col1 + col2
    } else if row1 == -1 {
        return col1 + this.sheet[row2][col2]
    } else if row2 == -1 {
        return this.sheet[row1][col1] + col2
    } else {
        return this.sheet[row1][col1] + this.sheet[row2][col2]
    }
}

func getCoord(cell string) (int, int, int) {
    if cell[0] == '=' || cell[0] == '+' {
        cell = cell[1:]
    }
    if len(cell) == 0 {
        return -1, -1, -1
    }
    if cell[0] >= '0' && cell[0] <= '9' {
        val := 0
        i := 0
        for ; i < len(cell) && cell[i] >= '0' && cell[i] <= '9'; i++ {
            val = val*10 + int(cell[i]-'0')
        }
        return -1, val, i
    }
    col := int(cell[0] - 'A')
    row := 0
    i := 1
    for ; i < len(cell) && cell[i] >= '0' && cell[i] <= '9'; i++ {
        row = row*10 + int(cell[i]-'0')
    }
    return row, col, i
}
