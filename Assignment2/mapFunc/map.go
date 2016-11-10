package mapFunc

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    )

//takes a slice of strings
// returns a word:frequency map
func word_count (words []string) map[string]int{
    
    word_freq := make(map[string]int)
    for _, word := range words{
        word_freq[word]++
    }
    return word_freq
}

 type word_struct struct {
        freq int
        word string
    }
    
    // word_struct will be displayed in this format
    func (p word_struct) String() string {
        return fmt.Sprintf("%3d   %s", p.freq, p.word)
    }
    
func main() {
    // reading string input from user
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter series of words ")
    text, _ := reader.ReadString('\n')
   fmt.Println(text)
   
   // change text to all lower characters
   text = strings.ToLower(text)
   
   //strings.Fields divide each word from the string
    words := strings.Fields(text)
    //calling word_count function 
     word_map := word_count(words)
     
     // convert the map to a slice of structures for sorting
    // create pointer to an instance of word_struct
     pws := new(word_struct)
        struct_slice := make([]word_struct, len(word_map))
        ix := 0
        for key, val := range word_map {
            pws.freq = val
            pws.word = key
            struct_slice[ix] = *pws
            ix++
        }
        
        for ix := 0; ix < len(struct_slice); ix++ {
            fmt.Printf("%v\n", struct_slice[ix])
        }
}
