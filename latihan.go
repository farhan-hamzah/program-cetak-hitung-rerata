package main
import "fmt"
func cetakHitungJumlahRata(n float64, m float64)float64{
    var i, j float64
    j = 0
    for n <= m{
        fmt.Println(n) 
        i = i+n
        n++
        j++
    }
    fmt.Print(i, " ")
    return(i/j)

}

func main(){
    var d1, d2, hasil float64
    fmt.Scan(&d1, &d2)
    hasil = cetakHitungJumlahRata(d1, d2)
    fmt.Println(hasil)
}