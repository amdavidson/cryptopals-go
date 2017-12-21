package main

import (
    "errors"
    "fmt"
)

func challenge6 () error {
    hammingTest1 := "this is a test"
    hammingTest2 := "wokka wokka!!!"
    distanceCheck := 37

    distance, err := hammingDistance([]byte(hammingTest1), []byte(hammingTest2))

    if err != nil {
        fmt.Println("Hamming function failed.")
        return err
    }

    if distance != distanceCheck {
        fmt.Println("Expected:",distanceCheck)
        fmt.Println("Received:",distance)
        return errors.New("Hamming distance did not match expected value")
    }

    fmt.Println("Challenge 6 success!")
    return nil

}
