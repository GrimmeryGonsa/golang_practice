package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b){
        return 0, errors.New("Strings dont have equal size, can't calculate Hamming Dist.")
    }
    distance := 0
    for i:=0; i<len(a); i += 1{
        if a[i] != b[i]{
            distance += 1
        }
    }
    return distance, nil
}
