ffffffffffffffffffffffffffffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
       fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Crefmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFfmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }ffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Creafmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
te a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Lfmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            ifmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            //fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

    }
    // Return the distances
    return distances
}
 Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
f newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
oop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}
fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ate a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

    }
    // Return the distances
    return distances
}
_, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
 // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffffffffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ffffffffffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
       fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Crefmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFfmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }ffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Creafmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
te a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Lfmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            ifmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            //fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

    }
    // Return the distances
    return distances
}
 Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
f newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
oop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}
fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ate a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

    }
    // Return the distances
    return distances
}
_, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
 // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffffffffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ffffffffffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
       fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Crefmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFfmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }ffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Creafmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
te a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Lfmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            ifmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            //fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

    }
    // Return the distances
    return distances
}
 Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
f newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
oop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}
fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ate a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

    }
    // Return the distances
    return distances
}
_, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
 // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffffffffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ffffffffffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
fffffffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
       fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Crefmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFfmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }ffmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Creafmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
te a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Lfmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            ifmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            //fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

    }
    // Return the distances
    return distances
}
 Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
f newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
oop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}
fmt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
ate a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}

    }
    // Return the distances
    return distances
}
_, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
 // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}
   // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

// Bellman-Ford algorithm
func bellmanFord(graph map[string]map[string]int, start string) map[string]int {
    // Create a map to store the distances
    distances := make(map[string]int)
    // Set the distance to the start node to 0
    distances[start] = 0
    // Loop through the graph
    for node := range graph {
        // If the node isn't the start node, set the distance to infinity
        if node != start {
            distances[node] = math.MaxInt32
        }
        // Loop through the edges
        for edge := range graph[node] {
            // Calculate the new distance to the edge
            newDistance := distances[node] + graph[node][edge]
            // If the new distance is less than the current distance, update the distance
            if newDistance < distances[edge] {
                distances[edge] = newDistance
            }
        }
    }
    // Return the distances
    return distances
}
mt.Println("Hello, playground")
// Function that takes in a string and returns a map of the number of times each letter appears in the string
func letterCount(s string) map[string]int {
    // Create a map to store the letter counts
    letterCounts := make(map[string]int)
    // Loop through the string
    for _, letter := range s {
        // Convert the letter to a string
        letterString := string(letter)
        // Check if the letter is already in the map
        if _, ok := letterCounts[letterString]; ok {
            // If it is, increment the count
            letterCounts[letterString]++
        } else {
            // If it isn't, add it to the map with a count of 1
            letterCounts[letterString] = 1
        }
    }
    // Return the map
    return letterCounts
}

