package hamming
import "errors"
func Distance(a, b string) (int, error) {
	runes1:=[]rune(a)
    runes2:=[]rune(b)
    if len(runes1)!=len(runes2){
        return 0,errors.New("The length should be same!")
    }
    m:=0
    for i:=0;i<len(runes1);i=i+1{
        if runes1[i]!=runes2[i]{
            m++
        }
    }
    return m,nil
}
