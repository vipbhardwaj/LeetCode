func reversePrefix(word string, ch byte) string {
    chi := -1

    for i, c := range word {
        if byte(c) == ch {
            chi = i
            break
        }
    }
    
    if chi == -1 {
        return word
    }

    bytes := []byte(word)
    for i := range chi/2 + 1 {
        temp := bytes[i]
        bytes[i] = bytes[chi-i]
        bytes[chi-i] = temp
    }

    return string(bytes)
}
