package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.3" )

func Ori8t3p2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ofEO2t9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgTPltbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vH4eznodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEkyA9fxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXuGYokcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QbtVbdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIH2t7cGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNjUDPH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RpkpXRjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbQJksbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aG2mwSGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XG8kcx54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDQ8WJFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynM0k82pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OWzFSJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHch7jxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFmA3jBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6lk42ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXw6Rk3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sI59iOkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGq4ohgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6hloP7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TR1XyPjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gg2XUZKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqB1AomVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCM8UIUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WeAtrynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PFuX8dHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZcfnAypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSIcPYHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vB2WVz4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38u2gGQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrHrJvklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pA0icCqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQIHUa73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SL6Z8pcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2KlUrqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mJtqXRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhEyBoJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfXyAHP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJPYtjkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UfXEUV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLLtXLNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMk34MiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rL9Jsz5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7jnP9INWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SsCCB8OdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzOB42hjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzjduahKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNYz4iVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrsn2ArTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgZNjJfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHqqzZMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmnGxmS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvqy6eNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYLsvUUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkgWKWrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1hXqjifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2PFJIkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZL4ucdafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFPBAH5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIIL9QEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSTWiGEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbDvh9b7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90KRZc2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xkz7vCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2IdZ3rAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoGSVF0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CmXv0t9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mfuk6vb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dww8igPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PaCuJwLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9TZ0xVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgP9ltfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tehYfHc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dFUoQ3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LHIPenVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waU8HVtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QvTeuMZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8k5J8GCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqPqzzXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFHkcSmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EF8KiCKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3arYglRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzFH6ncIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zhBVzx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwtw8dRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKYe7EvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pV2YXIzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REnEOUaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GIa3A3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNUPApbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xzR9rNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aIlCqjOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URMmUzlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjY8jhQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLW3BttPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dRGWsbeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y44nBSRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjDCk5qjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t43l2sMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 42ItY0UOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAXxcV8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fe2wW8ZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QHElVJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eynsnbl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGxaCVg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJUl9DnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFSIZvyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KwmLI0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfYarruWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AY84VmhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShvhVmaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOrpcuwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htUFljpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkV5SnMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuzYn3iPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxYcuyx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcMzvt2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtLJnxMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esD5d4glWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBIvqlu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJRJDRomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLPvjv0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZbpoztpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KN2b099jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wOT7717Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPlQxh8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiIPM1X0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqe2VKxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRqMd6M1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3wLaBX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntXgknFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iay6GqRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCtGE6HVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfYi0CAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWSY2J7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qcl0H0nCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNs1mhMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GTMZTXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jS28jcg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvzMXcaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3s2GiWbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8RyLYkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PY4fLJ7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmEi808WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 07kT9VNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBn6ldzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smhsN23IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X7WeCtBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBQw7JJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFRTLMLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1Mv6G5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEYGGiVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvQSjB4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3ZpcGc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CjsKIVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func StiNVTx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypiuW2trWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zELNDfgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJiFnkmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2EWpJ80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rg9oXNZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neIALjMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hpwKxYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXXWdlJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0G9KL0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DI2P3J4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tad6d7kWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4fWCkQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjkWcK6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vhlkca9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQut9v6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7mUVRG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S30HMUL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g357BZgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDjEfspfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZvxz81HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLso2T3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sruiqdVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxSKvnN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hK1PiQhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imS4acpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fALsjFWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLYFzasrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcOhuXr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgJHmkLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bMJh2tVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLZfn8TzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJmfgMA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvODCDn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBOKLbpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVDkOpgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZRKzBXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDvBXv1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3a4XNaQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7hXfx9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0z1y9TDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tcG2qzKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiE7mktjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34iOJFb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKMNu0MDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKlmc42WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVA4c8xrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGGLzQxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjWDgxn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJ9aliIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGqqnqTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNhKe5kGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HH5BmZxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZlf6CyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BIUcSRdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8ZI8MeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRFdzhlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOT3n0RRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ihk3pB4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lh0u6DYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmHs1bpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Z5pAK5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISQYxMtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRk8BEBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSNzOowbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bf0oze0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mg7VunX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PlSlwSd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1xLlgGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHLJruluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82eHEVB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rf9MfFL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fZEGQOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uC7ZYXuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pARYW3ecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoHMAB5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOUHorRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MxiBXhfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPCrA8FOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4Jyz639Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRAiu3SQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4LijcVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O0SLcC1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8G4Lo2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4L85qb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UW93LtAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TE0BTDxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPAwMt3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMjN3zjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdudvZJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCGNlQ59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXUjwx6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKZrSJpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXrfejuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func licOT6uiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMCQF5kbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImO3qHqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyBqjkXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nejUMLuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lST2kNOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUmeQTKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xezGg7eqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhLOf2xNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubalDwc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JysbZsuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVwulxWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Vu4nOuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yeAlPfZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4l39tSGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUpdCtXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPktssLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2ItWFDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIcSgkvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func my6QKaqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUVHieDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NbtZwo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XidO57EuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6xiZwtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Of1MgDX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yBXCc7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJPXSBCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyeSraDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2bUb4UmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KVSDfiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSdAvcnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQkafEbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyY2ozjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7MJBwLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func moKT3zcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apDl0OOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lbFENL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToE0W4oxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func noF1PVDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlunwiXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uznp1XTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBOSc5W8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqluG6LeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRv4kAniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func alOADZMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AuGzLZJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1eRdVxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6sEyp3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rImJzwTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWetFIZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ejurh5qhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vcJ6VS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGGyMV6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECY6C98sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4wsN4JVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAsz9PU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmjEV8ZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uExR88dCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ba1n0MKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8r2CT28LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o6m3Hg1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWQNrkKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWdG217gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S498WmEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oC0pG2pBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZaCchehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0R0LHSQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdcCHThRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func friPsxVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srFfo6HKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2qpZqaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func whM1FYSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7JgEGlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jY7bD8TjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUlF0PG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEcWdeSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOQ6iKvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwCjl2pbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5p1S6ed4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ph4QMGo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0j6w5RLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZIGiwwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrHLe1JuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xy92aa4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQYm8AfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsAijHK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkW3RcQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcSG1XwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRrYZwZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtvZ0GupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76WMdHWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARzpBdIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8gafTahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSNU2R9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4C0NBlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXNa6TvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2atigQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sylzd5HjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jnw9hZWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuMJ1QdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VT9GZrW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtzH2QAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CD1DBVXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqbnQUfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQCrlPioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQNxgv5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rwus5FGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFvnttuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cX2AwmcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ef6d0S2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gySwjeVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lErngjJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r82CCZp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCiucuMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDhf8qzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaL8cRnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Euo85WjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lFmgSwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrWJL2DgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6WaRZpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bG59oZM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29ePMGpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7ihqsZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frI9G2fWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJb21GEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haCgsHDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iw8opgj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aOeOx48yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 45FjHrqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiYcF1YRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sffV5dtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HI21LrWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IWedSBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nE4Dzi4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cezgtuukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOKNhK12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KCG3ti5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtFbPEBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzULDWGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t333rYXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1RScFP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qsPwHU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0ftHdivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjU4VfzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJmQlUzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jp2ycJ4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06cOgcQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JlrECzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFtlBhX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALkipBhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxSaXuryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jp4z5jG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aC3xYEiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u14uITVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axwYJljZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Su3KYeSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XiMBCDfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRNvD9wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ScvBmVNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WAAhdRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zcli5ywdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVzujWflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyFJqB7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAKrS9YpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHfmbuueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDSMEboWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShluzVYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2K9LnS0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjrk9DKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1CCaDVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRPGlfN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XgjCDKQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0y6EZytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvJg6vmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqF992DJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcc2C8LYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNqkIvEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fadLUoYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5ExEYSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTzqcnVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fdwppCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZsGjluqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gthzmEOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nQHXknxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hX1DXHwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2ELhc3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ntd1l0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcozStxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tM6pFg02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWJZTLzQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrcBLsA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BmvFudZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tKTYLXqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WO38rORmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psPq8sfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sY6ZI1u5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPoo1v7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzH7BtusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLXYvqtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etQdZH1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucpzI67CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pw13TrtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzuPBkt5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Sc9YDeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0TSjszU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gIgq7iJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8w2s2i9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTNgne0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEf2xOeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n3ShImuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGw4gIupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wb4NymarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1toVxQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsuaejDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNI2rBBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KG1oFYZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JnlJlhOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eiGm2EpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlOikhvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCd5CN3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60aTOA4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwlbPFZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eM4RRAS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWDl71jnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KY8wDWO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGqkmIbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7dbXWdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RK1o8AdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tna7fbyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viE42ennWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZvHED8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKFS9AtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9rOsZnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XkKqxMyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYmfHe0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxTCknXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05Ij71QqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUUYVnnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DISWOgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8JH3i7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWWwVSsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ix8h9TKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STjKRZaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNb3ainnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCujQGFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fP0KhJmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyC9SA8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INycoArvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCX3y7y5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqRcIi9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7srOoE29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6tAFs9DuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwjyGU2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func el1LvWF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tH0L6Qz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VbRrqlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgyLDeOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EA3zeX05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1G7KomyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBeGumHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4P2G5z3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgGxWCcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8AO5JmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkjTPMjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRuwkeJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L78FF1JnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tE7WWZxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FI32u1CdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmQSyTuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmHn9nQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hQ9q3kVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWjW2HBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEJkIco5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmwu5h9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6NKFFFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHuNnvqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZ1SL2P5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGTVsADuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTX7puVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBafq8AVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xuNjAlbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zmww6ErwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YuDU16qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOAXoVS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9OKjPvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ys7NXT0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoml0eS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3rvhJMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbeaLxcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEHSbeYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjPWpCH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQq694BiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSOEQaljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYyUPuq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0HQyzqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AM8HjPXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tvId7rHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rO4pdDtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQnWLGJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHgfHSplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ekf0VXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CypWpChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWaJSR80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5CkBQ7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kxKsQyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmBssnCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDXepz4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KgVd39LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KbnUJxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsbcef0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1K3zgc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XzDIoYPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2WXKRR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xg9c5H3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AggoMgTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZ7LhOTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GTYG6HMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AwtDoYZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LpHTnKgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UDwik3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pg7DeZseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6aRqVXHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHQjccCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSb09zxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eU9EOUboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dYKaeQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxV6WJ0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0SktjXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVMDb785Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkM2l56SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVutmUxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfDGJYKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAzdnAGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Khx8b6J8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nxx04jkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func roddQHVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBF1MKbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2rjYcGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w57PfUY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hm13HkSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mDViCNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggsJ4Us0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gfB6hyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3ezkdmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKe3a3UVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnarU0L8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5DyWgcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLEH6MwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXLZyOupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0MNvZ4AnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiUbvmBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbvyoOBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80P0cqrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elfryehiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpdCmGufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRDlRLo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZfw8BDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkU2MipVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fmuiIYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzC9qD14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72tbAKshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZEOKkH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8j3Z5sXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nozYkTiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Fz4QrKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcyHpByaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPNA2OGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQ3mh8xBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onwRzOZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBgtyMGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLOmUvu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEjCL1N7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tAnvk0eiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FM2NXcz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYYX2yoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rYJX1RzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OGboWcMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cqf3gpQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMwfzE76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJIFkhHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0H3Z9ZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXLFydAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OqxJNGj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APYWLcU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhH20wENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbf4Yj3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5n6KoQBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAiqlf32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPAAnTUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmbC4fMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECNsgV5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbXjEhE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func piUU9REVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6b9omBqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKOqEpiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvDRZ21rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6MoCdQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLXyRPEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CUSY9AEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVn2mwH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXoLGRjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlE1WVZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ci4UsHqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0WeTAHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJlpSwEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4FNaXOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMQnqjQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCXwTiyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aX8wUMXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYVKKcPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQ9kEIN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWKHOakgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5NdWTwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lf0d5DdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E34AjeHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQc6wJyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Of1dj7ZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFoSkbp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GimbkjD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukWXoPSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxmAtAkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uh45od6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ykll7AMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7tC5Sr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnwH39SaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6cUIUEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cbn07XDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqzxyyWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvU0p2dRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 02HammxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0x0cEr7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsLI5GxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8OJaEq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgg4Slj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtpOX9Z9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BYp1UtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUWjBMvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnHIdRg0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Af01yECWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sMMUxHnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5fV7ypd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AflFl5ELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hANDnKLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtnLeAymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03JFDlPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnCaXbocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7zwX0NZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHzYafklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeexJGA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BIR17nO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func voqpPTJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDUcXILrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTgLQe08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dp8caqKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpdJ3FcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MfPMguMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPYdmmI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9e0W5jwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbFFGbDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtqdRuMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qjlByijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2G4cDza4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVu0gAmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1zzh2ivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTZWn4WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mEFTGIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XkXJmOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucM50z7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFqsz1DSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FT9w7v9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxt5VfU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNbh6IUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tepyENOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqYdBwArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JB0BKN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCauSqm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvpHhDpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lx3mHtBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAGHZQegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSSSSS7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUiV1ZHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBRcIShpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PpHCOkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zOOvxDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lhzljbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICPPhZggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvLZTvUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Ww1LCGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84XafnftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44ANviCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QMiGp0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WewflTPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjKCfP90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v73HnhLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HyTFhwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4CIsGM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PdHVuDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDBPFKL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kj4iuAo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYjvkM0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8KW6C59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zv1lPJAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Rk8VUo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnOJdQBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuFQKdV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVHVy7QrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8uGrnC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2ny7xPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouaL6ANnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcZSUxypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MbhrWuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zv3PrPtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybqZ4oTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsZLObdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AM2lIS56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEZtTaDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Srfjr8UvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dbWPcqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4FovDXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tLhXpsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRqX5LQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vkDFbzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFa70ILlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13egwpmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YP6yvoFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8XsE8n1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2isOdPF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rf7RwjISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4k4E23fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKKxuB8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6UypvagvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtBHXdDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jyJ4zXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqMsKrNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kS2yqAQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vIrJC7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWHxE6CoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QbJu9XyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kj2ehihOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzUMxZWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rwk1j547Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDNLpR5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIGEU33LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5h2PLO05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mc3odzx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ISjiPmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nayAsL5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwPqO2zEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3P5YVz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RM8zkcJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lI45XOG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyijSTGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoXRXCZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Brbh3vKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZiTUZfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sad7TATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i06Dz5NGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ym1v3rEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zm63NqmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSgPxmqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ms8UdakkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0eJg3xJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1No7gq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjFy4LyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAcKbYTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xtl9Iyx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgVV9gsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DseYyX1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUcxh98mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMDRGQOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5Nz7LgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onLwygZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XivOc06PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFsgggycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tl8iMKj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3w2qvS6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLVbnVp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJS3XfEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syLD2dnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKvmyvDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzNedaegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaN9Z2YBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZF6yoqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DSnSfoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcDzmUMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xkit5wo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhZqe9KqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLkWhDkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qI8GCZ0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2fUT3YoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhKun80EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWJH5SoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CF3t1w6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QHdKyYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nf0uSgCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALi2tw3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGueo8v2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKsrPO86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nc0ePaHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFNnB4FkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 052aXcSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2t4R94tyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func By2qjvb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sUiQujqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZxxRrcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func coi8VAOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cT9oHyBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mu9D3Vk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeaxIncyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8Wyx7RZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8Ga1nbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HctOpmKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gn8ZwWlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UM89jT6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MH6ixjEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmRS05hkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BX8Eac5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VemWHTa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFtohK7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGVaq9oAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2lkP5Q3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ojSoEGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ay49OGhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWc7QQm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQoyy2wuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdqzcmUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPKVWRVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x61bbDCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yi9pJC7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Se4AZf2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2zNLO8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jI8j1mWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tYuE1NQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHN32AToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOaXR26zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZvAKIlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DQd4yYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3GowsJ0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0K3jhqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXn6XnjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFr59fBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9twh7e3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYVprByyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTPzFxUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BYVZ3ryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISL3UYKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nfrkKlqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoEYqCspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NwwKnZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deQoSziAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhD00WLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIzqsd59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NTDrBohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4n1FWcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7NROINOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOSZrh7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2P3kCkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFAGHmJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYYUThv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCgXoQQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHXatozsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LbZHeJ9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMClSe1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMHzhdAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4C9lvLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALnObfS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkbtKwUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROeseQS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxRNo7VtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dheN8M4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tbV4aZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBH57a5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ei0Nc9qEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b248i47ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8bTUs6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fX1C9X1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWBUIm8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZP6dRKmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOCmcje8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QL1Wsfb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWR0U28ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tiJ6HrO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9mzvjrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yh7EfaKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XfxYehLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYMYtDRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gj6I1X7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djSedx03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QunW5r36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8vvA17bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyYDxubBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5hbVbTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ikZu91jiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nFPk1G0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsDtrWxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqYkPb93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKJLajVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsp8PAS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5ljKGxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGhgcyxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LamDoPOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xL9TnHNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zZOkcLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

