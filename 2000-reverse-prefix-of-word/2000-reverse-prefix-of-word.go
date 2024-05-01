func reversePrefix(word string, ch byte) string {
    chi := -1
    // substr := ""

    for i, c := range word {
        // substr += string(c)
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
    // fmt.Println(string(bytes))
    // fmt.Println(word[chi:])
    return string(bytes)
}

// func reverse(str string) string {
//     n := len(str)
//     for i := range n {
//         if i == n/2 {
//             break
//         }
//         temp := str[i]
//         str[i] = str[n-i]
//         str[n-i] = temp
//     }
//     return str
// }