package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
		return 0, errors.New("input must be a positive integer")
	}
    count := 0
	for {
        if n == 1{
            return count, nil
        }
        count += 1
        if n %2 == 0{
            n = n/2
        } else{
            n = n *3 + 1
        }
        
    }
    return count, nil
}
