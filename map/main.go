// ...existing code...
package main

import "fmt"

func main() {
    fmt.Println("map in go lang")

    lang := make(map[string]string)

    lang["js"] = "javascript"
    lang["py"] = "python"
    lang["go"] = "golang"

    fmt.Println("before delete:", lang)

    // remove a key
    delete(lang, "py")
    fmt.Println("after delete py:", lang)

    // check if key exists
    v, ok := lang["py"]
    fmt.Println("py exists?", ok, "value:", v) // ok == false, v == ""

    // deleting a non-existent key is safe (no panic)
    delete(lang, "ruby")

    // deleting from a nil map is also safe (no-op)
    var nilMap map[string]int
    delete(nilMap, "any") // no panic

    // use-case: remove keys while iterating (allowed but iteration order is undefined)
    for k := range lang {
        if k == "js" {
            delete(lang, k)
        }
    }
    fmt.Println("after conditional delete during iteration:", lang)
}